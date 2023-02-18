# 概述

调用 `os/signal` 包即可。

# 例子

## 监听信号

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	fmt.Println("程序执行中... 按 Ctrl + C 终止执行")

	<-c // 等待信号被触发
	fmt.Println("程序执行终止")
}

// $ go run main.go
// 输出如下 
/**
  程序执行中... 按 Ctrl + C 终止执行
  ^C程序执行终止
*/
```

# 限制

`SIGKILL` 信号 和 `SIGSTOP` 信号无法被捕获，这是内核的限制， 目的是为了让操作系统控制进程的生命周期。