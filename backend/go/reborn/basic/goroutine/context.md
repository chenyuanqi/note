
### context
协程如何退出？  
一个协程启动后，大部分情况需要等待里面的代码执行完毕，然后协程会自行退出。但是如果有一种情景，需要让协程提前退出怎么办呢？
```go
func main() {
   var wg sync.WaitGroup
   wg.Add(1)
   go func() {
      defer wg.Done()
      watchDog("【监控狗1】")
   }()
   wg.Wait()
}
func watchDog(name string){
   //开启for select循环，一直后台监控
   for{
      select {
      default:
         // 一直在后台运行，每隔一秒就会打印"监控狗正在监控……"的文字
         fmt.Println(name,"正在监控……")
      }
      time.Sleep(1*time.Second)
   }
}
```
如果需要让监控狗停止监控、退出程序，一个办法是定义一个全局变量，其他地方可以通过修改这个变量发出停止监控狗的通知。然后在协程中先检查这个变量，如果发现被通知关闭就停止监控，退出当前协程。  
但是这种方法需要通过加锁来保证多协程下并发的安全，基于这个思路，有个升级版的方案：用 select+channel 做检测。  
```go
func main() {
   var wg sync.WaitGroup
   wg.Add(1)
   // stopCh 参数，用于接收停止指令
   stopCh := make(chan bool) //用来停止监控狗
   go func() {
      defer wg.Done()
      watchDog(stopCh,"【监控狗1】")
   }()
   time.Sleep(5 * time.Second) //先让监控狗监控5秒
   // 通过 stopCh<-true 发送停止指令让协程退出
   stopCh <- true //发停止指令
   wg.Wait()
}
func watchDog(stopCh chan bool,name string){
   //开启for select循环，一直后台监控
   for{
      select {
      case <-stopCh:
         fmt.Println(name,"停止指令已收到，马上停止")
         return
      default:
         fmt.Println(name,"正在监控……")
      }
      time.Sleep(1*time.Second)
   }
}
```

通过 select+channel 让协程退出的方式比较优雅，但是如果我们希望做到同时取消很多个协程呢？如果是定时取消协程又该怎么办？这时候 select+channel 的局限性就凸现出来了，即使定义了多个 channel 解决问题，代码逻辑也会非常复杂、难以维护。  
要解决这种复杂的协程问题，必须有一种可以跟踪协程的方案，只有跟踪到每个协程，才能更好地控制它们，这种方案就是 Go 语言标准库为我们提供的 Context。  
```go
func main() {
   var wg sync.WaitGroup
   wg.Add(1)
   ctx,stop:=context.WithCancel(context.Background())
   go func() {
      defer wg.Done()
      watchDog(ctx,"【监控狗1】")
   }()
   time.Sleep(5 * time.Second) //先让监控狗监控5秒
   stop() //发停止指令
   wg.Wait()
}
func watchDog(ctx context.Context,name string) {
   //开启for select循环，一直后台监控
   for {
      select {
      case <-ctx.Done():
         fmt.Println(name,"停止指令已收到，马上停止")
         return
      default:
         fmt.Println(name,"正在监控……")
      }
      time.Sleep(1 * time.Second)
   }
}
```

相比 select+channel 的方案，Context 方案主要有 4 个改动点：
- watchDog 的 stopCh 参数换成了 ctx，类型为 context.Context。
- 原来的 case <-stopCh 改为 case <-ctx.Done()，用于判断是否停止。
- 使用 context.WithCancel(context.Background()) 函数生成一个可以取消的 Context，用于发送停止指令。这里的 context.Background() 用于生成一个空 Context，一般作为整个 Context 树的根节点。 
- 原来的 stopCh <- true 停止指令，改为 context.WithCancel 函数返回的取消函数 stop()。

**什么是 Context**   
一个任务会有很多个协程协作完成，一次 HTTP 请求也会触发很多个协程的启动，而这些协程有可能会启动更多的子协程，并且无法预知有多少层协程、每一层有多少个协程。  
如果因为某些原因导致任务终止了，HTTP 请求取消了，那么它们启动的协程怎么办？该如何取消呢？因为取消这些协程可以节约内存，提升性能，同时避免不可预料的 Bug。  
Context 就是用来简化解决这些问题的，并且是并发安全的。Context 是一个接口，它具备手动、定时、超时发出取消信号、传值等功能，主要用于控制多个协程之间的协作，尤其是取消操作。一旦取消指令下达，那么被 Context 跟踪的这些协程都会收到取消信号，就可以做清理和退出操作。  

Context 接口只有四个方法，
- Deadline 方法可以获取设置的截止时间，第一个返回值 deadline 是截止时间，到了这个时间点，Context 会自动发起取消请求，第二个返回值 ok 代表是否设置了截止时间。
- Done 方法返回一个只读的 channel，类型为 struct{}。在协程中，如果该方法返回的 chan 可以读取，则意味着 Context 已经发起了取消信号。通过 Done 方法收到这个信号后，就可以做清理操作，然后退出协程，释放资源。
- Err 方法返回取消的错误原因，即因为什么原因 Context 被取消。
- Value 方法获取该 Context 上绑定的值，是一个键值对，所以要通过一个 key 才可以获取对应的值。

Context 接口的四个方法中最常用的就是 Done 方法，它返回一个只读的 channel，用于接收取消信号。当 Context 取消的时候，会关闭这个只读 channel，也就等于发出了取消信号。
```go
type Context interface {
   Deadline() (deadline time.Time, ok bool)
   Done() <-chan struct{}
   Err() error
   Value(key interface{}) interface{}
}
```

**Context 树**  
不需要自己实现 Context 接口，Go 语言提供了函数可以帮助我们生成不同的 Context，通过这些函数可以生成一颗 Context 树，这样 Context 才可以关联起来，父 Context 发出取消信号的时候，子 Context 也会发出，这样就可以控制不同层级的协程退出。

从使用功能上分，有四种实现好的 Context。
- 空 Context：不可取消，没有截止时间，主要用于 Context 树的根节点。
- 可取消的 Context：用于发出取消信号，当取消的时候，它的子 Context 也会取消。
- 可定时取消的 Context：多了一个定时的功能。
- 值 Context：用于存储一个 key-value 键值对。

Context 的衍生树中，最顶部的是空 Context，它作为整棵 Context 树的根节点，在 Go 语言中，可以通过 context.Background() 获取一个根节点 Context。有了根节点 Context 后，这颗 Context 树要怎么生成呢？需要使用 Go 语言提供的四个函数。
- WithCancel(parent Context)：生成一个可取消的 Context。
- WithDeadline(parent Context, d time.Time)：生成一个可定时取消的 Context，参数 d 为定时取消的具体时间。
- WithTimeout(parent Context, timeout time.Duration)：生成一个可超时取消的 Context，参数 timeout 用于设置多久后取消
- WithValue(parent Context, key, val interface{})：生成一个可携带 key-value 键值对的 Context。

以上四个生成 Context 的函数中，前三个都属于可取消的 Context，它们是一类函数，最后一个是值 Context，用于存储一个 key-value 键值对。  

**使用 Context 取消多个协程**  
取消多个协程也比较简单，把 Context 作为参数传递给协程即可。
```go
wg.Add(3)
go func() {
   defer wg.Done()
   watchDog(ctx,"【监控狗2】")
}()
go func() {
   defer wg.Done()
   watchDog(ctx,"【监控狗3】")
}()
```
增加了两个监控狗，也就是增加了两个协程，这样一个 Context 就同时控制了三个协程，一旦 Context 发出取消信号，这三个协程都会取消退出。

**Context 传值**  
Context 不仅可以取消，还可以传值，通过这个能力，可以把 Context 存储的值供其他协程使用。
```go
func main() {
   wg.Add(4) //记得这里要改为4，原来是3，因为要多启动一个协程
   
  //省略其他无关代码
   valCtx:=context.WithValue(ctx,"userId",2)
   go func() {
      defer wg.Done()
      getUser(valCtx)
   }()
   //省略其他无关代码
}
func getUser(ctx context.Context){
   for  {
      select {
      case <-ctx.Done():
         fmt.Println("【获取用户】","协程退出")
         return
      default:
         userId:=ctx.Value("userId")
         fmt.Println("【获取用户】","用户ID为：",userId)
         time.Sleep(1 * time.Second)
      }
   }
}
```
通过 context.WithValue 函数存储一个 userId 为 2 的键值对，就可以在 getUser 函数中通过 ctx.Value("userId") 方法把对应的值取出来，达到传值的目的。  

**Context 使用原则**   
Context 是一种非常好的工具，使用它可以很方便地控制取消多个协程。在 Go 语言标准库中也使用了它们，比如 net/http 中使用 Context 取消网络的请求。要更好地使用 Context，有一些使用原则需要尽可能地遵守：
- Context 不要放在结构体中，要以参数的方式传递。
- Context 作为函数的参数时，要放在第一位，也就是第一个参数。
- 要使用 context.Background 函数生成根节点的 Context，也就是最顶层的 Context。
- Context 传值要传递必须的值，而且要尽可能地少，不要什么都传。
- Context 多协程安全，可以在多个协程中放心使用。

以上原则是规范类的，Go 语言的编译器并不会做这些检查，要靠自己遵守。


核心是 `Context` 接口：
```go
// A Context carries a deadline, cancelation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
```

包含四个方法：
*   `Done()`：返回一个 channel，当 times out 或者调用 cancel 方法时。
*   `Err()`：返回一个错误，表示取消 ctx 的原因。
*   `Deadline()`：返回截止时间和一个 bool 值。
*   `Value()`：返回 key 对应的值。

有四个结构体实现了这个接口，分别是：`emptyCtx`, `cancelCtx`, `timerCtx` 和 `valueCtx`。

其中 `emptyCtx` 是空类型，暴露了两个方法：
```go
func Background() Context
func TODO() Context
```
一般情况下，会使用 `Background()` 作为根 ctx，然后在其基础上再派生出子 ctx。要是不确定使用哪个 ctx，就使用 `TODO()`。

另外三个也分别暴露了对应的方法：
```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```

**WithCancel**  

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```
`WithCancel` 返回带有新 `Done` 通道的父级副本。当调用返回的 `cancel` 函数或关闭父上下文的 `Done` 通道时，返回的 `ctx` 的 `Done` 通道将关闭。  
取消此上下文会释放与其关联的资源，因此在此上下文中运行的操作完成后，代码应立即调用 `cancel`。

举个例子：  
这段代码演示了如何使用可取消上下文来防止 goroutine 泄漏。在函数结束时，由 `gen` 启动的 goroutine 将返回而不会泄漏。
```go
package main

import (
    "context"
    "fmt"
)

func main() {
    // gen generates integers in a separate goroutine and
    // sends them to the returned channel.
    // The callers of gen need to cancel the context once
    // they are done consuming generated integers not to leak
    // the internal goroutine started by gen.
    gen := func(ctx context.Context) <-chan int {
        dst := make(chan int)
        n := 1
        go func() {
            for {
                select {
                case <-ctx.Done():
                    return // returning not to leak the goroutine
                case dst <- n:
                    n++
                }
            }
        }()
        return dst
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // cancel when we are finished consuming integers

    for n := range gen(ctx) {
        fmt.Println(n)
        if n == 5 {
            break
        }
    }
}
```
输出：
```go
1
2
3
4
5
```

**WithDeadline**  

```go
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
```
`WithDeadline` 返回父上下文的副本，并将截止日期调整为不晚于 `d`。如果父级的截止日期已经早于 `d`，则 `WithDeadline(parent, d)` 在语义上等同于 `parent`。  
当截止时间到期、调用返回的取消函数时或当父上下文的 `Done` 通道关闭时，返回的上下文的 `Done` 通道将关闭。  
取消此上下文会释放与其关联的资源，因此在此上下文中运行的操作完成后，代码应立即调用取消。

举个例子：  
这段代码传递具有截止时间的上下文，来告诉阻塞函数，它应该在到达截止时间时立刻退出。
```go
package main

import (
    "context"
    "fmt"
    "time"
)

const shortDuration = 1 * time.Millisecond

func main() {
    d := time.Now().Add(shortDuration)
    ctx, cancel := context.WithDeadline(context.Background(), d)

    // Even though ctx will be expired, it is good practice to call its
    // cancellation function in any case. Failure to do so may keep the
    // context and its parent alive longer than necessary.
    defer cancel()

    select {
    case <-time.After(1 * time.Second):
        fmt.Println("overslept")
    case <-ctx.Done():
        fmt.Println(ctx.Err())
    }
}
```
输出：
```go
context deadline exceeded
```

**WithTimeout**  

```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```
`WithTimeout` 返回 `WithDeadline(parent, time.Now().Add(timeout))`。  
取消此上下文会释放与其关联的资源，因此在此上下文中运行的操作完成后，代码应立即调用取消。

举个例子：  
这段代码传递带有超时的上下文，以告诉阻塞函数应在超时后退出。
```go
package main

import (
    "context"
    "fmt"
    "time"
)

const shortDuration = 1 * time.Millisecond

func main() {
    // Pass a context with a timeout to tell a blocking function that it
    // should abandon its work after the timeout elapses.
    ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
    defer cancel()

    select {
    case <-time.After(1 * time.Second):
        fmt.Println("overslept")
    case <-ctx.Done():
        fmt.Println(ctx.Err()) // prints "context deadline exceeded"
    }

}
```
输出：
```go
context deadline exceeded
```

**WithValue**  

```go
func WithValue(parent Context, key, val any) Context
```
`WithValue` 返回父级的副本，其中与 `key` 关联的值为 `val`。  
其中键必须是可比较的，并且不应是字符串类型或任何其他内置类型，以避免使用上下文的包之间发生冲突。 `WithValue` 的用户应该定义自己的键类型。  
为了避免分配给 `interface{}`，上下文键通常具有具体的 `struct{}` 类型。或者，导出的上下文键变量的静态类型应该是指针或接口。  

举个例子：  
这段代码演示了如何将值传递到上下文以及如何检索它（如果存在）。
```go
package main

import (
    "context"
    "fmt"
)

func main() {
    type favContextKey string

    f := func(ctx context.Context, k favContextKey) {
        if v := ctx.Value(k); v != nil {
            fmt.Println("found value:", v)
            return
        }
        fmt.Println("key not found:", k)
    }

    k := favContextKey("language")
    ctx := context.WithValue(context.Background(), k, "Go")

    f(ctx, k)
    f(ctx, favContextKey("color"))
}
```
输出：
```go
found value: Go
key not found: color
```
