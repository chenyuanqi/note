
### Go 变量
变量所绑定的内存区域是要有一个明确的边界的。也就是说，通过这样一个变量，我们究竟可以操作 4 个字节内存还是 8 个字节内存，又或是 256 个字节内存，编程语言的编译器或解释器需要明确地知道。

当一个计算机程序需要调用内存空间时，对内存发出的“占位”指令，称为：“声明”。  
在Go语言中，变量或常量的数据类型必须先声明，才能使用，且无法将不相关的数据赋值给它们。
```go
// 声明 var name type
var a int = 10
// 如果存在多个变量类型相同时，可以逗号分割排列
var a, b int
// 多变量声明，声明一组不同类型的变量，使用小括号包裹住
var (
    a int = 128
    b int8 = 6
    s string = "hello"
    c rune = 'A'
    t bool = true
)
// 一行变量声明中同时声明多个变量
var a, b, c int = 5, 6, 7
var (
    a, b, c int = 5, 6, 7
    c, d, e rune = 'C', 'D', 'E'
) 

// 省略类型信息的声明
// 使用 Go 语言可简化代码，类型推断体现了这一点。当声明与赋值一并进行时，如果数据为 Go 内置的基础类型，则可无需指定类型
var b = 13
var a, b, c = 12, 'A', "hello"

// 短变量声明
// 短变量声明只能在函数内部使用
a := 12
b := 'A'
c := "hello"
// 或
a, b, c := 12, 'A', "hello"
```

**包级变量**  
通常来说，Go 语言的变量可以分为两类：一类称为包级变量 (package varible)，也就是在包级别可见的变量。如果是导出变量（大写字母开头），那么这个包级变量也可以被视为全局变量；另一类则是局部变量 (local varible)，也就是 Go 函数或方法体内声明的变量，仅在函数或方法体内可见。而我们声明的所有变量都逃不开这两种。  
包级变量只能使用带有 var 关键字的变量声明形式，不能使用短变量声明形式，但在形式细节上可以有一定灵活度。  
```go
// 声明并同时显式初始化
var ErrShortWrite = errors.New("short write")
var ErrShortBuffer = errors.New("short buffer")
var EOF = errors.New("EOF")
// 第一种方式
var a = 13 // 使用默认类型
var b int32 = 17  // 显式指定类型
var f float32 = 3.14 // 显式指定类型
// 第二种方式
var a = 13 // 使用默认类型
var b = int32(17) // 显式指定类型
var f = float32(3.14) // 显式指定类型
// 更好的方式
var (
  a = 13
  b = int32(17)
  f = float32(3.14)
)

// 声明但延迟初始化
var a int32
var f float64
```

**声明聚类与就近原则**  
我们可以将延迟初始化的变量声明放在一个 var 声明块 (比如下面的第一个 var 声明块)，然后将声明且显式初始化的变量放在另一个 var 块中（比如下面的第二个 var 声明块），这里我称这种方式为 “声明聚类”，声明聚类可以提升代码可读性。
```go
var (
    netGo  bool 
    netCgo bool 
)

var (
    aLongTimeAgo = time.Unix(1, 0)
    noDeadline = time.Time{}
    noCancel   = (chan struct{})(nil)
)
```
使用静态编程语言的开发人员都知道，变量声明最佳实践中还有一条：就近原则。也就是说我们尽可能在靠近第一次使用变量的位置声明这个变量。就近原则实际上也是对变量的作用域最小化的一种实现手段。
```go
// ErrNoCookie 这个变量在整个包中仅仅被用在了 Cookie 方法中，因此它被声明在紧邻 Cookie 方法定义的地方
var ErrNoCookie = errors.New("http: named cookie not present")
func (r *Request) Cookie(name string) (*Cookie, error) {
    for _, c := range readCookies(r.Header, name) {
        return c, nil
    }
    return nil, ErrNoCookie
}
```

**局部变量的声明形式**  
和包级变量相比，局部变量又多了一种短变量声明形式，这是局部变量特有的一种变量声明形式，也是局部变量采用最多的一种声明形式。
```go
// 对于延迟初始化的局部变量声明，我们采用通用的变量声明形式
var err error
// 对于声明且显式初始化的局部变量，建议使用短变量声明形式
a := 17
f := 3.14
s := "hello, gopher!"
// 对于不接受默认类型的变量，我们依然可以使用短变量声明形式，只是在 ":=" 右侧要做一个显式转型，以保持声明的一致性
a := int32(17)
f := float32(3.14)
s := []byte("hello, gopher!")
```

`注意：尽量在分支控制时使用短变量声明形式`  
```go
func LastIndexAny(s, chars string) int {
    if chars == "" {
        // Avoid scanning all of s.
        return -1
    }
    if len(s) > 8 {
        // 作者注：在if条件控制语句中使用短变量声明形式声明了if代码块中要使用的变量as和isASCII
        if as, isASCII := makeASCIISet(chars); isASCII { 
            for i := len(s) - 1; i >= 0; i-- {
                if as.contains(s[i]) {
                    return i
                }
            }
            return -1
        }
    }
    for i := len(s); i > 0; { 
        // 作者注：在for循环控制语句中使用短变量声明形式声明了for代码块中要使用的变量c
        r, size := utf8.DecodeLastRuneInString(s[:i])
        i -= size
        for _, c := range chars {
            if r == c {
                return i
            }
        }
    }
    return -1
}
```
虽然良好的函数 / 方法设计都讲究 “单一职责”，所以每个函数 / 方法规模都不大，很少需要应用 var 块来聚类声明局部变量，但是如果你在声明局部变量时遇到了适合聚类的应用场景，你也应该毫不犹豫地使用 var 声明块来声明多于一个的局部变量。  
```go
func (r *Resolver) resolveAddrList(ctx context.Context, op, network, 
                            addr string, hint Addr) (addrList, error) {
    ... ...
    var (
        tcp      *TCPAddr
        udp      *UDPAddr
        ip       *IPAddr
        wildcard bool
    )
   ... ...
}
```

### 代码块和作用域
从 Go 变量遮蔽（Variable Shadowing）的问题说起。  
```go
var a = 11

func foo(n int) {
  a := 1
  a += n
}

func main() {
  fmt.Println("a =", a) // 11
  foo(5)
  fmt.Println("after calling foo, a =", a) // 11
}
```
变量遮蔽是 Go 开发人员在日常开发工作中最容易犯的编码错误之一，它低级又不容易查找，常常会让你陷入漫长的调试过程。变量遮蔽只是个引子，真正想说的是代码块（Block，也可译作词法块）和作用域（Scope）这两个概念。  
Go 语言中的代码块是包裹在一对大括号内部的声明和语句序列，如果一对大括号内部没有任何声明或其他语句，我们就把它叫做空代码块。Go 代码块支持嵌套，我们可以在一个代码块中嵌入多个层次的代码块。  
```go
func foo() { //代码块1
    { // 代码块2
        { // 代码块3
            { // 代码块4

            }
        }
    }
}
```
像代码块 1 到代码块 4 这样的代码块，它们都是由两个肉眼可见的且配对的大括号包裹起来的，我们称这样的代码块为显式代码块（Explicit Blocks）。  
隐式代码块没有显式代码块那样的肉眼可见的配对大括号包裹，我们无法通过大括号来识别隐式代码块。首当其冲的就是位于最外层的宇宙代码块（Universe Block），它囊括的范围最大，所有 Go 源码都在这个隐式代码块中；在宇宙代码块内部嵌套了包代码块（Package Block），每个 Go 包都对应一个隐式包代码块，每个包代码块包含了该包中的所有 Go 源码，不管这些代码分布在这个包里的多少个的源文件中；再往里面看，在包代码块的内部嵌套着若干文件代码块（File Block），每个 Go 源文件都对应着一个文件代码块，也就是说一个 Go 包如果有多个源文件，那么就会有多个对应的文件代码块；再下一个级别的隐式代码块就在控制语句层面了，包括 if、for 与 switch，我们可以把每个控制语句都视为在它自己的隐式代码块里；最后，位于最内层的隐式代码块是 switch 或 select 语句的每个 case/default 子句中，虽然没有大括号包裹，但实质上，每个子句都自成一个代码块。  
```go
{
	// 宇宙代码块
	{
		// 文件代码块
		{
			import xxx
		}
		// 包代码块
		{
			// switch 语句隐式代码块
			switch {
				case x: 
				{
					// case 子句代码块
				} 
			}
		}
	}
}
```

作用域的概念是针对标识符的，不局限于变量。每个标识符都有自己的作用域，而一个标识符的作用域就是指这个标识符在被声明后可以被有效使用的源码区域。  
显然，作用域是一个编译期的概念，也就是说，编译器在编译过程中会对每个标识符的作用域进行检查，对于在标识符作用域外使用该标识符的行为会给出编译错误的报错。不过，我们可以使用代码块的概念来划定每个标识符的作用域。这个划定原则是什么呢？原则就是*声明于外层代码块中的标识符，其作用域包括所有内层代码块*。而且，这一原则同时适于显式代码块与隐式代码块。  
```go
func (t T) M1(x int) (err error) {
    // 代码块1
    m := 13

    // 代码块1是包含m、t、x和err三个标识符的最内部代码块
    { // 代码块2
        
        // "代码块2"是包含类型bar标识符的最内部的那个包含代码块
        type bar struct {} // 类型标识符bar的作用域始于此
        { // 代码块3
            
            // "代码块3"是包含变量a标识符的最内部的那个包含代码块
            a := 5 // a作用域开始于此
            {  // 代码块4 
                //... ...
            }
            // a作用域终止于此
        }
        // 类型标识符bar的作用域终止于此
    }
    // m、t、x和err的作用域终止于此
}


func bar() {
    { // 等价于第一个if的隐式代码块
        a := 1 // 变量a作用域始于此
        if false {

        } else {
            { // 等价于第一个else if的隐式代码块
                b := 2 // 变量b的作用域始于此
                if false {

                } else {
                    { // 等价于第二个else if的隐式代码块
                        c := 3 // 变量c作用域始于此
                        if false {

                        } else {
                            println(a, b, c)
                        }
                        // 变量c的作用域终止于此
                    }
                }
                // 变量b的作用域终止于此
            }
        }
        // 变量a作用域终止于此
    }
}
```

变量遮蔽问题的根本原因，就是内层代码块中声明了一个与外层代码块同名且同类型的变量，这样，内层代码块中的同名变量就会替代那个外层变量，参与此层代码块内的相关计算，我们也就说内层变量遮蔽了外层同名变量。  
Go 官方提供了 go vet 工具可以用于对 Go 源码做一系列静态检查，在 Go 1.14 版以前默认支持变量遮蔽检查，Go 1.14 版之后，变量遮蔽检查的插件就需要我们单独安装了。  
```bash
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
go vet -vettool=$(which shadow) -strict complex.go 
```
