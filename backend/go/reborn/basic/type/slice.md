
### Go 切片
切片使用起来类似长度可变的数组，不像数组长度是固定的。但切片的底层使用的还是数组，切片只是保存了对数组的引用，帮着管理数组，实现可变的效果。  
切片格式： var 切片名称 [\]数据类型  
和数组声明的区别是，是否指明了长度，没有长度则为切片。  

```go
// 切片未初始化默认为 nil ，长度为 0 
var nums []int
// 清空切片可以赋值 nil，可以这样 nums = nil
// 重置切片，清空拥有的元素，也可以把切片的开始和结束位置都设为 0
a := []int{1, 2, 3}
fmt.Println(a[0:0])

// 初始化具体值，长度为 3 的切片，此时容量也为 3
nums := []int{1, 2, 3}
```

**make 函数**  
使用 make 函数初始化切片，容量参数可以省略，省略后长度和容量相等。  
```go
// 切片名称 := make([]数据类型，长度，容量)
nums := make([]int, 2, 5)

// 三种初始化切片的方式
arr[0:3] or slice[0:3]
slice := []int{1, 2, 3}
slice := make([]int, 10)
```

**操作具体元素**  
切片中元素的具体操作和数组的方式是一样的。如果获取元素时超出切片长度，即使没有超出容量，编译器也会报错。
```go
nums := []int{1, 2, 3}
// 设置索引 1 的元素为 4
nums[1] = 4
fmt.Println(nums[1]) // 4

// 读取最后一个元素
a := []string{"A", "B", "C"}
s := a[len(a)-1] // C
// 移除最后一个元素
a = a[:len(a)-1] // [A B]

data := []int{1, 2, 3}
// 修改原有元素的值，应该使用索引直接访问
for i, v := range data {
    data[i] = v * 10    
}
fmt.Println("data: ", data)    // data:  [10 20 30]
```

**迭代切片**  
迭代返回的变量是一个在迭代过程中根据切片依次赋值的新变量，所以 value 的地址总是相同的，要想获取每个元素的地址，需要使用切片变量和索引值。
```go
// 传统的 for 循环对切片进行迭代
// 创建一个整型切片，并赋值
slice := []int{10, 20, 30, 40}
// 从第三个元素开始迭代每个元素
for index := 2; index < len(slice); index++ {
    fmt.Printf("Index: %d Value: %d\n", index, slice[index])
}

// 创建一个整型切片，并赋值
slice := []int{10, 20, 30, 40}
// 迭代每个元素，并显示值和地址
for index, value := range slice {
    fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
}
// Value: 10 Value-Addr: 10500168 ElemAddr: 1052E100
// Value: 20 Value-Addr: 10500168 ElemAddr: 1052E104
// Value: 30 Value-Addr: 10500168 ElemAddr: 1052E108
// Value: 40 Value-Addr: 10500168 ElemAddr: 1052E10C
```

**获取子集**  
定义了一个切片或数组后，可以获取其中的一部分，即子集。  
格式： 切片或数组[开始索引:结束索引]  
获取从 “开始索引” 到 “结束索引” 的子集，`包含开始索引，但不包含结束索引`。如果是数组获取子集后，类型会转化为切片类型。
```go
// 切片
nums := []int{1, 2, 3, 4, 5}
// 获取切片子集
nums1 := nums[2:4]   // []int{3, 4}
// 数组
arr := [5]int{1, 2, 3, 4, 5}
// nums2 为切片类型
nums2 := arr[2:4]    // []int{3, 4}
```
“开始索引” 和 “结束索引” 都可以省略。
- 开始索引省略，表示子集从索引 0 开始到结束索引。
- 结束索引省略，表示子集从开始索引到最后结束。
- 都省略，如果是切片两者一样，如果是数组会转化为切片类型。

`注意：切片是一个引用类型，它在传递时不会进行拷贝。`


**追加和移除元素**  
往切片中追加元素，使用到 append 函数，此函数只能追加到切片末尾。  
如果想追加到切片开头，没有原生的函数，使用 append 变向的实现，这种方式其实就是合并两个切片。  
如何移除某个元素呢，使用切片子集和 append 函数变向实现。Go 语言中删除切片元素的本质是，以被删除元素为分界点，将前后两个部分的内存重新连接起来。  
```go
nums := []int{1, 2, 3}
nums = append(nums, 2)
nums = append(nums, 4, 5)
fmt.Println(nums) // [1 2 3 2 4 5]

nums := []int{1, 2, 3}
// ... 三个点表示将切片元素展开传递给函数（解包）
nums = append([]int{4}, nums...)
fmt.Println(nums) // [4 1 2 3]

nums := []int{1, 2, 3, 4, 5}
// 移除索引为 2 的元素值 3
nums = append(nums[:2], nums[3:]...)
fmt.Println(nums) // [1 2 4 5]

// 从开头位置删除-直接移动数据指针
var a = []int{1, 2, 3}
a = a[1:] // 删除开头1个元素
a = a[N:] // 删除开头N个元素
// 从开头位置删除-不移动数据指针
var a = []int{1, 2, 3}
a = append(a[:0], a[1:]...) // 删除开头1个元素
a = append(a[:0], a[N:]...) // 删除开头N个元素
// 用 copy() 函数来删除开头的元素，copy 返回长度/容量最小值
var a = []int{1, 2, 3}
a = a[:copy(a, a[1:])] // 删除开头1个元素
a = a[:copy(a, a[N:])] // 删除开头N个元素

// 从中间位置删除（对于删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用 append 或 copy 原地完成）
a = []int{1, 2, 3, ...}
a = append(a[:i], a[i+1:]...) // 删除中间1个元素
a = append(a[:i], a[i+N:]...) // 删除中间N个元素
a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素

// 从尾部删除
a = []int{1, 2, 3}
a = a[:len(a)-1] // 删除尾部1个元素
a = a[:len(a)-N] // 删除尾部N个元素

// 删除开头的元素和删除尾部的元素都可以认为是删除中间元素操作的特殊情况
seq := []string{"a", "b", "c", "d", "e"}
// 指定删除位置
index := 2
// 查看删除位置之前的元素和之后的元素
fmt.Println(seq[:index], seq[index+1:]) // [a b] [d e]
// 将删除点前后的元素连接起来
seq = append(seq[:index], seq[index+1:]...)
fmt.Println(seq) // [a b d e]
```

**切片清空**  
```go
// 移除所有元素
a := []string{"A", "B", "C", "D", "E"}
a = nil
fmt.Println(a, len(a), cap(a)) // [] 0 0

// 保留存储配置
a := []string{"A", "B", "C", "D", "E"}
a = a[:0]
fmt.Println(a, len(a), cap(a)) // [] 0 5

// 空切片和 nil
var a []int = nil
var a0 []int = make([]int, 0)
fmt.Println(a == nil)  // true
fmt.Println(a0 == nil) // false
fmt.Printf("%#v\n", a)  // []int(nil)
fmt.Printf("%#v\n", a0) // []int{}
```

**切片查找**  
```go
// 线性查找
// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) int {
    for i, n := range a {
        if x == n {
            return i
        }
    }
    return len(a)
}
// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
    for _, n := range a {
        if x == n {
            return true
        }
    }
    return false
}

// 二分查找，切片必须升序排列
// SearchType(a []Type, x Type) int 返回 x <= a[i] 最小索引值，不存在则返回插入排序后的索引有可能返回 len(a)
a := []string{"A", "C", "C"}
fmt.Println(sort.SearchStrings(a, "A")) // 0
fmt.Println(sort.SearchStrings(a, "B")) // 1，为啥返回 1？不存在则插入 B 按升序排列的位置是 1
fmt.Println(sort.SearchStrings(a, "C")) // 1
fmt.Println(sort.SearchStrings(a, "D")) // 3
// Search(n int, f func(int) bool) int 返回最小索引值，不存在返回 n
x := "C"
i := sort.Search(len(a), func(i int) bool { return x <= a[i] })
if i < len(a) && a[i] == x {
    fmt.Printf("Found %s at index %d in %v.\n", x, i, a)
} else {
    fmt.Printf("Did not find %s in %v.\n", x, a)
}
// Found C at index 1 in [A C C].
```

**切片拷贝**  
切片是一个引用类型，因此不能像数组一样直接赋值给一个新变量就会产生拷贝，可以使用 copy 函数完成切片的拷贝。
```go
// 将 nums 拷贝到 numsCopy：copy(destSlice, srcSlice []T) int
nums := []int{1, 2, 3}
numsCopy := make([]int, 3)
copy(numsCopy, nums)
// 修改了 numsCopy，不会对 nums 产生影响
numsCopy[0] = 2
fmt.Println("nums:", nums) // nums: [1 2 3]
fmt.Println("numsCopy:", numsCopy) // numsCopy: [2 2 3]

// numsCopy 长度可以小于或大于 nums 的长度，如果小于就会拷贝 nums 前面一部分，大于会保留 numsCopy 后面一部分
// numsCopy 长度小于 nums
nums := []int{1, 2, 3}
numsCopy := make([]int, 2)
// 前面两个元素 1 和 2 被复制
copy(numsCopy, nums)
fmt.Println("numsCopy(小于):", numsCopy) // numsCopy(小于): [1 2]
// numsCopy 长度大于 nums
numsCopy = []int{4, 5, 6, 7}
// 4,5,6 会被覆盖，保留 7
copy(numsCopy, nums)
fmt.Println("numsCopy(大于):", numsCopy) // numsCopy(大于): [1 2 3 7]
```

“长度> 容量” 会触发拷贝。  
使用 append 函数给切片追加元素时，如果追加的长度大于切片的容量时，切片的底层数组空间则重新开辟一块比原来大的地方，并把原来的数组值拷贝一份。切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充，例如 1、2、4、8、16...  

**排序**  
```go
sceneList := []int{1,3,2}
// 对切片进行排序
sort.Ints(sceneList)
```

### Go 多维切片
Go 多维切片和多维数组是类似的，唯一的不同点是切片没有指明长度。
```go
// 声明二维切片
var mult [][]int
// 初始化二维切片
students := [][]int{
    {2, 2, 0},
    {2, 2, 2},
    {2, 1, 2},
    {2, 2, 2},
}
```

### Go 切片字符串
Go 字符串可以使用上面的子集用法，来获取字符串中的一部分。
```go
str := "I'm laomiao."
fmt.Println(str[4:7]) // lao
```
