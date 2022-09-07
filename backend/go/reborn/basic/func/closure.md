
### 闭包函数
Go语言支持匿名函数，即在需要使用函数时再定义函数，匿名函数没有函数名只有函数体，函数可以作为一种类型被赋值给函数类型的变量，匿名函数也往往以变量方式传递，这与C语言的回调函数比较类似，不同的是，Go语言支持随时在代码里定义匿名函数。  
Go语言中闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量，因此，简单的说：`函数 + 引用环境 = 闭包`。  
一个函数类型就像结构体一样，可以被实例化，函数本身不存储任何信息，只有与引用环境结合后形成的闭包才具有“记忆性”，`函数是编译期静态的概念，而闭包是运行期动态的概念`。  

匿名函数的优越性在于可以直接使用函数内的变量，不必申明；匿名函数就是开辟一片密闭空间。  
匿名函数，顾名思义就是没有名字的函数，匿名函数最大的用途是来模拟块级作用域，避免数据污染的。  
闭包，说白了就是函数的嵌套，内层的函数可以使用外层函数的所有变量，即使外层函数已经执行完毕。  
匿名函数和闭包其实是一回事儿，匿名函数就是闭包。匿名函数给编程带来灵活性的同时也容易产生 bug，在使用过程当中要多注意函数的参数，及可接受的参数的问题。  

匿名函数是指不需要定义函数名的一种函数实现方式，由一个不带函数名的函数声明和函数体组成。
> 为什么需要闭包？  
> 使用闭包和使用普通函数的最大区别在于：  
> 如果是普通函数，那就是一次性买卖，函数执行完毕后就无法再更改函数中变量的值；  
> 使用闭包，函数就成为了一个变量的值。只要变量还在，函数就会一直处于存活并独享内部状态。方便后期更改函数中变量的值。  
> 另外，使用闭包还能起到一定的数据保护作用。

如果是普通函数，那就是一次性买卖，函数执行完毕后就无法再更改函数中变量的值；
使用闭包，函数就成为了一个变量的值。只要变量还在，函数就会一直处于存活并独享内部状态。方便后期更改函数中变量的值。

匿名函数的用途非常广泛，它本身就是一种值，可以方便地保存在各种容器中实现回调函数和操作封装。  
```go
// 匿名函数的定义格式
func(参数列表)(返回参数列表){
    函数体
}

// 在定义时调用匿名函数
func(data int) {
    fmt.Println("hello", data)
}(100)

// 将匿名函数赋值给变量
// 将匿名函数体保存到f()中
f := func(data int) {
    fmt.Println("hello", data)
}
// 使用f()调用
f(100)
```

**匿名函数用作回调函数**  
匿名函数作为回调函数的设计在 Go 语言的系统包中比较常见，strings 包中就有类似的设计。  
```go
func TrimFunc(s string, f func(rune) bool) string {
    return TrimRightFunc(TrimLeftFunc(s, f), f)
}
```

比如，对切片的遍历操作，遍历中访问每个元素的操作使用匿名函数来实现，用户传入不同的匿名函数体可以实现对元素不同的遍历操作。  
```go
package main
import (
    "fmt"
)
// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {
    for _, v := range list {
        f(v)
    }
}
func main() {
    // 使用匿名函数打印切片内容
    visit([]int{1, 2, 3, 4}, func(v int) {
        fmt.Println(v)
    })
}
```

**使用匿名函数实现操作封装**  
比如将匿名函数作为 map 的键值，通过命令行参数动态调用匿名函数。  
```go
package main
import (
    "flag"
    "fmt"
)
// 定义命令行参数 skill，从命令行输入 --skill 可以将=后的字符串传入 skillParam 指针变量
var skillParam = flag.String("skill", "", "skill to perform")
func main() {
    // 解析命令行参数，解析完成后，skillParam 指针变量将指向命令行传入的值
    flag.Parse()
    var skill = map[string]func(){
        "fire": func() {
            fmt.Println("chicken fire")
        },
        "run": func() {
            fmt.Println("soldier run")
        },
        "fly": func() {
            fmt.Println("angel fly")
        },
    }
    if f, ok := skill[*skillParam]; ok {
        f()
    } else {
        fmt.Println("skill not found")
    }
}
```

**在闭包内部修改引用的变量**  
闭包对它作用域上部的变量可以进行修改，修改引用的变量会对变量进行实际修改。  
被捕获到闭包中的变量让闭包本身拥有了记忆效应，闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在，闭包本身就如同变量一样拥有了记忆效应。  
```go
// 准备一个字符串
str := "hello world"
// 创建一个匿名函数
foo := func() {
    // 匿名函数中访问str
    str = "hello dude"
}
// 调用匿名函数
foo()
fmt.Println(str) // hello dude
```

又如累加器的实现，
```go
package main
import (
    "fmt"
)
// 提供一个值, 每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
    // 返回一个闭包
    return func() int {
        // 累加
        value++
        // 返回一个累加值
        return value
    }
}
func main() {
    // 创建一个累加器, 初始值为1
    accumulator := Accumulate(1)
    // 累加1并打印
    fmt.Println(accumulator()) // 2
    fmt.Println(accumulator()) // 3
    // 打印累加器的函数地址
    fmt.Printf("%p\n", &accumulator) // 0xc42000c028
    // 创建一个累加器, 初始值为1
    accumulator2 := Accumulate(10) 
    // 累加1并打印
    fmt.Println(accumulator2()) // 11
    // 打印累加器的函数地址
    fmt.Printf("%p\n", &accumulator2) // 0xc42000c038
}
```

再如闭包实现生成器，
```go
package main
import (
    "fmt"
)
// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen(name string) func() (string, int) {
    // 血量一直为150
    hp := 150
    // 返回创建的闭包
    return func() (string, int) {
        // 将变量引用到闭包中
        return name, hp
    }
}
func main() {
    // 创建一个玩家生成器
    generator := playerGen("high noon")
    // 返回玩家的名字和血量
    name, hp := generator()
    // 打印值
    fmt.Println(name, hp) // high noon 150
}
```
