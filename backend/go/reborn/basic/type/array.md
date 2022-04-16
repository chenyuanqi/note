
### Go 数组
数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。因为数组的长度是固定的，所以在 Go 语言中很少直接使用数组。  

数组的声明语法：var 数组变量名 [元素数量]Type  
- 数组变量名：数组声明及使用时的变量名
- 元素数量：数组的元素数量，可以是一个表达式，但最终通过编译期计算的结果必须是整型数值，元素数量不能含有到运行时才能确认大小的数值
- Type：可以是任意基本类型，包括数组本身，类型为数组本身时，可以实现多维数组

数组的每个元素都可以通过索引下标（数组每个元素的位置，称为索引）来访问，索引下标的范围是从 0 开始到数组长度减 1 的位置，内置函数 len() 可以返回数组中元素的个数。  
```go
var a [3]int             // 定义三个整数的数组
fmt.Println(a[0])        // 打印第一个元素
fmt.Println(a[len(a)-1]) // 打印最后一个元素
// 打印索引和元素
for i, v := range a {
    fmt.Printf("%d %d\n", i, v)
}
// 仅打印元素
for _, v := range a {
    fmt.Printf("%d\n", v)
}
```

默认情况下，数组的每个元素都会被初始化为元素类型对应的零值，对于数字类型来说就是 0，同时也可以使用数组字面值语法，用一组值来初始化数组：  
```go
var q [3]int = [3]int{1, 2, 3}
var r [3]int = [3]int{1, 2} // 1 2 0
fmt.Println(r[2]) // "0"

nums := [4]int{0: 3, 3: 4}
// 或
nums := [4]int{3, 3: 4}
// [3 0 0 4]
```

在数组的定义中，如果在数组长度的位置出现“...”省略号，则表示数组的长度是根据初始化值的个数来计算，因此，上面数组 q 的定义可以简化为：
```go
// 在编译期间就会被转换成前一种，这也就是编译器对数组大小的推导
q := [...]int{1, 2, 3}
fmt.Printf("%T\n", q) // "[3]int"

q := [3]int{1, 2, 3}
q = [4]int{1, 2, 3, 4} // 编译错误：无法将 [4]int 赋给 [3]int
```

**查看数组长度**  
使用内置的函数 len 获取数组的长度，还可以用于获取切片、map、字符串、通道的长度。  
```go
len(array)
```

**比较两个数组是否相等**  
如果两个数组类型相同（包括数组的长度，数组中元素的类型）的情况下，我们可以直接通过较运算符（==和!=）来判断两个数组是否相等，只有当两个数组的所有元素都是相等的时候数组才是相等的，不能比较两个类型不同的数组，否则程序将无法完成编译。  
```go
a := [2]int{1, 2}
b := [...]int{1, 2}
c := [2]int{1, 3}
fmt.Println(a == b, a == c, b == c) // "true false false"
d := [3]int{1, 2}
fmt.Println(a == d) // 编译错误：无法比较 [2]int == [3]int
```

**遍历数组——访问每一个数组元素**   
```go
// 迭代计数
nums := [...]int{3, 2, 1, 4}
for i := 0; i < len(nums); i++ {
    fmt.Println(nums[i])

}

// for-range
var team [3]string
team[0] = "hammer"
team[1] = "soldier"
team[2] = "mum"
for k, v := range team {
    fmt.Println(k, v)
}
```

**数组拷贝**  
在 Go 语言中，数组是值类型，也就是说在传递过程中会自动拷贝一份。
```go
nums := [...]int{3, 2, 1, 4}
numsCopy := nums
numsCopy[1] = 3
fmt.Println("nums:", nums) // nums: [3 2 1 4]
fmt.Println("numsCopy:", numsCopy) // numsCopy: [3 3 1 4]
```

**访问越界**  
无论是在栈上还是静态存储区，数组在内存中都是一连串的内存空间，如果我们不知道数组中元素的数量，访问时可能发生越界；数组访问越界是非常严重的错误，Go 语言可以在编译期间的静态类型检查阶段判断简单的数组越界：
- 访问数组的索引是非整数时，报错 “non-integer array index %v”；
- 访问数组的索引是负数时，报错 “invalid array index %v (index must be non-negative)"；
- 访问数组的索引越界时，报错 “invalid array index %v (out of bounds for %d-element array)"；

数组和字符串的一些简单越界错误都会在编译期间发现，例如：直接使用整数或者常量访问数组。但是，如果使用变量去访问数组或者字符串时，编译器就无法提前发现错误，需要 Go 语言运行时阻止不合法的访问。  
Go 语言运行时在发现数组、切片和字符串的越界操作会由运行时的 runtime.panicIndex 和 runtime.goPanicIndex 触发程序的运行时错误并导致崩溃退出。  
```go
// 编译错误
arr[4]: invalid array index 4 (out of bounds for 3-element array)
// 运行时错误
arr[i]: panic: runtime error: index out of range [4] with length 3
```

### Go 多维数组
Go 语言中允许使用多维数组，因为数组属于值类型，所以多维数组的所有维度都会在创建时自动初始化零值，多维数组尤其适合管理具有父子关系或者与坐标系相关联的数据。  

声明多维数组的语法：var array_name [size1][size2]...[sizen] array_type  
其中，array_name 为数组的名字，array_type 为数组的类型，size1、size2 等等为数组每一维度的长度。  
```go
// 二维数组是最简单的多维数组，二维数组本质上是由多个一维数组组成的
// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
var array [4][2]int
// 使用数组字面量来声明并初始化一个二维整型数组
array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
// 声明并初始化数组中索引为 1 和 3 的元素
array = [4][2]int{1: {20, 21}, 3: {40, 41}}
// 声明并初始化数组中指定的元素
array = [4][2]int{1: {0: 20}, 3: {1: 41}}

// 声明一个 2×2 的二维整型数组
var array [2][2]int
// 设置每个元素的整型值
array[0][0] = 10
array[0][1] = 20
array[1][0] = 30
array[1][1] = 40

// 声明两个二维整型数组
var array1 [2][2]int
var array2 [2][2]int
// 为array2的每个元素赋值
array2[0][0] = 10
array2[0][1] = 20
array2[1][0] = 30
array2[1][1] = 40
// 将 array2 的值复制给 array1
array1 = array2

// 将 array1 的索引为 1 的维度复制到一个同类型的新数组里
var array3 [2]int = array1[1]
// 将数组中指定的整型值复制到新的整型变量里
var value int = array1[1][0]

// 遍历
students := [4][3]int{
    {2, 2, 0},
    {2, 2, 2},
    {2, 1, 2},
    {2, 2, 2},
}
for i := 0; i < 4; i++ {
    for j := 0; j < 3; j++ {
        if students[i][j] == 0 {
            fmt.Printf("%d行%d列学生旷课", i+1, j+1)
        }
    }
}
// 1行3列学生旷课
```
