
### Go 数组
数组是一系列相同数据类型在内存中有序存储的数据集合。  
数组定义格式：var 数组名 [元素个数] 数据类型
```golang
// 定义完成数组 a 后，就在内存中开辟了 10 个连续的存储空间，每个数据都存储在相应的空间内
// 注意：数组的长度只能是常量
var a [10]int
fmt.Println(len(a))
// 定义数组时依次赋值
var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 或
arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 定义数组时部分赋值
var arr [10]int = [10]int{1, 2, 3, 4, 5}
// 定义数组时指定下标赋值
var arr [10]int = [10]int{1: 10, 4: 20, 6: 30}
// 定义时写 … 可以根据元素个数赋值
var arr [4]int = [...]int{1, 2, 3, 4}

// 数组赋值
arr[0] = 123
arr[1] = 110
arr[2] = 234
arr[5] = 567
// 打印数组
fmt.Println(arr)
// 数组读取
fmt.Println(arr[0])
fmt.Println(arr[1])
fmt.Println(arr[2])

// 遍历数组元素
arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
for i:=0;i<len(arr);i++{
    fmt.Println(arr[i])
}
// 也可以使用 range 数组名变量，按照最大的范围去遍历数组
arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
for _, v := range arr {
    fmt.Println(v)
}

// 数组也可以像变量一样，作为参数传递给函数
func modify(a [5]int) {
    a[0] = 666
    // 对传递过来的数组，惨改第—个元素的值
    fmt.Println("modify a = ", a)
}
a:=[5]int{ 1, 2,3, 4, 5}
modify(a) // 数组传递过去
fmt.Println("a 还是那个 a: a = ", a) // Go 变量作用域，所以 a 没有被改变

// 数组常见问题
arr[10] = 321 // err  数组下标越界
arr[-1] = 321 // err 数组下标越界
// 两个数组如果类型和元素个数相同可以赋值
arr2 = arr1
// 打印数组数据类型
fmt.Printf("%T\n", arr)
// 打印数组地址
fmt.Printf("数组地址：%p\n", &arr)
// 打印数组元素中地址
fmt.Printf("数组第一个元素地址：%p\n",&arr[0])
// fmt.Printf("数组第二个元素地址：%p\n",&arr[1])
// fmt.Printf("数组第三个元素地址：%p\n",&arr[2])

// 数组倒置
var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
i := 0            // 最小值下标
j := len(arr) - 1 // 最大值下标
for i < j {
    // 交换数据
    arr[i], arr[j] = arr[j], arr[i]
    // 改变下标
    i++
    j--
}
fmt.Println(arr)

// 数组冒泡排序
var arr [10]int = [10]int{9, 1, 5, 6, 8, 2, 10, 7, 4, 3}
for i := 0; i < len(arr)-1; i++ {
    for j := 0; j < len(arr)-1-i; j++ {
        if arr[j] < arr[j+1] {
            arr[j], arr[j+1] = arr[j+1], arr[j]
        }
    }
}
fmt.Println(arr)
```

### Go 二维数组
定义的数组只有一个下标，称之为一维数组，如果有两个下标，称之为二维数组。  
```golang
// 定义二维数组
b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}} // 可以理解数组 b 有 3 行 4 列构成，共能够存储 12 组数据
//部分初始化，没有初始化的值为 0
c := [3][4]int{{1, 2, 3}, {5, 6, 7, 8}, {9, 10}}
d := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
e := [3][4]int{1: {5, 6, 7, 8}} // 对第二列进行初始化，其它采用默认值

// 遍历二维数组
var a [3][4]int
k := 0
for i := 0; i < 3; i++ { // 对行进行循环
    for j:= 0;j < 4; j++ { // 对列进行循环
        k++
        a[i][j]= k
        fmt.Printf ("a[%d][%d] = %d, ", i, j, a[i][j])
    }
fmt.Printf("\n")
fmt.Println("a = ", a)
```
