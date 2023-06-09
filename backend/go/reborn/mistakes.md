
## 常见典型错误

### 注意 shadow 变量
这段代码中，声明了一个 client 变量，然后使用 tracing 控制变量的初始化，可能是因为没有声明 err 的缘故，使用的是 := 进行初始化，那么会导致外层的 client 变量永远是 nil。这个例子实际上是很容易发生在我们实际的开发中，尤其需要注意。
```go
var client *http.Client
if tracing {
	client, err := createClientWithTracing()
	if err != nil {
		return err
	}
	log.Println(client)
} else {
	client, err := createDefaultClient()
	if err != nil {
		return err
	}
	log.Println(client)
}
```
如果是因为 err 没有初始化的缘故，我们在初始化的时候可以这么做：
```go
var client *http.Client
var err error
if tracing {
    client, err = createClientWithTracing()
    if err != nil {
        return err
    }
    log.Println(client)
} else {
    client, err = createDefaultClient()
    if err != nil {
        return err
    }
    log.Println(client)
}
```
或者内层的变量声明换一个变量名字，这样就不容易出错了。

###  注意 init 函数执行顺序
1、init 函数会在全局变量之后被执行  
init 函数并不是最先被执行的，如果声明了 const 或全局变量，那么 init 函数会在它们之后执行。  

2、init 初始化按解析的依赖关系顺序执行  
如果有多个 init 函数，那么它们的执行顺序是按照它们的依赖关系进行的，比如下面这个例子，init 函数的执行顺序是 a -> b -> c -> main。  
```go
package main

import "fmt"

func init() {
    fmt.Println("a")
}

func init() {
    fmt.Println("b")
}

func init() {
    fmt.Println("c")
}

func main() {
    fmt.Println("main")
}
```
又比如 main 包里面有 init 函数，依赖了 redis 包，main 函数执行了 redis 包的 Store 函数，恰好 redis 包里面也有 init 函数，那么执行顺序会是：redis 包的 init 函数 -> main 包的 init 函数 -> redis 包的 Store 函数 -> main 包的 main 函数。  

3、扰乱单元测试  
比如我们在 init 函数中初始了一个全局的变量，但是单测中并不需要，那么实际上会增加单测得复杂度，比如：  
```go
var db *sql.DB

func init() {
	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME")
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	db = d
}
```
 init 函数初始化了一个 db 全局变量，那么在单测的时候也会初始化一个这样的变量，但是很多单测其实是很简单的，并不需要依赖这个东西。  

 ### embed types 优缺点
 embed types 指的是我们在 struct 里面定义的匿名的字段，如：  
```go
type User struct {
    Name string
    Age  int
}

type Student struct {
    User
    Class string
}
```
这里的 User 就是一个 embed type，它的优点是可以直接访问 User 的字段，比如：  
```go
s := Student{
    User: User{
        Name: "张三",
        Age:  18,
    },
    Class: "一年级",
}
fmt.Println(s.Name)
```
这里的 s.Name 其实是 s.User.Name，这样就可以直接访问 User 的字段，而不需要通过 s.User.Name 这样的方式访问。  

在很多时候可以增加我们使用的便捷性，如果没有使用 embed types 那么可能需要很多代码，如下：  
```go
type Logger struct {
	writeCloser io.WriteCloser
}

func (l Logger) Write(p []byte) (int, error) {
	return l.writeCloser.Write(p)
}

func (l Logger) Close() error {
	return l.writeCloser.Close()
}

func main() {
	l := Logger{writeCloser: os.Stdout}
	_, _ = l.Write([]byte("foo"))
	_ = l.Close()
}
```
如果使用了 embed types 我们的代码可以变得很简洁：
```go
type Logger struct {
	io.WriteCloser
}

func main() {
	l := Logger{WriteCloser: os.Stdout}
	_, _ = l.Write([]byte("foo"))
	_ = l.Close()
}
```
但是 embed types 也有它的缺点，有些字段我们并不想 export ，但是 embed types 可能给我们带出去，例如：
```go
type InMem struct {
	sync.Mutex
	m map[string]int
}

func New() *InMem {
	return &InMem{m: make(map[string]int)}
}
```
Mutex 一般并不想 export， 只想在 InMem 自己的函数中使用。但是这么写却可以让拿到 InMem 类型的变量都可以使用它里面的 Lock 方法。  

### Functional Options Pattern 传递参数
这种方法在很多 Go 开源库都有看到过使用，比如 zap、GRPC 等。
它经常用在需要传递和初始化校验参数列表的时候使用，比如我们现在需要初始化一个 HTTP server，里面可能包含了 port、timeout 等等信息，但是参数列表很多，不能直接写在函数上，并且我们要满足灵活配置的要求，毕竟不是每个 server 都需要很多参数。那么我们可以：  
`设置一个不导出的 struct 叫 options，用来存放配置参数； 创建一个类型 type Option func(options *options) error，用这个类型来作为返回值；`  
比如我们现在要给 HTTP server 里面设置一个 port 参数，那么我们可以这么声明一个 WithPort 函数，返回 Option 类型的闭包，当这个闭包执行的时候会将 options 的 port 填充进去：  
```go
type options struct {
	port *int
}

type Option func(options *options) error

func WithPort(port int) Option {
	// 所有的类型校验，赋值，初始化啥的都可以放到这个闭包里面做
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}
```
假如我们现在有一个这样的 Option 函数集，除了上面的 port 以外，还可以填充 timeout 等。然后我们可以利用 NewServer 创建我们的 server：
```go
func NewServer(addr string, opts ...Option) (*http.Server, error) {
	var options options
	// 遍历所有的 Option
	for _, opt := range opts {
		// 执行闭包
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}

	// 接下来可以填充我们的业务逻辑，比如这里设置默认的port 等等
	var port int
	if options.port == nil {
		port = defaultHTTPPort
	} else {
		if *options.port == 0 {
			port = randomPort()
		} else {
			port = *options.port
		}
	}

	// ...
}

// 初始化 server
server, err := NewServer("localhost", httplib.WithPort(8080), httplib.WithTimeout(time.Second))

// 这样写的话就比较灵活，如果只想生成一个简单的 server，我们的代码可以变得很简单
server, err := httplib.NewServer("localhost")
```

### 小心八进制整数  
你以为如下代码要输出 110，其实输出的是 108，因为在 Go 中以 0 开头的整数表示八进制。  
```go
sum := 100 + 010  
fmt.Println(sum)
```

经常用在处理 Linux 权限相关的代码上，如下面打开一个文件：
```go
f, err := os.OpenFile("foo.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
```
0666 表示的是文件的权限，这里的 6 是八进制，转换成十进制就是 110，表示的是文件的权限是 110，也就是 6，这里的 6 表示的是读写权限，如果是 7 表示的是读写执行权限。  
为了可读性，我们在用八进制的时候最好使用 "0o" 的方式表示，比如上面这段代码可以表示为：  
```go
f, err := os.OpenFile("foo.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
```

### float 的精度问题
在 Go 中 float 的精度问题是一个很大的坑，比如下面的代码，你以为会输出 0.1，其实输出的是 0.10000000000000002。  
```go
fmt.Println(0.1 + 0.2)
```
这是因为在计算机中，浮点数是以二进制的形式存储的，而二进制是无法精确表示 0.1 的，所以会有精度问题。

在 Go 中浮点数表示方式和其他语言一样，都是通过科学计数法表示，float 在存储中分为三部分：  
符号位（Sign）: 0代表正，1代表为负   
指数位（Exponent）:用于存储科学计数法中的指数数据，并且采用移位存储   
尾数部分（Mantissa）：尾数部分  

```go
a := 100000.001
b := 1.0001
c := 1.0002

fmt.Println(a * (b + c)) // 200030.00200030004
fmt.Println(a*b + a*c) // 200030.0020003
```
正确输出应该是 200030.0020003，所以它们实际上都有一定的误差，但是可以看到先乘再加精度丢失会更小。  
如果想要准确计算浮点的话，可以尝试 "github.com/shopspring/decimal" 库，换成这个库我们再来计算一下：
```go
a := decimal.NewFromFloat(100000.001)
b := decimal.NewFromFloat(1.0001)
c := decimal.NewFromFloat(1.0002)

fmt.Println(a.Mul(b.Add(c))) //200030.0020003
```



