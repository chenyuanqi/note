package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doTask(n int) {
	time.Sleep(time.Duration(n))
	fmt.Printf("Task %d Done\n", n)
	wg.Done()
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doTask(i + 1)
	}
	// 等待所有的子协程任务全部完成，所有子协程结束后，才会执行 wg.Wait() 后面的代码
	wg.Wait()
	fmt.Printf("All Task Done\n")
}
