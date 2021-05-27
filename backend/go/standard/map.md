
### Go 映射
Go 语言中的映射，就好比其他语言中的 hash 表或者字典。它们的工作方式就是：定义键和值，并且可以获取，设置和删除其中的值。  

```golang
// 映射和切片一样，使用 make 方法来创建
lookup := make(map[string]int)
// 传递第二个参数到 make 方法来设置一个初始大小
// 如果你事先知道映射会有多少键值，定义一个初始大小将会帮助改善性能
// lookup := make(map[string]int, 100)
lookup["goku"] = 9001
power, exists := lookup["vegeta"]
// prints 0, false
// 0 is the default value for an integer
fmt.Println(power, exists)

// 使用 len 方法类获取映射的键的数量
// returns 1
total := len(lookup)
// 使用 delete 方法来删除一个键对应的值
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

```

