
### Go 内置排序算法
排序操作和字符串格式化一样是很多程序经常使用的操作。尽管一个最短的快排程序只要 15 行就可以搞定，但是一个健壮的实现需要更多的代码，并且我们不希望每次我们需要的时候都重写或者拷贝这些代码。

幸运的是，sort 包内置的提供了根据一些排序函数来对任何序列排序的功能。它的设计非常独到。在很多语言中，排序算法都是和序列数据类型关联，同时排序函数和具体类型元素关联。

相比之下，Go 语言的 sort.Sort 函数不会对具体的序列和它的元素做任何假设。相反，它使用了一个接口类型 sort.Interface 来指定通用的排序算法和可能被排序到的序列类型之间的约定。这个接口的实现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片。
```go
package sort
type Interface interface {
    Len() int            // 获取元素数量
    Less(i, j int) bool // i，j是序列元素的指数。
    Swap(i, j int)        // 交换元素
}

// 为了对序列进行排序，我们需要定义一个实现了这三个方法的类型，然后对这个类型的一个实例应用 sort.Sort 函数
// 对一个字符串切片进行排序
type MyStringList  []string
func (p MyStringList ) Len() int { return len(m) }
func (p MyStringList ) Less(i, j int) bool { return m[i] < m[j] }
func (p MyStringList ) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
```

使用 sort.Interface 接口进行排序  
对一系列字符串进行排序时，使用字符串切片（[] string）承载多个字符串。使用 type 关键字，将字符串切片（[] string）定义为自定义类型 MyStringList。为了让 sort 包能识别 MyStringList，能够对 MyStringList 进行排序，就必须让 MyStringList 实现 sort.Interface 接口。
```go
package main
import (
    "fmt"
    "sort"
)
// 将[]string定义为MyStringList类型
type MyStringList []string
// 实现sort.Interface接口的获取元素数量方法
func (m MyStringList) Len() int {
    return len(m)
}
// 实现sort.Interface接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
    return m[i] < m[j]
}
// 实现sort.Interface接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}
func main() {
    // 准备一个内容被打乱顺序的字符串切片
    names := MyStringList{
        "3. Triple Kill",
        "5. Penta Kill",
        "2. Double Kill",
        "4. Quadra Kill",
        "1. First Blood",
    }
    // 使用sort包进行排序
    sort.Sort(names)
    // 遍历打印结果
    for _, v := range names {
        fmt.Printf("%s\n", v)
    }
}
// 1. First Blood
// 2. Double Kill
// 3. Triple Kill
// 4. Quadra Kill
// 5. Penta Kill
```

**常见类型的便捷排序**  
通过实现 sort.Interface 接口的排序过程具有很强的可定制性，可以根据被排序对象比较复杂的特性进行定制。例如，需要多种排序逻辑的需求就适合使用 sort.Interface 接口进行排序。但大部分情况中，只需要对字符串、整型等进行快速排序。Go 语言中提供了一些固定模式的封装以方便开发者迅速对内容进行排序。  
1) 字符串切片的便捷排序  
sort 包中有一个 StringSlice 类型，sort 包中的 StringSlice 的代码与 MyStringList 的实现代码几乎一样。因此，只需要使用 sort 包的 StringSlice 就可以更简单快速地进行字符串排序。
```go
type StringSlice []string
func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// Sort is a convenience method.
func (p StringSlice) Sort() { Sort(p) }

// 执行排序
names := sort.StringSlice{
    "3. Triple Kill",
    "5. Penta Kill",
    "2. Double Kill",
    "4. Quadra Kill",
    "1. First Blood",
}
sort.Sort(names)
```
2) 对整型切片进行排序  
除了字符串可以使用 sort 包进行便捷排序外，还可以使用 sort.IntSlice 进行整型切片的排序。
```go
type IntSlice []int
func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// Sort is a convenience method.
func (p IntSlice) Sort() { Sort(p) }

// 执行排序
names := []string{
    "3. Triple Kill",
    "5. Penta Kill",
    "2. Double Kill",
    "4. Quadra Kill",
    "1. First Blood",
}
sort.Strings(names)
// 遍历打印结果
for _, v := range names {
    fmt.Printf("%s\n", v)
}
```
3) sort 包内建的类型排序接口一览  
Go 语言中的 sort 包中定义了一些常见类型的排序方法。编程中经常用到的 int32、int64、float32、bool 类型并没有由 sort 包实现，使用时依然需要开发者自己编写。  
[![常见类型的排序方法](https://s1.ax1x.com/2022/04/22/LRSB1s.md.png)](https://imgtu.com/i/LRSB1s)

**对结构体数据进行排序**  
结构体比基本类型更为复杂，排序时不能像数值和字符串一样拥有一些固定的单一原则。结构体的多个字段在排序中可能会存在多种排序的规则，例如，结构体中的名字按字母升序排列，数值按从小到大的顺序排序。一般在多种规则同时存在时，需要确定规则的优先度，如先按名字排序，再按年龄排序等。

1) 完整实现 sort.Interface 进行结构体排序  
将一批英雄名单使用结构体定义，英雄名单的结构体中定义了英雄的名字和分类。排序时要求按照英雄的分类进行排序，相同分类的情况下按名字进行排序。    
```go
package main
import (
    "fmt"
    "sort"
)
// 声明英雄的分类
type HeroKind int
// 定义HeroKind常量, 类似于枚举
const (
    None HeroKind = iota
    Tank
    Assassin
    Mage
)
// 定义英雄名单的结构
type Hero struct {
    Name string  // 英雄的名字
    Kind HeroKind  // 英雄的种类
}
// 将英雄指针的切片定义为Heros类型
type Heros []*Hero
// 实现sort.Interface接口取元素数量方法
func (s Heros) Len() int {
    return len(s)
}
// 实现sort.Interface接口比较元素方法
func (s Heros) Less(i, j int) bool {
    // 如果英雄的分类不一致时, 优先对分类进行排序
    if s[i].Kind != s[j].Kind {
        return s[i].Kind < s[j].Kind
    }
    // 默认按英雄名字字符升序排列
    return s[i].Name < s[j].Name
}
// 实现sort.Interface接口交换元素方法
func (s Heros) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func main() {
    // 准备英雄列表
    heros := Heros{
        &Hero{"吕布", Tank},
        &Hero{"李白", Assassin},
        &Hero{"妲己", Mage},
        &Hero{"貂蝉", Assassin},
        &Hero{"关羽", Tank},
        &Hero{"诸葛亮", Mage},
    }
    // 使用sort包进行排序
    sort.Sort(heros)
    // 遍历英雄列表打印排序结果
    for _, v := range heros {
        fmt.Printf("%+v\n", v)
    }
}
// &{Name: 关羽 Kind:1}
// &{Name: 吕布 Kind:1}
// &{Name: 李白 Kind:2}
// &{Name: 貂蝉 Kind:2}
// &{Name: 妲己 Kind:3}
// &{Name: 诸葛亮 Kind:3}
```

2) 使用 sort.Slice 进行切片元素排序  
从 Go 1.8 开始，Go 语言在 sort 包中提供了 sort.Slice () 函数进行更为简便的排序方法。sort.Slice () 函数只要求传入需要排序的数据，以及一个排序时对元素的回调函数，类型为 func (i,j int) bool。  
使用 sort.Slice () 不仅可以完成结构体切片排序，还可以对各种切片类型进行自定义排序。  
```go
// sort.Slice () 函数的定义
func Slice(slice interface{}, less func(i, j int) bool)

package main
import (
    "fmt"
    "sort"
)
type HeroKind int
const (
    None = iota
    Tank
    Assassin
    Mage
)
type Hero struct {
    Name string
    Kind HeroKind
}
func main() {
    heros := []*Hero{
        {"吕布", Tank},
        {"李白", Assassin},
        {"妲己", Mage},
        {"貂蝉", Assassin},
        {"关羽", Tank},
        {"诸葛亮", Mage},
    }
    sort.Slice(heros, func(i, j int) bool {
        if heros[i].Kind != heros[j].Kind {
            return heros[i].Kind < heros[j].Kind
        }
        return heros[i].Name < heros[j].Name
    })
    for _, v := range heros {
        fmt.Printf("%+v\n", v)
    }
}
```

**冒泡排序**  
冒泡排序只会操作相邻的两个数据。每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求，如果不满足就让它俩互换。一次冒泡会让至少一个元素移动到它应该在的位置，重复 n 次，就完成了 n 个数据的排序工作。

```go
package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 6, 4}
	fmt.Println(BubbleSort(arr))
}

func BubbleSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arr
	}

	var flag bool
	for i := 0; i < arrLen; i++ {
		flag = false
		for j := 0; j < arrLen-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}

	return arr
}
```


