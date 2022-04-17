
### for-range
与经典的三段式 for 循环相比，范围循环在 Go 语言中更常见、实现也更复杂。这种循环同时使用 for 和 range 两个关键字，不过编译器会在编译期间将所有 for-range 循环变成经典循环。
```go
package main

import "fmt"

func main() {
    nums := []int{2, 3, 4}

    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    
    for k := range kvs {
        fmt.Println("key:", k)
    }
    
    for i, c := range "go" {
        fmt.Println(i, c)
    }

    numChannel := make(chan int)
    for  x := range numChannel {
            // 用 for range 读取通道中的值 （示例这会阻塞）
    }
}
```
range 可以对多种数据结构进行迭代。  
对序列类型 slice 或者 array 进行迭代时会提供元素的索引和值，如果程序中不需要序列元素的索引或者值可以用空标识符 _ 忽略。

range 使用在 map 上时会迭代其包含的键值对，也可以只迭代 map 中的键（range 赋值语句前只有一个接收变量，那么接收到的是 map 的键）。

range 使用在字符串上时会把其当做 Unicode 码点集合来进行迭代。

range 使用在通道上时，会在通道关闭后自动结束循环。
