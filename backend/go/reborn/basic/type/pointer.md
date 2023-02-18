### 概述

**Go 提供了指针操作，但是没有指针运算。**
也就是说，不像 C 语言中那般强大，毕竟 `指针是 C 语言的灵魂`。
即使如此，指针依然是非常重要的，在一些 `性能敏感` 的场景中，指针的身影随处可见。
如果是系统编程、操作系统或者网络应用，指针更是不可或缺的一部分。

**指针的值是一个变量的地址。当然了，指针也是变量的一种，但是一般称其为 `指针变量`。**

### 取地址

**关键字 `&` 表示取地址符**。

程序运行时，数据通常存储在内存中，每个内存块都有一个地址， 通常使用 `十六进制` 表示，比如 `0xc0000160a0`。

```shell
# 将 & 放到一个变量前，就会获得该变量对应的内存地址, 例如
x := 1024

# p 变量是一个指针变量，值对应着变量 x 的地址
p := &a
```

#### 例子

```go
package main

import "fmt"

func main() {
	pi := 3.1415
	fmt.Printf("%p\n", &pi) // 直接取地址, 输出的是变量 pi 的地址

	var p *float64        // 浮点型指针变量 
	p = &pi               // 通过变量取地址
	fmt.Printf("%p\n", p) // 输出的是指针的地址, 输出的是指针 p 的地址
}

// $ go run main.go
// 输出如下 
/**
  0xc0000160a0    // 这个是我电脑的内存地址，你的输入可能和这个不一样
  0xc0000b2000
*/
```

### 改变值

在刚才的例子中，获取到了变量的地址后，直接进行了输出。 那么，应该如何输出指针对应的变量的值呢？

**关键字 `*` 表示指针调用符**。

```shell
# 将 * 放到一个指针变量前，就会获得该指针变量对应的变量的值, 例如
x := 1024

# p 变量是一个指针变量，值对应着变量 x 的地址
p := &a     

# *p 表示 p 对应的变量的值，也就是 x 的值，也就是 1024，
# *p = 1025, 表示将 x 的值修改为 1025
*p = 1025   
```

#### 例子

```go
package main

import "fmt"

func main() {
	ok := true
	var p *bool            // 布尔型指针变量
	p = &ok                // 获取 ok 的地址
	fmt.Printf("%t\n", *p) // 输出指针变量 p 对应的变量 ok 的值

	*p = false             // 改变了变量 ok 的值
	fmt.Printf("%t\n", *p) // 输出指针变量 p 对应的变量 ok 的值
}

// $ go run main.go
// 输出如下 
/**
    true
    false
*/
```

### Go 指针
什么是内存地址？说通俗点就是电脑上数据存储位置的编号，就好比我们的身份证号一样。指针也就是所说的内存地址，内存地址保存在指针变量里。  
指针可以节省复制的开销，但同时要考虑解引用和垃圾回收带来的影响，所以不要把使用指针作为性能优化的首选方案。  

Go语言中的指针，它主要由两大核心概念构成：类型指针和切片指针。  
- 类型指针：在传递数据时直接使用指针，可以避免创建数据的副本，节约内存开销。类型指针不能进行偏移和运算，可以避免非法修改为其它数据的风险，也更有利于垃圾回收机制及时找到并回收它们；
- 切片指针：切片由指向起始元素的指针、元素数量和总容量构成。当访问切片发生越界时，会发生宕机并输出堆栈信息。宕机是可以恢复的，而崩溃只能导致程序停止运行。

使用指针更有利于程序运行的性能和稳定性。另外，在某些操作中，如使用反射修改变量的值，必须使用可寻址的变量（通过指针）。  
在实际应用中，最为常用的便是获取变量的内存地址，以及获取某个地址对应的值。在Go语言中，前者使用“&”运算符，后者使用“*”运算符。它们互为反向操作，操作的对象也不同。
```go
// 指针类型是在任意类型前增加星号
*BaseType
// *int 表示 int 类型变量的指针类型
// *string 表示 string 类型变量的指针类型

//exampleNumberA变量（整数型变量）声明和赋值
var exampleNumberA int = 10
//获取exampleNumberA的地址，并赋值给exampleNumberAPtr变量（exampleNumberAPtr的类型是指针类型）
exampleNumberAPtr := &exampleNumberA
//输出exampleNumberAPtr变量的值（将输出内存地址）
fmt.Println(exampleNumberAPtr) // 0xc00001a088
//获取exampleNumberAPtr（指针变量）表示的实际数据值，并赋值给exampleNumberAPtrValue变量（整数型变量）
exampleNumberAPtrValue := *exampleNumberAPtr
//输出exampleNumberAPtrValue变量（整数型变量）的值
fmt.Println(exampleNumberAPtrValue) // 10

// 创建一个 *int 指针类型的变量
var p *int
// 初始化
var num int =  11
p = &num // 获取变量 num 的地址并赋值给指针变量 p
// 输出指针变量信息
fmt.Println(p) // 0xc00000a088（0x 开头说明是十六进制，该十六进制就是变量 num 的内存地址）
```

我们还可以使用 new() 函数直接创建指针变量，相当于在内存中创建了没有变量名的某种类型的变量。这样做无需产生新的数据“代号”，取值和赋值转而通过指针变量完成。常用在无需变量名或必须要传递指针变量值的场景中。  
```go
//使用new()函数创建名为exampleNumberAPtr指针类型变量，表示int64型值
exampleNumberAPtr := new(int64)
//修改exampleNumberAPtr表示的实际数据值
*exampleNumberAPtr = 100
//获取exampleNumberAPtr表示的实际数据值
fmt.Println(*exampleNumberAPtr) // 100
```

空指针表示指针变量没有任何赋值，此时空指针变量等于 nil。  
nil 类似其它语言中的 null ，在 Go 语言中只能和指针类型、接口类型进行比较，也只能给指针类型变量和接口类型变量赋值。  
```go
var empty *int
fmt.Println(empty == nil) // true
```

指针取值  
声明了一个指针变量后，如果想从指针变量中取值那该如何做，指针的取值常常被称作解引用。  
如果指针变量是空指针，再从中取值时，编译器会报错。  
```go
var num int =  11
var p *int
p = &num
// 取值，取值时在指针变量前增加一个 * 符号
fmt.Println(*p) // 11
```

结构体  
如果指针变量是结构体指针类型时，获取结构体中的字段或调用方法时，无需在指针变量前增加 * 。
- 结构体指针输出的不是地址
- 调用结构体的字段或方法时无需添加 *
```go
p := &People{
            Name: "老苗",
            Age:  18,
        }
fmt.Pringln(p) // &{老苗 18}
fmt.Println(p.Name) // 或 fmt.Println((*p).Name) // 老苗

// 如果通过方法想修改结构体中的字段时，可以将接收者设置为指针类型
func (p *People) SetName(name string) {
    p.Name = name
}

func main() {
    people := People{
        Name: "老苗",
    }
    people.SetName("潇洒哥")
    fmt.Println(people.Name) // 潇洒哥
}
```

指针传递  
在 Go 语言中大部分的类型都是值传递，也就是说通过函数传值时，函数内的修改是不能影响外部的，如果想更改就使用指针类型。  
```go
func UpdateNum(num *int) {
    *num = 2
}

func main() {
    n := 1
    UpdateNum(&n)
    fmt.Println(n) // 2
}
```

对于 Go 语言中的个别类型本身就是引用类型，不需要使用指针就可以在函数内部修改值而影响外部。比如 map 和 通道、切片等。
```go
func SetCountry(countries map[string]string) {
    countries["china"] = "中国"
}
func main() {
    c := make(map[string]string)
    SetCountry(c)
    fmt.Println(c) // map[china:中国]
}

// 在切片传递时不会改变底层数组的引用，但如果对切片进行追加操作后，数组引用就会改变
func AppendAnimals(animals []string) {
    animals = append(animals, "老虎", "大象")
}
func main() {
    // AppendAnimals 函数给切片追加元素，但外部的变量 input 的值不受影响，因为 append 操作后底层数组会进行拷贝并改变引用
    input := []string{"猴子"}
    AppendAnimals(input)
    fmt.Println(input) // [猴子]
}
// 使用指针解决
func AppendAnimalsPointer(animals *[]string) {
    *animals = append(*animals, "老虎", "大象")
}
func main() {
    input := []string{"猴子"}
    AppendAnimalsPointer(&input)
    fmt.Println(input)  // [猴子 老虎 大象]
}
// 在传递切片时如果只修改切片内容，不追加元素，原切片数据将会受到影响，因为底层数组的引用没有改变
// UpdateAnimals 函数修改了切片内容，通过输出可以看出 updateInput 变量数据已改变
func UpdateAnimals(animals []string) {
    for i := range animals {
        animals[i] = "兔子"
    }
}
func main() {
    updateInput := []string{"老虎", "大象"}
    UpdateAnimals(updateInput)
    fmt.Println(updateInput) // [兔子 兔子]

}
```


