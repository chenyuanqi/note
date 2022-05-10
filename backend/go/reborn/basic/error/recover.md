
### recover：平息运行时恐慌
Go 语言的内建函数 recover 专用于恢复 panic，或者说平息运行时恐慌。recover 函数无需任何参数，并且会返回一个空接口类型的值。
如果用法正确，这个值实际上就是即将恢复的 panic 包含的值。并且，如果这个 panic 是因我们调用 panic 函数而引发的，那么该值同时也会是我们此次调用 panic 函数时，传入的参数值副本。请注意，这里强调用法的正确。  

Recover 是一个 Go 语言的内建函数，可以让进入恐慌流程中的 goroutine 恢复过来，recover 仅在延迟函数 defer 中有效，在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果，如果当前的 goroutine 陷入恐慌，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行。  
通常来说，不应该对进入 panic 恐慌的程序做任何处理，但有时，需要我们可以从恐慌中恢复，至少我们可以在程序崩溃前，做一些操作，举个例子，当 web 服务器遇到不可预料的严重问题时，在崩溃前应该将所有的连接关闭，如果不做任何处理，会使得客户端一直处于等待状态，如果 web 服务器还在开发阶段，服务器甚至可以将异常信息反馈到客户端，帮助调试。  
`提示：Go 语言没有异常系统，其使用 panic 触发恐慌类似于其他语言的抛出异常，recover 的恐慌恢复机制就对应其他语言中的 try/catch 机制。`

**让程序在崩溃时继续执行**  
```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")
	defer func() {
		fmt.Println("Enter defer function.")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()
	// 引发panic
	panic(errors.New("something wrong"))
	fmt.Println("Exit function main.")
}
```

**panic 和 recover 的关系**  
panic 和 recover 的组合有如下特性：
- 有 panic 没 recover，程序恐慌。
- 有 panic 也有 recover，程序不会恐慌，执行完对应的 defer 后，从恐慌点退出当前函数后继续执行。

虽然 panic/recover 能模拟其他语言的异常机制，但并不建议在编写普通函数时也经常性使用这种特性。  
在 panic 触发的 defer 函数内，可以继续调用 panic，进一步将错误外抛，直到程序整体崩溃。  
如果想在捕获错误时设置当前函数的返回值，可以对返回值使用命名返回值方式直接进行设置。  
