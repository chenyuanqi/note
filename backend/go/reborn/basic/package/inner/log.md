
### log package：日志包
在日常开发中，日志是必不可少的功能。虽然有时可以用 fmt 库输出一些信息，但是灵活性不够。Go 标准库提供了一个日志库 log。日志默认输出路径为临时路径，可通过执行命令时带上 -log_dir="路径"，指定输出，但路径必须已存在。
```go
package main
import (
        "log"
        "os"
        "time"
)
func init() {
        file := "./" +"message"+ ".txt"
        logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
        if err != nil {
                panic(err)
        }
        log.SetOutput(logFile) // 将文件设置为log输出的文件
        log.SetPrefix("[qSkipTool]")
        log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

        // 或者
        // loger = log.New(logFile, "[qSkiptool]",log.LstdFlags | log.Lshortfile | log.LUTC) // 将文件设置为loger作为输出
        return
}

func main() {
        log.Println("Hello Davis!") // log 还是可以作为输出的前缀
        return
}
```
[更多参考](https://pkg.go.dev/log)  

log 默认输出到标准错误（stderr），每条日志前会自动加上日期和时间。如果日志不是以换行符结尾的，那么 log 会自动加上换行符。即每条日志会在新行中输出。  
log 提供了三组函数（带 f 后缀的有格式化功能，带 ln 后缀的会在日志后增加一个换行符）：
- Print/Printf/Println：正常输出日志；
- Panic/Panicf/Panicln：输出日志后，以拼装好的字符串为参数调用 panic；
- Fatal/Fatalf/Fatalln：输出日志后，调用 os.Exit(1) 退出程序。

```go
package main

import (
  "log"
)

type User struct {
  Name string
  Age  int
}

func main() {
  u := User{
    Name: "dj",
    Age:  18,
  }

  log.Printf("%s login, age:%d", u.Name, u.Age)
  log.Panicf("Oh, system error when %s login", u.Name)
  // 由于调用 log.Panicf 会 panic，所以 log.Fatalf 并不会调用
  log.Fatalf("Danger! hacker %s login", u.Name)
}
```

log 库的核心是 Output 方法。如果设置了 Lshortfile 或 Llongfile，Ouput 方法中会调用 runtime.Caller 获取文件名和行号。runtime.Caller 的参数 calldepth 表示获取调用栈向上多少层的信息，当前层为 0。  
一般的调用路径是：
- 程序中使用 log.Printf 之类的函数；
- 在 log.Printf 内调用 std.Output。

```go
// src/log/log.go
func (l *Logger) Output(calldepth int, s string) error {
  now := time.Now() // get this early.
  var file string
  var line int
  l.mu.Lock()
  defer l.mu.Unlock()
  if l.flag&(Lshortfile|Llongfile) != 0 {
    // Release lock while getting caller info - it's expensive.
    l.mu.Unlock()
    var ok bool
    _, file, line, ok = runtime.Caller(calldepth)
    if !ok {
      file = "???"
      line = 0
    }
    l.mu.Lock()
  }
  l.buf = l.buf[:0]
  l.formatHeader(&l.buf, now, file, line)
  l.buf = append(l.buf, s...)
  if len(s) == 0 || s[len(s)-1] != '\n' {
    l.buf = append(l.buf, '\n')
  }
  _, err := l.out.Write(l.buf)
  return err
}
```
两个优化技巧：
- 由于 runtime.Caller 调用比较耗时，先释放锁，避免等待时间过长；
- 为了避免频繁的内存分配，logger 中保存了一个类型为 []byte 的 buf，可重复使用。前缀和日志内容先写到这个 buf 中，然后统一写入 Writer，减少 io 操作。

**前缀**  
调用 log.SetPrefix 为每条日志文本前增加一个前缀。调用 log.Prefix 可以获取当前设置的前缀。  
```go
log.SetPrefix("xxx: ")

fmt.Println(log.Prefix())
```

**选项**  
设置选项可在每条输出的文本前增加一些额外信息，如日期时间、文件名等。
- Ldate：输出当地时区的日期，如 2020/02/07；
- Ltime：输出当地时区的时间，如 11:45:45；
- Lmicroseconds：输出的时间精确到微秒，设置了该选项就不用设置 Ltime 了。如 11:45:45.123123；
- Llongfile：输出长文件名 + 行号，含包名，如 github.com/darjun/go-daily-lib/log/flag/main.go:50；
- Lshortfile：输出短文件名 + 行号，不含包名，如 main.go:50；
- LUTC：如果设置了 Ldate 或 Ltime，将输出 UTC 时间，而非当地时区。

调用 log.SetFlag 设置选项，可以一次设置多个。调用 log.Flags() 可以获取当前设置的选项。  
```go
log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

fmt.Println(log.Flags())
```
log 库还定义了一个 Lstdflag，为 Ldate | Ltime，这就是我们默认的选项。log 库还定义了一个 Lstdflag，为 Ldate | Ltime，这就是我们默认的选项。
`调用 log.SetFlag 之后，原有的选项会被覆盖掉。`  

**自定义**  
log 库为我们定义了一个默认的 Logger，名为 std，意为标准日志。我们直接调用的 log 库的方法，其内部是调用 std 的对应方法。  
```go
// src/log/log.go
var std = New(os.Stderr, "", LstdFlags)

func Printf(format string, v ...interface{}) {
  std.Output(2, fmt.Sprintf(format, v...))
}

func Fatalf(format string, v ...interface{}) {
  std.Output(2, fmt.Sprintf(format, v...))
  os.Exit(1)
}

func Panicf(format string, v ...interface{}) {
  s := fmt.Sprintf(format, v...)
  std.Output(2, s)
  panic(s)
}
```

我们也可以定义自己的 Logger，log.New 接受三个参数：
- io.Writer：日志都会写到这个 Writer 中；
- prefix：前缀，也可以后面调用 logger.SetPrefix 设置；
- flag：选项，也可以后面调用 logger.SetFlag 设置。

```go
package main

import (
  "bytes"
  "fmt"
  "log"
)

type User struct {
  Name string
  Age  int
}

func main() {
  u := User{
    Name: "dj",
    Age:  18,
  }

  buf := &bytes.Buffer{}
  logger := log.New(buf, "", log.Lshortfile|log.LstdFlags)

  logger.Printf("%s login, age:%d", u.Name, u.Age)

  fmt.Print(buf.String())

  // 可以使用 io.MultiWriter 实现多目的地输出
  writer1 := &bytes.Buffer{}
  writer2 := os.Stdout
  writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
  if err != nil {
    log.Fatalf("create file log.txt failed: %v", err)
  }

  logger := log.New(io.MultiWriter(writer1, writer2, writer3), "", log.Lshortfile|log.LstdFlags)
  logger.Printf("%s login, age:%d", u.Name, u.Age)
}
```

**重定向运行时 panic 到文件**  
```go
package main

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
)

const stdErrFile = "/tmp/go-app1-stderr.log"

var stdErrFileHandler *os.File

func RewriteStderrFile() error {
	if runtime.GOOS == "windows" {
		return nil
	}

	file, err := os.OpenFile(stdErrFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	stdErrFileHandler = file //把文件句柄保存到全局变量，避免被GC回收

	if err = syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd())); err != nil {
		fmt.Println(err)
		return err
	}
	// 内存回收前关闭文件描述符
	runtime.SetFinalizer(stdErrFileHandler, func(fd *os.File) {
		fd.Close()
	})

	return nil
}

func testPanic() {
	panic("test panic")
}

func main() {
	RewriteStderrFile()
	testPanic()
}
```
