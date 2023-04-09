
### go defer
Go 语言的 defer 语句会将其后面跟随的语句进行延迟处理，在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行，也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。延迟调用是在 defer 所在函数结束时进行，函数结束可以是正常返回时，也可以是发生宕机时。  

关键字 defer 的用法类似于面向对象编程语言 Java 和 C\# 的 finally 语句块，它一般用于释放某些已分配的资源，典型的例子就是对一个互斥解锁，或者关闭一个文件。  
`defer 是在函数结束时调用，但是 defer 函数参数确是立即求值的，参数求值顺序也是自然的从左向右的；函数内装饰对局部变量是引用，所以在 defer 里会受影响。`  
`panic 后依然会执行 defer，主动 os.Exist 不会执行 defer。`  

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

### Defer 必掌握的 7 个知识点  
defer 是在函数返回之前执行，defer 的执行顺序是优先于 return。return 的执行是一个两步操作，先对 return 返回的值进行赋值，然后执行 defer 语句，最后将结果进行返回给函数的调用者。  

即使函数内发生了 panic 异常，panic 之前定义的 defer 仍然会被执行。  

defer 中存在子函数，子函数会按照 defer 的定义顺序执行。

**知识点1：defer的执行顺序**  
多个 defer 出现的时候，它是一个“栈”的关系，也就是先进后出。一个函数中，写在前面的 defer 会比写在后面的 defer 调用的晚。
```go
package main

import "fmt"

func main() {
    defer func1()
    defer func2()
    defer func3()
    // C
    // B
    // A
}

func func1() {
    fmt.Println("A")
}

func func2() {
    fmt.Println("B")
}

func func3() {
    fmt.Println("C")
}
```

**知识点2: defer 与 return 谁先谁后**  
return 之后的语句先执行，defer 后的语句后执行。  
```go
package main

import "fmt"

func deferFunc() int {
    fmt.Println("defer func called")
    return 0
}

func returnFunc() int {
    fmt.Println("return func called")
    return 0
}

func returnAndDefer() int {
    defer deferFunc()
    return returnFunc()
}

func main() {
    returnAndDefer()
    // return func called
    // defer func called
}
```

**知识点3：函数的返回值初始化**  
该知识点不属于 defer 本身，但是调用的场景却与 defer 有联系，所以也算是 defer 必备了解的知识点之一。  
如：func DeferFunc1(i int) (t int) {} 其中返回值 t int，这个 t 会在函数起始处被初始化为对应类型的零值并且作用域为整个函数。  
证明，只要声明函数的返回值变量名称，就会在函数初始化时候为之赋值为0，而且在函数体作用域可见。  
```go
package main

import "fmt"

func DeferFunc1(i int) (t int) { // 初始化 i=10，t=0
    fmt.Println("t = ", t)

    return 2 // 赋值 t=2
}

func main() {
    DeferFunc11(10) 
    // t =  0
}
```

**知识点4: 有名函数返回值遇见 defer 情况**  
在没有 defer 的情况下，其实函数的返回就是与 return 一致的，但是有了 defer 就不一样了。“先 return，再 defer”的原则，所以在执行完 return 之后，还要再执行 defer 里的语句，依然可以修改本应该返回的结果。  
```go
package main

import "fmt"

func returnButDefer() (t int) {  //t初始化0，并且作用域为该函数全域

    defer func() {
        t = t * 10
    }()

    return 1 // t赋值1
}

func main() {
    fmt.Println(returnButDefer()) 
    // 10
}
```

**知识点5: defer 遇见 panic**  
遇到 panic 时，遍历本协程的 defer 链表，并执行 defer。  
在执行 defer 过程中：遇到 recover 则停止 panic，返回 recover 处继续往下执行；如果没有遇到 recover，遍历完本协程的 defer 链表后，向 stderr 抛出 panic 信息。  
defer 最大的功能是 panic 后依然有效，所以 defer 可以保证你的一些资源一定会被关闭，从而避免一些异常出现的问题。
```go
package main

import (
    "fmt"
)

func main() {
    defer_call()

    fmt.Println("main 正常结束")
    // defer: panic 之前2, 不捕获
    // defer: panic 之前1, 捕获异常
    // 异常内容
    // main 正常结束
}

func defer_call() {

    defer func() {
        fmt.Println("defer: panic 之前1, 捕获异常")
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()

    defer func() { fmt.Println("defer: panic 之前2, 不捕获") }()

    panic("异常内容")  //触发defer出栈

	defer func() { fmt.Println("defer: panic 之后, 永远执行不到") }()
}
```

**知识点6: defer 中包含 panic**  
panic 仅有最后一个可以被 revover 捕获。  
触发 panic("panic") 后 defer 顺序出栈执行，第一个被执行的 defer 中会有 panic("defer panic") 异常语句，这个异常将会覆盖掉 main 中的异常 panic("panic")，最后这个异常被第二个执行的 defer 捕获到。
```go
package main

import (
    "fmt"
)

func main()  {
    defer func() {
       if err := recover(); err != nil{
           fmt.Println(err)
       }else {
           fmt.Println("fatal")
       }
    }()

    defer func() {
        panic("defer panic")
    }()

    panic("panic")
    // defer panic
}
```

**知识点7: defer 下的函数参数包含子函数**  
如下 4 个函数的先后执行顺序是什么呢？  
这里面有两个 defer， 所以 defer 一共会压栈两次，先进栈 1，后进栈 2。 那么在压栈 function1 的时候，需要连同函数地址、函数形参一同进栈，那么为了得到 function1 的第二个参数的结果，所以就需要先执行function3 将第二个参数算出，那么 function3 就被第一个执行。同理压栈 function2，就需要执行 function4 算出 function2 第二个参数的值。然后函数结束，先出栈 fuction2、再出栈 function1。所以顺序如下：
● defer 压栈 function1，压栈函数地址、形参1、形参2(调用 function3) --> 打印3  
● defer 压栈 function2，压栈函数地址、形参1、形参2(调用 function4) --> 打印4  
● defer 出栈 function2, 调用 function2 --> 打印2   
● defer 出栈 function1, 调用 function1--> 打印1  
```go
package main

import "fmt"

func function(index int, value int) int {
    fmt.Println(index)

    return index
}

func main() {
    defer function(1, function(3, 0))
    defer function(2, function(4, 0))
    // 3
    // 4
    // 2
    // 1
}
```

[更多参考](https://www.yuque.com/aceld/golang/qnubsg)  

