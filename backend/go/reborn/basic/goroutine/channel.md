
### channel 通道
并发程序的几条优点：
- 并发能更客观地表现问题模型；
- 并发可以充分利用 CPU 核心的优势，提高程序的执行效率；
- 并发能充分利用 CPU 与其他硬件设备固有的异步性。

channel 是 Go 语言在语言级别提供的 goroutine 间的通信方式。我们可以使用 channel 在两个或多个 goroutine 之间传递消息。

channel 是进程内的通信方式，因此通过 channel 传递对象的过程和调用函数时的参数传递行为比较一致，比如也可以传递指针等。如果需要跨进程通信，我们建议用分布式系统的方法来解决，比如使用 Socket 或者 HTTP 等通信协议。Go 语言对于网络方面也有非常完善的支持。

channel 是类型相关的，也就是说，一个 channel 只能传递一种类型的值，这个类型需要在声明 channel 时指定。如果对 Unix 管道有所了解的话，就不难理解 channel，可以将其认为是一种类型安全的管道。

定义一个 channel 时，也需要定义发送到 channel 的值的类型，注意，必须使用 make 创建 channel。
```go
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
```

**无缓冲 channel**  
无缓冲 channel，它的容量是 0，不能存储任何数据。所以无缓冲 channel 只起到传输数据的作用，数据并不会在 channel 中做任何停留。这也意味着，无缓冲 channel 的发送和接收操作是同时进行的，它也可以称为同步 channel。 

**有缓冲 channel**  
有缓冲 channel 类似一个可阻塞的队列，内部的元素先进先出。通过 make 函数的第二个参数可以指定 channel 容量的大小，进而创建一个有缓冲 channel。  
一个有缓冲 channel 具备以下特点：
- 有缓冲 channel 的内部有一个缓冲队列；
- 发送操作是向队列的尾部插入元素，如果队列已满，则阻塞等待，直到另一个 goroutine 执行，接收操作释放队列的空间；
- 接收操作是从队列的头部获取元素并把它从队列中删除，如果队列为空，则阻塞等待，直到另一个 goroutine 执行，发送操作插入新的元素。

**关闭 channel**  
channel 还可以使用内置函数 close 关闭，代码：close(cacheCh)。如果一个 channel 被关闭了，就不能向里面发送数据了，如果发送的话，会引起 painc 异常。但是还可以接收 channel 里的数据，如果 channel 里没有数据的话，接收的数据是元素类型的零值。  
通过内置函数 cap 可以获取 channel 的容量，也就是最大能存放多少个元素，通过内置函数 len 可以获取 channel 中元素的个数。  
`小提示：无缓冲 channel 其实就是一个容量大小为 0 的 channel。比如 make(chan int,0)。`

**单向channel**  
我们有一些特殊的业务需求，比如限制一个 channel 只可以接收但是不能发送，或者限制一个 channel 只能发送但不能接收，这种 channel 称为单向 channel。  
单向 channel 的声明也很简单，只需要在声明的时候带上 <- 操作符即可，如下面的代码所示：
```go
onlySend := make(chan<- int)
onlyReceive:=make(<-chan int)
```
`注意，声明单向 channel <- 操作符的位置和上面讲到的发送和接收操作是一样的。`  
在函数或者方法的参数中，使用单向 channel 的较多，这样可以防止一些操作影响了 channel。
