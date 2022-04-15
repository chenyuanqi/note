
### for 语句
在 Go 语言中，重复地执行相似逻辑，可以使用循环结构实现。循环结构的格式如下：
```go
// for 表明接下来的代码是 for 循环结构
// init 是初始化语句
// condition 是关系或逻辑表达式，值为 false 时则会停止循环
// post 是每次循环结束后执行的语句
// 循环体代码块就是要重复执行的代码
for init; condition; post {
    // 循环体代码块
}
```
`注意： 使用循环时，务必确保有明确的可退出循环的条件，否则程序将陷入死循环，无法终止`

在 Go 的循环中，使用 break 语句打断相应循环的执行；使用 continue 语句提前终止本次循环，直接执行下一次循环。也可以指定标签跳转：  
```go
OuterLoop:
    for i := 0; i < 2; i++ {
        for j := 0; j < 5; j++ {
            switch j {
            case 2:
                fmt.Println(i, j)
                break OuterLoop
            case 3:
                fmt.Println(i, j)
                break OuterLoop
            }
        }
    }
    // 0 2

OuterLoop:
    for i := 0; i < 2; i++ {
        for j := 0; j < 5; j++ {
            switch j {
            case 2:
                fmt.Println(i, j)
                continue OuterLoop
            }
        }
    }
    // 0 2
    // 1 2
```

**for range 结构**  
for range 结构是 Go 语言特有的一种的迭代结构，在许多情况下都非常有用，for range 可以遍历数组、切片、字符串、map 及通道（channel），for range 语法上类似于其它语言中的 foreach 语句。
```go
// 遍历数据、切片——获得索引和值
for key, value := range []int{1, 2, 3, 4} {
    fmt.Printf("key:%d  value:%d\n", key, value)
}
// key:0  value:1
// key:1  value:2
// key:2  value:3
// key:3  value:4

// 遍历字符串——获得字符
var str = "hello 你好"
for key, value := range str {
    fmt.Printf("key:%d value:0x%x\n", key, value)
}
// key:0 value:0x68
// key:1 value:0x65
// ...

// 遍历 map——获得 map 的键和值
m := map[string]int{
    "hello": 100,
    "world": 200,
}
for key, value := range m {
    fmt.Println(key, value)
}
// hello 100
// world 200

// 在遍历中选择希望获得的变量
m := map[string]int{
    "hello": 100,
    "world": 200,
}
for _, value := range m {
    fmt.Println(value)
}

// 遍历通道（channel）——接收通道数据
c := make(chan int)
go func() {
    c <- 1
    c <- 2
    c <- 3
    close(c)
}()
for v := range c {
    fmt.Println(v)
}
```

