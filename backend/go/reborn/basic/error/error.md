
### Go 错误
在写代码时，不是所有情况都能处理，对于不能处理的逻辑，就需要使用错误机制告诉上层调用者。

在 Go 语言中，错误是被作为一个接口类型对待，它不像其它语言一样使用 try/catch 去捕捉，只需在函数或方法之间使用一个错误类型变量去传递。

**创建错误**  
创建错误，即实现错误接口。
```go
// Go 标准包内置的，所有创建的错误类型都需要实现此接口
type  error  interface {
    Error() string
}

// 创建方式1：Go 语言中内置了一个处理错误的标准包，不需要自己去实现 error 接口，它有函数帮你处理
import  "errors"
var  ErrNotFound = errors.New("not found")

// 创建方式2：fmt 标准包内也有一个创建错误的函数 Errorf ，该函数可以使用占位符设置错误信息，比 errors.New 函数更灵活
import  "fmt"
var  ErrHuman = fmt.Errorf("%s不符合我们人类要求", "老苗")

// 创建方式3：自定义错误类型
// ErrorPathNotExist 结构体
type  ErrorPathNotExist  struct {
    Filename string
}
// 实现 error 接口（方法的接收者没有被使用可以直接省略掉，也可以 func (e *ErrorPathNotExist) 或 func (_ *ErrorPathNotExist)）
func (*ErrorPathNotExist) Error() string {
    return  "文件路径不存在"
}
// 创建一个 ErrNotExist 错误类型变量
var  ErrNotExist  error = &ErrorPathNotExist{
    Filename: "./main.go",
}
```

**打印错误**  
在项目开发中，错误常常通过函数或方法的返回值携带，返回的位置也通常被放置在最后一位。
```go
package main
import (
    "fmt"
    "log"
    "io/ioutil"
)
// 读取文件内容
func  LoadConfig() (string, error) {
    filename := "./config.json"
    b, err := ioutil.ReadFile(filename)
    if err != nil {
        return  "", err
    }

    content := string(b)
    if  len(content) == 0 {
        return  "", errors.New("内容为空")
    }

    return content, nil
}

var  ErrEmpty = errors.New("内容为空")
func  LoadConfig() (string, error) {
    return  "", ErrEmpty
}

func  main() {
    content, err := LoadConfig()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("内容：", content) // xxx open ./config.json: The system cannot find the file specified.
}
```

**os.Exit**  
该函数通知程序退出，并且该函数之后的逻辑将不会被执行。在调用时需要指定退出码，为 0 时，表示正常退出程序。
```go
os.Exit(0)
// 不主动调用该函数，即程序从 main 函数自然结束时，默认的退出码为 0
// Process finished with exit code 0

// 如果不正常退出，退出码则为非 0，通常使用 1 表示未知错误（在使用 log.Fatal 函数时，内部就调用了 os.Exit(1)）
os.Exit(1)
```

**错误加工**  
1. 错误拼接  
在返回错误时，如果想携带附加的错误消息时，可以使用 fmt.Errorf。  
```go
func  LoadConfig() (string, error) {
    filename := "./config.json"
    b, err := ioutil.ReadFile(filename)
    if err != nil {
        return  "", fmt.Errorf("读取文件出错:%v", err) // %v 占位符表示获取数据的值，在这块表示错误消息
    }
}

content, err := LoadConfig()
if err != nil {
    log.Fatal(err) // 读取文件出错:open ./config.json: The system cannot find the file specified.
}
```

2. 错误嵌套和 errors.Unwrap  
错误消息杂糅在一块导致不能分离，err1 嵌套类 err2 ，err2 也可以继续嵌套。如果想从 err1 中获取 err2 就剥一层，类似洋葱一样，一层一层往里找。那怎么实现这种嵌套关系呢，还是使用 fmt.Errorf 函数，只是使用另外一个占位符 %w ，w 的英文全名就是 wrap 。  
在打印错误时，增加了一个 errors.Unwrap 函数，该函数就是用来取出嵌套的错误，再看看输出的结果，附加的错误信息” 读取文件出错:“已经没有了。  
```go
func  LoadConfig() (string, error) {
    filename := "./config.json"
    b, err := ioutil.ReadFile(filename)
    if err != nil {
        return  "", fmt.Errorf("读取文件出错:%w", err)
    }
}

func  main() {
    content, err := LoadConfig()
    if err != nil {
        log.Fatal(errors.Unwrap(err)) // open ./config.json: The system cannot find the file specified.
    }

    fmt.Println("内容：", content)
}
```

3. 自定义错误类型  
如何应用你自定义的错误.
```go
type  FileEmptyError  struct {
    Filename string
    Err      error
}
func (e *FileEmptyError) Error() string {
    return fmt.Sprintf("%s  %v", e.Filename, e.Err)
}
// 使用 errors.Unwrap 函数，就需要实现 Wrapper 接口，fmt.Errorf 函数中的 %w 占位符底层实现好了此接口
type  Wrapper  interface {
    // Unwrap returns the next error in the error chain.
    // If there is no next error, Unwrap returns nil.
    Unwrap() error
}
func (e *FileEmptyError) Unwrap() error {
    return e.Err
}

func  LoadConfig() (string, error) {
    filename := "./config.json"
    // ...
    if  len(content) == 0 {
        return  "", &FileEmptyError{
            Filename: filename,
            Err:      errors.New("内容为空"),
        }
    }

    return content, nil
}

content, err := LoadConfig()
if err != nil {
    if  v, ok := err.(*FileEmptyError); ok {
        fmt.Println("Filename:", v.Filename)
    }
    log.Fatal(err)
}
```

**错误判断**  
对于一个函数或方法，返回的错误常常不止一个错误结果，如果对于不同的错误结果你想有不同的处理逻辑，那这个时候就要对错误结果进行判断。
```go
var (
    ErrNotFoundRequest = errors.New("404")
    ErrBadRequest = errors.New("请求异常")
)

func  GetError() error {
    // 错误 1
    return ErrNotFoundRequest
    // ...
    // 错误 2
    return ErrBadRequest
    // ...
    // 错误 3
    path := "https://printlove.com"
    return fmt.Errorf("%s:%w", path, ErrNotFoundRequest)
}

 // 最简单的就是使用”==“判断错误结果
if err == ErrNotFoundRequest {
    // 错误 1
} else  if err == ErrBadRequest {
    // 错误 2
}
// ...

// errors.Is 函数的作用就是一层层的对错误进行剥离判断，直到成功或没有嵌套的错误为止（errors.Is 函数不仅判断类型，也要判断值（错误消息））
err := GetError()
if errors.Is(err, ErrNotFoundRequest) {
    // 错误 1,错误3
} else  if errors.Is(err, ErrBadRequest) {
    // 错误 2
}

// errors.As 和 errors.Is 函数类似，但是 errors.As 函数只判断错误类型
type  ErrorString  struct {
    s string
}
func (e *ErrorString) Error() string {
    return e.s
}
func  main() {
    // targetErr 变量：无需初始化；必须是指针类型，并且实现了 error 接口；As 函数不接受 nil ，因此不能直接使用 targetErr 变量，要使用其引用 &targetErr。
    var  targetErr *ErrorString
    err := fmt.Errorf("new error:[%w]", &ErrorString{s: "target err"})
    fmt.Println(errors.As(err, &targetErr))
}
```

### Go 异常
错误的出现不会导致程序异常退出，但是异常会导致程序退出。那么，什么情况会异常退出呢？比如：下标越界、除数为 0 等。通常异常是代码问题而抛出，也可以选择主动抛出。

**panic**  
使用 panic 函数可以主动抛出异常，该函数格式如下：func  panic(v interface{})。其中，v 参数为空接口类型，那就说明可以接受任意类型数据。
- 打印出具体的异常位置，这些信息称作堆栈信息。
- 程序终止，退出码为 2。
```go
panic("我是异常")
// panic: 我是异常
// goroutine 1 [running]:
// ...
```

**defer**  
defer 不是函数，只是一个关键字。该关键字后面所跟的语句将延迟执行，在所在函数或方法正常结束时或出现异常中断时，再提前执行。
- defer 关键字后面跟了一个匿名函数调用，有名字的函数当然也可以。
- 遇到 defer 关键字，后面的语句会延迟执行。
- panic 抛出异常，因此在退出前先执行 defer 语句。
- 如果调用了 os.Exit 函数，defer 后的语句将不会执行。
- 如果在一个函数或方法中出现了多个 defer 语句，那会采用先进后出原则，即先出现的 defer 语句后执行，后出现的先执行。
```go
package main
import  "fmt"

func  main() {
    defer func() {
        fmt.Println("defer")
    }()
    fmt.Println("main")
    panic("panic")
}
// main
// defer
// panic: panic
```

**处理异常**  
不管是主动抛出异常，还是程序的 bug 被动抛出异常，这些都是很严重的问题，因为它会导致程序异常退出。在其它语言中，通过 try/catch 机制可以捕捉异常来保证程序的正常运行，那在 Go 语言中使用 recover 函数捕捉异常。  
在程序出现异常之前 defer 语句先被执行，因此在 defer 后的语句就可以提前拦截异常。  
```go
package main
import  "fmt"

func  main() {
    defer  func() {
        // recover 返回异常信息
        if  err := recover(); err != nil {
            fmt.Println("我捕捉的：", err)
            fmt.Println("我好好的")
        }
    }()
    panic("我是异常")
}
// 我捕捉的： 我是异常
// 我好好的
```
不管是主动抛出异常，还是被动的，recover 函数都能捕捉，这样以保证程序的正常进行。  
假如函数 A 调用了很多函数，这些函数又调用了很多，只要下面被调用的函数出现异常，函数 A 中的 recover 函数将都可以捕捉到。但其中还是有个特例， Goroutine 中出现的异常是无法被捕捉到的，必须在新的 Goroutine 中重新使用 recover 函数。
