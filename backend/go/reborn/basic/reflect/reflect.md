
### Go 反射
反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。  
支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。

Go 程序在运行期使用 reflect 包访问程序的反射信息。  
Go 语言提供了一种机制在运行时更新和检查变量的值、调用变量的方法和变量支持的内在操作，但是在编译时并不知道这些变量的具体类型，这种机制被称为反射。反射也可以让我们将类型本身作为第一类的值类型处理。  
Go 程序的反射系统无法获取到一个可执行文件空间中或者是一个包中的所有类型信息，需要配合使用标准库中对应的词法、语法解析器和抽象语法树（AST）对源码进行扫描后获得这些信息。  

Go 语言中的反射是由 reflect 包提供支持的，它定义了两个重要的类型 Type 和 Value 任意接口值在反射中都可以理解为由 reflect.Type 和 reflect.Value 两部分组成，并且 reflect 包提供了 reflect.TypeOf 和 reflect.ValueOf 两个函数来获取任意对象的 Value 和 Type。

**反射的类型对象（reflect.Type）**  
在 Go 语言程序中，使用 reflect.TypeOf() 函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。
```go
package main
import (
    "fmt"
    "reflect"
)
func main() {
    var a int
    typeOfA := reflect.TypeOf(a)
    fmt.Println(typeOfA.Name(), typeOfA.Kind())
}
// int  int
```

**反射的类型（Type）与种类（Kind）**  
在使用反射时，需要首先理解类型（Type）和种类（Kind）的区别。编程中，使用最多的是类型，但在反射中，当需要区分一个大品种的类型时，就会用到种类（Kind）。例如需要统一判断类型中的指针时，使用种类（Kind）信息就较为方便。  

1) 反射种类（Kind）的定义  
Go 语言程序中的类型（Type）指的是系统原生数据类型，如 int、string、bool、float32 等类型，以及使用 type 关键字定义的类型，这些类型的名称就是其类型本身的名称。例如使用 type A struct{} 定义结构体时，A 就是 struct{} 的类型。

种类（Kind）指的是对象归属的品种，在 reflect 包中有如下定义：
```go
type Kind uint
const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)
```
Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr。type A struct{} 定义的结构体属于 Struct 种类，*A 属于 Ptr。  

2) 从类型对象中获取类型名称和种类  
Go 语言中的类型名称对应的反射获取方法是 reflect.Type 中的 Name() 方法，返回表示类型名称的字符串；类型归属的种类（Kind）使用的是 reflect.Type 中的 Kind() 方法，返回 reflect.Kind 类型的常量。
```go
package main
import (
    "fmt"
    "reflect"
)
// 定义一个Enum类型
type Enum int
const (
    Zero Enum = 0
)
func main() {
    // 声明一个空结构体
    type cat struct {
    }
    // 获取结构体实例的反射类型对象
    typeOfCat := reflect.TypeOf(cat{})
    // 显示反射类型对象的名称和种类
    fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
    // 获取Zero常量的反射类型对象
    typeOfA := reflect.TypeOf(Zero)
    // 显示反射类型对象的名称和种类
    fmt.Println(typeOfA.Name(), typeOfA.Kind())
}
// cat struct
// Enum int
```

**指针与指针指向的元素**  
Go 语言程序中对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型，这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作。  
```go
package main
import (
    "fmt"
    "reflect"
)
func main() {
    // 声明一个空结构体
    type cat struct {
    }
    // 创建cat的实例
    ins := &cat{}
    // 获取结构体实例的反射类型对象
    typeOfCat := reflect.TypeOf(ins)
    // 显示反射类型对象的名称和种类
    fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())
    // 取类型的元素
    typeOfCat = typeOfCat.Elem()
    // 显示反射类型对象的名称和种类
    fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
}
// name:'' kind:'ptr'
// element name: 'cat', element kind: 'struct'
```

**使用反射获取结构体的成员类型**  
任意值通过 reflect.TypeOf() 获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象 reflect.Type 的 NumField() 和 Field() 方法获得结构体成员的详细信息。

与成员获取相关的 reflect.Type 的方法如下表所示。
[![OKj7uD.md.png](https://s1.ax1x.com/2022/05/07/OKj7uD.md.png)](https://imgtu.com/i/OKj7uD)  

1) 结构体字段类型  
reflect.Type 的 Field() 方法返回 StructField 结构，这个结构描述结构体的成员信息，通过这个信息可以获取成员与结构体的关系，如偏移、索引、是否为匿名字段、结构体标签（StructTag）等，而且还可以通过 StructField 的 Type 字段进一步获取结构体成员的类型信息。

StructField 的结构如下：
```go
type StructField struct {
    Name string          // 字段名
    PkgPath string       // 字段路径
    Type      Type       // 字段反射类型对象
    Tag       StructTag  // 字段的结构体标签
    Offset    uintptr    // 字段在结构体中的相对偏移
    Index     []int      // Type.FieldByIndex中的返回的索引值
    Anonymous bool       // 是否为匿名字段
}
```

2) 获取成员反射信息  
下面代码中，实例化一个结构体并遍历其结构体成员，再通过 reflect.Type 的 FieldByName() 方法查找结构体中指定名称的字段，直接获取其类型信息。反射访问结构体成员类型及信息：
```go
package main
import (
    "fmt"
    "reflect"
)
func main() {
    // 声明一个空结构体
    type cat struct {
        Name string
        // 带有结构体tag的字段
        Type int `json:"type" id:"100"`
    }
    // 创建cat的实例
    ins := cat{Name: "mimi", Type: 1}
    // 获取结构体实例的反射类型对象
    typeOfCat := reflect.TypeOf(ins)
    // 遍历结构体所有成员
    for i := 0; i < typeOfCat.NumField(); i++ {
        // 获取每个成员的结构体字段类型
        fieldType := typeOfCat.Field(i)
        // 输出成员名和tag
        fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
    }
    // 通过字段名, 找到字段类型信息
    if catType, ok := typeOfCat.FieldByName("Type"); ok {
        // 从tag中取出需要的tag
        fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
    }
}
// name: Name  tag: ''
// name: Type  tag: 'json:"type" id:"100"'
// type 100
```

**结构体标签（Struct Tag）**  
通过 reflect.Type 获取结构体成员信息 reflect.StructField 结构中的 Tag 被称为结构体标签（StructTag）。结构体标签是对结构体字段的额外信息标签。

JSON、BSON 等格式进行序列化及对象关系映射（Object Relational Mapping，简称 ORM）系统都会用到结构体标签，这些系统使用标签设定字段在处理时应该具备的特殊属性和可能发生的行为。这些信息都是静态的，无须实例化结构体，可以通过反射获取到。  

1) 结构体标签的格式  
结构体标签由一个或多个键值对组成；键与值使用冒号分隔，值用双引号括起来；键值对之间使用一个空格分隔。Tag 在结构体字段后方书写的格式：
```
`key1:"value1" key2:"value2"`
```

2) 从结构体标签中获取值  
StructTag 拥有一些方法，可以进行 Tag 信息的解析和提取，如下所示：  
- func (tag StructTag) Get(key string) string：根据 Tag 中的键获取对应的值，例如\`key1:"value1" key2:"value2"\`的 Tag 中，可以传入“key1”获得“value1”。
- func (tag StructTag) Lookup(key string) (value string, ok bool)：根据 Tag 中的键，查询值是否存在。

3) 结构体标签格式错误导致的问题  
编写 Tag 时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误。  
```go
package main
import (
    "fmt"
    "reflect"
)
func main() {
    type cat struct {
        Name string
        Type int `json: "type" id:"100"` // 在 json: 和 "type" 之间增加了一个空格，这种写法没有遵守结构体标签的规则，因此无法通过 Tag.Get 获取到正确的 json 对应的值
    }
    typeOfCat := reflect.TypeOf(cat{})
    if catType, ok := typeOfCat.FieldByName("Type"); ok {
        fmt.Println(catType.Tag.Get("json"))
    }
}
// 输出一个空字符串，并不会输出期望的 type
```

### 反射规则
反射第一定律：反射可以将“接口类型变量”转换为“反射类型对象”  
`注：这里反射类型指 reflect.Type 和 reflect.Value。`  
从使用方法上来讲，反射提供了一种机制，允许程序在运行时检查接口变量内部存储的 (value, type) 对。  
reflect 包的两种类型 Type 和 Value，这两种类型使访问接口内的数据成为可能，它们对应两个简单的方法，分别是 reflect.TypeOf 和 reflect.ValueOf，分别用来读取接口变量的 reflect.Type 和 reflect.Value 部分。  

反射第二定律：反射可以将“反射类型对象”转换为“接口类型变量”  
使用 Interface 方法恢复其接口类型的值。事实上，这个方法会把 type 和 value 信息打包并填充到一个接口变量中，然后返回。  

反射第三定律：如果要修改“反射类型对象”其值必须是“可写的”  
首先我们要弄清楚什么是“可写性”，“可写性”有些类似于寻址能力，但是更严格，它是反射类型变量的一种属性，赋予该变量修改底层存储数据的能力。“可写性”最终是由一个反射对象是否存储了原始值而决定的。  
反射不太容易理解，reflect.Type 和 reflect.Value 会混淆正在执行的程序，但是它做的事情正是编程语言做的事情。只需要记住：只要反射对象要修改它们表示的对象，就必须获取它们表示的对象的地址。  

