
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

### 区分 slice 的 length 和 capacity
初始化一个带有 length 和 capacity 的 slice ：
```go
s := make([]int, 3, 6)
```
slice 的底层实际上指向了一个数组。当然，由于我们的 length 是 3，所以这样设置 s[4] = 0 会 panic 的。需要使用 append 才能添加新元素。  
当 appned 超过 cap 大小的时候，slice 会自动帮我们扩容，`在元素数量小于 1024 的时候每次会扩大一倍，当超过了 1024 个元素每次扩大 25% `。  

使用操作符从另一个 slice 上面创建一个新切片：
```go
s := []int{1, 2, 3, 4, 5}
s1 := s[1:3]
fmt.Println(s1) // [2 3]
fmt.Println(len(s1), cap(s1)) // 2 4
```
实际上这两个 slice 还是指向了底层同样的数组，只不过 s1 的 length 和 capacity 分别是 2 和 4。  
由于指向了同一个数组，那么当我们改变第一个槽位的时候，比如 s1[1]=100，实际上两个 slice 的数据都会发生改变。  
```go	
s1[1] = 100
fmt.Println(s) // [1 100 3 4 5]
fmt.Println(s1) // [2 100]
fmt.Println(len(s1), cap(s1)) // 2 4
```
但是当我们使用 append 的时候情况会有所不同：  
```go
s1 = append(s1, 200)
fmt.Println(s) // [1 100 3 4 5]
fmt.Println(s1) // [2 100 200]
fmt.Println(len(s1), cap(s1)) // 3 4
```
再继续 append s1 直到 s1 发生扩容，这个时候会发现 s1 实际上和 s 指向的不是同一个数组了。

###  slice 初始化
对于 slice 的初始化实际上有很多种方式：
```go
func main() {
    // 好处就是不用做任何的内存分配
	var s []string
	log(1, s) // 1: empty=true	nil=true

    // 较少使用，可能用它来进行 slice 的 copy：src := []int{0, 1, 2}; dst := append([]int(nil), src...)
	s = []string(nil)
	log(2, s)  // 2: empty=true	nil=true

    // 适合初始化一个已知元素的 slice
	s = []string{}
	log(3, s) // 3: empty=true	nil=false

    // 初始化 slice 的 length 和 capacity，如果我们能确定 slice 里面会存放多少元素，从性能的角度考虑最好使用 make 初始化好，因为对于一个空的 slice append 元素进去每次达到阈值都需要进行扩容
	s = make([]string, 0)
	log(4, s) // 4: empty=true	nil=false
}

func log(i int, s []string) {
    // 最好是使用 len(xxx) == 0来判断 slice 是不是空的
    fmt.Printf("%d: empty=%t\tnil=%t\n", i, len(s) == 0, s == nil)}
}
```
前两种方式会创建一个 nil 的 slice，后两种会进行初始化，并且这些 slice 的大小都为 0 。

### slice 的 copy
slice 的 copy 是浅拷贝，也就是说当我们拷贝一个 slice 的时候，实际上只是拷贝了 slice 的指针，而不是 slice 的内容。  
```go
func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s2[0] = 100
	fmt.Println(s1) // [100 2 3 4 5]
	fmt.Println(s2) // [100 2 3 4 5]
}
```

使用 copy 函数 copy slice 的时候需要注意，下面这种情况实际上会 copy 失败，因为对 slice 来说是由 length 来控制可用数据，copy 并没有复制这个字段。
```go
src := []int{0, 1, 2}
var dst []int
copy(dst, src)
fmt.Println(dst) // []
```
要想 copy 我们可以这么做：
```go
src := []int{0, 1, 2}
dst := make([]int, len(src))
copy(dst, src)
fmt.Println(dst) //[0 1 2]

// 或者这样
src := []int{0, 1, 2}
dst := append([]int(nil), src...)
```
而且，copy 函数只会复制两个 slice 中较小的那个的元素个数。如果 dst 的长度小于 src 的长度，只会复制 dst 中的元素；如果 dst 的长度大于 src 的长度，则只会复制 src 中的元素，多余的 dst 中的元素不会被修改。

### range 的 copy
range 会返回两个值，第一个是 index，第二个是 value。  
使用 range 的时候如果我们直接修改它返回的数据会不生效，因为返回的数据并不是原始数据：
```go
type account struct {
  balance float32
}

accounts := []account{
	{balance: 100.},
	{balance: 200.},
	{balance: 300.},
}
for _, a := range accounts {
	a.balance += 1000  
}
fmt.Printf("%+v\n", accounts) // [{balance:100} {balance:200} {balance:300}]
```
我们想要改变 range 中的数据可以这么做：
```go
for i := range accounts {
	accounts[i].balance += 1000
}
fmt.Printf("%+v\n", accounts) // [{balance:1100} {balance:1200} {balance:1300}]
```

range slice 的话也会 copy 一份：
```go
s := []int{0, 1, 2}
// 在 range 的时候会 copy 一份，因此只会调用 3 次 append 后停止
for range s {
    s = append(s, 10) 
	fmt.Printf("%+v", s) // [0 1 2 10] [0 1 2 10 10] [0 1 2 10 10 10]
}
```

### 注意指针的问题
比方我们想要 range slice 并将返回值存到 map 里面供后面业务使用，类似这样：
```go
type Customer struct {
    ID string
    Balance float64
}

test := []Customer{
      {ID: "1", Balance: 10},
      {ID: "2", Balance: -10},
      {ID: "3", Balance: 0},
} 

var m map[string]*Customer
for _, customer := range test {
    m[customer.ID] = &customer
}
```
这样做是不对的，因为 range 会 copy 一份，而我们存到 map 里面的是指针，这样的话 map 里面存的都是同一个指针，最后 map 里面的值都是最后一个元素的值。就是说当我们使用 range 遍历 slice 的时候，返回的 customer 变量实际上是一个固定的地址，而不是每次都是一个新的地址。  
所以我们可以这样在 range 里面获取指针：  
```go
for i := range test {
	m[test[i].ID] = &test[i]
}

// 或者
for _, customer := range test {
    current := customer // 使用局部变量
    fmt.Printf("%p\n", &current) // 这里获取的指针是 range copy 出来元素的指针  
}
```

### 注意 break 作用域
下面这个代码本来想 break 停止遍历，实际上只是 break 了 switch 作用域，print 依然会打印：0，1，2，3，4。
```go
for i := 0; i < 5; i++ {
	fmt.Printf("%d ", i)

	switch i {
	default:
	case 2:
		break
	}
}
// 0 1 2 3 4
```
正确做法应该是通过 label 的方式 break：
```go
loop:
for i := 0; i < 5; i++ {
	fmt.Printf("%d ", i)

	switch i {
	default:
	case 2:
		break loop
	}
}
// 0 1 2
```

同样的，select 也是一样的，break 只会 break select 作用域，而不会 break for 作用域：
```go
loop:
for {
	select {
	case <-ch:
		break loop
	}
}
```

