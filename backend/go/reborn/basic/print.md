
### Go 打印
print  
* 优点：内置函数，不需要引入额外的包，简单方便。
* 不足：无法进行格式化打印，无法完整打印复合数据结构 (如数组, Map 等)。

在 Go 语言中，标准包 fmt 有好多格式化的工具函数，函数名末尾通常以 f 结尾，列举如下：
- fmt.Printf 格式化字符串并打印到终端（标准输出）。
- fmt.Sprintf 格式化字符串并返回。
- fmt.Sscanf 从字符串中解析占位符的对应位置数据。
- fmt.Fscanf 从 io.Reader 类型中读取数据并解析占位符的对应位置数据，用于读取文件、终端（标准输入）。
- fmt.Fprintf 将格式化的字符串数据输出到 io.Writer 类型中，用于输出到文件。
- fmt.Errorf 格式化方式创建一个错误类型消息。

什么是占位符？你就从表面意思理解，就是占位置，只不过有很多不同种类的位置，而且这个位置不是谁都能坐，是有要求的，在程序层面用于格式化数据。  
占位符需要使用 % 符号表示。[更多查看](https://pkg.go.dev/fmt)  
```go
s := fmt.Sprintf("%s真帅", "老苗")
fmt.Println(s) // 老苗真帅

type  Example  struct {
    Content string
}
var  data = Example{Content: "例子"}
// %v：获取数据的值
fmt.Printf("%v", data) // {例子}
// 如果实现了 error 接口，仅表示错误消息
fmt.Printf("%v", errors.New("我错了")) // 我错了

// %+v：获取数据的值，如果结构体，会携带字段名
fmt.Printf("%+v", data) // {Content:例子}

// %#v：获取数据的值，如果是结构体，会携带结构体名和字段名
fmt.Printf("%#v", data) // main.Example{Content:"例子"}

// %T：获取数据类型
fmt.Printf("%T", data) // main.Example

// %%：字面上的一个百分号
fmt.Printf("%%") // %

// %t：bool占位符，true 或 false
fmt.Printf("%t", true) // true

// %b：二进制
fmt.Printf("%b", 4) // 100

// %c：Unicode 码转字符
fmt.Printf("%c", 0x82d7) // 苗

// %d、%5d（最小宽度为 5，右对齐，左边留白）、%-5d（左对齐，右边留白）、%05d（数字位数不足 5 位时，左边补零）：十进制整数表示
// 三个数据： 10 十进制，010 八进制，0x10 十六进制
fmt.Printf("%d,%d,%d", 10, 010, 0x10) // 10,8,16
fmt.Printf("|%5d|%-5d|%05d|", 1, 1, 1) // |    1|1 |00001|

// %o、%#o：八进制表示
fmt.Printf("%o,%o,%o", 10, 010, 0x10) // 12,10,20
// 在很多开发语言中，0 打头的数字都表示八进制。通过 %#o 输出带 0 开头
fmt.Printf("\n%#o\n", 10) // 012

// %x、%#x：十六进制表示，字母形式为小写 a-f，%#x 输出带 0x 开头
fmt.Printf("%x, %#x", 13, 13) // d, 0xd

// %X、%#X：十六进制表示，字母形式为小写 A-F，%#X 输出带 0X 开头
fmt.Printf("%X, %#X", 13, 13) // D, 0XD

// %U：转化为 Unicode 格式规范
fmt.Printf("%U", 0x82d7) // U+82D7
// %#U：转化为 Unicode 格式并带上对应的字符
fmt.Printf("%#U", 0x82d7) // U+82D7  '苗'

// %b：浮点数转化为 2 的幂的科学计数法
fmt.Printf("%b", 0.1) // 7205759403792794p-56

// %e、%E：10 的幂的科学计数法，区别：%e 与 %E 输出时的大小写之分
fmt.Printf("%e", 10.2) // 1.020000e+01
fmt.Printf("%E", 10.2) // 1.020000E+01

// % f、%.2f 等等：浮点数，%.2f 表示保留 2 位小数，%f 默认保留 6 位，%f 与 %F 等价
fmt.Printf("%f", 10.2) // 10.200000
fmt.Printf("%.2f|%.2f", 10.232, 10.235) // 10.23|10.23
// %9.2f 宽度最小为 9，包含小数位在内，精度为 2
// %9.f 或 %9f 宽度最小为 9

// %g、%.3g：根据情况选择 %e 或 %f ，但末尾去除 0
fmt.Printf("%g|%g", 10.20, 1.200000+3.400000i) // 10.2|(1.2+3.4i)
fmt.Printf("%g|%g", 2e2, 2E2) // 200|200
fmt.Printf("%.3g", 12.34) // 12.3

// %s：字符串或字节切片
fmt.Printf("%s|%s", "老苗", []byte("嘿嘿嘿")) // 老苗|嘿嘿嘿

// %q：有 Go 语言安全转义，双引号包裹
fmt.Printf("%q", "老苗") // "老苗"

// %x、%X：十六进制，%x 小写字母 a - f，%X 大写字母 A - F
fmt.Printf("%x|%X", "苗", "苗") // e88b97|E88B97

// %p、%#p：地址，使用十六进制表示，%p 带 0x，%#p 不带
num := 2
s := []int{1, 2}
fmt.Printf("%p|%p", &num, s) // 0xc00000a1d0|0xc00000a1e0

// +：打印数值的正负号，对于 %+q，只输出 ASCII 编码的字符，如果非 ASCII 编码，则转成 Unicode 编码输出
fmt.Printf("%+d|%+d", 2, -2) // +2|-2
fmt.Printf("%+q|%+q", "miao","苗") // "miao"|"\u82d7"

// -：在右侧填充空格，这块就不举例了，使用如 %-5d ，浮点 %-9.2f 也支持

// %+q 打印字符串时使用反引号包裹
fmt.Printf("%#q", "苗") // `苗`

// ‘ ‘ 空格：为正负号留出空白位置
fmt.Printf("|% d|", 2) // | 2|

// 填充前导的 0，对于数字会移到正负号之后，非数字也可使用
fmt.Printf("%05s", "a") // 0000a
fmt.Printf("%+05d", 1) // +0001

// 给字符串使用精度，用来截断字符串
fmt.Printf("%.2s", "abc") // ab
```

## println 函数

打印多个传入的参数，并自动加一个换行。

### 例子
```go
package main

func main() {
	println(1024, "hello world", true)
}
// $ go run main.go
// 输出如下 
/**
    1024 hello world true
*/
```

## print 函数

和 `println` 功能一样，但是不会自动加换行。

# 格式化打印

这里先介绍 2 个方法，分别是 `fmt` 包里面的 `Println()` 和 `Printf()`, 大多数场景下都适用。

## fmt.Println()

功能上和 `println 函数` 类似，但是可以打印复合数据结构 (如数组, Map 等)。

### 例子

```go
package main

import "fmt"

func main() {
	fmt.Println(1024, "hello world", true)
}

// $ go run main.go
// 输出如下 
/**
    1024 hello world true
*/
```

## fmt.Printf()

**最重要的格式化打印函数之一**，可以针对不同数据类型和数据结构进行打印，非常强大。

## 格式化规则

**和 `C 系列` 编程语言的 `printf()` 格式化规则差不多。**

### 通用

* `%v`   默认格式
* `%+v`  针对结构体，在 `%v` 的基础上输出结构体的键名
* `%#v`  Go 语言语法格式的值
* `%T`   Go 语言语法格式的类型和值
* `%%`   输出 `%`, 相当于转义

### 整型

* `%b`	 二进制格式 
* `%c`	 对应的 Unicode 码 
* `%d`	 十进制 
* `%o`	 八进制 
* `%O`	 八进制，加上 `0o` 前缀 
* `%q`	 Go 语言语法转义后的单引号字符 (很少使用) 例如 97 会输出 `'a'` 
* `%x`	 十六进制 (小写), 例如 `0xaf` 
* `%X`	 十六进制 (大写), 例如 `0xAF` 
* `%U`	 Unicode 例如 `"U+%04X"`

### Bool

* `%t`   true 或 false

### 浮点型

* `%b`	 指数为 2 的幂的无小数科学计数法，例如 -123456p-78 
* `%e`	 科学计数法, 例如 -1.234456e+78 
* `%E`	 科学计数法, 例如 -1.234456E+78 
* `%f`	 常规小数点表示法 (一般使用这个), 例如 123.456 
* `%F`	 和 `%f` 功能一样

### 字符串

* `%s`	 字符串
* `%q`	 将双引号 `"` 转义后的字符串
* `%x`	 将字符串作为小写的十六进制
* `%X`	 将字符串作为大写的十六进制

### 指针

* `%p`	 地址的十六进制，前缀为 `0x`

### 例子

```go
package main

import "fmt"

func main() {
	n := 1024
	fmt.Printf("n = %d\n", n) // 输出整型

	pi := 3.1415
	fmt.Printf("pi = %f\n", pi) // 输出浮点数

	str := "hello world"
	fmt.Printf("str = %s\n", str) // 输出字符串

	yes := true
	fmt.Printf("yes = %t\n", yes) // 输出布尔型

	x := 17
	fmt.Printf("yes = %b\n", x) // 输出二进制
}

// $ go run main.go
// 输出如下
/**
    n = 1024
    pi = 3.141500
    str = hello world
    yes = true
    x = 10001
*/
```

## fmt.Printf() 技巧

在打印中，如果一个变量打印多次，可以通过 `[1]` 来表示后续变量全部以第一个为准。

### 例子

```go
package main

import (
	"fmt"
)

func main() {
	n := 1024
	fmt.Printf("%T %d %v\n", n, n, n)

	fmt.Printf("%T %[1]d %[1]v\n", n) // 可以使用 [1] 来表示引用第一个变量，这样只需要一个变量就可以了
}
// $ go run main.go
// 输出如下
/**
    int 1024 1024
    int 1024 1024
*/
```