
### Go语言中的关键字和保留字
**Go语言中的关键字**  
| 包管理 | 声明与定义 | 流程控制 |  
| :---: | :---: | :---: |  
| import、package | chan、const、func、interface、map、struct、type、var | break、case、continue、default、defer、else、fallthrough、for、go、goto、if、range、return、select |  


**Go语言中的保留字**   
| 常量 | 变量类型 | 内置函数名 |  
| :---: | :---: | :---: |  
| true、false、iota、nil | Int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、uintptr、float32、float64、complex128、complex64、bool、byte、rune、string、error | make、len、cap、new、append、copy、close、delete、complex、real、imag、panic、recover |  

### Go 执行次序
Go 程序由一系列 Go 包组成，代码的执行也是在各个包之间跳转。  

Go 语言中有一个特殊的函数：main 包中的 main 函数，也就是 main.main，它是所有 Go 可执行程序的用户层执行逻辑的入口函数。main 函数的函数原型非常简单，没有参数也没有返回值。而且，Go 语言要求：可执行程序的 main 包必须定义 main 函数，否则 Go 编译器会报错。  
除了 main.main 函数之外，Go 语言还有一个特殊函数，它就是用于进行包初始化的 init 函数了。
> init 函数的第一个常用用途：重置包级变量值  
> init 函数的第二个常用用途，是实现对包级变量的复杂初始化  
> init 函数的第三个常用用途：在 init 函数中实现 “注册模式”  

关于Go程序初始化init函数的六个特点：  
- 包级别变量的初始化先于包内init函数的执行。
- 一个包下可以有多个init函数，每个文件也可以有多个init 函数。
- 多个 init 函数按照它们的文件名顺序逐个初始化。
- 应用初始化时初始化工作的顺序是，从被导入的最深层包开始进行初始化，层层递出最后到main包。
- 不管包被导入多少次，包内的init函数只会执行一次。
- 应用在所有初始化工作完成后才会执行main函数。

Go 包的初始化次序并不难，只需要记住这三点就可以了：  
1. 依赖包按 “深度优先” 的次序进行初始化；  
2. 每个包内按以 “常量 -> 变量 -> init 函数” 的顺序进行初始化；  
3. 包内的多个 init 函数按出现次序进行自动调用。  

### 最简单的 HTTP 服务  
```go
package main

import "net/http"

func main() {
  // 通过 http.HandleFunc 设置这个处理函数时，传入的模式字符串为 “/”
  // HTTP 服务器在收到请求后，会将请求中的 URI 路径与设置的模式字符串进行最长前缀匹配，并执行匹配到的模式字符串所对应的处理函数
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
      w.Write([]byte("hello, world"))
  })
  // 通过 http 包提供的 ListenAndServe 函数，建立起一个 HTTP 服务，这个服务监听本地的 8080 端口
  http.ListenAndServe(":8080", nil)
}
```
