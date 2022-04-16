
### Go 常量
常量的值在程序运行期间是不能改变的，而变量的值在运行期间是可以改变的。在使用时，只要你确定在程序运行期间不改变它的值，就可以使用常量。  
对于常量值的数据类型，只可以定义为布尔型、数字型（整数型、浮点型和复数）和字符串型。  

`注意： 声明常量时，必须为其赋值，且后续无法修改。`

```go
//常量声明 const name [type] = value
const WIDTH_OF_RECT int = 12
const ALLOW_DOWNLOAD_WHEN_WIFI bool = true

//声明一个名为 PI 的常量，类型为 float64（浮点数类型）
const PI float64 = 3.14

// 同时定义多个常量
// 第一种，常量块的形式
const (
    isOpen = true
    MyRune = 'r'
)
// 第二种，并行赋值
const limit, rate = 12, 29.8

// 隐式定义不限制，即省略了数据类型后，值的大小是不受限制，即不会产生溢出
const num = 111111111111111111111111111111111111111111111
```

### iota 
iota 是 Go 语言中的一个关键字。  
iota 从 0 开始，每增加新的一行就会自动加一，直到重新声明一个 const 时 iota 重置为 0，遇到新的赋值时 iota 将暂时不再应用（再次使用，iota 继续保持在增加新的一行时自增一）。  
```go
const (
    a = iota  // a = 0
    b         // b = 1
    c         // c = 2
    d = 5     // d = 5   
    e         // e = 5
)

const (
    a = iota  // a = 0
    b         // b = 1
    c         // c = 2
    d = 5     // d = 5   
    e = iota  // e = 4

)
```

iota 也可以参加运算。
```go
// 从1开始自动加一
const (
    Apple = iota + 1 // Apple=1 
    Cherimoya        // Cherimoya=2 
    Elderberry       // Elderberry=3
)

// 并行赋值两个常量，iota 只会在第一行增长一次
// 而不会因为使用了两次就增长两次
const (
    Apple, Banana = iota + 1, iota + 2 // Apple=1 Banana=2
    Cherimoya, Durian   // Cherimoya=2 Durian=3
    Elderberry, Fig     // Elderberry=3, Fig=4
)

// iota参与位运算
const (
    Open = 1 << iota  // 0001
    Close             // 0010
    Pending           // 0100
)
```

