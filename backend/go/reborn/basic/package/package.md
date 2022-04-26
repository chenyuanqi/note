
### Go 包
在 Go 源码中，package 的意思就是包，后面跟着的就是包名。Go 语言通过包来组织源码，拥有相同包名的 Go 源码属于同一个包。反过来，一个包内通常会包含一个或多个 Go 源码文件。“封装”和“复用”等就可以用包来实现。  
Go 语言有一个强制性要求，就是源码文件的第一行有效代码必须声明自己所在的包。  
需要特别指出的是：main 包是一个比较特殊的包。一个 Go 程序必须有main包，且只能有一个 main 包。  

和声明相对的，是导入。用通俗的话讲，包的声明就是要告诉大家：“我属于哪个包”；包的导入就是要提出要求：“我想要使用哪个包”。   
[Go 标准库参考](https://pkg.go.dev/std)
```go
// package 声明
package main

// package 导入
import "fmt"
import "github.com/miaogaolin/gobasic/pkgA"

// 多包导入
import (
   "fmt"
   "github.com/miaogaolin/gobasic/pkgA"
)

// 简化导入
import . "fmt"
// 不使用“点”导入
fmt.Println()
// 使用“点”导入
Println()

// 别名导入（如果导入的多个包时，名称一样出现冲突时，就可以取个别名；名称特别长时也可以考虑取别名）
import a "exmaple/pkgA"
// a 为别名
a.Func1()

// 匿名导入（导入包时，如果该包没有被使用，那编译器就会报错；为了不让报错，可以使用匿名导入；为何不直接删除呢？是因为想使用包内的 init() 函数，该函数在包被导入时自动调用）
import _ "github.com/go-sql-driver/mysql" // 该函数的意义表示注册 mysql 驱动
```

**main()函数**  
在 Go 语言中，main() 函数是程序的入口函数，它位于 main 包中。如果想要编译生成可执行文件，main() 函数则是必须的。如果将示例代码中的 main() 函数去掉直接编译，可以看到控制台会输出如下错误：  
`runtime.main_main·f: function main is undeclared in the main package，即 main() 函数没有在 main 包中声明`  
```go
func main(){
    fmt.Println("Hello World!")
}
```

**internal 目录**  
Go 语言中一个特殊的目录，如果源文件在 internal 目录中，那该目录的父级目录是不能访问 internal 目录下的内容的。如果 internal 目录下定义的包名不是 internal 名称，外部也是不能访问的。只要目录名称不是 internal（就算包名是），外部就能访问。

**Go源码的启动流程**  
main() 函数是 Go 程序的入口函数。实际上，Go 程序还有一个 init() 函数，被称为“初始化”函数。  
```go
package main
import "fmt"
func init() {
   fmt.Println("Hello")
}
func main() {
   fmt.Println("World")
}
// Hello
// World
```

Go 源码的启动流程是这样的：
- 程序开始运行后，首先来到 main 包，检索所有导入的包。发现代码中导入了 A 包，于是来到 A 包；
- 发现 A 包代码中导入了 B 包，于是又来到 B 包；
- B 包代码没有导入任何其它的包，于是开始声明 B 包内的常量和变量，并执行 B 包中的 init() 函数；
- 回到 A 包，进行 A 包内的常量和变量的声明，并执行 A 包中的 init() 函数；
- 回到 main 包，执行 main 包内的常量和变量的声明，并执行 main 包中的 init() 函数；
- 执行 main 包中的 main() 函数。  

`💡 提示：了解 Go 源码的启动加载过程，有助于编写更高效率的代码，排查程序启动缓慢等性能问题。`  
[![Go 源码的启动流程](https://s1.ax1x.com/2022/04/20/LsCH10.png)](https://imgtu.com/i/LsCH10)

### Go 自定义包
Go 语言借助文件系统树形结构来组织包。具体来说，
- 虽然 Go 语法没有强制要求包名与其所在的目录名相同，但习惯上我们还是会保持这二者相同 ；
- 包可以定义在多层级的目录中；
- 单个包的所有源码应存在相同的目录下，不同目录通常包含不同的包源码；
- 包名一般开头是小写的，采用小驼峰式命名法；
- 多个类似业务的公司可能会封装相同名称的包，为了确保唯一性，建议大家使用域名作为目录结构的一部分 。

```go
// 声明这个源码文件属于weather包
package weather
import (
   "fmt"
   "io/ioutil"
   "net/http"
)
func CurrentWeather(cityCode string) string {
   //使用net包发起Get请求
   resp, err := http.Get("https://devapi.qweather.com/v7/weather/now?location=" + cityCode + "&key=[您自己申请的AppKey]")
   if err != nil {
      fmt.Println("HTTP请求失败：", err)
      panic(err)
   }
   //使用断言关闭网络请求
   defer resp.Body.Close()
   //使用ioutil工具包获取服务端响应数据
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      fmt.Println("读取网络响应失败：", err)
      panic(err)
   }
   return string(body)
}

// 回到main.go中，尝试调用weather包中的函数CurrentWeather()
package main
import (
   "fmt"
   "go_juejin_weather/juejin.cn/weather"
)
func main() {
   result := weather.CurrentWeather("101010100")
   fmt.Println(result)
}
```
