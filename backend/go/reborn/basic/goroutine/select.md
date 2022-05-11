
### Select
对于通道发送方和接收方多对一的情况，Go 语言还内置了 Select 结构，用于在接收方识别发送方的角色，然后对其发送来的数据做分别处理。

用下载工具来举例，如果将呈现在用户面前的 UI 界面作为接收方，任务的调度（即下载开始、暂停、结束、删除等等）和下载进度的回传（即已完成的下载百分比）作为两个发送方。这两个发送方通过各自的通道同时向接收方发送数据，接收方则根据通道的不同，对数据做相应的处理和展示。
```go
package main

import (
	"fmt"
	"time"
)

type process struct {
	current int
	total   int
}

func main() {
	chan1 := make(chan process)
	chan2 := make(chan int)
	go recvFunc(chan1, chan2)
	go sendFunc1(chan1)
	go sendFunc2(chan2)
	time.Sleep(10 * time.Second)
	fmt.Println("下载完成")
}

func sendFunc1(chan1 chan process) {
	for i := 0; i < 10; i++ {
		chan1 <- process{
			current: i,
			total:   10,
		}
		time.Sleep(1 * time.Second)
	}
}

func sendFunc2(chan2 chan int) {
	time.Sleep(2 * time.Second)
	chan2 <- 1
	time.Sleep(2 * time.Second)
	chan2 <- 1
}

func recvFunc(chan1 chan process, chan2 chan int) {
	for {
		select {
		case processInfo := <-chan1:
			fmt.Printf("当前任务进度：%d\n", 100.0*processInfo.current/processInfo.total)
		case <-chan2:
			fmt.Println("添加了新任务")
		}
	}
}
// 当前任务进度：0
// 当前任务进度：10
// 当前任务进度：20
// 添加了新任务
// 当前任务进度：30
// 当前任务进度：40
// 添加了新任务
// 当前任务进度：50
// 当前任务进度：60
// 当前任务进度：70
// 当前任务进度：80
// 当前任务进度：90
// 下载完成
```

