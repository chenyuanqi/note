
### Go 接口
接口本身是调用方和实现方均需要遵守的一种协议，大家按照统一的方法命名参数类型和数量来协调逻辑处理的过程。  
Go 语言中使用组合实现对象特性的描述。对象的内部使用结构体内嵌组合对象应该具有的特性，对外通过接口暴露能使用的特性。  
Go 语言的接口设计是非侵入式的，接口编写者无须知道接口被哪些类型实现。而接口实现者只需知道实现的是什么样子的接口，但无须指明实现哪一个接口。编译器知道最终编译时使用哪个类型实现哪个接口，或者接口应该由谁来实现。  

使用接口的目的是什么？概括地说有两点：  
一是规范某个对象的行为，使其受控；  
二是接口的使用者和实现者各司其职，互不干涉。  

一般来说，接口分为两个步骤来准备，首先是定义（行为规范），然后是相关对象的实现（具体的操作），准备好后便可在后续的代码中使用接口了。

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
接口的实现，实际上就是指实现具体的行为，比如从缓存中加载或从网络上下载图片的具体方法。
```go
// 实现接口的格式
func (struct_variable struct_name) function_name([params]) [return_values] {
   // 方法实现 
}
```

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

**接口的调用**  
编写作用于*fileCache的接口实现，继续定义负责从网络下载图片的结构体以及作用于该结构体的接口实现。  
定义一个 ImageDownloader 类型的变量，然后通过 new(fileCache) 函数为其赋值，随后便可通过这个变量调用从缓存中加载图片的方法。  
```go
// ImageDownloader 图片加载接口
type ImageDownloader interface {
	// FetchImage 获取图片，需要传入图片地址，方法返回图片数据
	FetchImage(url string) string
}

type fileCache struct {
}
//FetchImage接口实现
func (f *fileCache) FetchImage(url string) string {
	return "从本地缓存中获取图片：" + url
}

//定义从网络下载图片的结构体
type netFetch struct {
}
//FetchImage接口实现
func (n *netFetch) FetchImage(url string) string {
	return "从网络下载图片：" + url
}

func main() {
	//从本地缓存中获取数据
	var imageLoader ImageDownloader
	imageLoader = new(fileCache)
	data := imageLoader.FetchImage("https://www.example.com/a.png")
	fmt.Println(data)
	if data == "" {
		// 当本地缓存中没有数据时，从网络下载
		var imageLoader2 ImageDownloader
		imageLoader2 = new(netFetch)
		data2 := imageLoader2.FetchImage("https://www.example.com/a.png")
		fmt.Println(data2)
	}
}
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
在 Go 语言中，不仅结构体与结构体之间可以嵌套，接口与接口间也可以通过嵌套创造出新的接口。  
一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。只要接口的所有方法被实现，则这个接口中的所有嵌套接口的方法均可以被调用。  

系统包中的接口嵌套组合  
Go 语言的 io 包中定义了写入器（Writer）、关闭器（Closer）和写入关闭器（WriteCloser）3 个接口。
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
type Closer interface {
    Close() error
}
type WriteCloser interface {
    Writer
    Closer
}
```

在代码中使用接口嵌套组合  
在代码中使用 io.Writer、io.Closer 和 io.WriteCloser 这 3 个接口时，只需要按照接口实现的规则实现 io.Writer 接口和 io.Closer 接口即可。而 io.WriteCloser 接口在使用时，编译器会根据接口的实现者确认它们是否同时实现了 io.Writer 和 io.Closer 接口。  
```go
package main
import (
    "io"
)
// 声明一个设备结构
type device struct {
}
// 实现io.Writer的Write()方法
func (d *device) Write(p []byte) (n int, err error) {
    return 0, nil
}
// 实现io.Closer的Close()方法
func (d *device) Close() error {
    return nil
}
func main() {
    // 声明写入关闭器, 并赋予device的实例
    var wc io.WriteCloser = new(device)
    // 写入数据
    wc.Write(nil)
    // 关闭设备
    wc.Close()
    // 声明写入器, 并赋予device的新实例
    var writeOnly io.Writer = new(device)
    // 写入数据
    writeOnly.Write(nil)
}
```

**接口和类型之间的转换**  
Go 语言中使用接口断言（type assertions）将接口转换成另外一个接口，也可以将接口转换为另外的类型。接口的转换在开发中非常常见，使用也非常频繁。

类型断言的格式  
类型断言是一个使用在接口值上的操作。语法上它看起来像 i.(T) 被称为断言类型，这里 i 表示一个接口的类型和 T 表示一个类型。一个类型断言检查它操作对象的动态类型是否和断言的类型匹配。  
类型断言的基本格式如下：t := i.(T)。其中，i 代表接口变量，T 代表转换的目标类型，t 代表转换后的变量。  

这里有两种可能。第一种，如果断言的类型 T 是一个具体类型，然后类型断言检查 i 的动态类型是否和 T 相同。如果这个检查成功了，类型断言的结果是 i 的动态值，当然它的类型是 T。换句话说，具体类型的类型断言从它的操作对象中获得具体的值。如果检查失败，接下来这个操作会抛出 panic。
```go
var w io.Writer
w = os.Stdout
f := w.(*os.File) // 成功: f == os.Stdout
c := w.(*bytes.Buffer) // 死机：接口保存*os.file，而不是*bytes.buffer
```
第二种，如果相反断言的类型 T 是一个接口类型，然后类型断言检查是否 i 的动态类型满足 T。如果这个检查成功了，动态值没有获取到；这个结果仍然是一个有相同类型和值部分的接口值，但是结果有类型 T。换句话说，对一个接口类型的类型断言改变了类型的表述方式，改变了可以获取的方法集合（通常更大），但是它保护了接口值内部的动态类型和值的部分。
```go
// 第一个类型断言后，w 和 rw 都持有 os.Stdout 因此它们每个有一个动态类型 *os.File，但是变量 w 是一个 io.Writer 类型只对外公开出文件的 Write 方法，然而 rw 变量也只公开它的 Read 方法
var w io.Writer
w = os.Stdout
rw := w.(io.ReadWriter) // 成功：*os.file具有读写功能
w = new(ByteCounter)
rw = w.(io.ReadWriter) // 死机：*字节计数器没有读取方法

//  i 没有完全实现 T 接口的方法，这个语句将会触发宕机。触发宕机不是很友好，换成这样
t,ok := i.(T)
```
如果断言操作的对象是一个 nil 接口值，那么不论被断言的类型是什么这个类型断言都会失败。几乎不需要对一个更少限制性的接口类型（更少的方法集合）做断言，因为它表现的就像赋值操作一样，除了对于 nil 接口值的情况。

将接口转换为其他接口  
实现某个接口的类型同时实现了另外一个接口，此时可以在两个接口间转换。
```go
package main
import "fmt"
// 定义飞行动物接口
type Flyer interface {
    Fly()
}
// 定义行走动物接口
type Walker interface {
    Walk()
}
// 定义鸟类
type bird struct {
}
// 实现飞行动物接口
func (b *bird) Fly() {
    fmt.Println("bird: fly")
}
// 为鸟添加Walk()方法, 实现行走动物接口
func (b *bird) Walk() {
    fmt.Println("bird: walk")
}
// 定义猪
type pig struct {
}
// 为猪添加Walk()方法, 实现行走动物接口
func (p *pig) Walk() {
    fmt.Println("pig: walk")
}
func main() {
// 创建动物的名字到实例的映射
    animals := map[string]interface{}{
        "bird": new(bird),
        "pig":  new(pig),
    }
    // 遍历映射
    for name, obj := range animals {
        // 判断对象是否为飞行动物
        f, isFlyer := obj.(Flyer)
        // 判断对象是否为行走动物
        w, isWalker := obj.(Walker)
        fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)
        // 如果是飞行动物则调用飞行动物接口
        if isFlyer {
            f.Fly()
        }
        // 如果是行走动物则调用行走动物接口
        if isWalker {
            w.Walk()
        }
    }
}
// name: pig isFlyer: false isWalker: true
// pig: walk
// name: bird isFlyer: true isWalker: true
// bird: fly
// bird: walk
```

将接口转换为其他类型  
将接口转换为普通的指针类型，例如将 Walker 接口转换为 *pig 类型。
```go
p1 := new(pig)
var a Walker = p1
p2 := a.(*pig)
fmt.Printf("p1=%p p2=%p", p1, p2)

// 如果尝试将上面这段代码中的 Walker 类型的 a 转换为 *bird 类型，将会发出运行时错误
p1 := new(pig)
var a Walker = p1
p2 := a.(*bird)
// panic: interface conversion: main.Walker is *main.pig, not *main.bird 即main.Walker 接口的内部保存的是 *main.pig，而不是 *main.bird
```
接口在转换为其他类型时，接口内保存的实例对应的类型指针，必须是要转换的对应的类型指针。  
接口断言类似于流程控制中的 if。但大量类型断言出现时，应使用更为高效的类型分支 switch 特性。  
