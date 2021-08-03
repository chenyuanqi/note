
### Go 错误处理
Go 语言错误处理机制非常简单明了，不需要学习了解复杂的概念、函数和类型，Go 语言为错误处理定义了一个标准模式，即 error 接口，该接口的定义非常简单。  
Go 语言的错误和其他语言的错误和异常不同，它们就是从函数或者方法中返回的、和其他返回值并没有什么区别的普通 Go 对象而已，如果程序出错，要如何处理程序下一步的动作，是退出程序还是警告后继续执行，决定权完全在开发者手上。  
```golang
// 创建自己的错误类型,要求是你必须实现内建 error 接口的契约
type error interface { 
	// 只声明了一个 Error() 方法，用于返回字符串类型的错误消息
    Error() string 
}
// 或通过导入 Go 标准错误包 errors，然后使用它的 New 函数创建我们自己的错误
if count < 1 {
    return errors.New("Invalid count")
}

// 对于大多数函数或类方法，如果要返回错误，基本都可以定义成如下模式 —— 将错误类型作为第二个参数返回
func Foo(param int) (n int, err error) { 
    // ...
}
// 然后在调用返回错误信息的函数 / 方法时，按照如下「卫述语句」模板编写处理代码即可
n, err := Foo(0)
if err != nil { 
    // 错误处理 
} else {
    // 使用返回值 n 
}
```

Go 首选错误处理方式是返回值，而不是异常。考虑 strconv.Atoi 函数，它将接受一个字符串然后将它转换为一个整数。  
```golang
if len(os.Args) != 2 {
    os.Exit(1)
}
n, err := strconv.Atoi(os.Args[1])
if err != nil {
	// 在打印错误信息时，直接传入了 err 对象实例，因为 Go 底层会自动调用 err 实例上的 Error() 方法返回错误信息并将其打印出来，就像普通类的 String() 方法一样
    fmt.Println(err)
	// 可以通过 fmt.Errorf() 格式化方法返回 error 类型错误，其底层调用的其实也是 errors.New 方法
	fmt.Errorf("not a valid number")) 
} else {
    fmt.Println(n)
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

除了上面这种最基本的、使用 errors.New() 方法返回包含错误信息的错误实例之外，Go 语言内置的很多包还封装了更复杂的错误类型。  
以 os 包为例，这个包主要负责与操作系统打交道，所以提供了 LinkError、PathError、SyscallError 这些实现了 error 接口的错误类型，以 PathError 为例，顾名思义，它主要用于表示路径相关的错误信息，比如文件不存在，其底层类型结构信息如下：
```golang
type PathError struct {
    Err  error
	// 额外的操作类型字段 Op 和文件路径字段 Path 以丰富错误信息，方便定位问题
    Op   string
    Path string
}
// 该类型的 Error() 方法实现
func (e *PathError) Error() string { 
    return e.Op + " " + e.Path + ": " + e.Err.Error() 
}
// 可以在调用 os 包方法出错时通过 switch 分支语句判定具体的错误类型，然后进行相应的处理
// 获取指定路径文件信息，对应类型是 FileInfo
// 如果文件不存在，则返回 PathError 类型错误
_, err := os.Stat("test.txt") 
if err != nil {
    switch err.(type) {
    case *os.PathError:
        // do something
    case *os.LinkError:
        // dome something
    case *os.SyscallError:
        // dome something
    case *exec.Error:
        // dome something
    }
} else {
    // ...
}
```
`注意：我们也可以仿照 PathError 的实现自定义一些复杂的错误类型，只需要组合 error 接口并实现 Error() 方法即可，然后按照自己的需要为自定义类型添加一些属性字段。`

### Go defer 语句
Go 语言的 defer 语句相当于 Java/PHP 中的析构函数和 finally 语句的功效。通过 defer 关键字声明兜底执行或者释放资源的语句可以轻松解决某个资源使用完毕后将其释放的问题。  

在函数执行完毕后或者运行抛出 panic 时执行，如果一个函数定义了多个 defer 语句，则按照先进后出的顺序执行。  
一个函数 / 方法中可以存在多个 defer 语句，defer 语句的调用顺序遵循先进后出的原则，即最后一个 defer 语句将最先被执行，相当于「栈」这个数据结构，如果在循环语句中包含了 defer 语句，则对应的 defer 语句执行顺序依然符合先进后出的规则。  
由于 defer 语句的执行时机和调用顺序，所以我们要尽量在函数 / 方法的前面定义它们，以免在后面编写代码时漏掉，尤其是运行时抛出错误会中断后面代码的执行，也就感知不到后面的 defer 语句。  
```golang
func ReadFile(filename string) ([]byte, error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    // defer 修饰的 f.Close() 方法会在函数执行完成后或读取文件过程中抛出错误时执行，以确保已经打开的文件资源被关闭，从而避免内存泄露
    defer f.Close()

    var n int64 = bytes.MinRead

    if fi, err := f.Stat(); err == nil {
        if size := fi.Size() + bytes.MinRead; size > n {
            n = size
        }
    }
    return readAll(f, n)
}

// 如果一条语句干不完清理的工作，也可以在 defer 后加一个匿名函数来执行对应的兜底逻辑
defer func() { 
    //  执行复杂的清理工作... 
} ()

// 函数 / 方法的前面定义 defer
func printError()  {
    fmt.Println("兜底执行")
}
defer printError()
// 遇到除数为 0，则抛出 panic，然后立即中断当前函数的执行（后续其他语句都不再执行），并按照先进后出顺序依次执行已经在当前函数中声明过的 defer 语句，最后打印出 panic 日志及错误信息
defer func() {
    fmt.Println("除数不能是 0!")
}()
var i = 1
var j = 1
var k = i / j
fmt.Printf("%d / %d = %d\n", i, j, k)
```


### Go panic 和 recover 
Go 语言没有像 Java、PHP 那样引入异常的概念，也没有提供 try...catch 这样的语法对运行时异常进行捕获和处理，当代码运行时出错，而又没有在编码时显式返回错误时，Go 语言会抛出 panic，中文译作「运行时恐慌」，我们也可以将其看作 Go 语言版的异常。  
Go 语言除了可以底层抛出 panic，我们还可以在代码中显式抛出 panic，以便对错误和异常信息进行自定义。  
```golang
defer func() {
    fmt.Println("代码清理逻辑")
}()
var i = 1
var j = 0
if j == 0 {
    panic("除数不能为 0！")
}
k := i / j
fmt.Printf("%d / %d = %d\n", i, j, k)
```
panic 函数支持的参数类型是 interface{}，即可以传入任意类型的参数。  
抛出 panic，处理机制：当遇到 panic 时，Go 语言会中断当前协程（即 main 函数）后续代码的执行，然后执行在中断代码之前定义的 defer 语句（按照先入后出的顺序），最后程序退出并输出 panic 错误信息，以及出现错误的堆栈跟踪信息。通过这些信息，可以帮助你快速定位问题并予以解决。  
```golang
panic(500)   // 传入数字
panic(errors.New("除数不能为0"))  // 传入 error 类型
```

通过 recover() 函数对 panic 进行捕获和处理，从而避免程序崩溃然后直接退出，而是继续可以执行后续代码，实现类似 Java、PHP 中 try...catch 语句的功能。  
由于执行到抛出 panic 的问题代码时，会中断后续其他代码的执行，所以，显然这个 panic 的捕获应该放到 defer 语句中完成，才可以在抛出 panic 时通过 recover 函数将其捕获，defer 语句执行完毕后，会退出抛出 panic 的当前函数，回调调用它的地方继续后续代码的执行。  
`可以类比为 panic、recover、defer 组合起来实现了传统面向对象编程异常处理的 try…catch…finally 功能。`  
```golang
func divide() {
    defer func() {
    	// 通过 recover() 函数捕获 panic，，并打印捕获到的错误信息
    	// 此时，程序会退出 divide() 函数而不是整个应用代码
        if err := recover(); err != nil {
            fmt.Printf("Runtime panic caught: %v\n", err)
        }
    }()
    var i = 1
    var j = 0
    k := i / j
    fmt.Printf("%d / %d = %d\n", i, j, k)
}
divide()
```
