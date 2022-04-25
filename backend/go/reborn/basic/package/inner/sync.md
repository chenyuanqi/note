
### sync package：限制线程对变量的访问
Go 语言中 sync 包里提供了互斥锁 Mutex 和读写锁 RWMutex 用于处理并发过程中可能出现同时两个或多个协程（或线程）读或写同一个变量的情况。  

**为什么需要锁**  
锁是 sync 包中的核心，它主要有两个方法，分别是加锁（Lock）和解锁（Unlock）。

在并发的情况下，多个线程或协程同时其修改一个变量，使用锁能保证在某一时间内，只有一个协程或线程修改这一变量。不使用锁时，在并发的情况下可能无法得到想要的结果。
```go
package main
import (
    "fmt"
    "time"
)
func main() {
    var a = 0
    for i := 0; i < 1000; i++ {
        go func(idx int) {
            a += 1
            fmt.Println(a)
        }(i)
    }
    time.Sleep(time.Second)
}
// 1
// 226
// 405
// 816
// 22
// 23
// 24
// 25
// 26
// 27
// 28
// 29
// 30
// 31
// 32
// 33
```
通过运行结果可以看出 a 的值并不是按顺序递增输出的，这是为什么呢？协程的执行顺序大致如下所示：
- 从寄存器读取 a 的值；
- 然后做加法运算；
- 最后写到寄存器。

按照上面的顺序，假如有一个协程取得 a 的值为 3，然后执行加法运算，此时又有一个协程对 a 进行取值，得到的值同样是 3，最终两个协程的返回结果是相同的。  
而锁的概念就是，当一个协程正在处理 a 时将 a 锁定，其它协程需要等待该协程处理完成并将 a 解锁后才能再进行操作，也就是说同时处理 a 的协程只能有一个，从而避免上面示例中的情况出现。   

**互斥锁 Mutex**  
什么是互斥锁呢 ？互斥锁中其有两个方法可以调用：
- func (m *Mutex) Lock()
- func (m *Mutex) Unlock()

```go
package main
import (
    "fmt"
    "sync"
    "time"
)
func main() {
    var a = 0
    var lock sync.Mutex
    for i := 0; i < 1000; i++ {
        go func(idx int) {
            lock.Lock()
            defer lock.Unlock()
            a += 1
            fmt.Printf("goroutine %d, a=%d\n", idx, a)
        }(i)
    }
    // 等待 1s 结束主程序
    // 确保所有协程执行完
    time.Sleep(time.Second)
}
// goroutine 0, a=1
// goroutine 76, a=2
// goroutine 772, a=3
// goroutine 773, a=4
// goroutine 774, a=5
// goroutine 775, a=6
// ...
```
需要注意的是`一个互斥锁只能同时被一个 goroutine 锁定，其它 goroutine 将阻塞直到互斥锁被解锁（重新争抢对互斥锁的锁定）`。  
```go
package main
import (
    "fmt"
    "sync"
    "time"
)
func main() {
    ch := make(chan struct{}, 2)
    var l sync.Mutex
    go func() {
        l.Lock()
        defer l.Unlock()
        fmt.Println("goroutine1: 我会锁定大概 2s")
        time.Sleep(time.Second * 2)
        fmt.Println("goroutine1: 我解锁了，你们去抢吧")
        ch <- struct{}{}
    }()
    go func() {
        fmt.Println("goroutine2: 等待解锁")
        l.Lock()
        defer l.Unlock()
        fmt.Println("goroutine2: 欧耶，我也解锁了")
        ch <- struct{}{}
    }()
    // 等待 goroutine 执行结束
    for i := 0; i < 2; i++ {
        <-ch
    }
}
// goroutine1: 我会锁定大概 2s
// goroutine2: 等待解锁
// goroutine1: 我解锁了，你们去抢吧
// goroutine2: 欧耶，我也解锁了
```

**读写锁**  
读写锁有如下四个方法：
- 写操作的锁定和解锁分别是 func (*RWMutex) Lock 和 func (*RWMutex) Unlock；
- 读操作的锁定和解锁分别是 func (*RWMutex) Rlock 和 func (*RWMutex) RUnlock。

读写锁的区别在于：
- 当有一个 goroutine 获得写锁定，其它无论是读锁定还是写锁定都将阻塞直到写解锁；
- 当有一个 goroutine 获得读锁定，其它读锁定仍然可以继续；
- 当有一个或任意多个读锁定，写锁定将等待所有读锁定解锁之后才能够进行写锁定。

所以说这里的读锁定（RLock）目的其实是告诉写锁定，有很多协程或者进程正在读取数据，写操作需要等它们读（读解锁）完才能进行写（写锁定）。

我们可以将其总结为如下三条：
- 同时只能有一个 goroutine 能够获得写锁定；
- 同时可以有任意多个 gorouinte 获得读锁定；
- 同时只能存在写锁定或读锁定（读和写互斥）。

```go
package main
import (
    "fmt"
    "math/rand"
    "sync"
)
var count int
var rw sync.RWMutex
func main() {
    ch := make(chan struct{}, 10)
    for i := 0; i < 5; i++ {
        go read(i, ch)
    }
    for i := 0; i < 5; i++ {
        go write(i, ch)
    }
    for i := 0; i < 10; i++ {
        <-ch
    }
}
func read(n int, ch chan struct{}) {
    rw.RLock()
    fmt.Printf("goroutine %d 进入读操作...\n", n)
    v := count
    fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
    rw.RUnlock()
    ch <- struct{}{}
}
func write(n int, ch chan struct{}) {
    rw.Lock()
    fmt.Printf("goroutine %d 进入写操作...\n", n)
    v := rand.Intn(1000)
    count = v
    fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
    rw.Unlock()
    ch <- struct{}{}
}
// goroutine 4 进入写操作...
// goroutine 4 写入结束，新值为：81
// goroutine 0 进入读操作...
// goroutine 0 读取结束，值为：81
// goroutine 1 进入读操作...
// goroutine 1 读取结束，值为：81
// goroutine 0 进入写操作...
// goroutine 0 写入结束，新值为：887
// goroutine 2 进入读操作...
// goroutine 2 读取结束，值为：887
// goroutine 3 进入读操作...
// goroutine 3 读取结束，值为：887
// goroutine 4 进入读操作...
// goroutine 4 读取结束，值为：887
// goroutine 2 进入写操作...
// goroutine 2 写入结束，新值为：847
// goroutine 1 进入写操作...
// goroutine 1 写入结束，新值为：59
// goroutine 3 进入写操作...
// goroutine 3 写入结束，新值为：81

// 多个读操作同时读取一个变量时，虽然加了锁，但是读操作是不受影响的。（读和写是互斥的，读和读不互斥）
package main
import (
    "sync"
    "time"
)
var m *sync.RWMutex
func main() {
    m = new(sync.RWMutex)
    // 多个同时读
    go read(1)
    go read(2)
    time.Sleep(2*time.Second)
}
func read(i int) {
    println(i,"read start")
    m.RLock()
    println(i,"reading")// println(i,"reading")
    time.Sleep(1*time.Second)
    m.RUnlock()
    println(i,"read over")
}
// 2 read start
// 2 reading
// 1 read start
// 1 reading
// 2 read over
// 1 read over

// 由于读写互斥，所以写操作开始的时候，读操作必须要等写操作进行完才能继续，不然读操作只能继续等待
package main
import (
    "sync"
    "time"
)
var m *sync.RWMutex
func main() {
    m = new(sync.RWMutex)
    // 写的时候啥也不能干
    go write(1)
    go read(2)
    go write(3)
    time.Sleep(2*time.Second)
}
func read(i int) {
    println(i,"read start")
    m.RLock()
    println(i,"reading")
    time.Sleep(1*time.Second)
    m.RUnlock()
    println(i,"read over")
}
func write(i int) {
    println(i,"write start")
    m.Lock()
    println(i,"writing")
    time.Sleep(1*time.Second)
    m.Unlock()
    println(i,"write over")
}
// 3 write start
// 3 writing
// 1 write start
// 2 read start
// 3 write over
// 2 reading
// 2 read over
// 1 writing
```

