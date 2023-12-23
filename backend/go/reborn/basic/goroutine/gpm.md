
### GPM 任务调度
Go 语言中的 GPM 任务调度模型充分利用了多核 CPU 的资源。需要时，将创建与之匹配的线程，并将用户态的协程任务 “智能” 地分配给多个线程执行。  
- M：内核级线程，即工作线程，在 Go 中称为 Machine，数量对应真实的 CPU 数（真正干活的对象）。  
- G：代表一个 goroutine，即 Go 协程，每个 go 关键字都会创建一个协程。  
- P：Processor，处理器，用来管理和执行 goroutine 的。包含运行 Go 代码的必要资源，用来调度 G 和 M 之间的关联关系，其数量可通过 GOMAXPROCS () 来设置，默认为核心数。

在 Go 程序启动时，会自动根据 CPU 的核心数设置线程的最大数量。当然，我们也可以通过编码手动设置。当一个线程发生阻塞时，新的线程便会创建。  
全局队列存放所有等待运行的协程任务（即 Goroutine）。当发起一个协程任务时，该任务会首先尝试加入到协程队列中。每个协程队列的最大任务数被限制在 256 个以内。 当协程队列满了之后，协程调度器会将一半数量的任务移动至全局队列中。至于一共能有多少个协程队列，在 Go 1.5 版本之后队列数默认为 CPU 核心数量，也可以通过编码来指定。从另一个角度讲，设置了队列数就意味着设置了程序能同时跑多少个 Goroutine 的数量。一般地，在该参数确定后，所有的队列便会一口气创建完成。  
在 Go 程序运行时，一个内核空间的线程若想获取某个协程任务来执行，就需要通过协程队列处理来获取特定的协程任务。当队列为空时，全局队列中的若干协程任务，或其它队列中的一半任务会被放到空队列中。如此循环往复，周而复始。  
另一方面，协程队列处理器的数量和线程在数量上并没有绝对关系。如果一个线程发生阻塞，协程队列处理器便会创建或切换至其它线程。因此，即使只有一个协程队列，也有可能会有多个线程。  

**Goroutine**  
Goroutine 就是代码中使用 go 关键词创建的执行单元，也是大家熟知的有“轻量级线程”之称的协程，协程是不为操作系统所知的，它由编程语言层面实现，上下文切换不需要经过内核态，再加上协程占用的内存空间极小，所以有着非常大的发展潜力。
```go
go func() {}()
```

在 Go 语言中，Goroutine 由一个名为 runtime.go 的结构体表示，该结构体非常复杂，有 40 多个成员变量，主要存储执行栈、状态、当前占用的线程、调度相关的数据。还有大家很想获取的 goroutine 标识，但是官方考虑到 Go 语言的发展设置成私有了。
```go
type g struct {
   stack struct {
      lo uintptr
      hi uintptr
   }                    // 栈内存：[stack.lo, stack.hi)
   stackguard0 uintptr
   stackguard1 uintptr

   _panic       *_panic
   _defer       *_defer
   m            *m            // 当前的 m
   sched        gobuf
   stktopsp     uintptr    // 期望 sp 位于栈顶，用于回溯检查
   param        unsafe.Pointer // wakeup 唤醒时候传递的参数
   atomicstatus uint32
   goid         int64
   preempt      bool          // 抢占信号，stackguard0 = stackpreempt 的副本
   timer        *timer         // 为 time.Sleep 缓存的计时器

   ...
}
```
Goroutine 调度相关的数据存储在 sched，在协程切换、恢复上下文的时候用到。
```go
type gobuf struct {
   sp   uintptr
   pc   uintptr
   g    guintptr
   ret  sys.Uintreg
   ...
}
```
M 就是对应操作系统的线程，最多会有 GOMAXPROCS 个活跃线程能够正常运行，默认情况下 GOMAXPROCS 被设置为内核数，假如有四个内核，那么默认就创建四个线程，每一个线程对应一个 runtime.m 结构体。线程数等于 CPU 个数的原因是，每个线程分配到一个 CPU 上就不至于出现线程的上下文切换，可以保证系统开销降到最低。
```go
type m struct {
   g0   *g 
   curg *g
   ...
}
```
M 里面存了两个比较重要的东西，一个是 g0，一个是 curg。
- g0：会深度参与运行时的调度过程，比如goroutine的创建、内存分配等
- curg：代表当前正在线程上执行的goroutine。

P 是负责 M 与 G 的关联，所以 M 里面还要存储与 P 相关的数据。
```go
type m struct {
  ...
   p             puintptr
   nextp         puintptr
   oldp          puintptr
}
```
- p：正在运行代码的处理器
- nextp：暂存的处理器
- oldp：系统调用之前的线程的处理器

**Processor**  
Proccessor 负责 Machine 与 Goroutine 的连接，它能提供线程需要的上下文环境，也能分配 G 到它应该去的线程上执行，有了它，每个 G 都能得到合理的调用。  
同样的，处理器的数量也是默认按照 GOMAXPROCS 来设置的，与线程的数量一一对应。
```go
type p struct {
   m           muintptr

   runqhead uint32
   runqtail uint32
   runq     [256]guintptr
   runnext guintptr
   ...
}
```
结构体 P 中存储了性能追踪、垃圾回收、计时器等相关的字段外，还存储了处理器的待运行队列，队列中存储的是待执行的 Goroutine 列表。

**GPM 三者关系**  
Goroutines（G）在 Processors（P）上运行，而 Processor（P）则被绑定到 Machine（M）上。每个 Processor 可以被视为一个本地调度器，负责调度附着在其上的 Goroutines。M 和 P 之间的关系是多对多的关系：一个 M 可以执行多个 P，而一个 P 可以在不同的 M 之间迁移。  
当一个 Goroutine 需要执行时，它会被放置在一个 P 的本地运行队列中。如果 P 已经绑定到一个 M，则该 Goroutine 可以被该 M 执行。如果所有的 M 都在忙，但还有空闲的 P，运行时系统可能会创建一个新的 M 来运行这个 P 上的 Goroutines。  
这种调度模型使得 Go 能够在保持高效的同时支持大量的并发 Goroutines。通过这种方式，Go 的运行时系统能够更好地在多核处理器上分配工作负载，从而实现高效的并发执行。

**Goroutine 调度策略**  
队列轮转：P 会周期性的将 G 调度到 M 中执行，执行一段时间后，保存上下文，将 G 放到队列尾部，然后从队列中再取出一个 G 进行调度。除此之外，P 还会周期性的查看全局队列是否有 G 等待调度到 M 中执行。  
系统调用：当 G0 即将进入系统调用时，M0 将释放 P，进而某个空闲的 M1 获取 P，继续执行 P 队列中剩下的 G。M1 的来源有可能是 M 的缓存池，也可能是新建的。  
当 G0 系统调用结束后，如果有空闲的 P，则获取一个 P，继续执行 G0。如果没有，则将 G0 放入全局队列，等待被其他的 P 调度。然后 M0 将进入缓存池睡眠。  


**动态调整系统资源**  
在 Go 程序运行时，可以根据需要设置程序要使用的 CPU 资源，也可以动态调整协程任务的执行方式，实现更灵活地运行。这些操作都是通过 runtime 包来实现的。

```go
// 获取运行当前程序的操作系统
// 在 macOS 中，操作系统名称为 darwin；在 Windows 中，操作系统名称即 windows；在 Linux 中，操作系统名称为 linux
fmt.Println(runtime.GOOS)
// 获取运行当前程序的CPU架构
fmt.Println(runtime.GOARCH)
// 获取运行当前程序的CPU核心数量
// 对于 32 位的 CPU，运行结果为 386；对于 64 位的 CPU，运行结果为 amd64；对于 arm 架构 32 位的 CPU，运行结果为 arm；对于 arm 架构 64 位的 CPU，运行结果为 arm64
fmt.Println(runtime.NumCPU())

// 设置程序只能使用一半数量的核心
if runtime.NumCPU() > 2 {
   runtime.GOMAXPROCS(runtime.NumCPU() / 2)
}
// 获取当前程序可用的CPU核心数
fmt.Println(runtime.GOMAXPROCS(0))
```

