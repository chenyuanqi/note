
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

