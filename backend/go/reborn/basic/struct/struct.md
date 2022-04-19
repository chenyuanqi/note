
### Go 结构体
Go 语言通过用自定义的方式形成新的类型，结构体是类型中带有成员的复合类型。Go 语言使用结构体和结构体成员来描述真实世界的实体和实体对应的各种属性。  
Go 语言中的类型可以被实例化，使用 new 或 & 构造的类型实例的类型是类型的指针。  

结构体成员是由一系列的成员变量构成，这些成员变量也被称为“字段”。字段有以下特性：  
- 字段拥有自己的类型和值。
- 字段名必须唯一。
- 字段的类型也可以是结构体，甚至是字段所在结构体的类型。

Go 语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。  
Go 语言的结构体与“类”都是复合结构体，但 Go 语言中结构体的内嵌配合接口比面向对象具有更高的扩展性和灵活性。  
Go 语言不仅认为结构体能拥有方法，且每种自定义类型也可以拥有自己的方法。  

**结构体定义**  
Go 语言可以通过自定义的方式形成新的类型，结构体就是这些类型中的一种复合类型，结构体是由零个或多个任意类型的值聚合成的实体，每个值都可以称为结构体的成员。  
结构体成员也可以称为“字段”，这些字段有以下特性：
- 字段拥有自己的类型和值；
- 字段名必须唯一；
- 字段的类型也可以是结构体，甚至是字段所在结构体的类型。

使用关键字 type 可以将各种基本类型定义为自定义类型，基本类型包括整型、字符串、布尔等。结构体是一种复合的基本类型，通过 type 定义为自定义类型后，使结构体更便于使用。  
`结构体的定义只是一种内存布局的描述，只有当结构体实例化时，才会真正地分配内存。`
```go
// 结构体的定义格式
type 类型名 struct {
    字段1 字段1类型
    字段2 字段2类型
    …
}

// 使用结构体可以表示一个包含 X 和 Y 整型分量的点结构
type Point struct {
    X int
    Y int
}
// 同类型的变量也可以写在一行，颜色的红、绿、蓝 3 个分量可以使用 byte 类型表示
type Color struct {
    R, G, B byte
}
```

**实例化结构体——为结构体分配内存并初始化**  
结构体的定义只是一种内存布局的描述，只有当结构体实例化时，才会真正地分配内存，因此必须在定义结构体并实例化后才能使用结构体的字段。  
实例化就是根据结构体定义的格式创建一份与格式一致的内存区域，结构体实例与实例间的内存是完全独立的。  
Go 语言可以通过多种方式实例化结构体，根据实际需要可以选用不同的写法。  

基本的实例化形式：var ins T。其中，T 为结构体类型，ins 为结构体的实例。  
```go
type Point struct {
    X int
    Y int
}
var p Point
// 使用.来访问结构体的成员变量
p.X = 10
p.Y = 20
```

Go语言中，还可以使用 new 关键字对类型（包括结构体、整型、浮点数、字符串等）进行实例化，结构体在实例化后会形成指针类型的结构体。  
创建指针类型的结构体格式：ins := new(T)。其中，T 为类型，可以是结构体、整型、字符串等；ins：T 类型被实例化后保存到 ins 变量中，ins 的类型为 *T，属于指针。
```go
type Player struct{
    Name string
    HealthPoint int
    MagicPoint int
}
tank := new(Player)
// 经过 new 实例化的结构体实例在成员赋值上与基本实例化的写法一致
// 访问结构体指针的成员变量时可以继续使用.，这是因为Go语言为了方便开发者访问结构体指针的成员变量，使用了语法糖（Syntactic sugar）技术，将 ins.Name 形式转换为 (*ins).Name
tank.Name = "Canon"
tank.HealthPoint = 300
```

取结构体的地址实例化。  
在 Go 语言中，对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作，取地址格式如下：ins := &T{}。其中：T 表示结构体类型；ins 为结构体的实例，类型为 *T，是指针类型。
```go
type Command struct {
    Name    string    // 指令名称
    Var     *int      // 指令绑定的变量
    Comment string    // 指令的注释
}
var version int = 1
cmd := &Command{}
cmd.Name = "version"
cmd.Var = &version
cmd.Comment = "show version"

// 取地址实例化是最广泛的一种结构体实例化方式，可以使用函数封装上面的初始化过程
func newCommand(name string, varref *int, comment string) *Command {
    return &Command{
        Name:    name,
        Var:     varref,
        Comment: comment,
    }
}
cmd = newCommand(
    "version",
    &version,
    "show version",
)
```

**初始化结构体的成员变量**  
结构体在实例化时可以直接对成员变量进行初始化，初始化有两种形式分别是以字段“键值对”形式和多个值的列表形式，键值对形式的初始化适合选择性填充字段较多的结构体，多个值的列表形式适合填充字段较少的结构体。  

使用“键值对”初始化结构体  
结构体可以使用“键值对”（Key value pair）初始化字段，每个“键”（Key）对应结构体中的一个字段，键的“值”（Value）对应字段需要初始化的值。  
键值对的填充是可选的，不需要初始化的字段可以不填入初始化列表中。    
结构体实例化后字段的默认值是字段类型的默认值，例如 ，数值为 0、字符串为 ""（空字符串）、布尔为 false、指针为 nil 等。  
```go
// 键值对初始化结构体的书写格式：键值之间以:分隔，键值对之间以,分隔
ins := 结构体类型名{ // 定义结构体时的类型名称
    // 结构体成员的字段名，结构体类型名的字段初始化列表中，字段名只能出现一次
    字段1: 字段1的值, // 结构体成员字段的初始值
    字段2: 字段2的值,
    …
}

type People struct {
    name  string
    child *People
}
relation := &People{
    name: "爷爷",
    // 结构体成员中只能包含结构体的指针类型，包含非指针类型会引起编译错误
    child: &People{
        name: "爸爸",
        child: &People{
                name: "我",
        },
    },
}
```

使用多个值的列表初始化结构体  
Go语言可以在“键值对”初始化的基础上忽略“键”，也就是说，可以使用多个值的列表初始化结构体的字段。使用这种格式初始化时，需要注意：
- 必须初始化结构体的所有字段。
- 每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致。
- 键值对与值列表的初始化形式不能混用。  
```go
// 多个值列表初始化结构体的书写格式
ins := 结构体类型名{
    字段1的值,
    字段2的值,
    …
}

type Address struct {
    Province    string
    City        string
    ZipCode     int
    PhoneNumber string
}
addr := Address{
    "四川",
    "成都",
    610000,
    "0",
}
fmt.Println(addr) // {四川 成都 610000 0}
```

初始化匿名结构体  
匿名结构体没有类型名称，无须通过 type 关键字定义就可以直接使用。  
匿名结构体的初始化写法由结构体定义和键值对初始化两部分组成，结构体定义时没有结构体类型名，只有字段和类型定义，键值对初始化部分由可选的多个键值对组成。  
匿名结构体的类型名是结构体包含字段成员的详细描述，匿名结构体在使用时需要重新定义，造成大量重复的代码，因此开发中较少使用。  
`注意：即使不为任何属性赋值，第二个大括号也是必不可少的，否则将引发编译时错误，程序无法被编译和运行。`
```go
// 匿名结构体定义格式和初始化写法
// 匿名结构体的初始化写法由结构体定义和键值对初始化两部分组成，结构体定义时没有结构体类型名，只有字段和类型定义，键值对初始化部分由可选的多个键值对组成
ins := struct {
    // 匿名结构体字段定义
    字段1 字段类型1
    字段2 字段类型2
    …
}{
    // 字段值初始化
    初始化字段1: 字段1的值,
    初始化字段2: 字段2的值,
    …
}

package main
import (
    "fmt"
)
// 打印消息类型, 传入匿名结构体
func printMsgType(msg *struct {
    id   int
    data string
}) {
    // 使用动词%T打印msg的类型
    fmt.Printf("%T\n", msg)
}
func main() {
    // 实例化一个匿名结构体
    msg := &struct {  // 定义部分
        id   int
        data string
    }{  // 值初始化部分
        1024,
        "hello",
    }
    printMsgType(msg) // *struct { id int; data string }
}
```

**构造函数**  
Go 语言的类型或结构体没有构造函数的功能，但是我们可以使用结构体初始化的过程来模拟实现构造函数。  

多种方式创建和初始化结构体——模拟构造函数重载  
由于 Go 语言中没有函数重载，为了避免函数名字冲突，使用 NewCatByName() 和 NewCatByColor() 两个不同的函数名表示不同的 Cat 构造过程。  
```go
type Cat struct {
    Color string
    Name  string
}
func NewCatByName(name string) *Cat {
    return &Cat{
        Name: name,
    }
}
func NewCatByColor(color string) *Cat {
    return &Cat{
        Color: color,
    }
}
```

带有父子关系的结构体的构造和初始化——模拟父级构造调用  
黑猫是一种猫，猫是黑猫的一种泛称，同时描述这两种概念时，就是派生，黑猫派生自猫的种类，使用结构体描述猫和黑猫的关系时，将猫（Cat）的结构体嵌入到黑猫（BlackCat）中，表示黑猫拥有猫的特性，然后再使用两个不同的构造函数分别构造出黑猫和猫两个结构体实例。  
```go
type Cat struct {
    Color string
    Name  string
}
type BlackCat struct {
    Cat  // 嵌入Cat, 类似于派生
}
// “构造基类”
func NewCat(name string) *Cat {
    return &Cat{
        Name: name,
    }
}
// “构造子类”
func NewBlackCat(color string) *BlackCat {
    cat := &BlackCat{}
    cat.Color = color
    return cat
}
```

**结构体方法**  
在 Go 语言中，方法和函数的定义格式非常像。由于方法和对象存在紧密的关系，因此在定义的格式上需要接收器。接收器变量和接收器类型共同构成了接收器；参数列表是可选的；返回参数也是可选的。  
```go
// 结构体方法格式
func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
    函数体
}

func (d *Dog) GetGender() string {
   if d.Gender == 0 {
      return "公"
   } else if d.Gender == 1 {
      return "母"
   }
   return ""
}
fatShibaInu := &Dog{
   Breed:  "Shiba Inu",
   Age:    2,
   Weight: 12.0,
   Gender: 0,
}
fmt.Println(fatShibaInu.GetGender()) // 公
```
为对象定义方法时，需要注意接收器的类型。使用指针与否，将决定了是否对原始变量产生影响。上面使用了*Dog，即指针类型，在方法中对该类型变量（d变量）的任何影响都将影响原始变量（fatShibaInu）；反之，若使用Dog类型，则不会影响。其原因是当不使用指针类型变量时，方法中的接收器变量实际上是对原始数据的“拷贝”，所做出的改变也仅仅会作用于这份“拷贝”的数据上，并不会影响到原始数据。  
```go
func (d *Dog) GrowUp() {
   d.Age++
}

func (d Dog) GrowUp2() {
   d.Age++
}

func main() {
   fatShibaInu := NewDog("Shiba Inu", 2, 12.0, "公")

   fatShibaInu.GrowUp()
   fmt.Println(fatShibaInu) // &{Shiba Inu 3 12 0}

   fatShibaInu.GrowUp2()
   fmt.Println(fatShibaInu) // &{Shiba Inu 3 12 0}
}
```

让方法操作对象动起来。
```go
func (d *Dog) Sport() {
   fmt.Println("做运动！")
   d.Weight -= 0.1
   fmt.Println("我减重到了", d.Weight)
}

func (d *Dog) Eat() {
   fmt.Println("多吃饭！")
   d.Weight += 0.1
   fmt.Println("我增重到了", d.Weight)
}
func main() {
   fatShibaInu := NewDog("Shiba Inu", 2, 12.0, "公")
   weakShibaInu := NewDog("Shiba Inu", 2, 7.0, "公")
   fatShibaInu.Sport() // 做运动！我减重到了 11.9
   weakShibaInu.Eat() // 多吃饭！我增重到了 7.1
}
```

**实现继承**  
从本质上说，Go 语言中继承，是通过结构体的嵌套来实现的。在 Go 语言中，相比较于继承，组合更受青睐。  
作为“子结构体”，不仅可以使用“父结构体”的属性，还拥有自己的属性。如果说 “父结构体”是概括的，抽象的，那么“子结构体”就是具体的，详细的。  
```go
type Animal struct {
   Name   int
   Age    int
   Gender string
}
type Bird struct {
   WingColor    string
   CommonAnimal Animal
}
// 创建Bird类型的构造函数
func NewBird(name string, age int, gender string, wingColor string) *Bird {
   return &Bird{
      WingColor: wingColor,
      CommonAnimal: Animal{
         Name:   name,
         Age:    age,
         Gender: gender,
      },
   }
}
// 创建Bird类型的“飞行”方法
func (b *Bird) Fly() {
   fmt.Println("我起飞啦！")
}
func main() {
   bird := *NewBird("小鸟", 1, "公", "绿色")
   fmt.Println(bird) // {绿色 {小鸟 1 公}}
   bird.Fly() // 我起飞啦！
}
```

**多态**  
继续定义子结构体狗（Dog），它拥有毛色（Color）属性。还有犬吠（Bark）动作。请读者参考上面小鸟（Bird）部分的代码，独立完成狗（Dog）部分的代码，要求依然使用构造函数（NewDog()）和方法（Bark()）。
```go
type Animal struct {
   Name   string
   Age    int
   Gender string
}

func (a *Animal) Eat() {
   fmt.Println(a.Name, "我要吃到饱！")
}

type Dog struct {
   Color        string
   CommonAnimal Animal
}

func NewDog(name string, age int, gender string, color string) *Dog {
   return &Dog{
      Color: color,
      CommonAnimal: Animal{
         Name:   name,
         Age:    age,
         Gender: gender,
      },
   }
}

func (d *Dog) Bark() {
   fmt.Println("汪汪汪！")
}

func main() {
   dog := *NewDog("小狗", 2, "公", "黄色")
   fmt.Println(dog) // {黄色 {小狗 2 公}}
   dog.Bark() // 汪汪汪！

   bird := *NewBird("小鸟", 1, "公", "绿色")
   bird.CommonAnimal.Eat() // 小鸟 我要吃到饱！
   dog.CommonAnimal.Eat() // 小狗 我要吃到饱！
}
```

匿名结构体嵌套  
Go 语言语法还允许开发者以一种更为简单的方式嵌套结构体使用，这种更简单的方式便是嵌套匿名结构体。在后期使用时，也会被简化。  
```go
type Animal struct {
   Name   string
   Age    int
   Gender string
}

func (a *Animal) Eat() {
   fmt.Println(a.Name, "我要吃到饱！")
}

type Bird struct {
   string
   Animal
}

func NewBird(name string, age int, gender string, wingColor string) *Bird {
   return &Bird{
      wingColor,
      Animal{
         name,
         age,
         gender,
      },
   }
}

func (b *Bird) Fly() {
   fmt.Println("我起飞啦！")
}

func main() {
   bird := *NewBird("小鸟", 1, "公", "绿色")
   //访问string类型成员
   fmt.Println(bird.string) // 绿色
   //访问Name成员
   fmt.Println(bird.Name) // 小鸟
   bird.Eat() // 鸟 我要吃到饱
}
```

