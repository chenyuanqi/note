
### Go 分片
切片与数组相比，切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大，所以可以将切片理解成 “动态数组”（切片不是数组）。

**数组的缺点**  
1、数组定义完，长度是固定的。  
2、使用数组作为函数参数进行传递时，如果实参为 5 个元素的整型数组，那么形参也必须 5 个元素的整型数组，否则报错。  

针对以上两个问题，可以使用切片来进行解决。  
切片如果定义了长度，长度是指已经初始化的空间（此时切片初始空间默认值都是 0）；  
切片如果定义了容量，容量是指已经开辟的空间，包括已经初始化的空间和空闲的空间。  
`注意：切片长度要小于容量`  

```golang
// 切片初始化：声明切片和声明数组一样，只是少了长度
s := []int{} // 空 (nil) 切片
// 或
var s1 []int // 空 (nil) 切片
// 或使用 make 函数定义切片（格式 make(切片类型，长度，{容量}), 容量缺省时=长度）
s = make([]int, 5, 8)  // s := make([]int, 5, 8)
fmt.Println(s) // [0 0 0 0 0]
fmt.Println("长度是:", len(s))
fmt.Println("容量是:", cap(s))

//初始化切片时添加部分数据
s := []int{1,2,3}
// 通过 append 函数向切片中追加数据中
s = append(s, 5, 6,7)
fmt.Println(s) // [1 2 3 5 6 7]

// 通过指定下标方式完成赋值
s = make([]int, 5,10)
s[0] = 1
s[1] = 2 // s[5] = 3 是不行的，必须 s = append(s, 3)
// 循环赋值
s := make ([]int, 5, 10)
// 注意：循环结束条件是小于切片的长度，而不是容量，因为切片的长度是指的是初始化的空间
for i := 0; i < len(s); i++ {
    s[i] = i
}
// 打印切片数据 - 通过下标
for i := 0; i < len(s); i++ {
    fmt.Println(s[i])
}
// 打印切片数据 - 通过 range
for k, v := range s {
    fmt.Println("下标:", k)
    fmt.Println( "值:", v)
}

// 切片
s = make([]int, 5,10)
s[0] = 1
s[1] = 2 // s[5] = 3 是不行的，必须 append(s, 3)
scores = s[0:10]
s[5] = 3
// 切片中除了最后一个元素的所有值
scores := []int{1, 2, 3, 4, 5}
scores = scores[:len(scores)-1]
```

