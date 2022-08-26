
### Golang 内建函数
内建函数主要参考 go 源码 go/src/builtin/builtin.go 中定义。

**make**  
make 内建函数，用于内建类型（map、slice 和channel）的内存分配，并且返回一个有初始值(非零)的 T 类型。  
> slice、map和channel 这三个类型在使用前必须被初始化。  
> slice 初始化默认是nil，map、channel类型 make默认初始化了内部的数据结构，填充适当的值（初始值）。  

```go
// 函数原型

func make(t Type, size ...IntegerType) Type 
// Type: 数据类型，必要参数，只能是 slice、 map、 channel 这三种数据类型
// size: 数据类型实际占用的内存空间长度，map、 channel 是可选参数，slice 是必要参数
// ...: 为数据类型提前预留的内存空间长度，可选参数。所谓的提前预留是当前为数据类型申请内存空间的时候，提前申请好额外的内存空间，这样可以避免二次分配内存带来的开销，大大提高程序的性能
// 这里Type、IntegerType 类型是 go 源码builtin.go中定义的type IntegerType int 别名

// slice 切片
num1 := make([]int, 2) // 指定长度为2的 slice
num2 := make([]int, 2, 4) // 指定长度为2, 指定预留的空间长度为4

// map 初始化
make(map[string]string)

// chan 声明
ch := make(chan int) // 声明初始化chan为int类型
```


**new**  
new 内建函数，用于各种类型的内存分配。  
new 内建函数用法 new(T) 分配了零值填充的T类型的内存空间，并且返回其地址，即 \*T 类型的值。用 Go 的术语说，它返回了一个指针，指向新分配的类型 T 的零值。  
`我们目前只需使用 new() 函数，无需担心其内存的生命周期或怎样将其删除，因为 Go 语言的内存管理系统会帮我们打理一切。`  
```go
// 函数原型
func new(Type) *Type 
// Type 是指类型，不是值（第一个参数是类型，不是值；返回的值是指向新分配的零的指针这种类型的值）

// 使用指针对象赋值
num := new(int) // 自动内存分配
*num = 100 // 赋值

// go 语言中只声明的指针变量不能直接赋值，需要手动分配空间才能使用
// 手动分配
var p *int
*p = new(int)
*p = 100

// 每次调用new函数都是返回一个新的变量的地址
num1 := new(int)
num2 := new(int)
fmt.Println(num1 == num2)   // "false"
```


**append**  
append 内建函数，主要用于切片(slice) 追加元素。  
如果该切片存储空间（cap）足够，就直接追加，长度（len）变长；如果空间不足，就会重新开辟内存，并将之前的元素和新的元素一同拷贝进去。  
`注意：append 返回值必须要有接收变量，不然编译器会报错；因为 Go 函数传递默认是值拷贝，将 slice 变量传入 append 函数相当于传了原 slice 变量的一个副本，注意不是拷贝底层数组，因为 slice 变量并不是数组，它仅仅是存储了底层数组的一些信息。`
```go
// 函数原型
func append(slice []Type, elems ...Type) []Type 
// slice: 切片的类型
// elems: 可以传入多个参数元素

slice1 := append([]int{1,2,3}, 4) // 单个元素传入
slice2 := append([]int{1,2,3}, 4, 5, 6) // 单个元素传入
// 将内容追加第一个[]int数组中
slice := append([]int{1,2,3}, []int{4,5,6}...) // 接收两个slice需要填写... 
// 使用[]byte类型，将字符串类型追加 []byte数组
slice := append([]byte("帽儿山的枪手 "), "分享技术文章"...)
```


**copy**  
copy 内建函数，只能用于数组切片内容赋值。返回结果为一个 int 型值，表示 copy 的长度。  
复制内置函数将元素从源切片复制到目标片（作为特例，它还将从字符串中复制字节）。源和目标可能重叠，副本返回复制的元素数，即 len（src）和len（dst）。  
> 如果 dst 长度小于 src 的长度，则 copy 部分  
> 如果大于，则全部拷贝过来，只是没占满 dst 的位子而已  
> 相等时刚好不多不少 copy 过来  
> 如果切片元素的类型是引用类型，那么 copy 的也将是个引用  

```go
// 函数原型
func copy(dst, src []Type) int 

// 长度相等时
s := []int{1, 2, 3}
copy(s, []int{4,5,6})
fmt.Println(s) // 输出 [4 5 6]

// 长度不相等时
s := []int{1, 2, 3, 4}
copy(s, []int{7, 8, 9, 10, 11})
fmt.Println(s) // 输出 [7 8 9 10]
```


**len**  
len 内置函数，用于计算数组(包括数组指针)、切片(slice)、map、channel、字符串等数据类型的长度。  
`注意：结构体(struct)、整型布尔等不能作为参数传给 len 函数。`  
> map和slice：元素个数  
> channel：通道中未读的元素个数  
> 字符串：字节数，并非字符串的字符数  
> 当 V 的值为 nil 值，len 返回 0  

```go
// 函数原型
func len(v Type) int

res := len([]int{1, 2, 3})
fmt.Println(res) // 3
```


**cap**  
cap 内建函数，返回指定类型的容量，根据不同类型，返回分配的空间大小。  
```go
// 函数原型
func cap(v Type) int 

slice := make([]int, 5, 10)  // 第三个参数, 预留内存空间 10
fmt.Println(cap(slice)) // 10
```

**print & println**  
print 内置函数将其参数格式化为实现特定的方法，并将结果写入标准错误。主要用于打印和调试；它不一定会留在实际生产代码中。  
println 内置函数相比 print 增加了默认换行，打印的每一项之间都会有空行。  
> 不能打印数组、结构体（复合类型）  
> 对于组合类型（除了基本类型都是）的值会打印底层引用值地址  

```go
// 函数原型
func print(args ...Type) 
func println(args ...Type)

print("print", "帽儿山的枪手\n") // print帽儿山的枪手
println("println", "帽儿山的枪手") // println 帽儿山的枪手
```

**close**  
close 内建函数，其功能是关闭一个通道，该通道必须是双向或仅发送。它只能由发送者执行，而不能由发送者执行接收器，并具有在最后一次发送值被接收。在从已关闭的通道 c，任何来自 c 的接收都将在不阻塞的情况下成功，返回通道元素的值为零。  
`注意: 对于值为 nil 的 channel 或者对同一个 channel 重复 close, 都会 panic, 关闭只读 channel 会报编译错误。`  
```go
// 函数原型
func close(c chan<- Type) 

ch1 := make(chan int, 1) // 双向通道
ch2 := make(chan<- int, 1) // 只写通道
close(ch1)
close(ch2)
```

**delete**  
delete内置函数，是删除指定键值(map)元素。如果 map 是 nil 或没有这样的元素，delete 是禁止删除的。  
```go
// 函数原型
func delete(m map[Type]Type1, key Type) 
// Type1 类型， 仅用于文档编制。这是一个替身对于任何 Go 类型，但表示任何给定函数的相同类型

mapInfo := make(map[string]int)
mapInfo["key1"] = 1
mapInfo["key2"] = 2
mapInfo["key3"] = 3
delete(mapInfo, "key1") //delete key1后输出map集合key和value
for k, v := range mapInfo {
   fmt.Printf("key:%s value:%d \n", k, v)
}
// key:key2 value:2
// key:key3 value:3
```

**complex**  
complex 内置函数，从指定的实部和虚部构建复数。  
```go
// 函数原型
func complex(r, i FloatType) ComplexType 
// Go 提供了两种大小的复数类型：complex64 和 complex128，分别由 float32 和 float64 组成

complex(1, 2) // 构建复数
```

**real&imag**  
内置函数 real 和 imag 用来获取复数的实部和虚部。
```go
// 函数原型
func real(c ComplexType) FloatType
func imag(c ComplexType) FloatType 

var x complex128 = complex(1, 2) // 1+2i  构建复数
var y complex128 = complex(3, 4) // 3+4i 构建复数
fmt.Println(x*y)                 // "(-5+10i)" 
fmt.Println(real(x*y))           // "-5"  获取实部
fmt.Println(imag(x*y))           // "10"  获取虚部
```

**panic**  
panic 内建函数，是在程序运行时才回抛出来的异常错误。在 panic 被抛出之后，如果没有在程序里添加任何保护措施的话，程序就会在打印出 panic 的详情，终止运行。  
Go 语言追求简洁优雅，所以，Go 语言不支持传统的 try…catch…finally 这种异常，因为 Go 语言的设计者们认为，将异常与控制结构混在一起会很容易使得代码变得混乱。因为开发者很容易滥用异常，甚至一个小小的错误都抛出一个异常。在 Go 语言中，使用多值返回来返回错误。不要用异常代替错误，更不要用来控制流程。
```go
// 函数原型
func panic(v interface{}) 

// 引发panic
panic("出现错误")
```

**recover**  
recover 内建函数，会捕获错误信息，使该错误信息终止报告。 (用来控制一个 goroutine 的 panicking 行为，捕获 panic，从而影响应用的行为)。
使用 recvoer 内建函数，通常在延迟函数中执行恢复调用。  
- 在 defer 函数中，通过 recever 来终止一个 goroutine 的 panicking 过程，从而恢复正常代码的执行  
- 可以获取通过 panic 传递的 error  
- 利用 recover 处理 panic 指令，defer 必须在 panic 之前声明，否则当 panic 时，recover 无法捕获到 panic  

```go
// 函数原型
func recover() interface{} 


func main() {
  defer func() { // 必须要先声明defer，否则不能捕获到panic异常
    if err := recover(); err != nil {
      fmt.Println(err) // panic内建函数传入的内容
    }
    fmt.Println("recover执行后") // 最后执行
  }()
  fmt.Println("正常执行")
  panic("panic错误")
}
// 正常执行
// panic错误
// recover执行后
```


