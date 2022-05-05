
### byte 和 rune 字符类型  
字符串中的每一个元素叫做 “字符”，在遍历或者单个获取字符串元素时可以获得字符。

Go 语言的字符有以下两种：
- 一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
- 另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。

byte 类型是 uint8 的别名，对于只占用 1 个字节的传统 ASCII 编码的字符来说，完全没有问题，例如 var ch byte = 'A'，字符使用单引号括起来。
```go
// 在 ASCII 码表中，A 的值是 65，使用 16 进制表示则为 41
var ch byte = 65 
// 或 var ch byte = '\x41'      //（\x 总是紧跟着长度为 2 的 16 进制数）
// 或 var ch byte = '\377' // (\3 后面紧跟着长度为 3 的八进制数)
```

Go 语言同样支持 Unicode（UTF-8），因此字符同样称为 Unicode 代码点或者 runes，并在内存中使用 int 来表示。在文档中，一般使用格式 U+hhhh 来表示，其中 h 表示一个 16 进制数。在书写 Unicode 字符时，需要在 16 进制数之前加上前缀 \u 或者 \U。因为 Unicode 至少占用 2 个字节，所以我们使用 int16 或者 int 类型来表示。如果需要使用到 4 字节，则使用 \u 前缀，如果需要使用到 8 个字节，则使用 \U 前缀。
```go
var ch int = '\u0041'
var ch2 int = '\u03B2'
var ch3 int = '\U00101234'
fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
// 格式化说明符 %c 用于表示字符，%v 或 %d 会输出用于表示该字符的整数，%U输出格式为 U+hhhh 的字符串
fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
// 65 - 946 - 1053236
// A - β - r
// 41 - 3B2 - 101234
// U+0041 - U+03B2 - U+101234
```

Unicode 包中内置了一些用于测试字符的函数，这些函数的返回值都是一个布尔值，如下所示（其中 ch 代表字符）：
- 判断是否为字母：unicode.IsLetter (ch)
- 判断是否为数字：unicode.IsDigit (ch)
- 判断是否为空白符号：unicode.IsSpace (ch)

**UTF-8 和 Unicode 有何区别？**  
Unicode 与 ASCII 类似，都是一种字符集。  
字符集为每个字符分配一个唯一的 ID，我们使用到的所有字符在 Unicode 字符集中都有一个唯一的 ID，例如上面例子中的 a 在 Unicode 与 ASCII 中的编码都是 97。汉字 “你” 在 Unicode 中的编码为 20320，在不同国家的字符集中，字符所对应的 ID 也会不同。而无论任何情况下，Unicode 中的字符的 ID 都是不会变化的。

UTF-8 是编码规则，将 Unicode 中字符的 ID 以某种方式进行编码，UTF-8 的是一种变长编码规则，从 1 到 4 个字节不等。编码规则如下：
- 0xxxxxx 表示文字符号 0～127，兼容 ASCII 字符集。
- 从 128 到 0x10ffff 表示其他字符。

根据这个规则，拉丁文语系的字符编码一般情况下每个字符占用一个字节，而中文每个字符占用 3 个字节。  
广义的 Unicode 指的是一个标准，它定义了字符集及编码规则，即 Unicode 字符集和 UTF-8、UTF-16 编码等。  

