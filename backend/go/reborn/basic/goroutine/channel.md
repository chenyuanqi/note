
### channel 通道
并发程序的几条优点：
- 并发能更客观地表现问题模型；
- 并发可以充分利用 CPU 核心的优势，提高程序的执行效率；
- 并发能充分利用 CPU 与其他硬件设备固有的异步性。

channel 是 Go 语言在语言级别提供的 goroutine 间的通信方式。我们可以使用 channel 在两个或多个 goroutine 之间传递消息。Go 语言中的通道（channel）是一种特殊的类型，在任何时候，同时只能有一个 goroutine 访问通道进行发送和获取数据。  

channel 是进程内的通信方式，因此通过 channel 传递对象的过程和调用函数时的参数传递行为比较一致，比如也可以传递指针等。如果需要跨进程通信，我们建议用分布式系统的方法来解决，比如使用 Socket 或者 HTTP 等通信协议。Go 语言对于网络方面也有非常完善的支持。

channel 是类型相关的，也就是说，一个 channel 只能传递一种类型的值，这个类型需要在声明 channel 时指定。如果对 Unix 管道有所了解的话，就不难理解 channel，可以将其认为是一种类型安全的管道。

定义一个 channel 时，也需要定义发送到 channel 的值的类型（指定将要被共享的数据的类型），注意，必须使用 make 创建 channel。可以通过通道共享内置类型、命名类型、结构类型和引用类型的值或者指针。channel 像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。

```go
// var 通道变量 chan 通道类型
var c chan int
// chan 类型的空值是 nil，声明后需要配合 make 后才能使用
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
```

**使用通道发送和接收数据**  
通道创建后，就可以使用通道进行发送和接收操作。  
1) 通道发送数据  
通道的发送使用特殊的操作符 <-，将数据通过通道发送的格式为：通道变量 <- 值  
- 通道变量：通过 make 创建好的通道实例。
- 值：可以是变量、常量、表达式或者函数返回值等。值的类型必须与 ch 通道的元素类型一致。  

使用 make 创建一个通道后，就可以使用 <- 向通道发送数据。  
```go
// 创建一个空接口通道
ch := make(chan interface{})
// 将0放入通道中
ch <- 0
// 将hello字符串放入通道中
ch <- "hello"
```

发送将持续阻塞直到数据被接收  
把数据往通道中发送时，如果接收方一直都没有接收，那么发送操作将持续阻塞。Go 程序运行时能智能地发现一些永远无法发送成功的语句并做出提示。  
```go
package main
func main() {
    // 创建一个整型通道
    ch := make(chan int)
    // 尝试将0通过通道发送
    ch <- 0
    // fatal error: all goroutines are asleep - deadlock!
}
```

2) 通道接收数据  
通道接收同样使用 <- 操作符，通道接收有如下特性：  
① 通道的收发操作在不同的两个 goroutine 间进行。由于通道的数据在没有接收方处理时，数据发送方会持续阻塞，因此通道的接收必定在另外一个 goroutine 中进行。  
② 接收将持续阻塞直到发送方发送数据。如果接收方接收时，通道中没有发送方发送数据，接收方也会发生阻塞，直到发送方发送数据为止。
③ 每次接收一个元素。通道一次只能接收一个数据元素。  

通道的数据接收一共有以下 4 种写法。
```go
// 1、阻塞接收数据：将接收变量作为 <- 操作符的左值。执行该语句时将会阻塞，直到接收到数据并赋值给 data 变量
data := <-ch

// 2、非阻塞接收数据：语句不会发生阻塞。非阻塞的通道接收方法可能造成高的 CPU 占用，因此使用非常少
// 如果需要实现接收超时检测，可以配合 select 和计时器 channel 进行
data, ok := <-ch // data：表示接收到的数据，未接收到数据时 data 为通道类型的零值；ok：表示是否接收到数据

// 3、接收任意数据，忽略接收的数据：执行该语句时将会发生阻塞，直到接收到数据，但接收到的数据会被忽略。这个方式实际上只是通过通道在 goroutine 间阻塞收发实现并发同步
<-ch
// 案例：使用通道做并发同步
// 构建一个通道
ch := make(chan int)
// 开启一个并发匿名函数
go func() {
    fmt.Println("start goroutine")
    // 通过通道通知main的goroutine
    ch <- 0
    fmt.Println("exit goroutine")
}()
fmt.Println("wait goroutine")
// 等待匿名goroutine
<-ch
fmt.Println("all done")
// wait goroutine
// start goroutine
// exit goroutine
// all done

// 4、for range 循环接收：通道 ch 是可以进行遍历的，遍历的结果就是接收到的数据。数据类型就是通道的数据类型
// 案例：遍历通道数据
// 构建一个通道
ch := make(chan int)
// 开启一个并发匿名函数
go func() {
    // 从3循环到0
    for i := 3; i >= 0; i-- {
        // 发送3到0之间的数值
        ch <- i
        // 每次发送完时等待
        time.Sleep(time.Second)
    }
}()
// 遍历接收通道数据
for data := range ch {
    // 打印通道数据
    fmt.Println(data)
    // 当遇到数据0时, 退出接收循环
    if data == 0 {
        break
    }
}
// 3
// 2
// 1
// 0
```


**无缓冲 channel**  
无缓冲 channel（即同步通道），它的容量是 0，不能存储任何数据。所以无缓冲 channel 只起到传输数据的作用，数据并不会在 channel 中做任何停留。这也意味着，无缓冲 channel 的发送和接收操作是同时进行的，它也可以称为同步 channel。 
同步 channel 有点类似于送外卖的过程。若外卖小哥和点餐顾客分别为协程 A 和协程 B，只有当协程 A 把数据（即外卖）送给协程 B（即顾客），协程 B 才能开始执行后续的操作（即吃外卖）。否则，协程 B 只能一直等待数据（即外卖）的到来。
`注意：使用同步通道时，要确保传出数据和获取数据必须成对出现。`

譬如我们正在饲养一只母鸡，等待其下蛋。每下一个蛋，我们就拿来做荷包蛋吃。
```go
var syncWait sync.WaitGroup
// 创建通道类型变量，该通道将传送int类型数据
var intChan = make(chan int)

func main() {
   // 执行2个协程任务
   syncWait.Add(2)
   // 开启下蛋任务
   go layEggs()
   // 开启吃荷包蛋任务
   go eatEggs(intChan)
   // 等待协程任务完成
   syncWait.Wait()
}

func layEggs() {
   // 使用断言确保协程任务正常结束
   defer syncWait.Done()
   fmt.Println("老母鸡生了一个鸡蛋")
   // 向通道传送int类型值
   intChan <- 1
   // 关闭通道
   close(intChan)
}

func eatEggs(intChan chan int) {
   // 使用断言确保协程任务正常结束
   defer syncWait.Done()
   // 从通道获取int类型值
   eggCounts := <-intChan
   // 输出结果
   fmt.Printf("吃%d个荷包蛋", eggCounts)
}
// 老母鸡生了一个鸡蛋
// 吃 1 个荷包蛋
```

**有缓冲 channel**  
有缓冲 channel 类似一个可阻塞的队列，内部的元素先进先出。通过 make 函数的第二个参数可以指定 channel 容量的大小，进而创建一个有缓冲 channel。  
一个有缓冲 channel 具备以下特点：
- 有缓冲 channel 的内部有一个缓冲队列；
- 发送操作是向队列的尾部插入元素，如果队列已满，则阻塞等待，直到另一个 goroutine 执行，接收操作释放队列的空间；
- 接收操作是从队列的头部获取元素并把它从队列中删除，如果队列为空，则阻塞等待，直到另一个 goroutine 执行，发送操作插入新的元素。

缓冲通道则有点类似于送快递的过程。若快递员和收件人分别为协程 A 和协程 B，协程 A 可以把数据（即快递）放到缓冲区（即菜鸟驿站）。当协程 B 需要时，只要去缓冲区（即菜鸟驿站）中取数据（即快递）即可。值得一提的是，缓冲区和菜鸟驿站真的很像，它们都有最大容量限制。一旦协程 A 发现缓冲区（即菜鸟驿站）满了，就不得不等待数据（即快递）被取走，才能将数据（即快递）放到空余的位置中。

**关闭 channel**  
channel 还可以使用内置函数 close 关闭，代码：close(cacheCh)。如果一个 channel 被关闭了，就不能向里面发送数据了，如果发送的话，会引起 painc 异常。但是还可以接收 channel 里的数据，如果 channel 里没有数据的话，接收的数据是元素类型的零值。  
通过内置函数 cap 可以获取 channel 的容量，也就是最大能存放多少个元素，通过内置函数 len 可以获取 channel 中元素的个数。  
`小提示：无缓冲 channel 其实就是一个容量大小为 0 的 channel。比如 make(chan int,0)。`
```go
intChan <- 1
close(intChan)
```

**单向channel**  
我们有一些特殊的业务需求，比如限制一个 channel 只可以接收但是不能发送，或者限制一个 channel 只能发送但不能接收，这种 channel 称为单向 channel。  
单向 channel 的声明也很简单，只需要在声明的时候带上 <- 操作符即可，如下面的代码所示：
```go
onlySend := make(chan<- int)
onlyReceive:=make(<-chan int)
```
`注意，声明单向 channel <- 操作符的位置和上面讲到的发送和接收操作是一样的。`  
在函数或者方法的参数中，使用单向 channel 的较多，这样可以防止一些操作影响了 channel。
