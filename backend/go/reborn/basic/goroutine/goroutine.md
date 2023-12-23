
### goroutine
进程是程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位；线程是进程的一个执行实体，是 CPU 调度和分派的基本单位，它是比进程更小的能独立运行的基本单位。一个进程可以创建和撤销多个线程，同一个进程中的多个线程之间可以并发执行。

并发的关键是你有处理多个任务的能力，不一定要同时；并行的关键是你有同时处理多个任务的能力。  
对应到 CPU 上，如果是多核它就有同时执行的能力，即有并行的能力。对于 Go 语言，它自行安排了我们的代码合适并发合适并行。  

协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似于用户级线程，这些用户级线程的调度也是自己实现的。  
线程：一个线程上可以跑多个协程，协程是轻量级的线程。  
优雅的并发编程范式，完善的并发支持，出色的并发性能是Go语言区别于其他语言的一大特色。使用Go语言开发服务器程序时，就需要对它的并发机制有深入的了解。  

Goroutine 是 Go 语言中的协程，其它语言称为的协程字面上叫 Coroutine，简单理解下就是比线程更轻量的一个玩意。说白了，就是可以异步执行函数。
- routine 英 [ru:ˈti:n] 美 [ruˈtin] n. [计算机] 程序
- goroutine 和 coroutine 的概念和运行机制都是脱胎于早期的操作系统；goroutine 间使用 channel 通信，coroutine 使用 yield 和 resume 操作
- goroutines 意味着并行（或者可以以并行的方式部署），coroutines 一般来说不是这样的，goroutines 通过通道来通信；coroutines 通过让出和恢复操作来通信，goroutines 比 coroutines 更强大，也很容易从 coroutines 的逻辑复用到 goroutines  

goroutine是Go语言实现的用户态线程，主要用来解决操作系统线程太“重”的问题，所谓的太重，主要表现在以下两个方面：
- 创建和切换太重：操作系统线程的创建和切换都需要进入内核，而进入内核所消耗的性能代价比较高，开销较大；
- 内存使用太重：一方面，为了尽量避免极端情况下操作系统线程栈的溢出，内核在创建操作系统线程时默认会为其分配一个较大的栈内存（虚拟地址空间，内核并不会一开始就分配这么多的物理内存），然而在绝大多数情况下，系统线程远远用不了这么多内存，这导致了浪费；另一方面，栈内存空间一旦创建和初始化完成之后其大小就不能再有变化，这决定了在某些特殊场景下系统线程栈还是有溢出的风险。

而相对的，用户态的goroutine则轻量得多：
- goroutine是用户态线程，其创建和切换都在用户代码中完成而无需进入操作系统内核，所以其开销要远远小于系统线程的创建和切换；
- goroutine启动时默认栈大小只有2k，这在多数情况下已经够用了，即使不够用，goroutine的栈也会自动扩大，同时，如果栈太大了过于浪费它还能自动收缩，这样既没有栈溢出的风险，也不会造成栈内存空间的大量浪费。

正是因为Go语言中实现了如此轻量级的线程，才使得我们在Go程序中，可以轻易的创建成千上万甚至上百万的goroutine出来并发的执行任务而不用太担心性能和内存等问题。

当启动 main 入口函数时，后台就自动跑了一个 main Goroutine，  
`所以，一个 goroutine 内也可以使用 goroutine。`
```go
package main
func  main() {
    panic("看这里")
}
// panic: 看这里
// goroutine 1 [running]: // 出现了一个 goroutine 字眼，它对应的索引为 1
// main.main()
```

**创建 Goroutine**  
创建 Goroutine 很简单，只需要在函数前增加一个 go 关键字，也支持匿名函数。
```go
go fun1(...)
// 匿名
go func(...){
    // ...
}(...)

//go 关键字放在方法调用前新建一个 goroutine 并执行方法体
go GetThingDone(param1, param2);
//新建一个匿名方法并执行
go func(param1, param2) {
}(val1, val2)
//直接新建一个 goroutine 并在 goroutine 中执行代码块
go {
    //do someting...
}
```
go 关键字后的函数可以写返回值，但无效。因为 Goroutine 是异步的，所以没法接收。
```go
package main

import (
	"fmt"
)

func PrintA() {
	fmt.Println("A")
}

func main() {
    // 创建一个 Goroutine，异步打印 “A” 字符串
	go PrintA()
    // 打印 “main” 字符串
	fmt.Println("main") 
}
// main
// go PrintA() 创建的 Goroutine 它是异步执行，main 函数执行完退出程序时，也不会管它
// 如何让 main 函数等待 Goroutine 执行完？

// 方法一：使用 time.Sleep 函数
func main() {
	go PrintA()
	fmt.Println("main")
	time.Sleep(time.Second) // main 函数退出前等一会
}
// main
// A

// 方法二：使用空的 select 语句
func main() {
	go PrintA()
	fmt.Println("main")
	select {}
}
// main
// A
// fatal error: all goroutines are asleep - deadlock!
// goroutine 1 [select (no cases)]:
// main.main()
// 	/tmp/sandbox905752354/prog.go:17 +0x6a
// 
// “A” 字符串是输出了，但程序也出现异常了
// 原因是，当程序中存在运行的 Goroutine，select{} 就会一直等待，如果 Goroutine 都执行结束了，没有什么可等待的了，就会抛出异常
// select{} 使用场景：爬虫项目，创建了 Goroutine，需要一直爬取数据，不需要停止

// 方法三：使用 WaitGroup 类型等待 Goroutine 结束，项目中常常使用
package main

import (
	"fmt"
	"sync"
)
// 声明 WaitGroup 类型变量 wg，使用时无需初始化
var wg sync.WaitGroup

func PrintA() {
	fmt.Println("A")
    // 当一个 Goroutine 运行完后使用 wg.Done() 通知
	wg.Done()
}

func main() {
    // wg.Add(1) 表示需要等待一个 Goroutine，如果有两个，使用 Add(2)
	wg.Add(1)
	go PrintA()
    // wg.Wait() 等待 Goroutine 执行完
	wg.Wait()
	fmt.Println("main")
}
// A
// main
```

**给其它任务 “让行” Vs 终止自身协程**  
在程序运行中，某些特定的情况下需要暂停当前协程，让其它协程任务先执行。
```go
func main() {
   go fmt.Println("Hello World")
   runtime.Gosched() // 使主线程中的任务让出资源，优先执行上面协程输出文本
   fmt.Println("程序运行结束")
   // Hello World
   // 程序运行结束
}
```

在某些条件下，我们还希望立即停止协程任务的执行。方法便是使用调用 runtime.Goexit () 函数。
```go
func main() {
    syncWait.Add(1)
    go testFunc()
    syncWait.Wait()
    fmt.Println("程序运行结束")
}

func testFunc() {
    defer syncWait.Done()
    for i := 1; i < 100; i++ {
        fmt.Println(i)
        if i >= 5 {
            runtime.Goexit()
        }
    }
}
```

**控制并发数**  
Go 语言中可以控制使用 CPU 的核心数量，从 Go1.5 版本开始，默认设置为 CPU 的总核心数。如果想自定义设置，使用如下函数：
```go
// num 如果大于 CPU 的核心数，也是允许的，Go 语言调度器会将很多的 Goroutine 分配到不同的处理器上
num := 2
runtime.GOMAXPROCS(num)
```

**通道**  
Goroutine 通信使用 “通道 (channel)”，如果 Goroutine1 想发送数据给 Goroutine2，就把数据放到通道里，Goroutine2 直接从通道里拿就行，反过来也是一样。  
在给通道放数据时，也可以指定通道放置的数据类型。`不要用共享内存来通信，要用通信来共享内存。`

创建通道时，分为无缓冲和有缓冲两种。  
1) 无缓冲  
无缓冲的通道，不要进行同步读写，不然会阻塞。
```go
// 定义了一个存储数据类型为 string 的无缓冲通道
strChan := make(chan string)
// 如果想存储任意类型，那数据类型设置为空接口
allChan := make(chan interface{})

// 给通道里放数据
strChan := make(chan string)
// 将 “老苗” 字符串送入 strChan 通道变量
// 这样放数据是会报错的，因为 strChan 变量是无缓冲通道，放入数据时 main 函数会一直等待，因此会造成死锁
strChan <- "老苗"
// 解决死锁情况，就要保证有地方在异步读通道，因此需要创建一个 Goroutine 来负责
package main
import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
// Read 函数负责读取通道数据，并打印
func Read(strChan chan string) {
    // <-strChan 表示从通道里拿数据，如果通道里没有数据它会进行阻塞
	data := <-strChan
	fmt.Println(data)
	wg.Done()
}
func main() {
	wg.Add(1)
	strChan := make(chan string)
	go Read(strChan)
    // 通道是引用类型，因此传递时无需使用指针
	strChan <- "老苗"
    // wg.Wait() 等待 Read 异步函数执行完
	wg.Wait()
}
// 老苗
```
2) 有缓冲  
对于无缓冲通道，它会产生阻塞。为了不让阻塞，必须创建一个 Goroutine 负责从通道读取才行。而有缓冲的通道，会有缓冲的余地。
```go
// 创建缓冲通道
bufferChan := make(chan  string, 3) // 创建一个存储数据类型为 string 的通道，可以缓冲 3 个数据，即给通道送入 3 个数据不会进行阻塞

package main
import "fmt"
func main() {
	bufferChan := make(chan string, 3)
	bufferChan <- "a"
	bufferChan <- "b"
	bufferChan <- "c"
	fmt.Println(<-bufferChan)
    // 当存入数量超过 3 时，就需要 Goroutine 异步读取
}
// a
```
缓冲通道何时使用（按照先入先出规则存取），例如：  
爬虫数据，第 1 个 Goroutine 负责爬取数据，第 2 个 Goroutine 负责处理和存储数据。 当第 1 个的处理速度大于第 2 个时，可以使用缓冲通道暂存起来。  
暂存起来后，第 1 个 Goroutine 就可以继续爬取，而不像无缓冲通道，放入数据时会阻塞，直到通道数据被读出，才能进行。  

**单向通道**  
上面说的是双向通道，指的就是即可以存，又可以取。  
单向通道，创建的通道是无法传递数据的，即如果只能读的通道，没法存数据。
```go
// readChan 只能读取数据
readChan := make(<-chan string)
// writeChan 只能存取数据
writeChan := make(chan<- string)

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 写通道
func write(data chan<- int) {
	data <- 520
	wg.Done()
}

// 读通道
func read(data <-chan int) {
	fmt.Println(<-data)
	wg.Done()
}

func main() {
	wg.Add(2)
    // 创建了两个 Goroutine，read 函数负责只读，write 函数负责只写
	dataChan := make(chan int)
    // 通道传递时，将双向通道转化为单向通道
	go write(dataChan)
	go read(dataChan)
	wg.Wait()
}
// 520
```

**遍历通道**  
在实际项目中，通道里会产生大量的数据，这时候就要循环的从通道里读取。
```go
// 1、给通道里循环写入数字
func write(data chan<- int) {
	for i := 0; i < 10; i++ {
		data <- i
	}
	wg.Done()
}

// 使用死循环读取数据
func read(data <-chan int) {
	for {
		d := <-data
		fmt.Println(d)
	}
	wg.Done()
}

// read 函数在读取通道时是不知道数据写入完了，如果读取不到数据，它会一直阻塞，因此，如果写数据完成时，需要使用 close 函数关闭通道
func write(data chan<- int) {
	for i := 0; i < 10; i++ {
		data <- i
	}
    close(data)
	wg.Done()
}
// 关闭后，读取通道时也需要检测判断
func read(data <-chan int) {
	for {
		d, ok := <-data
        // ok 变量为 false 时，表示通道已关闭
        // 关闭通道后，ok 变量不会立马变成 false，而是等已放入通道的数据都读取完
		if !ok {
			break
		}

		fmt.Println(d)
	}
	wg.Done()
}

// 2、for-range 使用 for-range 语句读取通道，这比死循环使用起来简单一点
func read(data <-chan int) {
    // 如果想退出 for-range 语句，也需要关闭通道
    // 如果关闭通道后，不需要增加 ok 判断，等通道数据读取完，自行会退出
	for d := range data {
		fmt.Println(d)
	}
	wg.Done()
}
```

**通道函数**  
使用 len 函数获取通道里还有多少个消息未读，cap 函数获取通道的缓冲大小。
```go
ch := make(chan  int, 3)
ch<-1
fmt.Println(len(ch)) // 1
fmt.Println(cap(ch)) // 3
```

**select 语句**  
select 语句 和 switch 语句类似，它也有 case 分支，也有 default 分支，但 select 语句的不同点有两个：
- case 分支只能是 “读通道” 或 “写通道”，如果读写成功，即不阻塞，则 case 分支就满足。
- fallthrough 关键字不能使用。

1）无 default 分支
select 语句会在 case 分支中选择一个可读写成功的通道。
```go
package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
    // 如果删除 ch1 <- 1 ，select 语句会在 main 函数中一直等待，因此会造成死锁
	ch1 <- 1

	select {
        // ch1 通道有数据，因此进入了第一个 case 分支
        // 读通道，也可以给通道写数据，例：case ch2<-2
        case v, ok := <-ch1:
            if ok {
                fmt.Println("ch1通道", v)
            }

        case v, ok := <-ch2:
            if ok {
                fmt.Println("ch2通道", v)
            }
        }
}
// ch1通道
```
2）有 default 分支  
为了防止 select 语句出现死锁，可以增加 default 分支。意思就是，当没有一个 case 分支可以进行通道读写，那就走 default 分支。  
```go
package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	select {
        case v, ok := <-ch1:
            if ok {
                fmt.Println("ch1通道", v)
            }

        case v, ok := <-ch2:
            if ok {
                fmt.Println("ch2通道", v)
            }

        default:
            fmt.Println("没有可读写通道")
	}
}
// 没有可读写通道
```

### 线程模型与调度器
goroutine建立在操作系统线程基础之上，它与操作系统线程之间实现了一个多对多(M:N)的两级线程模型。  
这里的 M:N 是指M个goroutine运行在N个操作系统线程之上，内核负责对这N个操作系统线程进行调度，而这N个系统线程又负责对这M个goroutine进行调度和运行。

所谓的对goroutine的调度，是指程序代码按照一定的算法在适当的时候挑选出合适的goroutine并放到CPU上去运行的过程，这些负责对goroutine进行调度的程序代码我们称之为goroutine调度器。用极度简化了的伪代码来描述goroutine调度器的工作流程大概是下面这个样子：
```go
// 程序启动时的初始化代码
......
for i := 0; i < N; i++ { // 创建N个操作系统线程执行schedule函数
    create_os_thread(schedule) // 创建一个操作系统线程执行schedule函数
}

//schedule函数实现调度逻辑
func schedule() {
   for { //调度循环
         // 根据某种算法从M个goroutine中找出一个需要运行的goroutine
         g := find_a_runnable_goroutine_from_M_goroutines()
         run_g(g) // CPU运行该goroutine，直到需要调度其它goroutine才返回
         save_status_of_g(g) // 保存goroutine的状态，主要是寄存器的值
    }
}
```
这段伪代码表达的意思是，程序运行起来之后创建了N个由内核调度的操作系统线程去执行shedule函数，而schedule函数在一个调度循环中反复从M个goroutine中挑选出一个需要运行的goroutine并跳转到该goroutine去运行，直到需要调度其它goroutine时才返回到schedule函数中通过save_status_of_g保存刚刚正在运行的goroutine的状态然后再次去寻找下一个goroutine。

内核对系统线程的调度简单的归纳为：在执行操作系统代码时，内核调度器按照一定的算法挑选出一个线程并把该线程保存在内存之中的寄存器的值放入CPU对应的寄存器从而恢复该线程的运行。  
万变不离其宗，系统线程对goroutine的调度与内核对系统线程的调度原理是一样的，实质都是通过保存和修改CPU寄存器的值来达到切换线程/goroutine的目的。

所谓的goroutine调度，是指程序代码按照一定的算法在适当的时候挑选出合适的goroutine并放到CPU上去运行的过程。这句话揭示了调度系统需要解决的三大核心问题：  
- 调度时机：什么时候会发生调度？  
- 调度策略：使用什么策略来挑选下一个进入运行的goroutine？  
- 切换机制：如何把挑选出来的goroutine放到CPU上运行？  

