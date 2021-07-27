
### 为什么是 Go
Go 作为 Google 开源的编程语言，近些年来，除 Docker 本身就是 Go 写的，知乎、Bilibili 都纷纷迁移到 Go，面对中国的人口优势，Go 的特性优势就被发挥的很充分。  
Go 语言的优势大抵有以下：  

- 如果有 C 语言的加持，可以轻松学会；
- 部署简单，实现高并发很轻松，内存占用也非常省；
- 代码简洁、格式清晰统一、方便协作与阅读；
- 具备性能强劲但是开发效率不输给一些动态语言，适合编写一些瓶颈业务。

**Go 可以做什么**  
云计算、DevOps、区块链、人工智能、存储引擎及 Web 服务器等。

### Go 安装
[安装包地址](https://golang.google.cn/dl/)，各环境下载对应的安装包。  

**Windows**  
下载最新的 zip 文件。如果你的电脑是 64 位的系统，你将需要 go#.#.#.windows-amd64.zip（这里的  #.#.# 是 Go 的最新版本号）。
解压缩  go#.#.#.windows-amd64.zip 文件到你选择的位置（c:\Go 这个位置是个不错的选择）。
下载最新的 zip 文件。如果你的电脑是 64 位的系统，你将需要 go#.#.#.windows-amd64.zip ，这里的  #.#.# 是 Go 的最新版本号。
解压缩  go#.#.#.windows-amd64.zip 文件到你选择的位置。 c:\Go 这个位置是个不错的选择。  

在系统中设置两个环境变量：
1、GOPATH 同样的指向的是你的工作目录（这个变量看起来像 c:\users\goku\work\go 这个样子）  
2、添加 c:\Go\bin  到系统的 PATH 环境变量  
打开一个 cmd 命令终端，输入 go version。  

**Linux/OSX**  
```bash
# Ubuntu 安装
# 下载
wget https://studygolang.com/dl/golang/go1.12.5.linux-amd64.tar.gz
tar -zxvf go1.12.5.linux-amd64.tar.gz
sudo mv go /usr/local/
# 配置 
export GOROOT=/usr/local/go # 安装目录
export GOPATH=$HOME/go # 工作环境
export GOBIN=$GOPATH/bin # 可执行文件存放
export PATH=$GOPATH:$GOBIN:$GOROOT/bin:$PATH # 添加 PATH 路径
# 测试
go version

# MacOS 安装
# 下载安装包并运行安装包
```

**文件树结构**  
Go 安装目录（$GOROOT）的文件夹结构应该如下所示：  

README.md, AUTHORS, CONTRIBUTORS, LICENSE

- /bin：包含可执行文件，如：编译器，Go 工具
- /doc：包含示例程序，代码工具，本地文档等
- /lib：包含文档模版
- /misc：包含与支持 Go 编辑器有关的配置文件以及 cgo 的示例
- /os_arch：包含标准库的包的对象文件（.a）
- /src：包含源代码构建脚本和标准库的包的完整源代码（Go 是一门开源语言）
- /src/cmd：包含 Go 和 C 的编译器和命令行脚本


### Hello Golang
```golang
package main

import "fmt"

func main() {
    fmt.Println("hello, golang")
}
```

### Go 基础
Go 基础包括环境安装、语言结构、基础结构和数据类型、数字和切片、map，流程控制、函数、struct 和 method、interface 和 反射、Goroutine、channel、常用包的使用（包括文件读取、时间和日期、Xml 和 Json 等格式解析、字符串处理、正则、锁和 sync 包、网络处理等）。

**注释**  
注释 是为了增加代码的可读性。  
```golang
// 单行注释

/*
多行注释 1
多行注释 2
多行注释 3
*/
```

**变量**  
变量的定义和声明不是一个概念。  
> 定义变量，将为期分配内存。  
> 声明则是告知编译器（或链接器）有这样一个符号（类型）。  

变量的定义需要明确分配内存，也就是同一个变量只能定义一次，不可能同一个变量分配两次内存。运行期修改是另外一回事，但定义不能指定两块内存。

变量的定义方式通常有两种，第一种称之为全局变量，第二种称之为局部变量。  
全局变量的作用域是整个包，局部变量的作用域是该变量所在的花括号内。

```golang
// 显式初始化
var x int = 100
// 隐式初始化
var y int

// 变量的定义
// 无法定义只读变量（readonly、const）
var x int // 自动初始化为零值
var s = "abc" // 可根据初始化值推断类型
var a, b = 1, 2.0 // 可一次定义多个不同类型的变量

// 分组的方式写法
var (
    x int
    s = "abc"
    a, b = 1, 200
)

// 自动推导类型：简短定义（必须显式提供初始化值，不能提供数据类型且只能用于函数内部）
x := 123

// 多变量赋值（先计算右边的值，然后批量对左边进行赋值）
// 注意：未使用的变量会引发错误
a, b := 1, 2
a, b = b+1, a+2

// 匿名变量（丢弃数据不进行处理，_ 匿名变量配合函数返回值使用才显示其价值）
_, _, c, d := 120, 110, "你好", "朋友"

// 动态修改字符串变量（仅支持字符串，可设置非导出成员）
var BuildTime string

func main() {
    println(BuildTime)
}
// go build -ldflags "-X main.BuildTime=$(date +'%Y.%m.%d')"
```

所有的变量都会被翻译成内存地址，因为符号名最终是没有任何意义的。所有的变量都是可寻址的，不管是全局变量还是局部变量最终是要求可寻址的，但是变量可以寻址并不代表一个计算中间结果可以寻址。  
变量代表着一段或者多段存储内存，变量实际上就是一种内存。  
定义一个变量，这个变量存储数据，但数据究竟存在哪，我们知道存储器有很多种，存储器体系来说寄存器、L1、L2、L3 三级缓存、主存。虚拟内存里有部分数据可能会交换到磁盘上，还有硬盘上存储、网络上存储。  

变量的命名规范：  

- 全局变量建议使用完整且有明确含义的单词；
- 局部变量建议使用短名和缩写，以便区分全局变量；
- 不要使用保留关键字、内置函数，以及常用标准库成员名称；
- 专有名词建议大写（escapeHTML）；
- 变量最关键的是变量的命名，需要满足可阅读性和可维护性；
> 1、全局变量建议使用完整且有明确含义的单词。可能是两个单词组成的，一个单词比较通用可能会引起一些误解，需要用两个单词明确的表达，但是不要太长。  
> 2、局部变量建议尽量使用短名和缩写，一个函数最好不要超过一个屏幕，用短名和缩写区分全局变量和局部变量避免歧义。  
- 关于注释的问题，假如只是一行加注释通常会建议写在后边，对一整块做加注释建议写在前面。  

**常量**  
常量的定义跟变量差不多，我们可以进行类型推断，可以同时定义多个，可以在函数内部定义，也可以在包块内定义。  
为什么使用常量？作为魔法数字（数字或者字符串），让代码具备更好的阅读性。  
一般定义常量使用大写字母，常量里面的值确定好后，后面是不允许修改的；常量可以参与程序的计算，不允许左值赋值。  
在程序开发中，我们用常量存储一直不会发生变化的数据，例如：π，身份证号码等。像这类的数据，在整个程序中运行中都是不允许发生改变的。
```golang
const x int32 = 100 // 指定常量类型，则左右类型必须一致
const s uintptr = unsafe.Sizeof(0) // 必要时，可进行类型转换
const n int = len("abc") // 支持编译期能计算结果的表达式
const (
    a int = 1 * int(unsafe.Sizeof("abc"))
    b
)

// 计算圆的面积和周长
// 面积 PI*r*r  math.Pow(r,2)
// 周长 2*PI*r
// 常量必须定义
const PI float64 = 3.14
var r float64
fmt.Scan(&r)
// 面积
s := PI * r * r
// 周长
l := 2 * PI * r
// fmt.Println(PI)
fmt.Printf("面积:%.2f\n", s)
fmt.Printf("周长:%.2f\n", l)

const MAX = "你瞅啥"
// fmt.Println(MAX)
// fmt.Printf("%T\n",MAX) // string
// go 语言常量的地址 不允许访问
// fmt.Printf("%p",&MAX) // err

// 所谓字面常量（literal），是指程序中硬编码的常量，比如
-12
3.14159265358979323846  // 浮点类型的常量
3.2+12i                 // 复数类型的常量
true                    // 布尔类型的常量
"foo"                   // 字符串常量
// 编程语言源程序中表示固定值的符号叫做字面量，也称字面常量
// Go 的字面量可以出现在两个地方：一是用于常量和变量的初始化，二是用在表达式中作为函数调用实参
// 变量初始化语句中如果没有显式地指定变量类型，则 Go 编译器会结合字面量的值自动进行类型推断
// Go 中的字面量只能表达基本类型的值，Go 不支持用户自定义字面量
```

常量有全局的，这样的好处就在于，我们在多个文本当中若想调整的话，只需要在定义常量的地方调整就行了。常量也可以是局部的，使用常量替换掉以后可阅读性就会好很多，同时我们在编码当中强调的观点是把逻辑和数据分离。  

严格意义上来说，没有运行期常量的概念，常量会被直接展开到你需要用的地方，既然没有运行期常量，所以它没有地址，不能会对常量取地址。换句话说，常量是数据，把数据放在某个地方才会有地址吧，那个地方有地址，也就是说虚拟空间有地址但数据本身没有地址。

**输入和输出**  
在 GO 语言中进行输出，用到两个函数：Print() 和 Println()。这两个函数的区别是 Print() 函数不换行，Println() 换行输出。
```golang
// 双引号内容原样输出
fmt.Print("a", a)

c:="你瞅啥"
// %s是一个占位符 表示输出一个字符串类型
fmt.Printf("%s",c)

a := 10
b := 3.14559
// %d是一个占位符 表示输出一个整型数据
// %f是一个占位符 表示输出一个浮点型数据
// %f默认保留六位小数  因为浮点型数据不是精准的数据 六位是有效的
// %.2f保留小数位数为两位  会对第三位小数进行四舍五入
// \n表示一个转义字符 换行
fmt.Printf("%d %.2f\n", a, b)
```

在 GO 中我们用到了 “fmt” 这个包中的 Scanf() 函数来接收用户键盘输入的数据。
当程序执行到 Scanf() 函数后，会停止往下执行，等待用户的输入，输入完成后程序继续往下执行。
```golang
func main0701() {
    var a int
    // 通过键盘为变量赋值
    // & 是一个运算符  取地址运算符
    fmt.Scan(&a)
    // 内存地址 0xc042058080  是一个十六进制整型数据
    // fmt.Println(&a)
    fmt.Println(a)
}
```

**枚举与 iota**  
枚举是非常常见的类型，通常情况下指的是一种一连串或者连续性的定义，它的总数是固定的，比如星期、月份、容量、颜色。它是有一定的规律并且可以用一连串顺序数字代替。  

- 没有明确意义上的 enum 定义
在 Go 语言里没有明确意义上的枚举定义。  
在 Go 语言里面枚举实质上是常量。  
```golang
type color byte //自定义类型

const (
    red color = iota //指定常量类型
    yellow
    blue
)
```

iota 枚举格式如果写在一行中值则相等，如果换行则值在上一行加一。  
常量声明可以使用 iota 常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。（注意：在一个 const 声明语句中，在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加一）
- 借助 iota 实现常量组自增值
- 可使用多个 iota，各自单独计数
iota 实际上是常量组里面实现自增的操作，严格来说和枚举没多大关系。  
iota 是编译器，为我们产生连续性数字。其实质是一个计数器，它从零开始计数，每行添加一。它是给编译器看的占位符，告诉编译器在一组里递增数字，每一组 iota 会重新进行计算。iota 可以作为表达式里面其中的操作数。  
```golang
const (
    a    = iota // 0
    b, c = iota, iota // 1 1
)

const (
    _ = 1 << (10 * iota)
    KB
    MB
    GB
)

const (
    _, _ = iota, iota * 10 // 0, 0 * 10
    a, b                   // 1, 1 * 10
    c, d                   // 2, 2 * 10
)
```

- 中断须显式恢复

**基本类型与空值**  
[![Golang 基本类型](https://z3.ax1x.com/2021/05/23/gXutZ8.png)](https://imgtu.com/i/gXutZ8)
相比 C 语言，Golang 有明确的基本类型。Golang 基本类型除了很明确的类型以外，其中比较特殊的有几个，其中 uintptr 用来存储地址的整数，rune 用来存储编码的码点，int32 有点类似 UCS-2 方式，但是不完全一样，毕竟不是对等关系。
```golang
x := 0x123456 // 56 34 12 00 00 00 00 00
p := (*[8]byte)(unsafe.Pointer(&x))
fmt.Printf("%x\n", p[1]) //打印第 1 位输出 34

// 布尔类型的变量取值要么是真（true）, 要么是假 (false)
// 作用：布尔类型主要用于条件判断
//布尔类型 值为 true 或者为 false
// var a bool // 默认值为false
// bool 类型一般用于条件判断
// a = true
// fmt.Println(a)
// 自动推到类型创建bool类型变量
a := false // bool

// 浮点型数据 分为 单精度浮点型 float32（小数位数为 7 位） 双精度浮点型 float64（小数位数为 15 位）
// float64 比 float32 更精准（通过自动推到类型创建的浮点型变量 默认类型为 float64）
a := 123.456
fmt.Printf("%T\n", a)

// byte 字符类型 同时也是 uint8 的别名
// 所有的字符都对应 ASCII 中的整型数据，比如 ‘0’对应的 48、 ‘A’对应的 65、‘a’ 对应的 97
// 所谓字符类型是用单引号括起来的单个字符
var a byte ='\t'
fmt.Println(a)
// 第二种方式
var a byte
a = 97
fmt.Printf("a=%c",a) // 输出必须 %c,否则还是整数输出

// 用单引号括起来的单个字符是字符类型，用双引号括起来的字符是字符串类型
// 定义字符串 str := "a"
// len 函数  用来计算字符串中字符个数（不包含 \0，返回值为 int 类型）
// a := "hello"
// 在 go 语言中一个汉字占 3 个字符 
a := "learnku函数"
var count int
count = len(a)
fmt.Println(count) // 13
// 字符串连接使用 + 
str1 := "马大师上线了"
str2 := "接化发，年轻人不讲武德，闪电五连鞭"
str3 := str1 + str2
fmt.Println(str3) // 马大师上线了接化发，年轻人不讲武德，闪电五连鞭
```
字符和字符串的区别：  
字符使用单引号，往往只包含一个字符（转义字符 \n 除外）；字符串使用双引号，由一个或多个字符组成，他们都是隐藏了一个结束符 \0。  


除指针外，函数、接口、字典、切片、通道默认值为 nil。

- nil 不是关键字，代表零值（zero），不仅仅是空引用。
nil 严格意义上来说有两层含义。在抽象层面代表的是空值，如果是指针的话可能表示没有指向任何地址；在实现层面它代表的是零值。

- 没有类型，不能作为简短赋值语句右值。
- 即便同为 nil，不同类型也不能直接比较。（不同类型零值含义不同）
- 值为 nil，不代表没有分配内存。  
```golang
var a []int = nil
println(unsafe.Sizeof(a)) // 24
```

fmt 格式化输入输出，使用格式如下：  
[![fmt 格式化输入输出格式](https://z3.ax1x.com/2021/05/23/gXl5tO.png)](https://imgtu.com/i/gXl5tO)

**复合（引用）类型**  
所谓引用类型，是指其内部结构，而非分配于托管堆。  

- slice、map、channel
从实现角度看，除 slice、interface 是结构体外，map、channel、function 都是指针。

- 使用 make 或初始化语句创建实例
Go 语言的引用类型只是一种行为上的概念，所谓的引用类型更多时候指的是它引用另外一块或者多块内存，用另外一块或者多块内存来存储或者处理一些相关的数据结构，至于这两块内存分配栈上还是堆上是由编译器决定的。任何时候编译器优先在栈上分配，避免对垃圾回收器造成负担。  
所以，所谓的引用类型，它内部需要引用另外一块内存，引用另外一块内存也就意味着必须有初始化的操作。切片引用另外一块数组，字典引用哈希桶。
```golang
m := make(map[string]int)
```

- 使用 new 无法有效初始化
对 new 来说，new 只分配一块被初始化为零值的内存，然后返回它的指针。如果是 new 字典，字典是个指针，它只是返回 8 字节内存，new 不初始化数据，那哈希桶的引用、指数的计算、哈希函数的处理等初始化操作根本不处理，这个字典肯定用不了。

new 只负责按照右边的类型来分配一块内存，这块内存有可能在栈上，也有可能在堆上。

new 返回指针，make 返回实例。  

**运算符与类型转换**   
[![算术运算符](https://z3.ax1x.com/2021/05/23/gXUP0J.png)](https://imgtu.com/i/gXUP0J)  
```golang
a := 10
b := 5
fmt.Println(a + b) //30
fmt.Println(a - b) //-10
fmt.Println(a * b) //200
// 两个整数相除等到的结果也是整型
// 在除法计算时 除数不能为0
fmt.Println(a / b)

a := 10
b := 2
// 取余运算符除数不能为 0
// 取余运算符不能对浮点型使用
c := a % b
fmt.Println(c)

func main() {

// 自增自减运算符
// 可以对浮点型进行自增自减运算，但是不能对常量进行自增自减运算
a := 10
// const a =10
// a = a + 1
// a++ // 自增 在变量本身加一
// a-- // 自减
// 自增自减不能出现在表达式中
// a = a++ + a--
// 二义性 
// 在不同操作系统中运算方式不同，结果可能会产生偏差
// a = a++ * a-- - a-- // err
// b := a-- // err
fmt.Println(a)
// fmt.Println(b)
```
Go 语言中不允许隐式转换，所有类型转换必须显式声明（强制转换），而且转换只能发生在两种相互兼容的类型之间。  
在类型转换时建议低类型转成高类型，保证数据精度；建议整型转成浮点型（int8 -> int16 ->int32 ->int64；float32 ->float64；int64 -> float64）。
```golang
a, b, c := 0, 0, 0
fmt.Scan(&a, &b, &c)
sum := a + b + c
fmt.Println(sum)
// 类型转换:数据类型(变量) / 数据类型(表达式)
// fmt.Println(float64(sum / 3))
fmt.Printf("%.2f", float64(sum)/3)

// 数据类型转换，数据溢出
var a int = 1234
fmt.Println(int8(a)) // -46
fmt.Println(int32(a)) // 1234

// 将浮点型转成整型：保留数据整数部分，丢弃小数部分
var a float64 = 3.999
b := int(a) // 3
fmt.Println(b)
```

赋值运算符  = ，var int num=9;num=num+1; 这里的 = 号是赋值运算符，不是数学义意上的相等。  
常见的赋值运算符如下，前面我们使用的 = 是普通赋值，+=，-= 等我们称为 “复合赋值运算符”。  
[![赋值运算符](https://z3.ax1x.com/2021/05/23/gX57TA.png)](https://imgtu.com/i/gX57TA)  
```golang
// a := 10
// b := 20
// c := a + b
// c += 20 // c = c + 20
// c -= 20
// c *= 20
// c /= 20 // 30
// c = 20
// c %= 3 // c = c % 3
var c int = 10
// 将表达式右侧进行结果计算在进行赋值运算符
c %= (2 + 3)
// c = c % 5 // ok
// c = c % 2 + 3 // err
fmt.Println(c)
```

关系运算符我们又称为比较运算符，关系运算的结果是布尔类型的。
[![关系运算符](https://z3.ax1x.com/2021/05/23/gXoAUA.png)](https://imgtu.com/i/gXoAUA)  
```golang
a := 'a'
b := 'A'
fmt.Println(a > b)
fmt.Println(a != b)
```

有逻辑运算符连接的表达式叫做逻辑表达式，逻辑表达式的结果是 bool 类型，逻辑运算符两边放的一般都是关系表达式或者 bool 类型的值。  
[![逻辑运算符](https://z3.ax1x.com/2021/05/23/gXoyP1.png)](https://imgtu.com/i/gXoyP1)  
```golang
a := 10
b := 20
// c := a > b // flase
// 逻辑非 !，非真为假，非假为真
fmt.Println(!(a > b))

a := 10
b := 20
//逻辑与  &&，同真为真，其余为假
c := a < b && false
fmt.Println(c)

a := 10
b := 20
// 逻辑或  ||，同假为假，其余为真
fmt.Println(a < b || a > b)

a := 10
b := 20
// 逻辑与高于逻辑或
fmt.Println(a > b && b > a || a > 0)
```

[![其他运算符](https://z3.ax1x.com/2021/05/23/gXTMz6.png)](https://imgtu.com/i/gXTMz6)  

在 Go 语言中，一元运算符（一些只需要一个操作数的运算符称为一元运算符（或单目运算符））拥有最高的优先级，二元运算符的运算方向均是从左至右。由上至下代表优先级由高到低：  
[![运算符优先级](https://z3.ax1x.com/2021/05/23/gXTYod.png)](https://imgtu.com/i/gXTYod)  


**流程控制**  
GO 语言有顺序结构、选择结构、循环结构。  
顺序结构：程序按顺序执行，不发生跳转；  
选择结构：依据是否满足条件，有选择的执行相应功能；  
循环结构：依据条件是否满足，循环多次执行某段代码。  
```golang
// 选择结构 if（if-else if: 可以处理范围，可以嵌套使用，执行效率低）
var score int
fmt.Scan(&score)
if score > 700 {
    fmt.Println("我要上清华")
} else if score > 680 {
    fmt.Println("我要上北大")
} else if score > 650 {
    fmt.Println("我要上人大")
} else {
    fmt.Println("我要上波大")
}
/*
注意：
条件语句不需要使用圆括号将条件包含起来 ()；
无论语句体内有几条语句，花括号 {} 都是必须存在的；
左花括号 { 必须与 if 或者 else 处于同一行；
在 if 之后，条件语句之前，可以添加变量初始化语句，使用 ; 间隔，比如 if score := 100; score > 90 {}
*/

// 选择结构 switch（switch: 一般用于等值比较，执行效率高、可以将多个满足相同条件的值放在一起，不建议嵌套使用）
var score int
fmt.Scan(&score)
switch score / 10 {
case 10:
    // case 后面跟着的代码执行完毕后，直接跳出整个 switch 结构，相当于每个 case 后面都跟着 break (终止)
    // 如果我们想执行完成某个 case 后，强制执行后面的 case, 可以使用 fallthrough
    fallthrough
case 9:
    fmt.Println("A")
case 8:
    fmt.Println("B")
case 7:
    fmt.Println("C")
case 6, 5:
    fmt.Println("D")
default:
    fmt.Println("E")
}
// 或者这样
score := 100
switch {
case score >= 90:
    fmt.Println("Grade: A")
case score >= 80 && score < 90:
    fmt.Println("Grade: B")
case score >= 70 && score < 80:
    fmt.Println("Grade: C")
case score >= 60 && score < 70:
    fmt.Println("Grade: D")
default:
    fmt.Println("Grade: F")
}
/*
注意：
和条件语句一样，左花括号 { 必须与 switch 处于同一行；
单个 case 中，可以出现多个结果选项（通过逗号分隔）；
与其它语言不同，Go 语言不需要用 break 来明确退出一个 case；
只有在 case 中明确添加 fallthrough 关键字，才会继续执行紧跟的下一个 case；
可以不设定 switch 之后的条件表达式，在这种情况下，整个 switch 结构与多个 if...else... 的逻辑作用等同。
*/


// 循环结构
sum := 0
for i := 1; i <= 100; i++ {
    // 计算 1-100 偶数的和
    if i%2 == 0 {
        sum+=i
    }
}
var i int = 0
for {
    // 在有些程序循环中，不知道程序执行次数，只有条件满足时程序停止
    if i >= 5 {
        // 跳出语句跳出当前循环
        break
    }
    fmt.Println(i)
    i++
}
sum := 0
for i := 0; i <= 100; i++ {
    if i%2 == 1 {
        // 结束本次循环，继续下次循环
        // 如果在程序中入到 continue 后剩余代码不会执行，会回到循环的位置
        continue
    }
    sum += i
}
fmt.Println(sum)
// 多重赋值
a := []int{1, 2, 3, 4, 5, 6} 
for i, j := 0, len(a) – 1; i < j; i, j = i + 1, j – 1 { 
    a[i], a[j] = a[j], a[i] 
}
fmt.Println(a)
// 迭代集合（数组、切片、字典）
for k, v := range a {
    fmt.Println(k, v)
}
// 忽略索引 / 键
for _, v := range a {
    fmt.Println(v)
}
// 忽略值
for k := range a {
    fmt.Println(k)
}
// 基于判断的循环，类似 while
sum := 0
i := 0
for i < 100 {
    i++
    sum += i
}
fmt.Println(sum)
/*
注意：
和条件语句、分支语句一样，左花括号 { 必须与 for 处于同一行；
不支持 whie 和 do-while 结构的循环语句；
可以通过 for-range 结构对可迭代集合进行遍历；
支持基于条件判断进行循环迭代；
允许在循环条件中定义和初始化变量，且支持多重赋值；
Go 语言的 for 循环同样支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break，可以选择中断哪一个循环，如下例：
*/
JLoop: 
for j := 0; j < 5; j++ { 
    for i := 0; i < 10; i++ { 
        if i > 5 { 
            break JLoop
        }
        fmt.Println(i)
    } 
} 

// goto 会跳到所定义的标志位
goto FLAG
fmt.Println("hello world3")
fmt.Println("hello world4")
FLAG:
fmt.Println("hello world5")
fmt.Println("hello world6")
```


### Go 函数式编程
**函数**  
函数就是将一堆代码进行重用的一种机制。函数就是一段代码，一个函数就像一个专门做这件事的人，我们调用它来做一些事情，它可能需要我们提供一些数据给它，它执行完成后可能会有一些执行结果给我们。要求的数据就叫参数，返回的执行结果就是返回值。  
```golang
// func 函数名(参数列表)(返回值列表){
//     代码体
// }
// 函数定义，只能定义一次
// 在整个项目中函数名是唯一的，不能重名
func Add(s1 int, s2 int) {
    sum := s1 + s2
    fmt.Println(sum)
}
// 调用函数
// 注意：在调用其他包定义的函数时，只有函数名首字母大写的函数才可以被访问（Go 语言中没有 public、protected、private 之类的关键字，它是通过首字母的大小写来区分可见性）
Add(1, 2)
```

**系统内置函数**  
日常开发中的常用功能提供了很多不需要导入任何包就可以直接调用的内置函数。  
```golang
// len 与 cap
str := "golang"
println(len(str))  // 6

arr := [3]int{1, 2, 3}
print(len(arr), "\n")  // 3
print(cap(arr), "\n")  // 3

slice := arr[1:]
println(len(slice)) // 2
println(cap(slice)) // 2

dict := map[string]int{"0":1, "1":2, "2":3}
println(len(dict))  // 3

// new 与 make
p1 := new(int)     // 返回 int 类型指针，相当于 var p1 *int
p2 := new(string)  // 返回 string 类型指针
p3 := new([3]int)  // 返回数组类型指针，数组长度是 3

type Student struct {
    id int
    name string
    grade string
}
p4 := new(Student)  // 返回对象类型指针

println("p1: ", p1)
println("p2: ", p2)
println("p3: ", p3)
println("p4: ", p4)

s1 := make([]int, 3)  // 返回初始化后的切片类型值，即 []int{0, 0, 0}
m1 := make(map[string]int, 2)  // 返回初始化的字典类型值，即散列化的 map 结构

println(len(s1))  // 3
for i, v := range s1 {
    println(i, v)
}

println(len(m1))   // 0
m1["test"] = 100
for k, v := range m1 {
    println(k, v)
}
```

**普通函数传参**  
Go 语言默认使用按值传参来传递参数，也就是传递参数值的一个副本。  
如果你想要实现在函数中修改形参值可以同时修改实参值，需要通过引用传参来完成，此时传递给函数的参数是一个指针，而指针代表的是实参的内存地址，修改指针引用的值即修改变量内存地址中存储的值，所以实参的值也会被修改（这种情况下，传递的是变量地址值的拷贝，所以从本质上来说还是按值传参）。  
`注意：在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型默认使用引用传参`  
```golang
// 按值传参
func add(a, b int) int  {
    a *= 2
    b *= 3
    return a + b
}

func main()  {
    x, y := 1, 2
    z := add(x, y)
    fmt.Printf("add(%d, %d) = %d\n", x, y, z)  // add(1, 2) = 8
}

// 引用传参
func add(a, b *int) int {
    *a *= 2
    *b *= 3
    return *a + *b
}

func main()  {
    x, y := 1, 2
    z := add(&x, &y)
    fmt.Printf("add(%d, %d) = %d\n", x, y, z) // add(2, 6) = 8
}

// 变长参数（同一类型）
func myfunc(numbers ...int) {
    for _, number := range numbers {
        fmt.Println(number)
    }
}
myfunc(1, 2, 3, 4, 5) 
// 变长参数还支持传递一个 []int 类型的切片，传递切片时需要在末尾加上 ... 作为标识，表示对应的参数类型是变长参数
slice := []int{1, 2, 3, 4, 5}
myfunc(slice...)
myfunc(slice[1:3]...) // 类型 ...type 本质上是一个切片，也就是 []type
// 任意类型的变长参数（泛型）
// 指定变长参数类型为 interface{}
func myPrintf(args ...interface{}) { // interface{} 是一个空接口，可以用于表示任意类型
    for _, arg := range args {
        // 通过反射获取类型
        // 在运行时通过反射对数据类型进行检查，以便让程序在预设的轨道内运行，避免因为类型问题导致程序崩溃
        switch reflect.TypeOf(arg).Kind() {
        case reflect.Int:
            fmt.Println(arg, "is an int value.")
        case reflect.String:
            fmt.Printf("\"%s\" is a string value.\n", arg)
        case reflect.Array:
            fmt.Println(arg, "is an array type.")
        default:
            fmt.Println(arg, "is an unknown type.")
        }
    }
}
myPrintf(1, "1", [1]int{1}, true)


// 多返回值
func add(a, b *int) (int, error) {
    if *a < 0 || *b < 0 {
        err := errors.New("只支持非负整数相加")
        return 0, err
    }
    *a *= 2
    *b *= 3
    // 通过 error 指定多返回一个表示错误信息的、类型为 error 的返回值，函数的多个返回值之间可以通过逗号分隔，并且在最外面通过圆括号包起来
    return *a + *b, nil
}
x, y := -1, 2
z, err := add(&x, &y)
if err != nil {
    fmt.Println(err.Error()) // 只支持非负整数相加
    return
}
fmt.Printf("add(%d, %d) = %d\n", x, y, z) 
// 返回值支持命名（不推荐）
// 这种机制避免了每次进行 return 操作时都要关注函数需要返回哪些返回值，为开发者节省了精力，尤其是在复杂的函数中
func add(a, b *int) (c int, err error) {
    if *a < 0 || *b < 0 {
        err = errors.New("只支持非负整数相加")
        return
    }
    *a *= 2
    *b *= 3
    c = *a + *b
    return
}
```

**匿名函数与闭包**  
匿名函数是一种没有指定函数名的函数声明方式。  

所谓闭包指的是引用了自由变量（未绑定到特定对象的变量，通常在函数外定义）的函数，被引用的自由变量将和这个函数一同存在，即使已经离开了创造它的上下文环境也不会被释放（比如传递到其他函数或对象中）。简单来说，「闭」的意思是「封闭外部状态」，即使外部状态已经失效，闭包内部依然保留了一份从外部引用的变量。  
显然，闭包只能通过匿名函数实现，我们可以把闭包看作是有状态的匿名函数，反过来，如果匿名函数引用了外部变量，就形成了一个闭包（Closure）。  
闭包的价值在于可以作为持有外部变量的函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示数据还要表示代码。支持闭包的语言都将函数作为第一类对象（firt-class object，有的地方也译作第一级对象、一等公民等，都是一个意思），Go 语言也不例外，这意味 Go 函数和普通 Go 数据类型（整型、字符串、数组、切片、字典、结构体等）具有同等的地位，可以赋值给变量，也可以作为参数传递给其他函数，还能够被函数动态创建和返回。  
> 注：所谓第一类对象指的是运行期可以被创建并作为参数传递给其他函数或赋值给变量的实体，在绝大多数语言中，数值和基本类型都是第一类对象，在支持闭包的编程语言中（比如 Go、PHP、JavaScript、Python 等），函数也是第一类对象，而像 C、C++ 等不支持匿名函数的语言中，函数不能在运行期创建，所以在这些语言中，函数不是不是第一类对象。  

```golang
// 匿名函数的支持
func(a, b int) int { 
    return a + b
}
// 1、将匿名函数赋值给变量
add := func(a, b int) int {
    return a + b
}
// 调用匿名函数 add
fmt.Println(add(1, 2))  
// 2、定义时直接调用匿名函数
func(a, b int) {
    fmt.Println(a + b)
} (1, 2) 

// 匿名函数的典型使用场景 - 保证局部变量的安全性
var j int = 1
f := func() {
    var i int = 1 // 闭包引用了局部变量 i 和 j，i 在闭包内部定义，其值被隔离，不能从外部修改
    fmt.Printf("i, j: %d, %d\n", i, j)
}
f() // i, j: 1, 1
j += 2 // 变量 j 在闭包外部定义，所以可以从外部修改，闭包持有的只是其引用
f() // i, j: 1, 3

// 匿名函数的典型使用场景 - 将匿名函数作为函数参数
add := func(a, b int) int {
    return a + b
}
// 将函数类型作为参数
func(call func(int, int) int) {
    fmt.Println(call(1, 2))
}(add)
// 将第二个匿名函数提取到 main 函数外，成为一个具名函数 handleAdd，然后定义不同的加法算法实现函数，并将其作为参数传入 handleAdd
// 通过一个函数执行多种不同加法实现算法，提升了代码的复用性，我们可以基于这个功能特性实现一些更复杂的业务逻辑，比如 Go 官方 net/http 包底层的路由处理器也是这么实现的
func main() {
    ...

    // 普通的加法操作
    add1 := func(a, b int) int {
        return a + b
    }

    // 定义多种加法算法
    base := 10
    add2 := func(a, b int) int {
        return a * base + b
    }

    handleAdd(1, 2, add1) // 3
    handleAdd(1, 2, add2) // 12
}
// 将匿名函数作为参数
func handleAdd(a, b int, call func(int, int) int) {
    fmt.Println(call(a, b))
}

// 匿名函数的典型使用场景 - 将匿名函数作为函数返回值
// 将函数作为返回值类型
func deferAdd(a, b int) func() int {
    return func() int {
        return a + b
    }
}
func main() {
    ...

    // 此时返回的是匿名函数
    addFunc := deferAdd(1, 2)
    // 这里才会真正执行加法操作
    fmt.Println(addFunc())
}
// 调用 deferAdd 函数返回的是一个匿名函数，但是这个匿名函数引用了外部函数传入的参数，因此形成闭包，只要这个闭包存在，这些持有的参数变量就一直存在，即使脱离了 deferAdd 函数的作用域，依然可以访问它们
// 另外调用 deferAdd 方法时并没有执行闭包，只有运行 addFunc() 时才会真正执行闭包中的业务逻辑（这里是加法运算），因此，我们可以通过将函数返回值声明为函数类型来实现业务逻辑的延迟执行，让执行时机完全掌握在开发者手中
```

**通过高阶函数实现装饰器模式**  
所谓高阶函数，就是接收其他函数作为参数传入，或者把其他函数作为结果返回的函数。  

装饰器模式（Decorator）是一种软件设计模式，其应用场景是为某个已经存在的功能模块（类或者函数）添加一些「装饰」功能，而又不会侵入和修改原有的功能模块。  
有过 Python、Java 编程经验的同学应该对这个模式很熟悉，在 Python、Java 中，我们可以通过注解非常优雅地实现装饰器模式，比如给某个功能模块添加日志功能、或者为路由处理器添加中间件功能，这些都可以通过装饰器实现。不过 Go 语言的设计哲学就是简单，没有提供「注解」之类的语法糖，在函数式编程中，要实现装饰器模式，可以借助高阶函数来实现。  

核心思路就是在被修饰的功能模块执行前后加上一些额外的业务逻辑，而又不影响原有功能模块的执行。显然，装饰器模式是遵循 SOLID 设计原则中的开放封闭原则的 —— 对代码扩展开放，对代码修改关闭。  

如下，原有的代码逻辑不需要做任何变动，只需要新增一个位运算版乘法实现函数 multiply2，然后套上装饰器函数 execTime 计算耗时。  
```golang
package main

import (
    "fmt"
    "time"
)

// 为函数类型设置别名提高代码可读性
type MultiPlyFunc func(int, int) int

// 乘法运算函数1（算术运算）
func multiply1(a, b int) int {
    return a * b
}

// 乘法运算函数2（位运算）
func multiply2(a, b int) int {
    return a << b
}

// 通过高阶函数在不侵入原有函数实现的前提下计算乘法函数执行时间
func execTime(f MultiPlyFunc) MultiPlyFunc {
    return func(a, b int) int {
        start := time.Now() // 起始时间
        c := f(a, b)  // 执行乘法运算函数
        end := time.Since(start) // 函数执行完毕耗时
        fmt.Printf("--- 执行耗时: %v ---\n", end)
        return c  // 返回计算结果
    }
}

func main() {
    a := 2
    b := 8
    fmt.Println("算术运算：")
    decorator1 := execTime(multiply1)
    c := decorator1(a, b)
    fmt.Printf("%d x %d = %d\n", a, b, c)

    fmt.Println("位运算：")
    decorator2 := execTime(multiply2)
    a = 1
    b = 4
    c = decorator2(a, b)
    fmt.Printf("%d << %d = %d\n", a, b, c)
}
```

**递归函数**  
递归函数指的是在函数内部调用函数自身的函数，从数学解题思路来说，递归就是把一个大问题拆分成多个小问题，再各个击破，在实际开发过程中，某个问题满足以下条件就可以通过递归函数来解决：  
> 一个问题的解可以被拆分成多个子问题的解  
> 拆分前的原问题与拆分后的子问题除了数据规模不同，求解思路完全一样  
> 子问题存在递归终止条件  

`注意：编写递归函数时，这个递归一定要有终止条件，否则就会无限调用下去，直到内存溢出`  
```golang
// 实现斐波那契
func fibonacci(n int) int {
    if n == 1 {
        return 0
    }
    if n == 2 {
        return 1
    }

    return fibonacci(n-1) + fibonacci(n-2)
}
n := 5
num := fibonacci(n)
fmt.Printf("The %dth number of fibonacci sequence is %d\n", n, num) // The %dth number of fibonacci sequence is 3

// 通过内存缓存技术优化递归函数性能（内存缓存技术 - 优化计算成本相对昂贵的函数调用时非常有用）
const MAX = 50
// 通过预定义数组 fibs 保存已经计算过的斐波那契序号对应的数值
var fibs [MAX]int
func fibonacci(n int) int {
    if n == 1 {
        return 0
    }

    if n == 2 {
        return 1
    }

    index := n - 1
    if fibs[index] != 0 {
        return fibs[index]
    }

    num := fibonacci(n-1) + fibonacci(n-2)
    fibs[index] = num
    return num
}
```
函数调用底层是通过栈来维护的，对于递归函数而言，如果层级太深，同时保存成百上千的调用记录，会导致这个栈越来越大，消耗大量内存空间，严重情况下会导致栈溢出（stack overflow），为了优化这个问题，可以引入*尾递归优化技术*来重用栈，降低对内存空间的消耗，提升递归函数性能。  
在计算机科学里，*尾调用*是指一个函数的最后一个动作是调用一个函数（只能是一个函数调用，不能有其他操作，比如函数相加、乘以常量等）。该调用位置为尾位置，若这个函数在尾位置调用自身，则称这种情况为*尾递归*，它是尾调用的一种特殊情形。尾调用的一个重要特性是它不是在函数调用栈上添加一个新的堆栈帧 —— 而是更新它，尾递归自然也继承了这一特性，这就使得原来层层递进的调用栈变成了线性结构，因而可以极大优化内存占用，提升程序性能，这就是尾递归优化技术。  
以计算斐波那契数列的递归函数为例，简单来说，就是处于函数尾部的递归调用前面的中间状态都不需要再保存了，这可以节省很大的内存空间，在此之前的代码实现中，递归调用 fibonacci(n-1) 时，还有 fibonacci(n-2) 没有执行，因此需要保存前面的中间状态，内存开销很大。  
一些编程语言的编译器提供了对尾递归优化的支持，但是 Go 目前并不支持，为了使用尾递归优化技术，需要手动编写实现代码。  
尾递归的实现需要重构之前的递归函数，确保最后一步只调用自身，要做到这一点，就要把所有用到的内部变量 / 中间状态变成函数参数。  
```golang
func fibonacci(n int) int {
    return fibonacciTail(n, 0, 1) // F(1) = 0, F(2) = 1
}
// 当前 first + second 的和赋值给下次调用的 second 参数，当前 second 值赋值给下次调用的 first 参数，就等同于实现了 F(n) = F(n-1) + F(n-2) 的效果，循环往复，不断累加，直到 n 值等于 1（F (1) = 0，无需继续迭代下去），则返回 first 的值，也就是最终的 F(n) 的值
// 简单来说，就是把原来通过递归调用计算结果转化为通过外部传递参数初始化，再传递给下次尾递归调用不断累加，这样就可以保证 fibonacciTail 调用始终是线性结构的更新，不需要开辟新的堆栈保存中间函数调用
func fibonacciTail(n, first, second int) int {
    if n < 2 {
        return first
    }
    return fibonacciTail(n-1, second, first+second)
}
```

**Map-Reduce-Filter 模式处理集合元素**  
日常开发过程中，要处理数组、切片、字典等集合类型，常规做法都是循环迭代进行处理。比如将一个字典类型用户切片中的所有年龄属性值提取出来，然后求和，常规实现是通过循环遍历所有切片，然后从用户字典键值对中提取出年龄字段值，再依次进行累加，最后返回计算结果。  
在函数式编程中，我们可以通过 Map-Reduce 技术让这个功能实现变得更优雅，代码复用性更好。  
Map-Reduce 并不是一个整体，而是要分两步实现：Map 和 Reduce，Map-Reduce 模型：先将字典类型切片转化为一个字符串类型切片（Map，字面意思就是映射），再将转化后的切片元素转化为整型后累加起来（Reduce，字面意思就是将多个集合元素通过迭代处理减少为一个）。
```golang
// 常规做法
func ageSum(users []map[string]string) int {
    var sum int
    for _, user := range users {
        num, _ := strconv.Atoi(user["age"])
        sum += num
    }
    return sum
}
var users = []map[string]string{
    {
        "name": "张三",
        "age": "18",
    },
    {
        "name": "李四",
        "age": "22",
    },
    {
        "name": "王五",
        "age": "20",
    },
}
fmt.Printf("用户年龄累加结果: %d\n", ageSum(users)) // 用户年龄累加结果: 60

// Map-Reduce 模式
// Map 映射转化函数
func mapToString(items []map[string]string, f func(map[string]string) string) []string {
    newSlice := make([]string, len(items))
    for _, item := range items {
        newSlice = append(newSlice, f(item))
    }
    return newSlice
}
// Reduce 求和函数
func fieldSum(items []string, f func(string) int) int {
    var sum int
    for _, item := range items{
        sum += f(item)
    }
    return sum
}
// 调用
ageSlice := mapToString(users, func(user map[string]string) string {
    return user["age"]
})
sum := fieldSum(ageSlice, func(age string) int {
    intAge, _ := strconv.Atoi(age)
    return intAge
})
fmt.Printf("用户年龄累加结果: %d\n", sum)
```
为了让 Map-Reduce 代码更加健壮（排除无效的字段值），或者只对指定范围的数据进行统计计算，还可以在 Map-Reduce 基础上引入 Filter（过滤器），对集合元素进行过滤。
```golang
func itemsFilter(items []map[string]string, f func(map[string]string) bool) []map[string]string {
    newSlice := make([]map[string]string, len(items))
    for _, item := range items {
        if f(item) {
            newSlice = append(newSlice, item)
        }
    }
    return newSlice
}
var users = []map[string]string{
    {
        "name": "张三",
        "age": "18",
    },
    {
        "name": "李四",
        "age": "22",
    },
    {
        "name": "王五",
        "age": "20",
    },
    {
        "name": "赵六",
        "age": "-10",
    },
    {
        "name": "孙七",
        "age": "60",
    },
    {
        "name": "周八",
        "age": "10",
    },
}

validUsers := itemsFilter(users, func(user map[string]string) bool {
    age, ok := user["age"]
    if !ok {
        return false
    }
    intAge, err := strconv.Atoi(age)
    if err != nil {
         return false
    }
    if intAge < 18 || intAge > 35 {
        return false
    }
    return true
})
ageSlice := mapToString(validUsers, func(user map[string]string) string {
    return user["age"]
})
sum := fieldSum(ageSlice, func(age string) int {
    intAge, _ := strconv.Atoi(age)
    return intAge
})
fmt.Printf("用户年龄累加结果: %d\n", sum)
```

**基于管道技术实现函数的流式调用**  
管道（Pipeline）这一术语来源是 Unix 的 Shell 命令行，我们可以使用管道连接符 | 通过组合简单的命令实现强大的功能。
```bash
ps -ef | grep nginx 
```
在函数式编程中，我们也可以借助管道的思想串联一些简单的函数构建更加强大的功能，比如最常见的流式函数调用（水流一样，在面向对象编程中对应的是流接口模式，可以实现链式处理）。  
这样一来，每个函数就可以专注于自己要处理的事情，把它做到极致，然后通过组合方式（管道）构建更加复杂的业务功能，这也是符合 SOLID 设计原则的单一职责原则。  
```golang
// 通过管道模式实现 Map-Reduce-Filter 模式处理集合元素的流式调用
package main

import (
    "log"
)

type user struct {
    name string
    age  int
}

func filterAge(users []user) interface{} {
    var slice []user
    for _, u := range users {
        if u.age >= 18 && u.age <= 35 {
            slice = append(slice, u)
        }
    }
    return slice
}

func mapAgeToSlice(users []user) interface{} {
    var slice []int
    for _, u := range users {
        slice = append(slice, u.age)
    }
    return slice
}

func sumAge(users []user, pipes ...func([]user) interface{}) int {
    var ages []int
    var sum int
    // 由于这些处理函数的返回值类型被声明为了空接口，所以需要在运行时动态对它们的返回值类型做检测
    for _, f := range pipes {
        result := f(users)
        switch result.(type) {
        case []user:
            users = result.([]user)
        case []int:
            ages = result.([]int)
        }
    }
    if len(ages) == 0 {
        log.Fatalln("没有在管道中加入 mapAgeToSlice 方法")
    }
    for _, age := range ages {
        sum += age
    }
    return sum
}
var users = []user{
    {
        name: "张三",
        age: 18,
    },
    {
        name: "李四",
        age: 22,
    },
    {
        name: "王五",
        age: 20,
    },
    {
        name: "赵六",
        age: -10,
    },
    {
        name: "孙七",
        age: 60,
    },
    {
        name: "周八",
        age: 10,
    },
}
sum := sumAge(users, filterAge, mapAgeToSlice)
log.Printf("用户年龄累加结果: % d\n", sum)
```


### Go 面向对象
Go 语言面向对象编程设计得简洁而优雅。  
简洁之处在于，Go 语言并没有沿袭传统面向对象编程中的诸多概念，比如类的继承、接口的实现、构造函数和析构函数、隐藏的 this 指针等，也没有 public、protected、private 之类的访问修饰符。  
优雅之处在于，Go 语言对面向对象编程的支持是语言类型系统中的天然组成部分，整个类型系统通过接口串联，浑然一体。  

**类型系统**  
类型系统才是一门编程语言的地基，它的地位至关重要。类型系统是指一个语言的类型体系结构。一个典型的类型系统通常包含如下基本内容：  
- 基本类型，如 byte、int、bool、float、string 等；  
- 复合类型，如数组、切片、字典、指针、结构体等；  
- 可以指向任意对象的类型（Any 类型）；  
- 值语义和引用语义；  
- 面向对象，即所有具备面向对象特征（比如成员方法）的类型；  
- 接口。  

类型系统描述的是这些内容在一个语言中如何被关联。Go 语言中的大多数类型都是值语义，包括：  
- 基本类型，如布尔类型、整型、浮点型、字符串等；  
- 复合类型，如数组、结构体等（切片、字典、指针和通道都是引用语义）；  

`这里的值语义和引用语义等价于值类型和引用类型。`

**为值类型定义成员方法**  
所有值语义类型都支持定义成员方法，包括内置基本类型。  
```golang
// 需要将基本类型通过 type 关键字设置为新的类型（这是一个新类型，不是类型别名）
type Integer int

func (a Integer) Equal(b Integer) bool {
    return a == b
}

var x Integer
var y Integer
x, y = 10, 15
fmt.Println(x.Equal(y))
```
在实现某个接口时，只需要实现该接口要求的所有方法即可，无需显式声明实现的接口（实际上，Go 语言根本就不支持传统面向对象编程中的继承和实现语法）。  
```golang
type Math interface {
    Add(i Integer) Integer
    Multiply(i Integer) Integer
}
```
任何类型都可以被 Any 类型引用。在 Go 语言中，Any 类型就是空接口，即 interface{}。  

**类的定义、初始化和成员方法**  
Go 语言的面向对象编程与我们之前所熟悉的 PHP、Java 那一套完全不同，没有 class、extends、implements 之类的关键字和相应的概念，而是借助*结构体来*实现类的声明。  

Go 语言中也不支持构造函数、析构函数，取而代之地，可以通过定义形如 NewXXX 这样的全局函数（首字母大写）作为类的初始化函数。  
`注意：在 Go 语言中，未进行显式初始化的变量都会被初始化为该类型的零值，例如 bool 类型的零值为 false，int 类型的零值为 0，string 类型的零值为空字符串，float 类型的零值为 0.0`

由于 Go 语言不支持 class 这样的代码块，要为 Go 类定义成员方法，需要在 func 和方法名之间声明方法所属的类型（有的地方将其称之为接收者声明）。  
在类的成员方法中，可以通过声明的类型变量来访问类的属性和其他方法（Go 语言不支持隐藏的 this 指针，所有的东西都是显式声明）。因为 Go 语言面向对象编程不像 PHP、Java 那样支持隐式的 this 指针，所有的东西都是显式声明的，在 GetXXX 方法中，由于不需要对类的成员变量进行修改，所以不需要传入指针，而 SetXXX 方法需要在函数内部修改成员变量的值，并且该修改要作用到该函数作用域以外，所以需要传入指针类型（结构体是值类型，不是引用类型，所以需要显式传入指针）。  
需要声明的是，在 Go 语言中，当我们将成员方法 SetName 所属的类型声明为指针类型时，严格来说，该方法并不属于 Student 类，而是属于指向 Student 的指针类型，所以，归属于 Student 的成员方法只是 Student 类型下所有可用成员方法的子集，归属于 \*Student 的成员方法才是 Student 类完整可用方法的集合。  
我们在调用方法时，之所以可以直接在 student 实例上调用 SetName 方法，是因为 Go 语言底层会自动将 student 转化为对应的指针类型 &student，所以真正调用的代码是 (&student).SetName("小七2号")。  

PHP、Java 支持默认调用类的 toString 方法以字符串格式打印类的实例，Go 语言也有类似的机制，只不过这个方法名是 String。  

在 Go 语言中，有意弱化了传统面向对象编程中的类概念，这也符合 Go 语言的简单设计哲学，基于结构体定义的「类」就是和内置的数据类型一样的普通数据类型而已，内置的数据类型也可以通过 type 关键字转化为可以包含自定义成员方法的「类」。  
一个数据类型关联的所有方法，共同组成了该类型的方法集合，和其他支持面向对象编程的语言一样，同一个方法集合中的方法也不能出现重名，并且，如果它们所属的是一个结构体类型，那么它们的名称与该类型中任何字段的名称也不能重复。  
```golang
// 声明类的结构体
type Student struct {
    id uint
    name string
    male bool
    score float64
}

// 初始化函数
func NewStudent(id uint, name string, male bool, score float64) *Student {
    return &Student{id, name, male, score}
    // 初始化指定字段
    // return &Student{id: id, name:name, score:score}
}

// 定义成员方法 - 值方法（接收者类型为非指针的成员方法）
func (s Student) GetName() string  {
    return s.name
}
// 指针方法（接收者类型为指针的成员方法）
func (s *Student) SetName(name string) {
    s.name = name
}
// 以字符串格式打印类的实例
func (s Student) String() string {
    return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}",
        s.id, s.name, s.male, s.score)
}

// 使用
student := NewStudent(1, "小七", 100)
fmt.Println(student)
student.SetName("小七2号")
fmt.Println("Name:", student.GetName())
```



### Go 其他
**工程管理**  
GO 语言规定通用管理：  
为了更好的管理项目中的文件，要求将文件都要放在相应的文件夹中。  

- src 目录：用于以代码包的形式组织并保存 Go 源码文件。（比如：.go .c .h .s 等）
- pkg 目录：用于存放经由 go install 命令构建安装后的代码包（包含 Go 库源码文件）的 “.a” 归档文件。
- bin 目录：与 pkg 目录类似，在通过 go install 命令完成安装后，保存由 Go 命令源码文件生成的可执行文件。

以上目录称为工作区，工作区其实就是一个对应于特定工程的目录。  
目录 src 用于包含所有的源代码，是 Go 命令行工具一个强制的规则，而 pkg 和 bin 则无需手动创建，如果必要 Go 命令行工具在构建过程中会自动创建这些目录。  

只要配置了 gopath，同一个 packge 的方法，是可以调用的（注意：同一个目录下不能定义不同的 package）。  

包中成员以名称⾸字母⼤⼩写决定访问权限（*注意：同一个目录下不能定义不同的 package*）：  

- public: ⾸字母⼤写，可被包外访问
- private: ⾸字母⼩写，仅包内成员可以访问

```golang
// test.go
package main
import "fmt"
func Test() {
    fmt.Println("test file.")
}
// main.go
package main
import "fmt"
func main() {
    fmt.Println("main file.")
    Test()
}
```
要使用包，必须要进行导入，可以通过关键字进行 import 进行导入，它会告诉编译器你想引用该包内的代码。如果导入的是标准库中的包（GO 语言自带，例如:”fmt” 包）会在安装 Go 的位置找到。 Go 开发者创建的包会在 GOPATH 环境变量指定的目录里查找。所以，import 关键字的作用就是查找包所在的位置。  
注意：  
1、如果编译器查遍 GOPATH 也没有找到要导入的包，那么在试图对程序执行 run 或者 build 的时候就会出错；  
2、如果导入包之后，未调用其中的函数或者类型将会报出编译错误。  
```golang
// 导入单个包
import "fmt"
// 导入多个包
import (
    "users"
    "goods"
)
// 调用 src/users.go 中的方法
users.GetInfo()
```

**Go Runtime**  
有一个扩展库叫做 runtime （运行时），每一个 Go 程序都会使用它。运行时库实现了垃圾回收，并发，栈管理等重要的语言特性。尽管它对于 Go 语言很重要，但是它更类似于 C 语言的 libc 库。   
值得注意的是，Go 的运行时不包括 JVM 那样的虚拟机。Go 代码会被预先编译成原生的机器码（某些特别的编译器也可以把它编译为 JavaScript 和 WebAssembly ）。因此，尽管「运行时」这个词通常指程序运行的虚拟环境，在 Go 语言中它指的只是一个支持语言重要特性的库。  

