
### switch 语句
Go语言改进了 switch 的语法设计，case 与 case 之间是独立的代码块，不需要通过 break 语句跳出当前 case 代码块以避免执行到下一行。  
```go
switch [var1] {
    case: var2
        // todo
    case: var3,var4
        // todo
    default:
        // todo
}
```
switch 语句有以下特点：
- var1 可以是任意类型，也可以是 “赋值语句”，也可以省略。
- case 后可以是变量（数量不限）、也可以是判断语句。
- switch 进入 case 后，默认只执行当前 case 分支，不用写 break。
- 如果 case 分支没有一个满足的，最终则执行 default 语句 ，类似 if 语句中的 else 分支。
- 使用 fallthrough 关键字，执行下一个 case 分支。
- case 分支也可以为空， 什么都不写。

```go
var a = "hello"
switch a {
case "hello":
    fmt.Println(1)
case "world":
    fmt.Println(2)
default:
    fmt.Println(0)
}

// 当出现多个 case 要放在一起的时候
var a = "mum"
switch a {
case "mum", "daddy":
    fmt.Println("family")
}

// case 后不仅仅只是常量，还可以和 if 一样添加表达式
var r int = 11
switch {
case r > 10 && r < 20:
    fmt.Println(r)
}

// 跨越 case 的 fallthrough——兼容C语言的 case 设计
var s = "hello"
switch {
case s == "hello":
    fmt.Println("hello")
    fallthrough
case s != "world":
    fmt.Println("world")
}

switch num1, num2 := 1, 2; {
case num1 >= num2:
    fmt.Println("num1 > num2")
case num1 < num2:
    fmt.Println("num1 < num2")
}
// 输出 num1 < num2
```

### type-switch 语句
type-switch 用于获取接口实际类型，对于不同的类型通过分支判断。
```go
var data interface{}
data = "111"

// data 是接口类型， .(type) 获取实际类型
// 将实际类型的值赋给 d 变量
switch d := data.(type) {
case string:
    // 进入分支后，d 是 string 类型
    fmt.Println(d + "str")
case int:
    // 进入分支后， d 是 int 类型
    fmt.Println(d + 1)
}
// 输出 111str
```

