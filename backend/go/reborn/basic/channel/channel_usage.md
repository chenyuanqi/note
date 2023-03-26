以下是几个常见的使用 channel 的场景及相应的示例代码：

1. 用于协调多个 goroutine 的执行顺序

```go
package main

import "fmt"

func main() {
    ch1 := make(chan bool)
    ch2 := make(chan bool)

    go func() {
        fmt.Println("goroutine 1")
        ch1 <- true // 向 ch1 发送数据
    }()

    go func() {
        <-ch1 // 从 ch1 中接收数据
        fmt.Println("goroutine 2")
        ch2 <- true // 向 ch2 发送数据
    }()

    <-ch2 // 从 ch2 中接收数据
    fmt.Println("main goroutine")
}
```

上述代码中，通过使用两个 channel ch1 和 ch2，来协调三个 goroutine 的执行顺序。首先启动两个 goroutine，当第一个 goroutine 执行完毕时，向 ch1 发送数据；第二个 goroutine 从 ch1 中接收数据，执行完毕后向 ch2 发送数据；最后在主 goroutine 中从 ch2 中接收数据，完成整个流程。

2. 用于限制 goroutine 的并发数量

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    maxGoroutines := 5
    goroutines := make(chan struct{}, maxGoroutines)

    for i := 0; i < 10; i++ {
        wg.Add(1)
        goroutines <- struct{}{} // 向 goroutines 中发送空结构体
        go func(id int) {
            defer func() {
                <-goroutines // 从 goroutines 中接收数据
                wg.Done()
            }()
            fmt.Printf("goroutine %d\n", id)
        }(i)
    }

    wg.Wait()
}
```

上述代码中，通过使用一个带缓冲的 channel goroutines，来限制 goroutine 的并发数量。在启动每个 goroutine 之前，向 goroutines 中发送一个空结构体，如果 goroutines 中已经有了 maxGoroutines 个空结构体，则会阻塞，等待其他 goroutine 执行完毕后再继续执行。在每个 goroutine 执行完毕后，从 goroutines 中接收一个空结构体，以便让其他 goroutine 继续执行。

3. 用于传递数据

```go
package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        ch <- 1 // 向 ch 中发送数据
    }()

    data := <-ch // 从 ch 中接收数据
    fmt.Println(data)
}
```

上述代码中，通过使用一个 channel ch，来传递数据。在启动 goroutine 之前，创建一个 channle ch，然后在 goroutine 中向 ch 中发送数据。在主 goroutine 中从 ch 中接收数据，并输出。

4. 用于实现同步操作

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    ch := make(chan bool)

    go func() {
        fmt.Println("goroutine")
        ch <- true // 向 ch 中发送数据
    }()

    wg.Add(1)
    go func() {
        <-ch // 从 ch 中接收数据
        fmt.Println("main goroutine")
        wg.Done()
    }()

    wg.Wait()
}
```

上述代码中，通过使用一个 channel ch，来实现 goroutine 和主 goroutine 的同步操作。在启动 goroutine 之前，创建一个 channle ch，然后在 goroutine 中向 ch 中发送数据。在主 goroutine 中从 ch 中接收数据，等待 goroutine 执行完毕后再继续执行。

5. 用于处理信号

```go
package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

    fmt.Println("waiting for signal...")
    sig := <-sigCh // 从 sigCh 中接收信号
    fmt.Printf("received signal: %v\n", sig)
}
```

上述代码中，通过使用一个 channel sigCh，来处理操作系统的信号。使用 signal.Notify 函数将需要处理的信号注册到 sigCh 中，然后在主 goroutine 中等待信号的到来，使用 <- 操作符从 sigCh 中接收信号，并输出相应的信息。

