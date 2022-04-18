
### 计算函数执行时间
函数的运行时间的长短是衡量这个函数性能的重要指标，特别是在对比和基准测试中，要得到函数的运行时间，最简单的办法就是在函数执行之前设置一个起始时间，并在函数运行结束时获取从起始时间到现在的时间间隔，这个时间间隔就是函数的运行时间。  
`由于计算机 CPU 及一些其他因素的影响，在获取函数运行时间时每次的结果都有些许不同，属于正常现象。`

```go
package main
import (
    "fmt"
    "time"
)
func test1() {
    start := time.Now() // 获取当前时间
    sum := 0
    for i := 0; i < 100000000; i++ {
        sum++
    }
    elapsed := time.Since(start)
    fmt.Println("该函数执行完成耗时：", elapsed)
}
func test2() {
    start := time.Now() // 获取当前时间
    sum := 0
    for i := 0; i < 100000000; i++ {
        sum++
    }
    elapsed := time.Now().Sub(start)
    fmt.Println("该函数执行完成耗时：", elapsed)
}
func main() {
    test1()
    test2()
}

```
