
### if 语句
不同的条件执行不同的逻辑，Go 语言的条件分支，格式如下：  
```go
if [assignment1;]condition1 {
    // 条件condition1成立时要执行的语句
} else if [assignment2;]condition2 {
    // 条件condition2成立时要执行的语句
} else if [assignment3;]condition3 {
    // 条件condition3成立时要执行的语句
} else {
    // 以上三种条件都不成立时要执行的语句
}
```

if 语句有如下特点：
- if 后面不需要小括号包裹，后面 switch 和 for 语句也是一样
- 可以在条件判断前增加赋值语句，用赋值的结果进行条件判断

```go
num := 10
if num > 12 {
    fmt.Println("分支1")
} else if num > 9 {
    fmt.Println("分支2") // 10 > 9 为 true, 进入此分支
} else {
    fmt.Println("分支3")
}

// 判断函数错误并打印
//  “赋值语句” 的结果只在当前 if 语句中使用
if err := fun1(); err != nil {
    // 程序退出，并打印出错误
    panic(err)
}
```

### 三元运算
go 没有这样的三元运算 expr ? x : y，但是可以这样：  
```go
if expr {
    res = x
} else {
    res = y
}

func Min(x, y int) int {
    if x <= y {
        return x
    }
    return y
}
```