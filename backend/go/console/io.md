

### Go 命令行输入
从键盘和标准输入 os.Stdin 读取输入，最简单的办法是使用 fmt 包提供的 Scan 和 Sscan 开头的函数。  
```golang
fmt.Println("Please enter your full name: ")
// Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行
fmt.Scanln(&firstName, &lastName)

// Scanf 的第一个参数用作格式字符串，用来决定如何读取
fmt.Scanf("%s %s", &firstName, &lastName)
fmt.Printf("Hi %s %s!\n", firstName, lastName)

// 从字符串中读取
fmt.Sscanf(input, format, &f, &i, &s)
mt.Println("From the string we read: ", f, i, s)
```

也可以使用 bufio 包提供的缓冲读取（buffered reader）来读取数据。  
```golang
var inputReader *bufio.Reader
var input string
var err error
// inputReader 是一个指向 bufio.Reader 的指针 (创建一个读取器，并将其与标准输入绑定)
inputReader = bufio.NewReader(os.Stdin)
fmt.Println("Please enter some input: ")
// ReadString(delim byte)，该方法从输入中读取内容，直到碰到 delim 指定的字符，然后将读取到的内容连同 delim 字符一起放到缓冲区
input, err = inputReader.ReadString('\n') // 或者省略变量声明 input, err := inputReader.ReadString('\n')
if err == nil {
    fmt.Printf("The input was: %s\n", input)
}
switch input {
case "Philip\r\n", "Ivo\r\n":   fmt.Printf("Welcome %s\n", input)
default: fmt.Printf("You are not welcome here! Goodbye!\n")
}
```

`注意：Unix 和 Windows 的行结束符是不同的。`
