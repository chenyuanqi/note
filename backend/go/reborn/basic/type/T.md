
## T 泛型
一、泛型简介

泛型是一种在编程中实现代码复用的技术，它允许你编写一种通用的数据结构或函数，而不需要为每种可能的数据类型都编写重复的代码。在 Go 1.18 版本中，Golang 引入了泛型特性，从而为 Go 语言提供了更强大的类型抽象能力。

二、泛型的应用

1. 泛型类型

泛型类型是一种特殊的数据类型，它可以在运行时根据需要实例化为不同的具体类型。在 Golang 中，你可以使用 `type` 关键字来定义一个泛型类型，如下所示：

```go
type MyGeneric[T any] struct {
    data T
}
```

在这个例子中，`MyGeneric` 是一个泛型类型，`T` 是一个类型参数，`any` 是一个类型约束，表示 `T` 可以是任何类型。

2. 泛型函数

泛型函数是一种可以处理不同类型数据的函数。在 Golang 中，你可以使用 `func` 关键字来定义一个泛型函数，如下所示：

```go
// [T any] 是泛型参数列表
// T 是一个类型参数，它是一个占位符，表示在实际使用时，将由具体的类型替换。
// any 是一个类型约束，它表示 T 可以是任意类型
func Swap[T any](a, b *T) {
    *a, *b = *b, *a
}
```

在这个例子中，`Swap` 是一个泛型函数，`T` 是一个类型参数，`any` 是一个类型约束，表示 `T` 可以是任何类型。该函数接受两个指向 `T` 类型的指针，并交换它们的值。

三、泛型示例代码

1. 泛型切片反转

```go
package main

import "fmt"

// 由于 T 可以是任意类型，因此这个函数可以接受和返回任意类型的切片
func ReverseSlice[T any](s []T) []T {
    result := make([]T, len(s))
    for i := range s {
        result[len(s)-1-i] = s[i]
    }
    return result
}

func main() {
    intSlice := []int{1, 2, 3, 4, 5}
    fmt.Println(ReverseSlice(intSlice))

    stringSlice := []string{"a", "b", "c", "d", "e"}
    fmt.Println(ReverseSlice(stringSlice))
}
```

2. 泛型 Map 函数

```go
package main

import "fmt"

func Map[T, U any](data []T, mapper func(T) U) []U {
    result := make([]U, len(data))
    for i, v := range data {
        result[i] = mapper(v)
    }
    return result
}

func main() {
    intSlice := []int{1, 2, 3, 4, 5}
    intMapper := func(a int) int
    {
        return a * 2
    }
    fmt.Println(Map(intSlice, intMapper))

    stringSlice := []string{"a", "b", "c", "d", "e"}
    stringMapper := func(s string) string {
        return s + s
    }
    fmt.Println(Map(stringSlice, stringMapper))
}
```

3. 泛型冒泡排序

```go
package main

import (
    "fmt"
    "sort"
)

func BubbleSort[T any](arr []T, less func(T, T) bool) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-1-i; j++ {
            if less(arr[j+1], arr[j]) {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    intSlice := []int{5, 3, 8, 2, 1}
    BubbleSort(intSlice, func(a, b int) bool {
        return a < b
    })
    fmt.Println(intSlice)

    stringSlice := []string{"e", "c", "b", "a", "d"}
    BubbleSort(stringSlice, func(a, b string) bool {
        return a < b
    })
    fmt.Println(stringSlice)
}
```
