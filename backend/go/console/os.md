
### 从命令行读取参数
os 包中有一个 string 类型的切片变量 os.Args，用来处理一些基本的命令行参数，它在程序启动后读取命令行输入的参数。  
```golang
// os_args.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}
```

flag 包有一个扩展功能用来解析命令行选项。但是通常被用来替换基本常量，例如，在某些情况下我们希望在命令行给常量一些不一样的值。  
```golang
// 在 flag 包中有一个 Flag 被定义成一个含有如下字段的结构体
type Flag struct {
	Name     string // name as it appears on command line
	Usage    string // help message
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message
}

// 模拟了 Unix 的 echo 功能
package main

import (
	"flag" // command line option parser
	"os"
)

var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	// flag.PrintDefaults() 打印 flag 的使用帮助信息
	flag.PrintDefaults()
	// flag.Parse() 扫描参数列表（或者常量列表）并设置 flag, flag.Arg(i) 表示第 i 个参数
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	// flag.Narg() 返回参数的数量
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}

```
