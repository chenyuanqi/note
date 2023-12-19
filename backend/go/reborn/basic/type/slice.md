
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

**Go 语言 Slice 的扩容策略**
在 Go 语言中，`slice` 是一个非常重要的数据结构，它是对底层数组的抽象，提供了一种动态和便利的方式来管理数组。但是，当我们向 `slice` 中添加元素时，可能会遇到需要扩容的情况。了解 `slice` 的扩容策略可以帮助我们更好地理解 Go 语言的内部工作原理，并编写更高效的代码。

扩容机制：当 `slice` 的容量不足以容纳更多的元素时，Go 语言会自动进行扩容。扩容策略如下：  
1. 如果当前 `slice` 的容量小于 1024 个元素，那么新的容量将是当前容量的两倍。  
2. 如果当前 `slice` 的容量大于或等于 1024 个元素，那么新的容量将增加 25%，直到容量足够容纳新的元素。  
```go
package main

import "fmt"

func main() {
    s := make([]int, 0, 1000) // 创建一个容量为 1000 的 slice
    fmt.Println(cap(s))       // 输出：1000

    s = append(s, make([]int, 500)...) // 添加 500 个元素
    fmt.Println(cap(s))                // 输出：1000

    s = append(s, make([]int, 600)...) // 添加 600 个元素，触发扩容
    fmt.Println(cap(s))                // 输出：2000
}
```
在上面的示例中，我们可以看到 `slice` 在添加元素后进行了扩容，新的容量是原来的两倍。

优化建议：  
1. 尽可能预估 `slice` 的容量需求，通过 `make` 函数预先分配足够的容量，以减少运行时的扩容操作。  
2. 避免频繁的小批量添加元素，可以通过一次性添加多个元素来减少扩容次数。  
理解 Go 语言的 `slice` 扩容策略可以帮助我们编写更高效和可维护的代码。通过合理的预分配容量和减少不必要的扩容操作，我们可以优化程序的性能。  


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

### Go 切片问题
Q、数组和切片有什么区别？  
1.数组是一个长度固定的数据类型，其长度在定义时就已经确定，不能动态改变；切片是一个长度可变的数据类型，其长度在定义时可以为空，也可以指定一个初始长度。  
2.数组的内存空间是在定义时分配的，其大小是固定的；切片的内存空间是在运行时动态分配的，其大小是可变的。  
3.当数组作为函数参数时，函数操作的是数组的一个副本，不会影响原始数组；当切片作为函数参数时，函数操作的是切片的引用，会影响原始切片。  
4.切片还有容量的概念，它指的是分配的内存空间。  

Q、切片是如何扩容的？  
当切片的长度超过其容量时，切片会自动扩容。这通常发生在使用 `append` 函数向切片中添加元素时。  
扩容时，Go 运行时会分配一个新的底层数组，并将原始切片中的元素复制到新数组中。然后，原始切片将指向新数组，并更新其长度和容量。  
需要注意的是，由于**扩容会分配新数组并复制元素，因此可能会影响性能**。如果你知道要添加多少元素，可以使用 `make` 函数预先分配足够大的切片来避免频繁扩容。  
切片扩容策略有两个阶段，go1.18 之前和之后是不同的，这一点在 go1.18 的 release notes 中有说明。  
go1.17 扩容调用的是 `growslice` 函数，我复制了其中计算新容量部分的代码。在分配内存空间之前需要先确定新的切片容量，运行时根据切片的当前容量选择不同的策略进行扩容：  
1.如果期望容量大于当前容量的两倍就会使用期望容量；  
2.如果当前切片的长度小于 1024 就会将容量翻倍；  
3.如果当前切片的长度大于等于 1024 就会每次增加 25% 的容量，直到新容量大于期望容量；  
```go
// src/runtime/slice.go

func growslice(et *_type, old slice, cap int) slice {
    // ...

    newcap := old.cap
    doublecap := newcap + newcap
    if cap > doublecap {
        newcap = cap
    } else {
        if old.cap < 1024 {
            newcap = doublecap
        } else {
            // Check 0 < newcap to detect overflow
            // and prevent an infinite loop.
            for 0 < newcap && newcap < cap {
                newcap += newcap / 4
            }
            // Set newcap to the requested cap when
            // the newcap calculation overflowed.
            if newcap <= 0 {
                newcap = cap
            }
        }
    }

    // ...

    return slice{p, old.len, newcap}
}
```
go1.18 和之前版本的区别，主要在扩容阈值，以及这行代码：`newcap += (newcap + 3*threshold) / 4`。在分配内存空间之前需要先确定新的切片容量，运行时根据切片的当前容量选择不同的策略进行扩容：  
1.如果期望容量大于当前容量的两倍就会使用期望容量；  
2.如果当前切片的长度小于阈值（默认 256）就会将容量翻倍；  
3.如果当前切片的长度大于等于阈值（默认 256），就会每次增加 25% 的容量，基准是 `newcap + 3*threshold`，直到新容量大于期望容量；  
```go
// src/runtime/slice.go

func growslice(et *_type, old slice, cap int) slice {
    // ...

    newcap := old.cap
    doublecap := newcap + newcap
    if cap > doublecap {
        newcap = cap
    } else {
        const threshold = 256
        if old.cap < threshold {
            newcap = doublecap
        } else {
            // Check 0 < newcap to detect overflow
            // and prevent an infinite loop.
            for 0 < newcap && newcap < cap {
                // Transition from growing 2x for small slices
                // to growing 1.25x for large slices. This formula
                // gives a smooth-ish transition between the two.
                newcap += (newcap + 3*threshold) / 4
            }
            // Set newcap to the requested cap when
            // the newcap calculation overflowed.
            if newcap <= 0 {
                newcap = cap
            }
        }
    }

    // ...

    return slice{p, old.len, newcap}
}
```
扩容之后的容量其实并不是严格按照这个策略的。为什么呢？  
实际上，`growslice` 的后半部分还有更进一步的优化（内存对齐等），靠的是 `roundupsize` 函数，在计算完 `newcap` 值之后，还会有一个步骤计算最终的容量：
```go
capmem = roundupsize(uintptr(newcap) * ptrSize)
newcap = int(capmem / ptrSize)
```
