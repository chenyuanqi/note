package main

import (
	"fmt"
	"time"
)

func dump(str string) {
	fmt.Println(str)
	go fmt.Println("dump中调用goroutine")
}

func main() {
	go dump("我在main里调用goroutine")
	fmt.Println("我在main中")
	time.Sleep(2 * time.Second)
}
