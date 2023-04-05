
### 常见问题
- make 和 new 的区别
> make 会进行初始化，new 会返回一个零值的指针。  
> 
> new (T) 和 make (T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。
> new (T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T 的值。换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。
> make (T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make () 只适用于 slice、map 和 channel.

- 在多行 array、slice、map 语句中为什么最后一项需要，号
> 
```go
x := []int {
    1,
    2    // syntax error: unexpected newline, expecting comma or }
}
```

- Go 的 Slice 如何扩容？
> 在使用 append 向 slice 追加元素时，若 slice 空间不足则会发生扩容，扩容会重新分配一块更大的内存，将原 slice 拷贝到新 slice ，然后返回新 slice。扩容后再将数据追加进去。  
> 扩容操作只对容量，扩容后的 slice 长度不变，容量变化规则如下：  
> 若 slice 容量小于 1024 个元素，那么扩容的时候 slice 的 cap 就翻番，乘以 2；一旦元素个数超过 1024 个元素，增长因子就变成 1.25，即每次增加原来容量的四分之一。  
> 若 slice 容量够用，则将新元素追加进去，slice.len++，返回原 slice。  
> 若 slice 容量不够用，将 slice 先扩容，扩容得到新 slice，将新元素追加进新 slice，slice.len++，返回新 slice。  

- 使用值为 nil 的 slice、map会发生什么
> 允许对值为 nil 的 slice 添加元素，但对值为 nil 的 map 添加元素，则会造成运行时 panic。
```go
// map 错误示例
func main() {
    var m map[string]int
    m["one"] = 1  // error: panic: assignment to entry in nil map
    // m := make(map[string]int)// map 的正确声明，分配了实际的内存
}    

// slice 正确示例
func main() {
 var s []int
 s = append(s, 1)
}
```

- 访问 map 中的 key，需要注意什么
> 当访问 map 中不存在的 key 时，Go 则会返回元素对应数据类型的零值，比如 nil、’’ 、false 和 0，取值操作总有值返回，故不能通过取出来的值，来判断 key 是不是在 map 中。
> 检查 key 是否存在可以用 map 直接访问，检查返回的第二个参数即可。
```go
// 错误的 key 检测方式
func main() {
 x := map[string]string{"one": "2", "two": "", "three": "3"}
 if v := x["two"]; v == "" {
  fmt.Println("key two is no entry") // 键 two 存不存在都会返回的空字符串
 }
}

// 正确示例
func main() {
 x := map[string]string{"one": "2", "two": "", "three": "3"}
 if _, ok := x["two"]; !ok {
  fmt.Println("key two is no entry")
 }
}
```

- string 类型的值可以修改吗
> 不能，尝试使用索引遍历字符串，来更新字符串中的个别字符，是不允许的。
> string 类型的值是只读的二进制 byte slice，如果真要修改字符串中的字符，将 string 转为 []byte 修改后，再转为 string 即可。
```go
// 修改字符串的错误示例
func main() {
 x := "text"
 x[0] = "T"  // error: cannot assign to x[0]
 fmt.Println(x)
}


// 修改示例
func main() {
 x := "text"
 xBytes := []byte(x)
 xBytes[0] = 'T' // 注意此时的 T 是 rune 类型
 x = string(xBytes)
 fmt.Println(x) // Text
}
```

- switch 中如何强制执行下一个 case 代码块
> switch 语句中的 case 代码块会默认带上 break，但可以使用 fallthrough 来强制执行下一个 case 代码块。
```go
func main() {
 isSpace := func(char byte) bool {
  switch char {
  case ' ': // 空格符会直接 break，返回 false // 和其他语言不一样
  fallthrough // 返回 true
  case '\t':
   return true
  }
  return false
 }
 fmt.Println(isSpace('\t')) // true
 fmt.Println(isSpace(' ')) // false
}
```

- 如何关闭 HTTP 的响应体
> 直接在处理 HTTP 响应错误的代码块中，直接关闭非 nil 的响应体；手动调用 defer 来关闭响应体。
```go
// 正确示例
func main() {
 resp, err := http.Get("http://www.baidu.com")

    // 关闭 resp.Body 的正确姿势
    if resp != nil {
  defer resp.Body.Close()
 }

 checkError(err)
 defer resp.Body.Close()

 body, err := ioutil.ReadAll(resp.Body)
 checkError(err)

 fmt.Println(string(body))
}
```

- 主动关闭过 HTTP 连接
不关闭会程序可能会消耗完 socket 描述符。有如下2种关闭方式：
```go
// 方式一：直接设置请求变量的 Close 字段值为 true，每次请求结束后就会主动关闭连接。设置 Header 请求头部选项 Connection: close，然后服务器返回的响应头部也会有这个选项，此时 HTTP 标准库会主动断开连接
// 主动关闭连接
func main() {
 req, err := http.NewRequest("GET", "http://golang.org", nil)
 checkError(err)

 req.Close = true
 //req.Header.Add("Connection", "close") // 等效的关闭方式

 resp, err := http.DefaultClient.Do(req)
 if resp != nil {
  defer resp.Body.Close()
 }
 checkError(err)

 body, err := ioutil.ReadAll(resp.Body)
 checkError(err)

 fmt.Println(string(body))
}

// 创建一个自定义配置的 HTTP transport 客户端，用来取消 HTTP 全局的复用连接
func main() {
 tr := http.Transport{DisableKeepAlives: true}
 client := http.Client{Transport: &tr}

 resp, err := client.Get("https://golang.google.cn/")
 if resp != nil {
  defer resp.Body.Close()
 }
 checkError(err)

 fmt.Println(resp.StatusCode) // 200

 body, err := ioutil.ReadAll(resp.Body)
 checkError(err)

 fmt.Println(len(string(body)))
}
```

- 解析 JSON 数据时，默认将数值当做哪种类型
> 默认将数值解析为 float64 类型，如果需要解析为 int 类型，需要使用 json.Unmarshal 的第二个参数，传入一个指向 int 类型的指针。
```go
// 解析 JSON 数据时，默认将数值当做 float64 类型
func main() {
 jsonStr := `{"num":6.13}`
 var data map[string]interface{}
 err := json.Unmarshal([]byte(jsonStr), &data)
 checkError(err)

 fmt.Println(data["num"]) // 6.13
 fmt.Printf("%T\n", data["num"]) // float64
}
```

- 如何从 panic 中恢复
> 在一个 defer 延迟执行的函数中调用 recover ，它便能捕捉/中断 panic。
```go
// 错误的 recover 调用示例
func main() {
 recover() // 什么都不会捕捉
 panic("not good") // 发生 panic，主程序退出
 recover() // 不会被执行
 println("ok")
}

// 正确的 recover 调用示例
func main() {
 defer func() {
  fmt.Println("recovered: ", recover())
 }()
 panic("not good")
}
```

- 简短声明的变量需要注意什么
> 简短声明的变量只能在函数内部使用，不能在函数外部使用
> struct 的变量字段不能使用 := 来赋值
> 不能用简短声明方式来单独为一个变量重复声明， := 左侧至少有一个新变量，才允许多变量的重复声明

- range 迭代 map 是有序的吗
> map 是无序的，range 迭代 map 时，每次迭代的顺序都是不一样的。
> Go 的运行时是有意打乱迭代顺序的，这样可以避免程序依赖于 map 的迭代顺序，从而导致程序的不可预测性。但也并不总会打乱，得到连续相同的 5 个迭代结果也是可能的。
> 如果需要有序的迭代，可以将 map 的 key 排序后再迭代。
```go
// map 的 key 排序后再迭代
func main() {
 m := map[int]string{3: "c", 1: "a", 2: "b"}
 keys := make([]int, 0, len(m))
 for k := range m {
  keys = append(keys, k)
 }
 sort.Ints(keys)
 for _, k := range keys {
  fmt.Println(k, m[k])
 }
}
```

- recover 的执行时机
> recover 只能在 defer 延迟调用的函数中执行，否则会导致程序 panic。
```go
// recover 只能在 defer 延迟调用的函数中执行
func main() {
 recover() // 什么都不会捕捉
 panic("not good") // 发生 panic，主程序退出
 recover() // 不会被执行
 println("ok")
}

// 必须在 defer 函数中直接调用才有效
func main() {
    defer func() {
        recover()
    }()
    panic(1)
}
```

- 闭包错误引用同一个变量问题怎么处理
> 在每轮迭代中生成一个局部变量 i 。如果没有 i := i 这行，将会打印同一个变量。
> 或者是通过函数参数传入 i 。
```go
func main() {
    for i := 0; i < 5; i++ {
        i := i
        defer func() {
            println(i)
        }()
    }
}

func main() {
    for i := 0; i < 5; i++ {
        defer func(i int) {
            println(i)
        }(i)
    }
}
```

- 在循环内部执行defer语句会发生什么
> defer 语句会在函数返回时执行，而不是在循环结束时执行。
> 在 for 执行 defer 会导致资源延迟释放。
```go
func main() {
    for i := 0; i < 5; i++ {
        defer func() {
            println(i)
        }()
    }
}
```

- 避免 goroutine 泄露的措施
> golang 中的 Goroutine 泄露指的是指程序创建了 Goroutine，但没有正确地关闭或管理它，导致 Goroutine 无法正常退出，继续占用系统资源
> 1. 通过 context 控制 Goroutine 的生命周期
> 2. 通过 sync.WaitGroup 等待 Goroutine 执行完毕
> 3. 通过 channel 通知 Goroutine 退出
> 4. 通过 sync.Once 确保 Goroutine 只执行一次
```go
// 通过 context 控制 Goroutine 的生命周期
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go func() {
        defer cancel()
        for {
            select {
            //  for 循环停止取数据时，就用 cancel 函数，让另一个协程停止写数据
            case <-ctx.Done():
                return
            default:
                println("hello")
            }
        }
    }()
    time.Sleep(time.Second)
}
```

- 如何跳出 for select 循环
> 通常在 for 循环中，使用 break 可以跳出循环，但是注意在 go 语言中，for select 配合时，break 并不能跳出循环，需要使用 return 语句或 goto 语句。
```go
func testSelectFor2(chExit chan bool){
 EXIT:
    for  {
        select {
        case v, ok := <-chExit:
            if !ok {
                fmt.Println("close channel 2", v)
                // break EXIT
                goto EXIT2
            }

            fmt.Println("ch2 val =", v)
        }
    }

    //EXIT2:
    fmt.Println("exit testSelectFor2")
}
```

- 如何在切片中查找
> 1. 使用 for 循环遍历切片，查找元素
> 2. 使用 sort.Search 函数查找元素
> 3. 使用 sort.searchXXX 方法，在排序好的切片中查找指定的方法，但是其返回是对应的查找元素不存在时，返回待插入的位置下标(元素插入在返回下标前)
```go
// 使用 for 循环遍历切片，查找元素
func main() {
    s := []int{1, 2, 3, 4, 5}
    for i, v := range s {
        if v == 3 {
            fmt.Println(i)
            break
        }
    }
}

// 使用 sort.Search 函数查找元素
func main() {
    s := []int{1, 2, 3, 4, 5}
    i := sort.Search(len(s), func(i int) bool { return s[i] >= 3 })
    if i < len(s) && s[i] == 3 {
        fmt.Println(i)
    }
}
// 使用 sort.searchXXX 方法
func IsExist(s []string, t string) (int, bool) {
    iIndex := sort.SearchStrings(s, t)
    bExist := iIndex!=len(s) && s[iIndex]==t

    return iIndex, bExist
}
```

- 如何初始化带嵌套结构的结构体
> go 的哲学是组合优于继承，使用 struct 嵌套即可完成组合，内嵌的结构体属性就像外层结构的属性即可，可以直接调用。
> 注意初始化外层结构体时，必须指定内嵌结构体名称的结构体初始化，如下看到 s1方式报错，s2 方式正确。
```go
type stPeople struct {
    Gender bool
    Name string
}

type stStudent struct {
    stPeople
    Class int
}

//尝试4 嵌套结构的初始化表达式
//var s1 = stStudent{false, "JimWen", 3}
var s2 = stStudent{stPeople{false, "JimWen"}, 3}
fmt.Println(s2.Gender, s2.Name, s2.Class)
```

- 切片和数组的区别
> 1. 数组是值类型，切片是引用类型
> 2. 数组的长度是固定的，切片的长度是不固定的
> 3. 数组的长度是数组类型的一部分，切片的长度不是切片类型的一部分
> 4. 数组可以使用 == 或 != 操作符比较，切片不可以
> 5. 数组可以使用 len() 方法获取长度，切片可以使用 len() 方法获取长度
> 6. 切片可以使用 cap() 方法获取容量
> 7. 切片可以使用 append() 方法追加元素
> 8. 切片可以使用 copy() 方法复制元素
> 9. 切片可以使用 make() 方法创建切片
> 10. 切片可以使用 new() 方法创建切片
> 
> 数组是具有固定长度，且拥有零个或者多个，相同数据类型元素的序列。数组的长度是数组类型的一部分，所以 [3]int 和 [4]int 是两种不同的数组类型。数组需要指定大小，不指定也会根据初始化的自动推算出大小，不可改变；数组是值传递。数组是内置类型，是一组同类型数据的集合，它是值类型，通过从0开始的下标索引访问元素值。在初始化后长度是固定的，无法修改其长度。
> 当作为方法的参数传入时将复制一份数组而不是引用同一指针。数组的长度也是其类型的一部分，通过内置函数 len(array) 获取其长度。
> 
> 切片表示一个拥有相同类型元素的可变长度的序列。切片是一种轻量级的数据结构，它有三个属性：指针、长度和容量。切片不需要指定大小；切片是地址传递；切片可以通过数组来初始化，也可以通过内置函数 make() 初始化 。  
> 初始化时 len=cap,在追加元素时如果容量 cap 不足时将按 len 的 2 倍扩容，直到超过 cap 时，再按 cap 的 2 倍扩容。
```go
var array [10]int
var array =[5]int{1,2,3,4,5}

var slice []type = make([]type, len)
```

- Printf()、Sprintf()、Fprintf() 函数的区别用法是什么？
> Printf() 函数用于格式化并输出字符串，格式化的字符串会写入到标准输出 os.Stdout 中，返回写入的字节数和遇到的任何错误。
> Sprintf() 函数用于格式化并返回一个字符串而不带任何输出。
> Fprintf() 函数用于格式化并输出字符串，格式化的字符串会写入到 io.Writer 接口类型的变量 w 中，返回写入的字节数和遇到的任何错误。Fprintf() 是格式化输出到一个stream，通常是到文件。

- 关于 for 循环
> 1. for 循环的三个组成部分都是可选的，但是分号是必须的
> 2. for 循环的第一个组成部分是一个简单的语句，该语句通常是一个短变量声明，该变量声明仅在 for 语句的作用域中可见
> 3. for 循环的第二个组成部分是一个布尔表达式，该表达式用于判断循环是否继续执行
> 4. for 循环的第三个组成部分是一个语句，该语句在每次迭代后执行
> 5. for 循环的第二个组成部分和第三个组成部分都是可选的，如果省略第二个组成部分，则表示该循环不会在执行前检查条件，也就是说该循环会无限次地执行下去，直到在循环体内使用了 break 或者 return 语句跳出循环为止
> 6. 如果省略第三个组成部分，则表示在每次迭代后都不会执行任何操作
> 7. 如果省略了 for 循环的所有组成部分，则表示该循环会无限次地执行下去，直到在循环体内使用了 break 或者 return 语句跳出循环为止
> 8. for 循环的第一个组成部分还可以是一个并行的赋值语句，该语句可以在一次迭代中声明多个变量，这些变量的作用域仅在 for 循环的作用域中；需要注意的是，不支持以逗号为间隔的多个赋值语句
> 9. for 循环的第二个组成部分还可以是一个 range 表达式，该表达式用于迭代数组、切片、字符串、map 或者通道（channel）
```go
// for 循环的第一个组成部分还可以是一个并行的赋值语句
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```

- switch 语句
> 单个 case 中，可以出现多个结果选项。只有在 case 中明确添加 fallthrough关键字，才会继续执行紧跟的下一个 case。

- Array 类型的值作为函数参数
> 数组是值类型，因此在函数调用时会将数组的副本传递给函数，而不是数组的指针
> 想改变数组，直接传递指向这个数组的指针类型
> 直接使用 slice：即使函数内部得到的是 slice 的值拷贝，但依旧会更新 slice 的原始数据（底层 array）

- go 语言中有没有隐藏的 this 指针
> 没有，go 语言中没有 this 指针，但是有一个隐式的参数，就是接收者 receiver，它是一个指针类型，指向调用该方法的对象
> go 面向对象表达更直观，对于面向过程只是换了一种语法形式来表达方法施加的对象不需要非得是指针，也不用非得叫 this

- go 语言中的引用类型包含哪些
> 切片(slice)、字典(map)、通道（channel）、接口（interface）。

- go 语言的 main 函数
> go 语言的 main 函数是一个特殊的函数，它是程序的入口
> main 函数不能带参数；main 函数不能定义返回值
> main 函数所在的包必须为 main 包
> main 函数中可以使用 flag 包来获取和解析命令行参数

- go 语言触发异常的场景有哪些
> 1. 数组越界
> 2. 空指针引用
> 3. 类型断言失败
> 4. 除数为 0
> 5. 关闭已关闭的 channel
> 6. 向已关闭的 channel 发送数据
> 7. 从已关闭的 channel 接收数据
> 8. 调用 panic

- go 语言的 select 机制
> select 语句使一个 goroutine 在多个通信操作上等待，select 常用于 goroutine 的完美退出
> select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支
> 当多个条件分支都可以执行时，select 会随机地选择一个执行
> select 语句中的每个 case 都必须是一个通信操作，要么是发送要么是接收
> select 语句中的 default 分支总是可运行的
> select 语句中的 default 分支可以省略
> 
> golang 的 select 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作每个 case 语句里必须是一个 IO 操作，确切的说，应该是一个面向 channel 的 IO 操作
> select 机制用来处理异步 IO 问题，比如 socket 编程中的读写操作
> select 机制最大的一条限制就是每个 case 语句里必须是一个 IO 操作
> golang 在语言级别支持 select 关键字

- go 语言中的静态类型声明
> 静态类型声明是告诉编译器不需要太多的关注这个变量的细节
> 静态类型声明是告诉编译器这个变量的类型，编译器会在编译的时候检查类型是否正确
> 静态变量的声明，只是针对于编译的时候, 在连接程序的时候，编译器还要对这个变量进行实际的声明
```go
// 静态类型声明
var a int
// 动态类型声明
var a = 1
```

- go 的接口是什么
> 在 go 语言中，interface 也就是接口，被用来指定一个对象
> 接口是一组方法签名的集合，接口类型的变量可以保存任何实现了这些方法的值
> 接口类型的变量可以保存任何实现了这些方法的值
> 在 go 中使用 interface 来实现多态
```go
// 使用 interface 来实现多态
type Animal interface {
    Eat()
    Move()
    Speak()
}

type Cow struct {
    name string
}

func (c Cow) Eat() {
    fmt.Println("grass")
}
func (c Cow) Move() {
    fmt.Println("walk")
}
func (c Cow) Speak() {
    fmt.Println("moo")
}

type Bird struct {
    name string
}

func (b Bird) Eat() {
    fmt.Println("worms")
}

func (b Bird) Move() {
    fmt.Println("fly")
}

func (b Bird) Speak() {
    fmt.Println("peep")
}
```

- go 类型断言是怎么回事
> 类型断言是用来从一个接口里面读取数值给一个具体的类型变量。类型转换是指转换两个不相同的数据类型。
> 类型断言的语法格式如下：
```go
x.(T)
```

- go 局部变量和全局变量的缺省值是什么
> 缺省值是与这个类型相关的零值
> 类型的零值，比如 int 的零值是 0，string 的零值是 ""，bool 的零值是 false，指针的零值是 nil，map 的零值是 nil，slice 的零值是 nil，channel 的零值是 nil，interface 的零值是 nil

- 模块化编程是怎么回事
> 模块化编程是一种编程思想，它将一个复杂的系统分解成多个模块，每个模块都是一个独立的单元，它们之间相互依赖，但是又相互独立，这样就可以很方便地对系统进行维护和扩展。

- go 函数和方法
> go 函数的定义声明没有接收者
> 方法的声明和函数类似，他们的区别是：方法在定义的时候，会在func和方法名之间增加一个参数，这个参数就是接收者，这样我们定义的这个方法就和接收者绑定在了一起，称之为这个接收者的方法
> go 有两种类型的接收者：值接收者和指针接收者。使用值类型接收者定义的方法，在调用的时候，使用的其实是值接收者的一个副本，所以对该值的任何操作，不会影响原来的类型变量。——-相当于形式参数。
> 如果我们使用一个指针作为接收者，那么就会其作用了，因为指针接收者传递的是一个指向原值指针的副本，指针的副本，指向的还是原来类型的值，所以修改时，同时也会影响原来类型变量的值。
> 
> 方法的调用和函数的调用非常类似，只是在调用的时候，需要在方法名前面加上接收者的变量名

- 关于 go 的可变参数
> 可变参数是指函数的最后一个参数是以 ...type 的形式出现的，这种形式的参数可以接受任意多个 type 类型的参数
> 可变参数的类型必须是同一类型
> 比如常用的 fmt.Println() 就是一个可变参数的函数
```go
func main() {
 print("1","2","3")
}


func print (a ...interface{}){
 for _,v:=range a{
  fmt.Print(v)
 }
 fmt.Println()
}
```

- go slice 底层实现
> 切片是基于数组实现的，它的底层是数组，它自己本身非常小，可以理解为对底层数组的抽象。因为基于数组实现，所以它的底层的内存是连续分配的，效率非常高，还可以通过索引获得数据，可以迭代以及垃圾回收优化。
> 切片本身并不是动态数组或者数组指针。它内部实现的数据结构通过指针引用底层数组，设定相关属性将数据读写操作限定在指定的区域内。切片本身是一个只读对象，其工作机制类似数组指针的一种封装。
> 切片对象非常小，是因为它是只有3个字段的数据结构：指向底层数组的指针、切片的长度、切片的容量
> 
> slice 是一个引用类型，它的内部结构包含地址、长度和容量
> slice 的底层是一个数组，slice 通过指针、长度和容量来描述一个数组的子集
> slice 的长度就是它所包含的元素个数
> slice 的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数
> slice 的零值是 nil
> slice 的 len 和 cap 可以通过内置函数 len() 和 cap() 来获取
> slice 的赋值操作会复制整个 slice，包括它所引用的底层数组
> slice 的 == 操作符比较的是底层数组的指针和长度
> slice 的每个元素都可以通过索引号来访问，索引号从 0 开始
> slice 可以通过内置的 append() 函数来进行追加元素
> slice 可以通过内置的 copy() 函数来进行拷贝
> slice 可以通过内置的 delete() 函数来进行删除元素
> slice 可以通过内置的 make() 函数来进行创建
> slice 可以通过内置的 new() 函数来进行创建

- go slice 扩容机制，有什么注意点
> 首先判断，如果新申请容量大于 2 倍的旧容量，最终容量就是新申请的容量。否则判断，如果旧切片的长度小于 1024，则最终容量就是旧容量的两倍
> 否则判断，如果旧切片长度大于等于 1024，则最终容量从旧容量开始循环增加原来的 1/4 , 直到最终容量大于等于新申请的容量。如果最终容量计算值溢出，则最终容量就是新申请容量
> 情况一：原数组还有容量可以扩容（实际容量没有填充完），这种情况下，扩容以后的数组还是指向原来的数组，对一个切片的操作可能影响多个指针指向相同地址的 Slice
> 情况二：原来数组的容量已经达到了最大值，再想扩容， Go 默认会先开一片内存区域，把原来的值拷贝过来，然后再执行 append() 操作。这种情况丝毫不影响原数组
> 要复制一个 Slice，最好使用 Copy 函数

- go map 底层实现
> Golang 中 map 的底层实现是一个散列表，因此实现 map 的过程实际上就是实现散表的过程。
> 在这个散列表中，主要出现的结构体有两个，一个叫 hmap(a header for a go map)，一个叫 bmap(a bucket for a Go map，通常叫其 bucket)。
> 哈希表的主要思想是，通过计算键的哈希值，将键和值映射到表中的一个桶（Bucket）中。当查询一个键的值时，系统可以通过计算键的哈希值，快速定位到该键所在的桶，从而获取该键对应的值
> 
> map 是一种无序的基于 key-value 的数据结构，Go 语言中的 map 是引用类型，必须初始化才能使用
> map 的声明必须带上 key 和 value 的类型
> map 的初始化可以使用内建函数 make()，也可以使用 map 关键字
> map 的长度是不固定的，也就是和 slice 一样，也是一种引用类型
> map 的 value 可以很方便的修改，通过 key 来修改
> map 的 value 可以通过 key 来获取，语法格式：map[key]
> map 的 value 也可以通过另外一个 map 来获取，语法格式：map1[map2[key]]
> map 的 value 可以通过 make() 函数来创建，语法格式：make(map[keytype]valuetype)
> map 的 value 可以通过 new() 函数来创建，语法格式：new(map[keytype]valuetype)
> map 的 value 可以通过字面量来创建，语法格式：map[keytype]valuetype{key1:val1,key2:val2}
> map 的 value 可以通过 map 关键字来创建，语法格式：map[keytype]valuetype{key1:val1,key2:val2}
> map 的 value 可以通过 map 的内置函数 make() 来创建，语法格式：make(map[keytype]valuetype)
> map 的 value 可以通过 map 的内置函数 new() 来创建，语法格式：new(map[keytype]valuetype)
> map 的 value 可以通过 map 的内置函数 delete() 来删除，语法格式：delete(map, key)
> map 的 value 可以通过 map 的内置函数 len() 来获取长度，语法格式：len(map)
> map 的 value 可以通过 map 的内置函数 cap() 来获取容量，语法格式：cap(map)
[具体参考](https://golang.design/go-questions/map/principal/)

- JSON 标准库对 nil slice 和 空 slice 的处理是一致的吗
> 不一致，nil slice 序列化后是 null，空 slice 序列化后是 []，这是因为 nil slice 底层没有数组，而空 slice 底层有数组，只是数组长度为 0
> 通常错误的用法，会报数组越界的错误，因为只是声明了slice，却没有给实例化的对象
```go
// 错误用法
var slice []int
slice[1] = 0

// 正确用法
slice := make([]int,0）
// slice := []int{}
```

- go 的内存模型，为什么小对象多了会造成 gc 压力
> Go 语言的内存模型是一个抽象的概念，它描述了程序在执行过程中如何管理内存，以及内存管理的一些规则
> 通常小对象过多会导致 GC 三色法消耗过多的 GPU。优化思路是，减少对象分配。

- Data Race 问题怎么解决？能不能不加锁解决这个问题
> Golang 中的数据竞争是指在多个 Goroutine 同时访问和修改同一个变量时可能导致的问题。由于 Goroutine 是并发执行的，因此它们可以在任意时刻同时对同一个变量进行读写操作，导致结果不可预测和不一致。
> 同步访问共享数据是处理数据竞争的一种有效的方法
> golang 在 1.1 之后引入了竞争检测机制，可以使用 go run -race 或者 go build -race 来进行静态检测。其在内部的实现是,开启多个协程执行同一个命令， 并且记录下每个变量的状态
> 竞争检测器基于 C/C++ 的 ThreadSanitizer 运行时库，该库在 Google 内部代码基地和 Chromium 找到许多错误。这个技术在 2012 年 9 月集成到 Go 中，从那时开始，它已经在标准库中检测到 42 个竞争条件。现在，它已经是我们持续构建过程的一部分，当竞争条件出现时，它会继续捕捉到这些错误
> 竞争检测器已经完全集成到 Go 工具链中，仅仅添加 -race 标志到命令行就使用了检测器
> 
> 数据竞争是一个常见的并发问题，如果没有正确的处理方式，可能导致程序的错误行为和数据不一致。因此，在编写并发代码时，要特别注意数据竞争的问题，并采用合适的同步手段来避免它。
> 要想解决数据竞争的问题可以使用互斥锁 sync.Mutex,解决数据竞争(Data race),也可以使用管道解决,使用管道的效率要比互斥锁高

- 在 range 迭代 slice 时，怎么修改值
> 在 range 迭代中，得到的值其实是元素的一份值拷贝，更新拷贝并不会更改原来的元素，即是拷贝的地址并不是原有元素的地址
> 如果要修改原有元素的值，应该使用索引直接修改
```go
for i := range slice {
    slice[i] = 0
}
```
> 如果你的集合保存的是指向值的指针，需稍作修改。依旧需要使用索引访问元素，不过可以使用 range 出来的元素直接更新原有值
```go
func main() {
 data := []*struct{ num int }{{1}, {2}, {3},}
 for _, v := range data {
  v.num *= 10 // 直接使用指针更新
 }
 fmt.Println(data[0], data[1], data[2]) // &{10} &{20} &{30}
}
```

- nil interface 和 nil interface 的区别
> nil interface 是指 interface 的动态类型和动态值都为 nil 的 interface
> nil interface 的动态类型和动态值都为 nil，因此不能调用任何方法，也不能赋值给任何类型的变量
> nil interface 的零值是 nil，因此可以使用 == 操作符来判断是否为 nil
> nil interface 的类型是 nil，因此不能使用类型断言来获取动态类型
> 虽然 interface 看起来像指针类型，但它不是。interface 类型的变量只有在类型和值均为 nil 时，才为 nil；如果你的 interface 变量的值是跟随其他变量变化的，与 nil 比较相等时小心。如果你的函数返回值类型是 interface，更要小心如下这个坑
```go
func main() {
   var data *byte
   var in interface{}

   fmt.Println(data, data == nil) // <nil> true
   fmt.Println(in, in == nil) // <nil> true

   in = data
   fmt.Println(in, in == nil) // <nil> false // data 值为 nil，但 in 值不为 nil
}

// 正确示例
func main() {
  doIt := func(arg int) interface{} {
    var result *struct{} = nil

    if arg > 0 {
        result = &struct{}{}
    } else {
        return nil // 明确指明返回 nil
    }

    return result
  }


  if res := doIt(-1); res != nil {
    fmt.Println("Good result: ", res)
  } else {
    fmt.Println("Bad result: ", res) // Bad result: <nil>
  }
}
```





