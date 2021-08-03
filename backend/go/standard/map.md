
### Go 映射（字典）
Go 语言中的映射，就好比其他语言中的 hash 表或者字典。它们的工作方式就是：定义键和值，并且可以获取，设置和删除其中的值。  
对于强类型的 Go 语言来说，需要在声明时指定键和值的类型，此外，Go 映射也是个无序集合，底层不会按照元素添加顺序维护元素的存储顺序。

Go 语言中的映射和 Redis 一样，底层也是通过哈希表实现的，添加键值对到字典时，实际是将键转化为哈希值进行存储，在查找时，也是先将键转化为哈希值去哈希表中查询，从而提高性能。  
但是哈希表存在哈希冲突问题，即不同的键可能会计算出同样的哈希值，这个时候 Go 底层还会判断原始键的值是否相等。也正因如此，我们在声明字典的键类型时，要求数据类型必须是支持通过 == 或 != 进行判等操作的类型，比如数字类型、字符串类型、数组类型、结构体类型等，不过为了提高字典查询性能，类型长度越短约好，通常，我们会将其设置为整型或者长度较短的字符串类型。  

```golang
// 声明
// 字典初始化之后才能进行赋值操作，如果仅仅是声明，此时 testMap 的值为 nil，在 nil 上进行操作编译期间会报 panic（运行时恐慌），导致编译不通过
var testMap map[string]int
// 初始化
testMap := map[string]int{
  "one": 1,
  "two": 2,
  "three": 3,
}
fmt.Println(testMap) // map[one:1 three:3 two:2]

// 映射和切片一样，使用 make 方法来创建
lookup := make(map[string]int)
// 传递第二个参数到 make 方法来设置一个初始大小
// 如果你事先知道映射会有多少键值，定义一个初始大小将会帮助改善性能
// lookup := make(map[string]int, 100)
lookup["goku"] = 9001 // var lookup map[string]int 不能这么操作，否则编译期间会抛出 panic
power, exists := lookup["vegeta"]
// prints 0, false
// 0 is the default value for an integer
fmt.Println(power, exists)

// 从字典中查找一个特定的键对应的值
value, ok := testMap["one"] 
// 从字典中查找指定键时，会返回两个值，第一个是真正返回的键值，第二个是是否找到的标识，判断是否在字典中成功找到指定的键，不需要检查取到的值是否为 nil，只需查看第二个返回值 ok，这是一个布尔值，如果查找成功，返回 true，否则返回 false
if ok { // 找到了
  // 处理找到的value 
}

// 使用 len 方法类获取映射的键的数量
// returns 1
total := len(lookup)
// 使用 delete 方法来删除一个键对应的值（键不存在/未初始化，都不会报错）
// has no return, can be called on a non-existing key
delete(lookup, "goku")

// 将映射作为结构体字段的时候，你可以这样定义
type Saiyan struct {
  Name string
  Friends map[string]*Saiyan
}
// 或
goku := &Saiyan{
  Name: "Goku",
  Friends: make(map[string]*Saiyan),
}
goku.Friends["krillin"] = ... //加载或者创建 Krillin
// 或定义为复合方式
lookup := map[string]int{
  "goku": 9001,
  "gohan": 2044,
}

// 使用 for 组合 range 关键字迭代映射
// 注意：迭代映射是没有顺序的，每次迭代查找将会随机返回键值对
for key, value := range lookup {
  ...
}
// 可以借助匿名变量只获取字典的值
for _, value := range lookup {
    fmt.Println(value)
}

// 键值对调，即交换字典的键和值
testMap := map[string]int{
    "one": 1,
    "two": 2,
    "three": 3,
}
invMap := make(map[int] string, 3)
for k, v := range testMap {
    invMap[v] = k
}
for k, v := range invMap {
    fmt.Println(k, v)
}
/*
3 three
1 one
2 two
*/

// 排序
// Go 语言的字典是一个无序集合，如果你想要对字典进行排序，可以通过分别为字典的键和值创建切片，然后通过对切片进行排序来实现
// 按照键进行排序
keys := make([]string, 0)
for k, _ := range testMap {
    keys = append(keys, k)
}
sort.Strings(keys)  // 对键进行排序，按照键名在字母表中的排序进行升序排序的结果
fmt.Println("Sorted map by key:")
for _, k := range keys {
    fmt.Println(k, testMap[k])
}
// 按照值进行排序
values := make([]int, 0)
for _, v := range testMap {
    values = append(values, v)
}
sort.Ints(values)   // 对值进行排序，按照键值对应数字大小进行升序排序的结果
fmt.Println("Sorted map by value:")
for _, v := range values  {
    fmt.Println(invMap[v], v) // 借助了之前创建的 invMap 通过字典的值反查对应的键
}
```

