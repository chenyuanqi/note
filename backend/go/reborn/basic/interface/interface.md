
### Go 接口
接口本身是调用方和实现方均需要遵守的一种协议，大家按照统一的方法命名参数类型和数量来协调逻辑处理的过程。  
Go 语言中使用组合实现对象特性的描述。对象的内部使用结构体内嵌组合对象应该具有的特性，对外通过接口暴露能使用的特性。  
Go 语言的接口设计是非侵入式的，接口编写者无须知道接口被哪些类型实现。而接口实现者只需知道实现的是什么样子的接口，但无须指明实现哪一个接口。编译器知道最终编译时使用哪个类型实现哪个接口，或者接口应该由谁来实现。  

**接口声明（定义）**  
Go语言里有非常灵活的接口概念，通过它可以实现很多面向对象的特性。很多面向对象的语言都有相似的接口概念，但Go语言中接口类型的独特之处在于它是满足隐式实现的。也就是说，我们没有必要对于给定的具体类型定义所有满足的接口类型；简单地拥有一些必需的方法就足够了。  
这种设计可以让你创建一个新的接口类型满足已经存在的具体类型却不会去改变这些类型的定义；当我们使用的类型来自于不受我们控制的包时这种设计尤其有用。  
接口类型是对其它类型行为的抽象和概括；因为接口类型不会和特定的实现细节绑定在一起，通过这种抽象的方式我们可以让我们的函数更加灵活和更具有适应能力。  
接口是双方约定的一种合作协议。接口实现者不需要关心接口会被怎样使用，调用者也不需要关心接口的实现细节。接口是一种类型，也是一种抽象结构，不会暴露所含数据的格式、类型及结构。  

Go 语言的每个接口中的方法数量不会很多。Go 语言希望通过一个接口精准描述它自己的功能，而通过多个接口的嵌入和组合的方式将简单的接口扩展为复杂的接口。
- 接口类型名：使用 type 将接口定义为自定义的类型名。Go 语言的接口在命名时，一般会在单词后面添加 er，如有写操作的接口叫 Writer，有字符串功能的接口叫 Stringer，有关闭功能的接口叫 Closer 等。
- 方法名：当方法名首字母是大写时，且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
- 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以被忽略
```go
// 每个接口类型由数个方法组成。接口的形式代码
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}

// 例如 io 包中提供的 Writer 接口
// 调用 Write() 方法写入一个字节数组（[]byte），返回值告知写入字节数（n int）和可能发生的错误（err error）
type Writer interface {
    Write(p []byte) (n int, err error)
}
// 以字符串形式展现的接口，只要实现了这个接口的类型，在调用 String() 方法时，都可以获得对象对应的字符串
// Stringer 接口在 Go 语言中的使用频率非常高，功能类似于 Java 或者 C# 语言里的 ToString 的操作
type Stringer interface {
    String() string
}
```

**实现接口的条件**   
如果一个任意类型 T 的方法集为一个接口类型的方法集的超集，则我们说类型 T 实现了此接口类型。T 可以是一个非接口类型，也可以是一个接口类型。  
实现关系在 Go 语言中是隐式的。两个类型之间的实现关系不需要在代码中显式地表示出来。Go 语言中没有类似于 implements 的关键字。 Go 编译器将自动在需要的时候检查两个类型之间的实现关系。  
接口定义后，需要实现接口，调用方才能正确编译通过并使用接口。接口的实现需要遵循两条规则才能让接口可用。  
- 接口被实现的条件一：接口的方法与实现接口的类型方法格式一致  
- 接口被实现的条件二：接口中所有方法均被实现，即当一个接口中有多个方法时，只有这些方法都被实现了，接口才能被正确编译并使用

Go 语言的接口实现是隐式的，无须让实现接口的类型写出实现了哪些接口。这个设计被称为非侵入式设计。  
实现者在编写方法时，无法预测未来哪些方法会变为接口。一旦某个接口创建出来，要求旧的代码来实现这个接口时，就需要修改旧的代码的派生部分，这一般会造成雪崩式的重新编译。  
```go
package main
import (
    "fmt"
)
// 定义一个数据写入器
type DataWriter interface {
    WriteData(data interface{}) error
}
// 定义文件结构，用于实现DataWriter
type file struct {
}
// 实现DataWriter接口的WriteData方法
func (d *file) WriteData(data interface{}) error {
    // 模拟写入数据
    fmt.Println("WriteData:", data)
    return nil
}
func main() {
    // 实例化file
    f := new(file)
    // 声明一个DataWriter的接口
    var writer DataWriter
    // 将接口赋值f，也就是*file类型
    writer = f
    // 使用DataWriter接口进行数据写入
    writer.WriteData("data")
}
```

当类型无法实现接口时，编译器会报错，下面列出常见的几种接口无法实现的错误。  
1) 函数名不一致导致的报错  
在以上代码的基础上尝试修改部分代码，造成编译错误，通过编译器的报错理解如何实现接口的方法。  
```go
// 报错，不能将 f 变量（类型*file）视为 DataWriter 进行赋值
// 原因：*file 类型未实现 DataWriter 接口（丢失 WriteData 方法）
func (d *file) WriteDataX(data interface{}) error {}
```
2) 实现接口的方法签名不一致导致的报错  
将修改的代码恢复后，再尝试修改 WriteData() 方法，把 data 参数的类型从 interface{} 修改为 int 类型。  
```go
// 未实现 DataWriter 的理由变为（错误的 WriteData() 方法类型）发现 WriteData(int)error，期望 WriteData(interface{})error
func (d *file) WriteData(data int) error {}
```

**类型与接口的关系**  
在 Go 语言中类型和接口之间有一对多和多对一的关系。  

一个类型可以实现多个接口  
一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。  
网络上的两个程序通过一个双向的通信连接实现数据的交换，连接的一端称为一个 Socket。Socket 能够同时读取和写入数据，这个特性与文件类似。因此，开发中把文件和 Socket 都具备的读写特性抽象为独立的读写器概念。  
Socket 和文件一样，在使用完毕后，也需要对资源进行释放。  
使用 Socket 实现的 Writer 接口的代码，无须了解 Writer 接口的实现者是否具备 Closer 接口的特性。同样，使用 Closer 接口的代码也并不知道 Socket 已经实现了 Writer 接口。usingWriter () 和 usingCloser () 完全独立，互相不知道对方的存在，也不知道自己使用的接口是 Socket 实现的。  
```go
// 把 Socket 能够写入数据和需要关闭的特性使用接口来描述
type Socket struct {
}
func (s *Socket) Write(p []byte) (n int, err error) {
    return 0, nil
}
func (s *Socket) Close() error {
    return nil
}

// Socket 结构的 Write () 方法实现了 io.Writer 接口
type Writer interface {
    Write(p []byte) (n int, err error)
}
// Socket 结构也实现了 io.Closer 接口
type Closer interface {
    Close() error
}

// 使用io.Writer的代码, 并不知道Socket和io.Closer的存在
func usingWriter( writer io.Writer){
    writer.Write( nil )
}
// 使用io.Closer, 并不知道Socket和io.Writer的存在
func usingCloser( closer io.Closer) {
    closer.Close()
}
func main() {
    // 实例化Socket
    s := new(Socket)
    usingWriter(s)
    usingCloser(s)
}
```

多个类型可以实现相同的接口  
一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。也就是说，使用者并不关心某个接口的方法是通过一个类型完全实现的，还是通过多个结构嵌入到一个结构体中拼凑起来共同实现的。  
Service 接口定义了两个方法：一个是开启服务的方法（Start ()），一个是输出日志的方法（Log ()）。使用 GameService 结构体来实现 Service，GameService 自己的结构只能实现 Start () 方法，而 Service 接口中的 Log () 方法已经被一个能输出日志的日志器（Logger）实现了，无须再进行 GameService 封装，或者重新实现一遍。所以，选择将 Logger 嵌入到 GameService 能最大程度地避免代码冗余，简化代码结构。  
```go
// 一个服务需要满足能够开启和写日志的功能
type Service interface {
    Start()  // 开启服务
    Log(string)  // 日志输出
}
// 日志器
type Logger struct {
}
// 实现Service的Log()方法
func (g *Logger) Log(l string) {
}
// 游戏服务
type GameService struct {
    Logger  // 嵌入日志器
}
// 实现Service的Start()方法
func (g *GameService) Start() {
}

// 实例化 GameService，并将实例赋给 Service
var s Service = new(GameService)
// s 就可以使用 Start () 方法和 Log () 方法
// 其中，Start () 由 GameService 实现，Log () 方法由 Logger 实现
s.Start()
s.Log(“hello”)
```

**接口嵌套组合**  




