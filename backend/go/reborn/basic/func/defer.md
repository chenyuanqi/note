
### go defer
Go 语言的 defer 语句会将其后面跟随的语句进行延迟处理，在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行，也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。延迟调用是在 defer 所在函数结束时进行，函数结束可以是正常返回时，也可以是发生宕机时。  

关键字 defer 的用法类似于面向对象编程语言 Java 和 C\# 的 finally 语句块，它一般用于释放某些已分配的资源，典型的例子就是对一个互斥解锁，或者关闭一个文件。  

**多个延迟执行语句的处理顺序**  
当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）。
```go
package main
import (
    "fmt"
)
func main() {
    fmt.Println("defer begin")
    // 将defer放入延迟调用栈
    defer fmt.Println(1)
    defer fmt.Println(2)
    // 最后一个放入, 位于栈顶, 最先调用
    defer fmt.Println(3)
    fmt.Println("defer end")
}
// defer begin
// defer end
// 3
// 2
// 1
```

**使用延迟执行语句在函数退出时释放资源**  
处理业务或逻辑中涉及成对的操作是一件比较烦琐的事情，比如打开和关闭文件、接收请求和回复请求、加锁和解锁等。在这些操作中，最容易忽略的就是在每个函数退出处正确地释放和关闭资源。  
defer 语句正好是在函数退出时执行的语句，所以使用 defer 能非常方便地处理资源释放问题。  

1) 使用延迟并发解锁
```go
var (
    // 一个演示用的映射
    valueByKey = make(map[string]int)
    // 保证使用映射时的并发安全的互斥锁
    valueByKeyGuard sync.Mutex
)
// 根据键读取值
func readValue(key string) int {
    valueByKeyGuard.Lock()
   
    // defer后面的语句不会马上调用, 而是延迟到函数结束时调用
    defer valueByKeyGuard.Unlock()
    return valueByKey[key]
}
```

2) 使用延迟释放文件句柄  
文件的操作需要经过打开文件、获取和操作文件资源、关闭资源几个过程，如果在操作完毕后不关闭文件资源，进程将一直无法释放文件资源，在下面的例子中将实现根据文件名获取文件大小的函数，函数中需要打开文件、获取文件大小和关闭文件等操作，由于每一步系统操作都需要进行错误处理，而每一步处理都会造成一次可能的退出，因此就需要在退出时释放资源，而我们需要密切关注在函数退出处正确地释放文件资源。  
```go
func fileSize(filename string) int64 {
    f, err := os.Open(filename)
    if err != nil {
        return 0
    }
    // 延迟调用Close, 此时Close不会被调用
    defer f.Close()
    info, err := f.Stat()
    if err != nil {
        // defer机制触发, 调用Close关闭文件
        return 0
    }
    size := info.Size()
    // defer机制触发, 调用Close关闭文件
    return size
}
```

