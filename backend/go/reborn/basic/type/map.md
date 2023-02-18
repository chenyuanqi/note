
### Go map
map 是一种键 (key)/ 值 (value) 对的无序集合，在其它语言中称为字典、关联数组、哈希表等。当给定了键可以快速定位到值，而且键必须唯一的，不能出现相同。  

**声明**  
声明 map 时，键不是所有类型都支持，它只支持可以使用 != 或 == 操作符比较的类型。哪些类型不能进行比较？  
- 函数
- map
- 切片
- 元素是函数、map、切片的数组
- 字段中包含函数、map、切片的结构体

声明格式： var 变量名 map[键类型\][值类型\]
```go
// 未初始化，dic 为 nil
var dic map[int]string
```

**初始化**  
初始化 map 有两种方式，第一种使用 make 函数，第二种是声明 map 时，初始化具体的键和值。如果 map 未初始化是不能存取值的，不然编译器报错。  
`注意：可以使用 make()，但不能使用 new() 来构造 map，如果错误的使用 new() 分配了一个引用对象，会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址。`  

```go
// make 初始化
dic := make(map[int]string)

// 初始化后的 map 会根据新增的键值动态伸缩
dic := make(map[int]string)
dic[1] = "lao"
dic[3] = "chen"
fmt.Println("dic长度:", len(dic))

// 在初始化时，可以提前定义好 map 所需要的容量（空间大小），当添加的键值超过容量时自动加一
dic := make(map[int]string, 10)
// 容量为 10 ，存了 1 个
dic[1] = "lao"
fmt.Println("dic长度:", len(dic)) // dic长度: 1

// map 声明时初始化
m := map[string]int{
    "a": 2,
    "b": 3,
    "c": 4,
}
fmt.Println("b:", m["b"]) // b: 3

// 初始化时也可以不指定键和值，这种情况和不指定容量的 make 函数是相同的
m := map[string]int{}
// 等价于 m := make(map[string]int)

// nil 的 map 不可赋值
var m map[string]float64
m["pi"] = 3.1416 // panic: assignment to entry in nil map
// 必须使用 make 初始化
m := make(map[string]float64)
m["pi"] = 3.1416
```
当 map 增长到容量上限的时候，如果再增加新的 key-value，map 的大小会自动加 1，所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。

**键是否存在**  
从初始化的 map 中获取一个没有存储的键时，编译器是会报错的。那么，怎么判断键是否存在呢？格式：v，ok := map[key\]  
```go
dic := map[int]string{}
dic[0] = "a"
if v, ok := dic[0]; ok {
    fmt.Println(v) // a

}
```

**删除键值对**  
使用 delete 函数可以删除 map 中的键值对，格式：delete(map, 键)  
`注：如果键不存在时，编译器也可以通过。`  
```go
m := map[string]int{
    "a": 2,
    "b": 3,
    "c": 4,
}
delete(m, "b")
fmt.Println(m) // map[a:2 c:4]
```

**清空 map 中的所有元素**  
有意思的是，Go 语言中并没有为 map 提供任何清空所有元素的函数、方法，清空 map 的唯一办法就是重新 make 一个新的 map，不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。
```go
m := map[string]int{
    "a": 2,
    "b": 3,
    "c": 4,
}
// 重新赋值
m = map[string]int{}
```

**遍历**  
遍历 map 需要用到 for-range 语法。  
```go
m := map[string]int{
    "a": 2,
    "b": 3,
    "c": 4,
}
for k, v := range m {
    fmt.Println("key:", k, ",value:", v)
}

// 只遍历键
for k := range m {
    // ...
}

// 只遍历值
for _, v := range m {
    // ...
}
```

**引用类型**  
map 是引用类型，因此在传递过程中它只存在一份。  
拷贝 map，没有类似 copy 的函数，需要新创建一个 map，手动遍历赋值。  
```go
m := map[string]int{
    "a": 2,
    "b": 3,
}

// 遍历拷贝
mCopy := map[string]int{}
for k, v := range m {
    mCopy[k] = v
}
```

**排序**  
`Map` 的遍历是无序的，这意味着不能依赖遍历的键值顺序。如果想实现 Map 遍历时顺序永远一致，
一个折中的方案时预先给 Map 的 `键` 排序，然后根据排序后的键序列遍历 Map, 这样可以保证每次遍历顺序都是一样的。  
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var m = make(map[int]string)

	m[0] = "zero"
	m[1] = "one"
	m[2] = "two"

	keys := make([]int, len(m)) // 将所有的键放入一个切片中
	index := 0
	for k, _ := range m {
		keys[index] = k
		index++
	}

	sort.Ints(keys) // 将所有的键进行排序

	for i := 0; i < 5; i++ {
		for _, key := range keys { // 根据排序后的键遍历 Map
			fmt.Printf("key = %d, val = %s\n", key, m[key])
		}
		fmt.Printf("第 %d 次遍历完成\n", i+1)
	}
}

// $ go run main.go
// 输出如下 
/**
  key = 0, val = zero
  key = 1, val = one
  key = 2, val = two
  第 1 次遍历完成
  key = 0, val = zero
  key = 1, val = one
  key = 2, val = two
  第 2 次遍历完成
  key = 0, val = zero
  key = 1, val = one
  key = 2, val = two
  第 3 次遍历完成
  key = 0, val = zero
  key = 1, val = one
  key = 2, val = two
  第 4 次遍历完成
  key = 0, val = zero
  key = 1, val = one
  key = 2, val = two
  第 5 次遍历完成
*/
```
