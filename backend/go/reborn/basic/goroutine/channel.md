
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
`注意：使用同步通道时，要确保传出数据和获取数据必须成对出现。这种类型的通道要求发送 goroutine 和接收 goroutine 同时准备好，才能完成发送和接收操作。`

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

使用无缓冲的通道，在 goroutine 之间同步数据，来模拟接力比赛。在接力比赛里，4 个跑步者围绕赛道轮流跑。第二个、第三个和第四个跑步者要接到前一位跑步者的接力棒后才能起跑。比赛中最重要的部分是要传递接力棒，要求同步传递。在同步接力棒的时候，参与接力的两个跑步者必须在同一时刻准备好交接。
```go
// 这个示例程序展示如何用无缓冲的通道来模拟
// 4 个goroutine 间的接力比赛
package main

import (
    "fmt"
    "sync"
    "time"
)

// wg 用来等待程序结束
var wg sync.WaitGroup

// main 是所有Go 程序的入口
func main() {
    // 创建一个无缓冲的通道
    baton := make(chan int)

    // 为最后一位跑步者将计数加1
    wg.Add(1)

    // 第一位跑步者持有接力棒
    go Runner(baton)

    // 开始比赛
    baton <- 1

    // 等待比赛结束
    wg.Wait()
}

// Runner 模拟接力比赛中的一位跑步者
func Runner(baton chan int) {
    var newRunner int

    // 等待接力棒
    runner := <-baton

    // 开始绕着跑道跑步
    fmt.Printf("Runner %d Running With Baton\n", runner)

    // 创建下一位跑步者
    if runner != 4 {
        newRunner = runner + 1
        fmt.Printf("Runner %d To The Line\n", newRunner)
        go Runner(baton)
    }

    // 围绕跑道跑
    time.Sleep(100 * time.Millisecond)

    // 比赛结束了吗？
    if runner == 4 {
        fmt.Printf("Runner %d Finished, Race Over\n", runner)
        wg.Done()
        return
    }

    // 将接力棒交给下一位跑步者
    fmt.Printf("Runner %d Exchange With Runner %d\n",
        runner,
        newRunner)

    baton <- newRunner
}
// Runner 1 Running With Baton
// Runner 1 To The Line
// Runner 1 Exchange With Runner 2
// Runner 2 Running With Baton
// Runner 2 To The Line
// Runner 2 Exchange With Runner 3
// Runner 3 Running With Baton
// Runner 3 To The Line
// Runner 3 Exchange With Runner 4
// Runner 4 Running With Baton
// Runner 4 Finished, Race Over
```

**有缓冲 channel**  
有缓冲 channel 类似一个可阻塞的队列，内部的元素先进先出。通过 make 函数的第二个参数可以指定 channel 容量的大小，进而创建一个有缓冲 channel。这种类型的通道并不强制要求 goroutine 之间必须同时完成发送和接收。通道会阻塞发送和接收动作的条件也会不同。只有在通道中没有要接收的值时，接收动作才会阻塞。只有在通道没有可用缓冲区容纳被发送的值时，发送动作才会阻塞。  
所以，有缓冲的通道和无缓冲的通道之间的一个很大的不同：无缓冲的通道保证进行发送和接收的 goroutine 会在同一时间进行数据交换；有缓冲的通道没有这种保证。一个有缓冲 channel 具备以下特点：
- 有缓冲 channel 的内部有一个缓冲队列；
- 发送操作是向队列的尾部插入元素，如果队列已满，则阻塞等待，直到另一个 goroutine 执行，接收操作释放队列的空间；
- 接收操作是从队列的头部获取元素并把它从队列中删除，如果队列为空，则阻塞等待，直到另一个 goroutine 执行，发送操作插入新的元素。

在无缓冲通道的基础上，为通道增加一个有限大小的存储空间形成带缓冲通道。带缓冲通道在发送时无需等待接收方接收即可完成发送过程，并且不会发生阻塞，只有当存储空间满时才会发生阻塞。同理，如果缓冲通道中有数据，接收时将不会发生阻塞，直到通道中没有数据可读时，通道将会再度阻塞。  
缓冲通道则有点类似于送快递的过程。若快递员和收件人分别为协程 A 和协程 B，协程 A 可以把数据（即快递）放到缓冲区（即菜鸟驿站）。当协程 B 需要时，只要去缓冲区（即菜鸟驿站）中取数据（即快递）即可。值得一提的是，缓冲区和菜鸟驿站真的很像，它们都有最大容量限制。一旦协程 A 发现缓冲区（即菜鸟驿站）满了，就不得不等待数据（即快递）被取走，才能将数据（即快递）放到空余的位置中。
```go
// 创建带缓冲的通道
// 通道实例 := make (chan 通道类型，缓冲大小)

// 创建一个3个元素缓冲大小的整型通道
ch := make(chan int, 3)
// 查看当前通道的大小
fmt.Println(len(ch)) // 0
// 发送3个整型元素到通道
ch <- 1
ch <- 2
ch <- 3
// 查看当前通道的大小
fmt.Println(len(ch)) // 3
```

带缓冲通道在很多特性上和无缓冲通道是类似的。无缓冲通道可以看作是长度永远为 0 的带缓冲通道。因此根据这个特性，带缓冲通道在下面列举的情况下依然会发生阻塞：  
- 带缓冲通道被填满时，尝试再次发送数据时发生阻塞。
- 带缓冲通道为空时，尝试接收数据时发生阻塞。

为什么 Go 语言对通道要限制长度而不提供无限长度的通道？  
通道（channel）是在两个 goroutine 间通信的桥梁。使用 goroutine 的代码必然有一方提供数据，一方消费数据。当提供数据一方的数据供给速度大于消费方的数据处理速度时，如果通道不限制长度，那么内存将不断膨胀直到应用崩溃。因此，限制通道的长度有利于约束数据提供方的供给速度，供给数据量必须在消费方处理量 + 通道长度的范围内，才能正常地处理数据。  

**关闭 channel**  
channel 还可以使用内置函数 close 关闭，代码：close(cacheCh)。如果一个 channel 被关闭了，就不能向里面发送数据了，如果发送的话，会引起 painc 异常。但是还可以接收 channel 里的数据，如果 channel 里没有数据的话，接收的数据是元素类型的零值。  
通过内置函数 cap 可以获取 channel 的容量，也就是最大能存放多少个元素，通过内置函数 len 可以获取 channel 中元素的个数。  
`小提示：无缓冲 channel 其实就是一个容量大小为 0 的 channel。比如 make(chan int,0)。`
```go
intChan <- 1
close(intChan)

// 判断一个 channel 是否已经被关闭：只需要看第二个 bool 返回值即可，如果返回值是 false 则表示 ch 已经被关闭
x, ok := <-ch
```

**单向channel**  
我们有一些特殊的业务需求，比如限制一个 channel 只可以接收但是不能发送，或者限制一个 channel 只能发送但不能接收，这种 channel 称为单向 channel。  
单向 channel 的声明也很简单，只需要在声明的时候带上 <- 操作符即可，如下面的代码所示：
```go
// var 通道实例 chan<- 元素类型    // 只能写入（发送）数据的通道
// var 通道实例 <-chan 元素类型    // 只能读取（接收）数据的通道
onlySend := make(chan<- int)
onlyReceive:=make(<-chan int)

// 或
ch := make(chan int)
// 声明一个只能写入数据的通道类型, 并赋值为ch
var chSendOnly chan<- int = ch
//声明一个只能读取数据的通道类型, 并赋值为ch
var chRecvOnly <-chan int = ch
```
`注意，声明单向 channel <- 操作符的位置和上面讲到的发送和接收操作是一样的。一个不能写入数据只能读取的通道是毫无意义的。`  
在函数或者方法的参数中，使用单向 channel 的较多，这样可以防止一些操作影响了 channel。所以，单向通道有利于代码接口的严谨性。  

time 包中的单向通道，time 包中的计时器会返回一个 timer 实例。
```go
// timer 的 Timer 类型定义
type Timer struct {
    // C 通道的类型就是一种只能读取的单向通道。如果此处不进行通道方向约束，一旦外部向通道写入数据，将会造成其他使用到计时器的地方逻辑产生混乱
    C <-chan Time
    r runtimeTimer
}

timer := time.NewTimer(time.Second)
```

**channel 超时机制**  
Go 语言没有提供直接的超时处理机制，所谓超时可以理解为当我们上网浏览一些网站时，如果一段时间之后不作操作，就需要重新登录。那么我们应该如何实现这一功能呢，这时就可以使用 select 来设置超时。  
虽然 select 机制不是专门为超时而设计的，却能很方便的解决超时问题，因为 select 的特点是只要其中有一个 case 已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况。超时机制本身虽然也会带来一些问题，比如在运行比较快的机器或者高速的网络上运行正常的程序，到了慢速的机器或者网络上运行就会出问题，从而出现结果不一致的现象，但从根本上来说，解决死锁问题的价值要远大于所带来的问题。  

select 的用法与 switch 语言非常类似，由 select 开始一个新的选择块，每个选择条件由 case 语句来描述。与 switch 语句相比，select 有比较多的限制，其中最大的一条限制就是每个 case 语句里必须是一个 IO 操作。  
在一个 select 语句中，Go 语言会按顺序从头至尾评估每一个发送和接收的语句。  
如果其中的任意一语句可以继续执行（即没有被阻塞），那么就从那些可以执行的语句中任意选择一条来使用。如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有如下两种可能的情况：
- 如果给出了 default 语句，那么就会执行 default 语句，同时程序的执行会从 select 语句后的语句中恢复；
- 如果没有 default 语句，那么 select 语句将被阻塞，直到至少有一个通信可以进行下去。

```go
select {
    case <-chan1:
    // 如果 chan1 成功读到数据，则进行该 case 处理语句
    case chan2 <- 1:
    // 如果成功向 chan2 写入数据，则进行该 case 处理语句
    default:
    // 如果上面都没有成功，则进入 default 处理流程
}


ch := make(chan int)
quit := make(chan bool)
//新开一个协程
go func() {
    for {
        select {
        case num := <-ch:
            fmt.Println("num = ", num)
        case <-time.After(3 * time.Second):
            fmt.Println("超时")
            quit <- true
        }
    }
}() //别忘了()
for i := 0; i < 5; i++ {
    ch <- i
    time.Sleep(time.Second)
}
<-quit
fmt.Println("程序结束")
// num =  0
// num =  1
// num =  2
// num =  3
// num =  4
// 超时
// 程序结束
```



