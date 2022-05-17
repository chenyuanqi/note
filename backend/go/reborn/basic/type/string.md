
### Go 字符串
虽然字符串往往被看做一个整体，但是它实际上是一片连续的内存空间，我们也可以将它理解成一个由字符组成的数组。C 语言中的字符串使用字符数组 char[] 表示。数组会占用一片连续的内存空间，而内存空间存储的字节共同组成了字符串，Go 语言中的字符串只是一个只读的字节数组。  

Go 语言只是不支持直接修改 string 类型变量的内存空间，我们仍然可以通过在 string 和 []byte 类型之间反复转换实现修改这一目的：
- 先将这段内存拷贝到堆或者栈上；
- 将变量的类型转换成 []byte 后并修改字节数据；
- 将修改后的字节数组转换回 string；

字符串在 Go 语言中的接口其实非常简单，每一个字符串在运行时都会使用如下的 reflect.StringHeader 表示，其中包含指向字节数组的指针和数组的大小：  
```go
type StringHeader struct {
	Data uintptr
	Len  int
}
```

与切片的结构体相比，字符串只少了一个表示容量的 Cap 字段，而正是因为切片在 Go 语言的运行时表示与字符串运行时表示高度相似，所以我们经常会说字符串是一个只读的切片类型。  
```go
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
```

因为字符串作为只读的类型，我们并不能直接向字符串直接追加元素改变其本身的内存空间，所有在字符串上的写入操作都是通过拷贝实现的。

**常用字符串操作**  
详见[「常用字符串操作」](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/)  
```go
// 连接字符串
"Ja" + "pan" // Japan
// 高性能连接
var b strings.Builder
b.Grow(32)
for i, p := range []int{2, 3, 5, 7, 11, 13} {
    fmt.Fprintf(&b, "%d:%d, ", i+1, p)
}
s := b.String()   // no copying
s = s[:b.Len()-2] // no copying (removes trailing ", ")
fmt.Println(s) // 1:2, 2:3, 3:5, 4:7, 5:11, 6:13

// 多行字符串
const s = `First line
Second line`
fmt.Println(s)
// First line
// Second line
// 或
const s = "\tFirst line\n" +
"Second line"
fmt.Println(s)
//    First line
// Second line
// 输出双引号
fmt.Println("\"foo\"") // "foo"

// 转义html
const s = `"Foo's Bar" <foobar@example.com>`
fmt.Println(html.EscapeString(s)) // &#34;Foo&#39;s Bar&#34; &lt;foobar@example.com&gt;
// url编码
const s = `Foo's Bar?`
fmt.Println(url.PathEscape(s)) // Foo%27s%20Bar%3F

// 等值比较
"Japan" == "Japan" // true
// 忽略大小写比较
strings.EqualFold("Japan", "JAPAN")	// true
// 字典顺序比较
"Japan" < "japan" // true

// 长度
len("日") // 字节长度 3
utf8.RuneCountInString("日") // 1 in runes
utf8.ValidString("日") // true 是否 utr-8

// 切片
"Japan"[2] // 'p'	Byte at position 2
"Japan"[1:3] // ap	Byte indexing
"Japan"[:2] // Ja	
"Japan"[2:] // pan

// 遍历字符串
for i, ch := range "Japan 日本" {
    fmt.Printf("%d:%q ", i, ch)
}
// Output: 0:'J' 1:'a' 2:'p' 3:'a' 4:'n' 5:' ' 6:'日' 9:'本'
s := "Japan 日本"
for i := 0; i < len(s); i++ {
    fmt.Printf("%q ", s[i])
}
// Output: 'J' 'a' 'p' 'a' 'n' ' ' 'æ' '\u0097' '¥' 'æ' '\u009c' '¬'
// 非 ascii 转为乱码

// 是否包含
strings.Contains("Japan", "abc") // false
strings.ContainsAny("Japan", "abc") // true 包含任意一个
strings.Index("Japan", "abc") // -1

// 替换
strings.Replace("foo", "o", ".", 2) // f..
// 转为大写
strings.ToUpper("Japan") // JAPAN
// 转为小写
strings.ToLower("Japan") // japan

// 两端去除空格/指定字符
strings.Trim("foo", "fo") // 
// 去除两端空白符
s := strings.TrimSpace("\t Goodbye hair!\n ")
fmt.Printf("%q", s) // "Goodbye hair!"
// 去除所有空白符
space := regexp.MustCompile(`\s+`)
s := space.ReplaceAllString("Hello  \t \n world!", " ")
fmt.Printf("%q", s) // "Hello world!"

// 切分
strings.Fields(" a\t b\n") // ["a" "b"]
strings.Split("a,b", ",") // ["a" "b"]
// 联合
strings.Join([]string{"a", "b"}, ":") // a:b

// 重复
strings.Repeat("da", 2) // dada

// 格式化
s := fmt.Sprintf("%.4f", math.Pi) // s == "3.1416"

// url 编码
url.PathEscape("A B") // A%20B
// html 实体化
html.EscapeString("<>")

// 字符串转浮点型
f := "3.14159265"
if s, err := strconv.ParseFloat(f, 32); err == nil {
    fmt.Println(s) // 3.1415927410125732，精度问题
}
if s, err := strconv.ParseFloat(f, 64); err == nil {
    fmt.Println(s) // 3.14159265
}
// 浮点型转字符串
s := fmt.Sprintf("%f", 123.456) // s == "123.456000"

// 整形转字符串
s := strconv.Itoa(97) // s == "97"
s := string(97) // s == "a"
var n int64 = 97
s := strconv.FormatInt(n, 10) // s == "97" (decimal)，十进制
var n int64 = 97
s := strconv.FormatInt(n, 16) // s == "61" (hexadecimal)，十六进制
// 字符串转整形
s := "97"
if n, err := strconv.Atoi(s); err == nil {
    fmt.Println(n+1) // 98
} else {
    fmt.Println(s, "is not an integer.")
}
// 转int64
n, err := strconv.ParseInt(s, 10, 64)
if err == nil {
    fmt.Printf("%d of type %T", n, n) // 97 of type int64
}
var n int = 97
m := int64(n) // safe
// 格式化转换指定长度
s := fmt.Sprintf("%+8d", 97) // s == "     +97" (width 8, right justify, always show sign)

// 接口转字符串
var x interface{} = "abc"
str := fmt.Sprintf("%v", x)
var x interface{} = []int{1, 2, 3}
str := fmt.Sprintf("%v", x)
fmt.Println(str) // "[1 2 3]"
```

**拼接字符串**  
在Go中字符串是原生类型，是只读的，所以用"+"操作符进行字符串时会创建一个新的字符串。如果在循环中使用"+"进行字符串拼接则会创建多个字符串，比如下面这样：  
```go
var s string
for i := 0; i < 1000; i++ {
    s += "a"
}

str := ""
// 1
str += "test-string"
// 2
str = fmt.Sprintf("%s%s", str, "test-string")
// 3
str = strings.Join([]string{str, "test-string"}, "")
// 4
buf := new(bytes.Buffer)
buf.WriteString("test-string")
str := buf.String()
// 5
var b []byte
s := "test-string"
b = append(b, s...)
str := string(b)
// 6
ts := "test-string"
n := 5
tsl := len(ts) * n
bs := make([]byte, tsl)
bl := 0
for bl < tsl {
    bl += copy(bs[bl:], ts)
}
str := string(bs)
// 7
var builder strings.Builder
builder.WriteString("test-string")
str := builder.String()
```
怎么更高效得进行字符串拼接呢？在早先 Go1.10 以前使用的是 bytes.Buffer，
```go
package main

import (
    "bytes"
    "fmt"
)

func main() {
    var buffer bytes.Buffer

    for i := 0; i < 1000; i++ {
        buffer.WriteString("a")
    }

    fmt.Println(buffer.String())
}
```
Go 1.10 版本后引入了一个 strings.Builder 类型。
```go
package main

import (
    "strings"
    "fmt"
)

func main() {
    var str strings.Builder

    for i := 0; i < 10; i++ {
        str.WriteString("a")
    }

    fmt.Println(str.String())
}
```

**操作包含中文的字符串**  
在 Golang 中，如果字符串中出现中文字符不能直接调用 len 函数来统计字符串字符长度，这是因为在 Go 中，字符串是以 UTF-8 为格式进行存储的，在字符串上调用 len 函数，取得的是字符串包含的 byte 的个数。

正确的做法是将字符串转换为 [\]rune，统计 [\]rune切片的长度。同样截取包含字符串也是一样，先将其转为 [\]rune，再截取后，转为 string。  
```go
package main

import (
	"fmt"
)

// 截取姓名的前四位
func NormalizeRealName(name string) (realName string) {
	realNameRune := []rune(name)
	if len(realNameRune) <= 4 {
		realName = name
		return
	}

	realName = string(realNameRune[0:4])
	return
}

func main() {
	name := "欧阳讷讷野木木"
	
	fmt.Println(NormalizeRealName(name))
}
```
针对统计字符串长度我们可以使用内置库 unicode/utf8 的 utf8.RuneCountInString(s) 方法，更方便。

rune 代表的 unicode 码点是固定的 4 个字节长度(等同于 int32)，但是很多时候常用字符的 Unicode 码点只用 2-3 个字节，这就造成了很多空间浪费所以才用的 UTF-8 这种变长字符编码。

### Go 字符串格式化
更多详阅[「字段串格式化」](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/)  

### Go 正则表达式
更多详阅[「正则表达式」](https://yourbasic.org/golang/regexp-cheat-sheet/)  

###

