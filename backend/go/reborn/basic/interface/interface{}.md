
### 空接口
空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无须实现空接口。从实现的角度看，任何值都满足这个接口的需求。因此空接口类型可以保存任何值，也可以从空接口中取出原值。

`空接口类型类似于 C# 或 Java 语言中的 Object、C 语言中的 void*、C++ 中的 std::any。在泛型和模板出现前，空接口是一种非常灵活的数据抽象保存和使用的方法。`

空接口的内部实现保存了对象的类型和指针。使用空接口保存一个数据的过程会比直接用数据对应类型的变量保存稍慢。因此在开发中，应在需要的地方使用空接口，而不是在所有地方使用空接口。
```go
// 将值保存到空接口
var any interface{}
any = 1
fmt.Println(any) // 1
any = "hello"
fmt.Println(any) // hello
any = false
fmt.Println(any) // false

// 从空接口获取值
// 声明a变量, 类型int, 初始值为1
var a int = 1
// 声明i变量, 类型为interface{}, 初始值为a, 此时i的值变为1
var i interface{} = a
// 声明b变量, 尝试赋值i
var b int = i // cannot use i (type interface {}) as type int in assignment: need type assertion 不能将 i 变量视为 int 类型赋值给 b，编译器提示我们得使用 type assertion，意思就是类型断言

// 空接口取值
// 修改后，代码可以编译通过，并且 b 可以获得 i 变量保存的 a 变量的值：1
var b int = i.(int)
```

**空接口的值比较**  
空接口在保存不同的值后，可以和其他变量值一样使用 == 进行比较操作。  
1) 类型不同的空接口间的比较结果不相同  
保存有类型不同的值的空接口进行比较时，Go 语言会优先比较值的类型。因此类型不同，比较结果也是不相同的。  
```go
// a保存整型
var a interface{} = 100
// b保存字符串
var b interface{} = "hi"
// 两个空接口不相等
fmt.Println(a == b) // false

var a interface{} = 10
var b interface{} = "10"
fmt.Println(a == b) // false
```
2) 不能比较空接口中的动态值  
当接口中保存有动态类型的值时，运行时将触发错误。比如 Map 和 Slice，强行比较会引发如上宕机错误。  
```go
var a interface{} = [5]int{1, 2, 3, 4, 5}
var b interface{} = [5]int{1, 2, 3, 4, 5}
fmt.Println(a == b) // true

// c保存包含10的整型切片
var c interface{} = []int{10}
// d保存包含20的整型切片
var d interface{} = []int{20}
// 这里会发生崩溃
fmt.Println(c == d) // panic: runtime error: comparing uncomparable type []int
```
[![类型的可比较性](https://s1.ax1x.com/2022/04/22/L25SoV.md.png)](https://imgtu.com/i/L25SoV)