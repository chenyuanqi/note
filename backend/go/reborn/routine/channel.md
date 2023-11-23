
### Channel

```go
package main

import "fmt"

func worker(c chan int) {
	fmt.Println(" I am a worker...")

	//从c中得到 main中传递过来的数据
	num := <-c
	fmt.Println("得到了 从main中 传递过来的数据是", num)
}

func main() {
	//创建一个chennel
	c := make(chan int)

	//开辟一个协程 去执行worker函数
	go worker(c)

	//main向c中写一个数字2
	c <- 2

	fmt.Println("I am main")
}
```