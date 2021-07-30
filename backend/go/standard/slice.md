
### Go 分片
切片与数组相比，切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大，所以可以将切片理解成 “动态数组”（切片不是数组）。

[Go 切片使用技巧](https://ueokande.github.io/go-slice-tricks/)  

**数组的缺点**  
1、数组定义完，长度是固定的。  
2、使用数组作为函数参数进行传递时，如果实参为 5 个元素的整型数组，那么形参也必须 5 个元素的整型数组，否则报错。  

针对以上两个问题，可以使用切片来进行解决。  
切片如果定义了长度，长度是指已经初始化的空间（此时切片初始空间默认值都是 0）；  
切片如果定义了容量，容量是指已经开辟的空间，包括已经初始化的空间和空闲的空间。  
`注意：切片长度要小于容量`  

如果在 append 后，没有超过切片的容量大小，哪么容量不会发生变化；如果 append 后，超过了容量大小，则底层会重新分配一块 “够大” 的内存。
正如前面我们在使用 make () 函数创建切片时，如果我们能够预计出合理的容量大小（太大浪费内存空间，太小会不断的扩容），哪么我们在进行切片的 append 时，可能不会发生扩容，也就避免了切片元素的复制，减少了开销。  

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

// 先定义一个数组
months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
// 基于数组创建切片
q2 := months[3:6]    // 第二季度
summer := months[5:8]  // 夏季
fmt.Println(q2) // [April May June]
fmt.Println(summer)  // [June July August]

// 切片的容量和实际长度
var oldSlice = make([]int, 5, 10)
fmt.Println("len(oldSlice):", len(oldSlice))
fmt.Println("cap(oldSlice):", cap(oldSlice))
// 如果追加的元素个数超出 oldSlice 的默认容量，则底层会自动进行扩容
newSlice := append(oldSlice, 1, 2, 3, 4, 5, 6)
fmt.Println(newSlice)
fmt.Println(len(newSlice))
fmt.Println(cap(newSlice))

// 通过指定下标方式完成赋值
s = make([]int, 5,10)
s[0] = 1
s[1] = 2 // s[5] = 3 是不行的，必须 s = append(s, 3)
// 循环赋值
s := make ([]int, 5, 10)
// 注意：循环结束条件是小于切片的长度，而不是容量，因为切片的长度是指的是初始化的空间
for i := 0; i < len(s); i++ {
    s[i] = i // 初始化的值必须为 len 的范围内，超过则需要 append（超过容量自动扩容，可修改初始化的值）
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

var a = [5]int{1, 2, 3, 4, 5}
s1 := a[1:3]
fmt.Println(s1, len(s1), cap(s1)) // [2 3] 2 5
s2 := a[2:]
fmt.Println(s2, len(s2), cap(s2)) // [3 4 5] 3 3
s2[0] = 8 // 注意：由于 s1 和 s2 底层的数组为同一个，所以修改 s2 也会影响 s1
fmt.Println(a)  // [1,2,8,4,5]
fmt.Println(s1) // [2,8]
fmt.Println(s2) // [8,4,5]

// copy(目标，复制的切片) 切片拷贝，返回复制的个数（在 len 范围内产生覆盖）
var slice1 = []int{1, 2, 3, 4, 5, 6}
var slice2 = make([]int, 3, 5)
num := copy(slice2, slice1) // 把 slice1 的前 3 个值复制到 slice2
fmt.Println(slice1)  // [1 2 3 4 5 6]
fmt.Println(slice2)  // [1 2 3]
fmt.Println(num)  // 3

// 动态删除元素
slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
slice3 = slice3[:len(slice3) - 5]  // 删除 slice3 尾部 5 个元素
slice3 = slice3[5:]  // 删除 slice3 头部 5 个元素
```

**数据共享问题**  
切片底层是基于数组实现的，对应的结构体对象如下所示：
```golang
type slice struct {
    array unsafe.Pointer //指向存放数据的数组指针
    len   int            //长度有多大
    cap   int            //容量有多大
}
```

在结构体中使用指针存在不同实例的数据共享问题，slice2 是基于 slice1 创建的，它们的数组指针指向了同一个数组，因此，修改 slice2 元素会同步到 slice1，因为修改的是同一份内存数据，这就是数据共享问题。
```golang
slice1 := []int{1, 2, 3, 4, 5}

slice2 := slice1[1:3]
slice2[1] = 6

fmt.Println("slice1:", slice1) // slice1: [1 2 6 4 5]
fmt.Println("slice2:", slice2) // slice2: [2 6]
```

解决这个问题，使用 append 函数会重新分配新的内存，然后将结果赋值。  
```golang
slice1 := make([]int, 4)
slice2 := slice1[1:3]
slice1 = append(slice1, 0)
slice1[1] = 2
slice2[1] = 6

fmt.Println("slice1:", slice1) // slice1: [0 2 0 0 0]
fmt.Println("slice2:", slice2) // slice2: [0 6]
```

但是这里有个需要注意的地方，就是一定要重新分配内存空间，如果没有重新分配，依然存在数据共享问题。  
```golang
slice1 := make([]int, 4, 5)
slice2 := slice1[1:3]
slice1 = append(slice1, 0)
slice1[1] = 2
slice2[1] = 6

fmt.Println("slice1:", slice1) // slice1: [0 2 6 0 0]
fmt.Println("slice2:", slice2) // slice2: [2 6]
```
这里就发生了数据共享问题，因为初始化的容量是 5，比长度大，执行 append 的时候没有进行扩容，也就不存在重新分配内存操作。
