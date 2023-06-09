
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

