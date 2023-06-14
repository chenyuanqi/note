
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

### 注意 defer 的调用时机
有时候我们会像下面一样使用 defer 去关闭一些资源：
```go
func readFiles(ch <-chan string) error {
	for path := range ch {
		file, err := os.Open(path)
		if err != nil {
				return err
		}

		defer file.Close()

		// Do something with file
	}
	return nil
}
```
因为 defer 会在方法结束的时候调用，但是如果上面的 readFiles 函数永远没有 return，那么 defer 将永远不会被调用，从而造成内存泄露。并且 defer 写在 for 循环里面，编译器也无法做优化，会影响代码执行性能。  
为了避免这种情况，我们可以 wrap 一层：
```go
func readFiles(ch <-chan string) error {
	for path := range ch { 
		if err := readFile(path); err != nil {
				return err
		} 
	}
	return nil
} 

func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
			return err
	}

	defer file.Close()

	// Do something with file
	return nil
}
```

### 注意 defer 的参数
defer 声明时会先计算确定参数的值。
```go
func a() {
	i := 0
	defer notice(i) // 0
	i++
	return
}

func notice(i int) {
	fmt.Println(i)
}
```
在这个例子中，变量 i 在 defer 被调用的时候就已经确定了，而不是在 defer执行的时候，所以上面的语句输出的是 0。所以我们想要获取这个变量的真实值，应该用引用：
```go
func a() {
	i := 0
	defer notice(&i) // 1
	i++
	return
}
```

### defer 下的闭包
```go
func a() int {
	i := 0
	defer func() {
		fmt.Println(i + 1) //12
	}()
	i++
	return i+10  
}

func TestA(t *testing.T) {
	fmt.Println(a()) //11
}
```
如果换成闭包的话，实际上闭包中对变量 i 是通过指针传递的，所以可以读到真实的值。但是上面的例子中 a 函数返回的是 11 是因为执行顺序是：  
`先计算（i+10）-> (call defer) -> (return)`

### string 迭代带来的问题
在 Go 语言中，字符串是一种基本类型，**默认是通过 utf8 编码的字符序列**，当字符为 ASCII 码时则占用 1 个字节，其他字符根据需要占用 2-4 个字节，比如中文编码通常需要 3 个字节。  
那么我们在做 string 迭代的时候可能会产生意想不到的问题：
```go
s := "hêllo"
for i := range s {
	fmt.Printf("position %d: %c\n", i, s[i])
}
fmt.Printf("len=%d\n", len(s))
// position 0: h
// position 1: Ã
// position 3: l
// position 4: l
// position 5: o
// len=6
```
上面的输出中发现第二个字符是 Ã，不是 ê，并且位置2的输出”消失“了，这其实就是因为 ê 在 utf8 里面实际上占用 2 个 byte，而我们在迭代的时候是按照 byte 来迭代的，所以导致了这个问题。  
根据上面的分析，我们就可以知道在迭代获取字符的时候不能只获取单个 byte，应该使用 range 返回的 value 值：
```go
s := "hêllo"
for i, v := range s {
	fmt.Printf("position %d: %c\n", i, v)  
}
// position 0: h
// position 1: ê
// position 3: l
// position 4: l
// position 5: o
```
或者我们可以把 string 转成 rune 数组，在 go 中 rune 代表 Unicode码位，用它可以输出单个字符：
```go
s := "hêllo"
runes := []rune(s)
for i, _ := range runes {
	fmt.Printf("position %d: %c\n", i, runes[i])  
}
// position 0: h
// position 1: ê
// position 2: l
// position 3: l
// position 4: o
```

### string 截断带来的问题
**在对slice使用 ：操作符进行截断的时候，底层的数组实际上指向同一个**，在 string 里面也需要注意这个问题，比如下面：
```go
func (s store) handleLog(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	uuid := log[:36]
	s.store(uuid)
	// Do something    
 }
```
这段代码用了 ：操作符进行截断，但是如果 log 这个对象很大，比如上面的 store 方法把 uuid 一直存在内存里，可能会造成底层的数组一直不释放，从而造成内存泄露。  
为了解决这个问题，我们可以先复制一份再处理： 
```go
func (s store) handleLog(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	uuid := strings.Clone(log[:36]) // copy一份
	s.store(uuid)
	// Do something    
}
```

### interface 类型返回的非 nil 问题
假如我们想要继承 error 接口实现一个自己的 MultiError：
```go
type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}
```
然后在使用的时候返回 error，并且想通过 error 是否为 nil 判断是否有错误：
```go
func Validate(age int, name string) error {
	var m *MultiError
	if age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}

	return m
}

func Test(t *testing.T) {
	if err := Validate(10, "a"); err != nil {
		t.Errorf("invalid")
	}
}
```
实际上 Validate 返回的 err 会总是为非 nil 的，也就是上面代码只会输出 invalid。
这是因为在 Go 语言中，interface 类型的变量在赋值的时候，如果没有赋值，那么它的值就是 nil，但是它的类型是有值的，所以在上面的代码中，m 的类型是 *MultiError，所以它的值就不是 nil。
解决这个问题的方法是在返回的时候判断 errs 是否为空：
```go
if len(m.errs) == 0 {
	return nil
}
```

### error wrap
对于 err 的 return 我们一般可以这么处理：
```go
err := xxx()
if err != nil {
	return err 
}
```
但是这样处理只是简单地将原始的错误抛出去了，无法知道当前处理的这段程序的上下文信息，这个时候我们可能会自定义个 error 结构体，继承 error 接口：
```go
type MyError struct {
	Err error
	Msg string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s: %s", e.Msg, e.Err.Error())
}
```
然后在处理的时候：
```go
err := xxx()
if err != nil {
	return &MyError{Err: err, Msg: "xxx"}
}
```
但是这样虽然可以添加一些上下文信息，但是每次都需要创建一个特定类型的 error 类会变得很麻烦，那么在 1.13 之后，我们可以使用 %w 进行 wrap。
```go
err := xxx()
if err != nil {
	return fmt.Errorf("xxx: %w", err)
}
```
当然除了上面这种做法以外，我们还可以直接 %v 直接格式化我们的错误信息：
```go
err := xxx()
if err != nil {
	return fmt.Errorf("xxx: %v", err)
}
```
这样做的缺点就是我们会丢失这个 err 的类型信息，如果不需要这个类型信息，只是想往上抛打印一些日志当然也无所谓。

### error Is & As
因为我们的 error 可以会被 wrap 好几层，那么使用 == 是可能无法判断我们的 error 究竟是不是我们想要的特定的 error，那么可以用 errors.Is：
```go
var BaseErr = errors.New("base error")

func main() {
	err1 := fmt.Errorf("wrap base: %w", BaseErr)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	println(err2 == BaseErr)

	if !errors.Is(err2, BaseErr) {
		panic("err2 is not BaseErr")
	}
	println("err2 is BaseErr")
}
// false
// err2 is BaseErr
```
在上面，我们通过 errors.Is 就可以判断出 err2 里面包含了 BaseErr 错误。errors.Is 里面会递归调用 Unwrap 方法拆包装，然后挨个使用 == 判断是否和指定类型的 error 相等。  
errors.As 主要用来做类型判断，原因也是和上面一样，error 被 wrap 之后我们通过 err.(type) 无法直接判断，errors.As 会用 Unwrap 方法拆包装，然后挨个判断类型。使用如下：
```go
type TypicalErr struct {
	e string
}

func (t TypicalErr) Error() string {
	return t.e
}

func main() {
	err := TypicalErr{"typical error"}
	err1 := fmt.Errorf("wrap err: %w", err)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	var e TypicalErr
	if !errors.As(err2, &e) {
		panic("TypicalErr is not on the chain of err2")
	}
	println("TypicalErr is on the chain of err2")
	println(err == e)
}
// TypicalErr is on the chain of err2
// true
```
上面的代码中，我们通过 errors.As 将 err2 里面的 TypicalErr 取出来了，然后和 e 进行了比较，发现是相等的。

### 处理 defer 中的 error
我们如果在调用 Close 的时候报错是没有处理的：
```go
func getBalance(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	// Use rows    
}
```
那么也许我们可以在 defer 中打印一些 log，但是无法 return，defer 不接受一个 err 类型的返回值：
```go
defer func() {
	err := rows.Close()
	if err != nil {
		log.Printf("failed to close rows: %v", err)
	}
	return err //无法通过编译    
}()
```
那么我们可能想通过默认 err 返回值的方式将 defer 的 error 也返回了：
```go
func getBalance(db *sql.DB, clientID string) (balance float32, err error) {
	rows, err = db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer func() {
		err = rows.Close()
	}()

	// Use rows    
}
```
上面代码看起来没问题，那么假如 Query 的时候和 Close 的时候同时发生异常呢？其中有一个 error 会被覆盖，那么我们可以根据自己的需求选择一个打印日志，另一个 error 返回：
```go
defer func() {
	closeErr := rows.Close()
	if err != nil {
		if closeErr != nil {
			log.Printf("failed to close rows: %v", err)
		}
		return
	}
	err = closeErr    
}()
```

### 错误使用 sync.WaitGroup
sync.WaitGroup 通常用在并发中等待 goroutines 任务完成，用 Add 方法添加计数器，当任务完成后需要调用 Done 方法让计数器减一。等待的线程会调用 Wait 方法等待，直到 sync.WaitGroup 内计数器为零。
需要注意的是 Add 方法是怎么使用的，如下：
```go
wg := sync.WaitGroup{}
var v uint64

for i := 0; i < 3; i++ {
	go func() {
		wg.Add(1)
		atomic.AddUint64(&v, 1)
		wg.Done()
	}()
}

wg.Wait()    
fmt.Println(v)
```
这样使用可能会导致 v 不一定等于3，因为在 for 循环里面创建的 3 个 goroutines 不一定比外面的主线程先执行，从而导致在调用 Add 方法之前可能 Wait 方法就执行了，并且恰好 sync.WaitGroup 里面计数器是零，然后就通过了。
正确的做法应该是在创建 goroutines 之前就将要创建多少个 goroutines 通过 Add 方法添加进去。
```go
wg := sync.WaitGroup{}
var v uint64

wg.Add(3)
for i := 0; i < 3; i++ {
	go func() {
		atomic.AddUint64(&v, 1)
		wg.Done()
	}()
}

wg.Wait()
fmt.Println(v)
```

### 不要拷贝 sync 类型
sync 包里面提供一些并发操作的类型，如 mutex、condition、wait gorup 等等，这些类型都不应该被拷贝之后使用。  
有时候我们在使用的时候拷贝是很隐秘的，比如下面：
```go
type Counter struct {
	mu sync.Mutex
	counters map[string]int
}

func (c Counter) Increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func NewCounter() Counter {
  	return Counter{counters: map[string]int{}}
}

func main() {
	counter := NewCounter()
	go counter.Increment("aa")
	go counter.Increment("bb")
}
```
receiver 是一个值类型，所以调用 Increment 方法的时候实际上拷贝了一份 Counter 里面的变量。这里我们可以将 receiver 改成一个指针，或者将 sync.Mutex 变量改成指针类型。
所以如果：
`receiver 是值类型； 函数参数是 sync 包类型； 函数参数的结构体里面包含了 sync 包类型；`  
遇到这种情况需要注意检查一下，我们可以借用 go vet 来检测。

### time.After 内存泄露
我们用一个简单的例子模拟一下：
```go
package main

import (
	"fmt"
	"time"
)

//define a channel
var chs chan int

func Get() {
	for {
		select {
		case v := <-chs:
			fmt.Printf("print:%v\n", v)
		case <-time.After(3 * time.Minute):
			fmt.Printf("time.After:%v", time.Now().Unix())
		}
	}
}

func Put() {
	var i = 0
	for {
		i++
		chs <- i
	}
}

func main() {
	chs = make(chan int, 100)
	go Put()
	Get()
}
```
逻辑很简单就是先往 channel 里面存数据，然后不停地使用 for select case 语法从 channel 里面取数据，为了防止长时间取不到数据，所以在上面加了 time.After 定时器，这里只是简单打印一下。  
发现不一会儿 Timer 的内存占用很高了。这是因为在计时器触发之前，垃圾收集器不会回收 Timer，但是在循环里面每次都调用 time.After都会实例化一个一个新的定时器，并且这个定时器会在激活之后才会被清除。  
为了避免这种情况我们可以使用下面代码：  
```go
func Get() {
	delay := time.NewTimer(3 * time.Minute)

	defer delay.Stop()

	for {
		delay.Reset(3 * time.Minute)

		select {
		case v := <-chs:
			fmt.Printf("print:%v\n", v)
		case <-delay.C:
			fmt.Printf("time.After:%v", time.Now().Unix())
		}
	}
}
```

### HTTP body 忘记 Close 导致的泄露
```go
type handler struct {
	client http.Client
	url    string
}

func (h handler) getBody() (string, error) {
	resp, err := h.client.Get(h.url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
```
上面这段代码看起来没什么问题，但是 resp 是 *http.Response 类型，里面包含了 Body io.ReadCloser 对象，它是一个 io 类，必须要正确关闭，否则是会产生资源泄露的。一般我们可以这么做：
```go
defer func() {
	err := resp.Body.Close()
	if err != nil {
		log.Printf("failed to close response: %v\n", err)
	}
}()
```

### byte slice 和 string 的转换优化
直接通过强转 string(bytes) 或者 []byte(str) 会带来数据的复制，性能不佳，所以在追求极致性能场景使用 unsafe 包的方式直接进行转换来提升性能：
```go
// toBytes performs unholy acts to avoid allocations
func toBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// toString performs unholy acts to avoid allocations
func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
```
在 Go 1.12 中，增加了几个方法 String、StringData、Slice 和 SliceData ,用来做这种性能转换。
```go
// String returns the string representation of the underlying data.
func (b *Builder) String() string {
	return *(*string)(unsafe.Pointer(&b.buf))
}

// StringData returns a string that shares the underlying data buffer with the builder.
// If the builder is not empty, the result is undefined.
func (b *Builder) StringData() string {
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: b.buf.Data,
		Len:  b.buf.Len,
	}))
}

// Slice returns the contents of the builder's buffer.
func (b *Builder) Slice() []byte {
	return b.buf.Bytes()
}

// SliceData returns a byte slice that shares the underlying data buffer with the builder.
// If the builder is not empty, the result is undefined.
func (b *Builder) SliceData() []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: b.buf.Data,
		Len:  b.buf.Len,
		Cap:  b.buf.Cap,
	}))
}
```

### 容器中的 GOMAXPROCS
自 Go 1.5 开始， Go 的 GOMAXPROCS 默认值已经设置为 CPU 的核数，但是在 Docker 或 k8s 容器中 runtime.GOMAXPROCS() 获取的是宿主机的 CPU 核数 。这样会导致 P 值设置过大，导致生成线程过多，会增加上下文切换的负担，导致严重的上下文切换，浪费 CPU。  
所以可以使用 uber 的 automaxprocs 库，大致原理是读取 CGroup 值识别容器的 CPU quota，计算得到实际核心数，并自动设置 GOMAXPROCS 线程数量。
```go
package main

import _ "go.uber.org/automaxprocs"

func main() {
  // Your application logic here
}
```

### 逃逸分析
逃逸分析是指编译器分析程序的代码，确定程序中的对象是否在函数栈上分配还是在堆上分配。如果对象被分配到函数栈上，它就不需要进行垃圾回收，因为当函数返回时，栈上的内存自动被释放。而如果对象被分配到堆上，它需要进行垃圾回收，这会对程序的性能产生影响。

逃逸分析可以帮助编译器优化程序的性能，例如将一些对象从堆上移动到栈上，从而减少垃圾回收的开销。同时，逃逸分析还可以帮助开发者更好地理解程序的内存使用情况，从而提高程序的可维护性。  

在 Go 中，编译器会自动进行逃逸分析，以优化程序的性能。  
Go 是通过在编译器里做逃逸分析（escape analysis）来决定一个对象放栈上还是放堆上，不逃逸的对象放栈上，可能逃逸的放堆上。对于 Go 来说，我们可以通过下面指令来看变量是否逃逸：
```go
go build -gcflags "-m -l" main.go
```
`-m 会打印出逃逸分析的优化策略，实际上最多总共可以用 4 个 -m，但是信息量较大，一般用 1 个就可以了。 -l 会禁用函数内联，在这里禁用掉内联能更好的观察逃逸情况，减少干扰。`

Go 语言的编译器会对代码进行逃逸分析，如果发现变量在函数内部定义，但是在函数外部被引用，那么这个变量就会被分配在堆上，而不是栈上。  
比如下面这段代码：
```go
func main() {
	var s string
	s = "hello world"
	fmt.Println(s)
}
```
这里的 s 变量是在 main 函数内部定义的，但是在函数外部被引用，所以会被分配在堆上。

**指针逃逸**  
指针逃逸是指指针指向的变量逃逸到堆上。  
在函数中创建了一个对象，返回了这个对象的指针。这种情况下，函数虽然退出了，但是因为指针的存在，对象的内存不能随着函数结束而回收，因此只能分配在堆上。
```go
type Demo struct {
	name string
}

func createDemo(name string) *Demo {
	d := new(Demo) // 局部变量 d 逃逸到堆
	d.name = name
	return d
}

func main() {
	demo := createDemo("demo")
	fmt.Println(demo)
}

// go run -gcflags '-m -l'  .\main\main.go
// # command-line-arguments
// main\main.go:12:17: leaking param: name
// main\main.go:13:10: new(Demo) escapes to heap
// main\main.go:20:13: ... argument does not escape&{demo}
```

**interface{}/any 动态类型逃逸**  
因为编译期间很难确定其参数的具体类型，也会发生逃逸，例如这样：
```go
func createDemo(name string) any {
	d := new(Demo) // 局部变量 d 逃逸到堆
	d.name = name
	return d
}
```

**切片长度或容量没指定逃逸**  
如果使用局部切片时，已知切片的长度或容量，请使用常量或数值字面量来定义，否则也会逃逸：
```go
func main() {
	number := 10
	s1 := make([]int, 0, number)
	for i := 0; i < number; i++ {
		s1 = append(s1, i)
	}

	s2 := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		s2 = append(s2, i)
	}
}

// go run -gcflags '-m -l'  .\main\main.go
// ./main.go:65:12: make([]int, 0, number) escapes to heap
// ./main.go:69:12: make([]int, 0, 10) does not escape
```

**闭包引用逃逸**  
Increase() 返回值是一个闭包函数，该闭包函数访问了外部变量 n，那变量 n 将会一直存在，直到 in 被销毁。很显然，变量 n 占用的内存不能随着函数 Increase() 的退出而回收，因此将会逃逸到堆上。
```go
func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	in := Increase()
	fmt.Println(in()) // 1
	fmt.Println(in()) // 2
}

// go run -gcflags '-m -l'  main.go  
//  
// ./main.go:64:5: moved to heap: n
// ./main.go:65:12: func literal escapes to heap
```




