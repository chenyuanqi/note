
### Go 打印
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
