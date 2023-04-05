

### 最佳实践
Go 箴言
- 不要通过共享内存进行通信，通过通信共享内存
- 并发不是并行
- 管道用于协调；互斥量（锁）用于同步
- 接口越大，抽象就越弱
- 利用好零值
- 空接口 interface{} 没有任何类型约束
- Gofmt 的风格不是人们最喜欢的，但 gofmt 是每个人的最爱
- 允许一点点重复比引入一点点依赖更好
- 系统调用必须始终使用构建标记进行保护
- 必须始终使用构建标记保护 Cgo
- Cgo 不是 Go
- 使用标准库的 unsafe 包，不能保证能如期运行
- 清晰比聪明更好
- 反射永远不清晰
- 错误是值
- 不要只检查错误，还要优雅地处理它们
- 设计架构，命名组件，（文档）记录细节
- 文档是供用户使用的
- 不要（在生产环境）使用 panic()

Go 之禅
- 每个 package 实现单一的目的
- 显式处理错误
- 尽早返回，而不是使用深嵌套
- 让调用者处理并发（带来的问题）
- 在启动一个 goroutine 时，需要知道何时它会停止
- 避免 package 级别的状态
- 简单很重要
- 编写测试以锁定 package API 的行为
- 如果你觉得慢，先编写 benchmark 来证明
- 适度是一种美德
- 可维护性

一、代码规范
1. 命名规范
- 变量、函数、方法使用驼峰命名法。
- 全局常量使用全大写字母加下划线的方式命名。
- 缩写尽量避免使用，如果必须使用则按照大驼峰命名法。

2. 代码格式
- 使用 gofmt 工具格式化代码。
- 使用 4 个空格代替制表符缩进。
- 使用 UTF-8 编码。

3. 注释规范
- 使用 // 或 /* */ 注释。
- 注释应该清晰简洁，避免过多无用的注释。
- 函数、方法需要编写文档注释。

二、错误处理
1. 错误处理
- 函数、方法应该返回错误信息。
- 使用 errors 包或自定义错误类型处理错误。
- 在处理错误时，应该记录日志并返回错误信息。

2. 错误日志
- 使用 log 包记录错误日志。
- 日志应该包含错误信息、堆栈信息和错误发生的时间等关键信息。

三、并发处理
1. 并发模型
- 使用 goroutine 和 channel 实现并发模型。
- 避免使用共享变量，使用 channel 传递数据。
- 使用 sync 包提供的锁实现互斥访问。

2. 并发安全
- 避免出现竞态条件。
- 使用 sync 包提供的锁实现互斥访问。
- 使用 atomic 包提供的原子操作实现线程安全。

四、性能优化
1. 内存管理
- 避免过多的内存分配和释放。
- 使用 sync.Pool 重用对象池。
- 使用标准库提供的内存分配和释放函数。

2. 并发优化
- 避免过多的 goroutine 创建和销毁。
- 使用 sync.WaitGroup 等待 goroutine 完成任务。
- 使用 sync.Mutex 避免竞态条件。

3. 数据结构优化
- 使用 map 代替 slice 实现索引访问。
- 使用 bytes.Buffer 代替字符串拼接。
- 使用 sync.Map 提供并发安全的 map。

五、其他  
- 30 * time.Second 比 time.Duration(30) * time.Second 更好
- 按类型分组 const 声明，按逻辑和/或类型分组 var
```go
// BAD
const (
    foo = 1
    bar = 2
    message = "warn message"
)

// MOSTLY BAD
const foo = 1
const bar = 2
const message = "warn message"

// GOOD
const (
    foo = 1
    bar = 2
)

const message = "warn message"
```
- 多行字符串用反引号(`)  
- 用 _ 来跳过不用的参数
- 用 range 循环来进行数组或 slice 的迭代
- 如果你要比较时间戳，请使用 time.Before 或 time.After
- 在 Go 里面要小心使用 range: for i := range a and for i, v := range &a ，都不是 a 的副本；但是 for i, v := range a 里面的就是 a 的副本
- 要 marshal 任意的 JSON， 你可以 marshal 为 map[string]interface{}{}
- **尽量减少从string转化为[]byte的过程**，因为string是只读的，改变string只能先转化为一个slice ，因为string是只读，所以说它转化为[]byte必须复制数据，以及，当字符串的量过大的时候，分配的slice就很有可能分配到堆上，所以说字符串的添加或者是各种处理都是一件非常耗费资源的事情，要尽量的减少字符串的操作。如果真的要拼接字符串，**不要使用 `+`,尽量使用bytes.Buffer**
```go
var b bytes.Buffer // A Buffer needs no initialization.
b.WriteString("a")
```
- 数字转化为字符串的时候，使用 `strconv.Iota` 效率比 `fmt.Sprintf` 高
- slice的append过程，如果cap不够，尽量自己计算好，人工分配新的cap，也不要让系统分配，因为系统很有可能直接扩展到两倍的容量，会造成内存的浪费。
- 对于在 for 循环中的的固定正则表达式，要使用 regexp.Compile 编译正则表达式
- 如果要获得更好的编码速度，就不要使用 go 的 json 包，因为这个包内使用了反射，可以使用 protonbuf 来代替，protobuf 是 google 出品，go 完美支持。`google.golang.org/protobuf` 引入此包即可
```go
import (
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// ExampleMessage 是演示用的消息类型
type ExampleMessage struct {
	Label *string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	Type  *int32  `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	Reps  []int64 `protobuf:"varint,3,rep,packed,name=reps" json:"reps,omitempty"`
}

// 编码
// 创建一个消息
message := &ExampleMessage{
    Label: proto.String("hello"),
    Type:  proto.Int32(17),
    Reps:  []int64{1, 2, 3},
}
// 将消息编码为 JSON
marshaler := &jsonpb.Marshaler{}
data, err := marshaler.MarshalToString(message)
if err != nil {
    fmt.Println("error marshaling JSON:", err)
    return
}

// 解码
// 创建一个消息的 JSON 表示
data := `{"label":"hello","type":17,"reps":[1,2,3]}`

// 将 JSON 解码为消息
var message ExampleMessage
if err := jsonpb.UnmarshalString(data, &message); err != nil {
    fmt.Println("error unmarshaling JSON:", err)
    return
}

fmt.Println(message.GetLabel(), message.GetType(), message.GetReps())
```

### 设计模式

#### 面向接口编程
在 Golang 中，接口是一种定义对象行为的抽象类型，它定义了一组方法集合，任何实现了这个方法集合的类型都可以被认为是这个接口的实现。接口的定义非常简单，只需要定义方法集合，而不需要实现这些方法。
```go
type Animal interface {
    Move() string
    Speak() string
}
```
这个接口定义了两个方法：`Move()` 和 `Speak()`。任何实现了这两个方法的类型都可以被认为是 `Animal` 接口的实现。
下面是一个实现了 `Animal` 接口的类型的例子：
```go
type Dog struct {
    Name string
}

func (d Dog) Move() string {
    return "Dog " + d.Name + " is running"
}

func (d Dog) Speak() string {
    return "Dog " + d.Name + " is barking"
}
```
在这个例子中，`Dog` 类型实现了 `Animal` 接口的两个方法，它可以被认为是 `Animal` 接口的一个实现。
使用接口编程可以让我们编写更加灵活的代码，因为我们可以使用接口来描述对象的行为，而不是关注对象的具体类型。这样，我们就可以将对象的实现与接口分离开来，从而提高代码的可维护性和可测试性。
举个例子，假设我们有一个 `Zoo` 类型，它包含了多个动物，我们可以使用接口来描述这些动物的行为：
```go
type Zoo struct {
    Animals []Animal
}

func (z Zoo) MoveAll() {
    for _, animal := range z.Animals {
        fmt.Println(animal.Move())
    }
}

func (z Zoo) SpeakAll() {
    for _, animal := range z.Animals {
        fmt.Println(animal.Speak())
    }
}
```
在这个例子中，我们可以将多个不同类型的动物放在一个 `Zoo` 中，然后通过接口来调用它们的行为。这样，我们就可以将动物的实现与 `Zoo` 类型分离开来，从而提高代码的可维护性和可测试性。  
总之，使用接口编程可以让我们编写更加灵活、可维护、可测试的代码。

#### 函数式编程
函数式编程（Functional Programming，简称 FP）是一种编程范式，它强调使用函数来表示程序的逻辑，并避免在程序中使用可变状态和副作用。  
在 Go 语言中实现函数式编程，主要有以下几个方面：  
1. 高阶函数：高阶函数是指接受一个或多个函数作为参数，或者返回一个函数的函数。在 Go 语言中，可以使用函数类型作为参数类型或返回类型。例如：
```go
func apply(f func(int, int) int, a int, b int) int {
    return f(a, b)
}
```
2. 匿名函数和闭包：在 Go 语言中，可以使用匿名函数（也称为 Lambda 函数）来创建一个没有名字的函数。当一个匿名函数引用了外部作用域的变量时，就形成了闭包。例如：  
```go
func main() {
    add := func(a int, b int) int {
        return a + b
    }

    result := apply(add, 3, 4)
    fmt.Println(result) // 输出：7
}
```
> 匿名函数（Anonymous Function）：匿名函数是指没有函数名的函数。在 Go 语言中，匿名函数通常用于定义简单的、只在一个地方使用的函数  
> 闭包是指一个函数捕获了其外部作用域中的变量，并且这些变量可以在闭包中被访问和修改。闭包是由匿名函数和外部作用域中的变量共同组成的  

3. 递归：函数式编程中，递归是一种常见的处理方式。在 Go 语言中，可以使用递归函数解决问题。例如，计算斐波那契数列：
```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```
4. 不可变数据结构：函数式编程鼓励使用不可变数据结构，这可以避免程序中的副作用。在 Go 语言中，可以使用只读的数据结构来实现不可变性，例如使用 `const` 定义常量。  
5. 函数组合：在函数式编程中，可以将多个函数组合起来，形成新的函数。例如，将两个函数 `f(x)` 和 `g(x)` 组合成一个新函数 `h(x) = f(g(x))`。  
```go
func compose(f func(int) int, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}
```
总之，虽然 Go 语言并非专为函数式编程设计，但在实际开发中，我们仍然可以借鉴函数式编程的思想，编写简洁、高效和可维护的代码。

再看一个例子，
```go
type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}
// 要用的函数类型
type Option func(*Server)

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func Tls(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}
func NewServer(addr string, port int, options ...Option) *Server {
	// 给定默认值
	ser := &Server{
		Addr:     addr,
		Port:     port,
		Protocol: "xx",
		Timeout:  time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	// 二次赋值，比如 Timeout、Tls
	for _, option := range options {
		option(ser)
	}
	return ser
}
```

#### 修饰器模式
修饰器模式其实就是函数式编程的一种，它的主要思想就是传入一个函数，然后返回的还是一个函数，我们将传入的这个函数进行二次修饰，然后再返回，进而调用使用。
```go
func decorator(fn func(s string)string)func(string)string {
	return func(s string) string {
		return fn(s) + "。。。"
	}
}

hello  := func(s string)string {
    return s
}

defn := decorator(hello)
defn("你好")
// print: 你好。。。
```

#### pipeline 管道模式
```go
func A(des ...fn()){
	for _,v := range des {
		v()
	}
}

A(fn1, fn2, fn3, fn4)
```

#### k8s 的 visitor 模式
K8s 的 visitor 模式是一种特殊的工作方式，它可以让一个应用程序在不安装在容器集群中的情况下，仍然可以通过容器集群访问资源。  
比如说，你有一个应用程序需要访问数据库，但是你不想将数据库安装在容器集群中。在 visitor 模式下，你可以创建一个访问者，让它连接到容器集群，然后通过它访问数据库。  
这样做的优点是可以隔离应用程序和数据库，使得它们更容易管理和维护。它还可以让你更方便地控制应用程序的访问权限，以保护数据安全。    
总的来说，K8s 的 visitor 模式是一种很方便的工具，可以帮助你在不把所有东西都放到容器集群中的情况下，仍然可以利用容器集群的优势。  
```go
package main
import "fmt"

func main() {
	p := new(Peo)
	p.year = 10
	p.name = "a"
    // 调用它的 Did 方法，并将 Run 函数作为参数传递给它
    // 这样，Run 函数就可以通过 Peo 类型的对象来访问它所需要的资源
	p.Did(Run)
	p.year = 100
    // Run 函数就可以再次通过 Peo 类型的对象来访问它所需要的资源
	p.Did(Run)
}

// 定义一个 Visitor 函数类型，它接收一个 Do 接口类型的参数
type Visitor func(Do)
// Do 接口类型定义了一个 Did 方法，它接收一个 Visitor 类型的参数
type Do interface {
	Did(Visitor)
}

// Peo 类型实现了 Did 方法，它接受一个 Visitor 类型的参数，并将自己作为参数传递给 Visitor 函数
type Peo struct {
	name string
	year int
}
func (p *Peo) Did(v Visitor) {
	v(p)
}


// 算法，它接受一个 Do 接口类型的参数，并通过接口对象来进行一系列的操作，真实的数据结构和这里的算法完全解除耦合
func Run(do Do) {
	fmt.Println(do)
}
```

#### 泛型
Golang 泛型是 Go 语言在 1.18 版本中引入的一个新特性，它让我们能够编写更通用、更灵活的代码。简单来说，泛型就像是代码的模板，允许我们为不同的数据类型编写相同的逻辑，而不需要重复编写代码。  
举个例子，如果你想要实现一个可以对任意类型的切片进行排序的函数，而不仅仅是 int 类型，你可以使用泛型。在使用泛型之前，你可能需要为每种类型编写一个排序函数，如 `sortInts`、`sortFloats` 等。而有了泛型，你可以用一个函数来实现所有的排序需求。  

要使用泛型，你需要了解两个关键概念：  
1. 类型参数（Type Parameters）：类型参数是用来表示泛型中的未知类型，通常用大写字母 T、U、V 等表示。你可以把它看作是一个占位符，表示我们还不知道它是什么具体类型，但我们希望代码可以适应不同类型的数据。  
2. 约束（Constraints）：约束用于限制泛型中可以使用的类型，以确保我们的代码可以安全地使用这些类型。例如，我们可以要求类型参数 T 必须实现某个接口，这样我们就可以在泛型代码中使用接口中定义的方法。  
```go
package main

import (
	"fmt"
	"sort"
)

type SortableSlice[T any] []T

func (s SortableSlice[T]) Sort(less func(T, T) bool) {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
}

func main() {
    // 分别为整数切片和字符串切片提供了两个不同的 less 函数实现
	ints := SortableSlice[int]{3, 1, 4, 2}
	ints.Sort(func(a, b int) bool { return a < b })
	fmt.Println(ints) // 输出：[1 2 3 4]

	strings := SortableSlice[string]{"banana", "apple", "cherry", "orange"}
	strings.Sort(func(a, b string) bool { return a < b })
	fmt.Println(strings) // 输出：[apple banana cherry orange]
}
```
再看一个例子，
```go
func Map[T any](data []T, fn func(T)T) []T {
	var ma []T
	for _, value := range data {
		ma = append(ma, fn(value))
	}
	return ma
}

Map([]string{"1"}, func(i string)string {
    return i + i
})

Map([]int{1}, func(i int)int {
    return i + i
})
```

