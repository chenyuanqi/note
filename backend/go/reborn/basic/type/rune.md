
### rune [ruːn]
ASCII 码只需要 7 bit 就可以完整地表示，但只能表示英文字母在内的 128 个字符，为了表示世界上大部分的文字系统，发明了 Unicode， 它是 ASCII 的超集，包含世界上书写系统中存在的所有字符，并为每个代码分配一个标准编号（称为 Unicode CodePoint），在 Go 语言中称之为 rune，是 int32 类型的别名。

Go 语言中，字符串的底层表示是 byte (8 bit) 序列，而非 rune (32 bit) 序列。例如下面的例子中 语 和 言 使用 UTF-8 编码后各占 3 个 byte，因此 len("Go语言") 等于 8，当然我们也可以将字符串转换为 rune 序列。
```go
fmt.Println(len("Go语言")) // 8
fmt.Println(len([]rune("Go语言"))) // 4
```

rune 类型是和 int32 类型等价，在所有方面都等同于 int32，按照约定，它用于区分字符值和整数值。  
rune 常用来处理 unicode 或 utf-8 字符，通常用于表示一个 Unicode 码点，这两个名称可以互换使用。即 rune 一个值代表的就是一个 Unicode 字符，它的最大特点就是可变长。它可以使用 1-4 个字节表示一个字符，根据字符的不同变换长度。所以使用 int32 类型范围就可以完美适配。 单个中文占 2 个字节，单个英文占 2 个字节。因为 Go 语言中字符串编码为 UTF-8 ，英文占 1 个字节，中文占 3 个字节。 占用空间相比之下会更大。  
```go
package main

import (
    "fmt"
)

func main() {
	// 计算中文字符
    var data = "帽儿山的枪手"
    fmt.Println("data length", len(data)) // data length 18
    fmt.Println("data word length", len([]rune(data))) // data word length 6
}
```