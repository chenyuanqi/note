
### Go 错误处理
Go 首选错误处理方式是返回值，而不是异常。考虑 strconv.Atoi 函数，它将接受一个字符串然后将它转换为一个整数。  
```golang
if len(os.Args) != 2 {
    os.Exit(1)
}
n, err := strconv.Atoi(os.Args[1])
if err != nil {
    fmt.Println("not a valid number")
} else {
    fmt.Println(n)
}

// 创建自己的错误类型,要求是你必须实现内建 error 接口的契约
type error interface {
    Error() string
}
// 或通过导入 errors 包然后使用它的 New 函数创建我们自己的错误
if count < 1 {
    return errors.New("Invalid count")
}

// Go 标准库中有一个使用 error 变量的通用模式
// io 包中有一个 EOF 变量（包级别的变量（被定义在函数之外））
var EOF = errors.New("EOF")
// 当我们从一个文件或者 STDIN 读取时，使用 io.EOF 
var input int
_, err := fmt.Scan(&input)
if err == io.EOF {
    fmt.Println("no more input!")
}

```

Go 确实有  panic  和  recover  函数。 panic  就像抛出异常，而 recover 就像 catch，它们很少使用。  
