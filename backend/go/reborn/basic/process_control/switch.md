
### switch 语句
Go语言改进了 switch 的语法设计，case 与 case 之间是独立的代码块，不需要通过 break 语句跳出当前 case 代码块以避免执行到下一行。  
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
```