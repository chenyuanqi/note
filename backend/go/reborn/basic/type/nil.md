
### Go nil
在 Go 语言中，布尔类型的零值（初始值）为 false，数值类型的零值为 0，字符串类型的零值为空字符串 ""，而指针、切片、映射、通道、函数和接口的零值则是 nil。

nil 是 Go 语言中一个预定义好的标识符，有过其他编程语言开发经验的开发者也许会把 nil 看作其他语言中的 null（NULL），其实这并不是完全正确的，因为 Go 语言中的 nil 和其他语言中的 null 有很多不同点。  
nil 是一个特殊的值，它只能赋值给指针类型和接口类型。  

**nil 标识符是不能比较的**  
== 对于 nil 来说是一种未定义的操作。  
```go
var a interface{} = nil
fmt.Println(a == nil) // true

fmt.Println(nil==nil) // invalid operation: nil == nil (operator == not defined on nil)

// 不同类型的 nil 是不能比较的
var m map[int]string
var ptr *int
fmt.Printf(m == ptr) // invalid operation: arr == ptr (mismatched types []int and *int)

// 两个相同类型的 nil 值也可能无法比较
var s1 []int
var s2 []int
fmt.Printf(s1 == s2) // invalid operation: s1 == s2 (slice can only be compared to nil)

// 可以，不可比较类型的空值直接与 nil 标识符进行比较
var s1 []int
fmt.Println(s1 == nil) // true
```

将一个带有类型的 nil 赋值给接口时，只有值为 nil，而类型不为 nil。此时，接口与 nil 判断将不相等。  
```go
type Person struct {
   name   string
   age    int
   gender int
}
type SayHello interface {
   sayHello()
}
func (p *Person) sayHello() {
   fmt.Println("Hello!")
}
func getSayHello() SayHello {
   var p *Person = nil
   return p
}
func main() {
   var person = new(Person)
   person.name = "David"
   person.age = 18
   person.gender = 0
   var sayHello SayHello
   sayHello = person
   fmt.Println(reflect.TypeOf(sayHello)) // *main.Person
   fmt.Println(sayHello == nil) // false
   fmt.Println(getSayHello() == nil) // false
}

// 不妨在 getSayHello () 函数值做些特殊处理。当函数体中的 p 变量为 nil 时，直接返回 nil 即可
func getSayHello() SayHello {
   var p *Person = nil
   if p == nil {
      return nil
   } else {
      return p
   }
}
fmt.Println(getSayHello() == nil) // true
```

**nil 不是关键字或保留字**  
nil 并不是 Go 语言的关键字或者保留字，也就是说我们可以定义一个名称为 nil 的变量。虽然声明语句可以通过编译，但是并不提倡这么做。
```go
var nil = errors.New("my god")
```

**nil 没有默认类型**  
```go
fmt.Printf("%T", nil)
print(nil) // use of untyped nil
```

**不同类型 nil 的指针是一样的**  
```go
var arr []int
var num *int
fmt.Printf("%p\n", arr) // 0x0
fmt.Printf("%p", num) // 0x0
```

**nil 是 map、slice、pointer、channel、func、interface 的零值**  
零值是 Go 语言中变量在声明之后但是未初始化被赋予的该类型的一个默认值。
```go
var m map[int]string
var ptr *int
var c chan int
var sl []int
var f func()
var i interface{}
fmt.Printf("%#v\n", m) // map[int]string(nil)
fmt.Printf("%#v\n", ptr) // (*int)(nil)
fmt.Printf("%#v\n", c) // (chan int)(nil)
fmt.Printf("%#v\n", sl) // []int(nil)
fmt.Printf("%#v\n", f) // (func())(nil)
fmt.Printf("%#v\n", i) // <nil>
```

**不同类型的 nil 值占用的内存大小可能是不一样的**   
具体的大小取决于编译器和架构，下面打印的结果是在 64 位架构和标准编译器下完成的，对应 32 位的架构的，打印的大小将减半。
```go
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var p *struct{}
    fmt.Println( unsafe.Sizeof( p ) ) // 8
    var s []int
    fmt.Println( unsafe.Sizeof( s ) ) // 24
    var m map[int]bool
    fmt.Println( unsafe.Sizeof( m ) ) // 8
    var c chan string
    fmt.Println( unsafe.Sizeof( c ) ) // 8
    var f func()
    fmt.Println( unsafe.Sizeof( f ) ) // 8
    var i interface{}
    fmt.Println( unsafe.Sizeof( i ) ) // 16
}
```