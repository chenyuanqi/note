
### Go 函数
函数是基本的代码块，它负责将一个复杂问题分解为不同的函数提供调用与复用。编写函数时，无需关注顺序，因为 Go 语言是编译型的。

在 Go 语言中有三种函数类型：
- 基本格式：有命名的函数，直接调用完事。
- 匿名函数或者 lambda 函数：没有名字的函数。
- 结构体携带的函数：也可以称之为方法。

Go 语言支持多返回值，多返回值能方便地获得函数执行后的多个返回参数，Go 语言经常使用多返回值中的最后一个返回参数返回函数执行中可能发生的错误。

Go 函数为什么没有支持默认参数？Go 语言追求显式的表达，避免隐含。默认函数参数，可能导致调用者不知道默认参数是什么出现一些问题。
```go
// 基本格式
func Fun1(arg1 T, arg2 T) T {
    ...
    return r1
}
// Go 语言函数中有个特点，可以多个值返回。在声明返回值类型时，可以不指定名称，也可以指定名称
// 无名称
func Fun1(arg1 T, arg2 T) (T, T) {
    ...
    return r1, r2
}
// 有名称
func Fun1(arg1 T, arg2 T) (n1 T, n2 T) {
    ...
    return
}
// 构造好一个函数后，如何调用，格式如下
r1, r2 := Fun1(param1, param2)
// 如果 r2 不想使用
r1, _ := Fun1(param1, param2)


// 计算两个数之和并且返回
func AddNum(n1 int, n2 int) int {
    return n1 + n2
}
```

**匿名函数**  
匿名函数就是在构造函数时，函数没有名称，想调用时，需要把匿名函数赋值给一个变量，或者在构造时直接调用。
```go
// 匿名函数的定义就是没有名字的普通函数定义
func(参数列表)(返回参数列表){
    函数体
}

// 赋值给变量
fun1 := func (arg1 T, arg2 T) T {
    ...
    return r1
}
// 赋值后， fun1 就是一个函数类型的变量，这样调用
fun1(param1, param2)

// 构造时调用
func (arg1 T, arg2 T) T {
    ...
    return r1
}(param1, param2)

// 使用匿名函数实现操作封装
// 定义命令行参数 skill，从命令行输入 --skill 可以将=后的字符串传入 skillParam 指针变量
var skillParam = flag.String("skill", "", "skill to perform")
// 解析命令行参数，解析完成后，skillParam 指针变量将指向命令行传入的值
flag.Parse()
var skill = map[string]func(){
    "fire": func() {
        fmt.Println("chicken fire")
    },
    "run": func() {
        fmt.Println("soldier run")
    },
    "fly": func() {
        fmt.Println("angel fly")
    },
}
if f, ok := skill[*skillParam]; ok {
    f()
} else {
    fmt.Println("skill not found")
}
```

**传递函数**  
在 Go 语言中，函数是 “一等公民”，它和 int 、string 等等，都是一个级别，可以作为参数进行传递。
```go
package main

// callback 是一个函数类型参数
func Calc(callback func(n1 int, n2 int) int) int {
    x, y := 3, 4
    return callback(x, y)
}

// 计算两数之积
func Mul(n1 int, n2 int) int {
    return n1 * n2
}

func main() {
    // 第一个：传递一个匿名函数
    Calc(func(n1 int, n2 int) int {
        return n1 + n2
    })

    // 第二个：传递一个定义好的函数
    Calc(Mul)
}
```

**声明函数类型**  
声明函数类型，意思就是可以自定义一个函数类型，给这个函数取一个别名，像例如 int 一样很方便的去声明变量或者参数类型
```go
type CallbackFunc func(n1 int, n2 int) int

func Calc(callback CallbackFunc) int {
    ...
}
```

**函数参数**  
1. 参数类型省略  
在声明函数参数时，有时候会遇到连续声明多个相同类型，这个时候，就可以只保留一个类型名称。
```go
// 没精简的
func Fun1(arg1 string, arg2 int, arg3 int)

// 精简后
func Fun1(arg1 string, arg2, arg3 int)
```

2. 值传递与引用传递  
函数调用时传递的参数称为实参，构造函数时的参数称为形参。  
在 Go 语言中，切片（slice）、map、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递，在函数内修改形参是会改变实参的值。  
对于切片，有种情况会打破引用传递这个规律。  
```go

```
对于其它剩下的类型，默认都是值传递，函数接收到的形参只是副本，函数内对形参的更改是不会影响到实参的。

如果希望更改实参的值，可以传递指针，在实参前增加 “&” 符号，表示取实参的地址，例如： Fun1(&param) 。

3. 变长参数  
当构造函数时，函数的最后一个参数是 ...T 形式时，称为变长参数，它可以接受至少 0 个数据。  
```go
// nums 实际是一个切片
func Func1(str string, nums ...int) {
    ...
}
// 没传递变长参数
Func1("miao")
// 给变长参数传递不同数量的值
Func1("miao", 1)
Func1("miao", 1, 2)
// 当把一个切片类型传递给可变参数时，在切片后跟着 ... 三个点，传递给变长参数，表示将切片元素展开
nums := []int{1, 2, 3}
Func1("miao", nums...)
```

**把函数作为接口来调用**  
```go
// 结构体类型
type Struct struct {
}
// 实现Invoker的Call
func (s *Struct) Call(p interface{}) {
    fmt.Println("from struct", p)
}

// 调用器接口
type Invoker interface {
    // 需要实现一个Call()方法
    Call(interface{})
}

// 声明接口变量
var invoker Invoker
// 实例化结构体
s := new(Struct)
// 将实例化的结构体赋值到接口
invoker = s
// 使用接口调用实例化结构体的方法Struct.Call
invoker.Call("hello") // from struct hello

// 函数的声明不能直接实现接口，需要将函数定义为类型后，使用类型实现结构体，当类型方法被调用时，还需要调用函数本体
// 函数定义为类型
type FuncCaller func(interface{})
// 实现Invoker的Call
// 将 func(v interface{}){} 匿名函数转换为 FuncCaller 类型（函数签名才能转换）
func (f FuncCaller) Call(p interface{}) {
    // 调用f()函数本体
    f(p)
}
// 声明接口变量
var invoker Invoker
// 将匿名函数转为FuncCaller类型, 再赋值给接口
// FuncCaller 类型实现了 Invoker 的 Call() 方法，赋值给 invoker 接口
invoker = FuncCaller(func(v interface{}) {
    fmt.Println("from function", v)
})
// 使用接口调用FuncCaller.Call, 内部会调用函数本体
invoker.Call("hello") // from function hello
```

**内置函数**  
在 Go 语言中，有一些函数无需导入任何包就可以使用，比如：  
- make：为切片，map、通道类型分配内存并初始化对象。
- len：计算数组、切片、map、通道的长度。
- cap：计算数组、切片、通道的容量。
- delete：删除 map 中对应的键值对。
- append：将数据添加到切片的末尾。
- copy：将原切片的数据复制到新切片中。
- new：除切片、map、通道类型以外的类型分配内存并初始化对象，返回的类型为指针。
- complex：生成一个复数。
- real：获取复数的实部。
- imag：获取复数的虚部
- print：将信息打印到标准输出，没有换行。
- println：将信息打印到标准输出并换行。
- close：关闭通道。
- panic：触发程序异常。
- recover：捕捉 panic 的异常信息。
