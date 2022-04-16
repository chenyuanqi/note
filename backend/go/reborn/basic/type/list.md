
### Go list
列表是一种非连续的存储容器，由多个节点组成，节点通过一些变量记录彼此之间的关系，列表有多种实现方法，如单链表、双链表等。  
在 Go 语言中，列表使用 container/list 包来实现，内部的实现原理是双链表，列表能够高效地进行任意位置的元素插入和删除操作。  

**初始化**  
list 的初始化有两种方法：分别是使用 New() 函数和 var 关键字声明，两种方法的初始化效果都是一致的。 
```go
// 1) 通过 container/list 包的 New() 函数初始化 list
变量名 := list.New()

// 通过 var 关键字声明初始化 list
var 变量名 list.List
```
列表与切片和 map 不同的是，列表并没有具体元素类型的限制，因此，列表的元素可以是任意类型，这既带来了便利，也引来一些问题，例如给列表中放入了一个 interface{} 类型的值，取出值后，如果要将 interface{} 转换为其他类型将会发生宕机。

**在列表中插入元素**  
双链表支持从队列前方或后方插入元素，分别对应的方法是 PushFront 和 PushBack。
`这两个方法都会返回一个 *list.Element 结构，如果在以后的使用中需要删除插入的元素，则只能通过 *list.Element 配合 Remove() 方法进行删除，这种方法可以让删除更加效率化，同时也是双链表特性之一。`  
- InsertAfter(v interface {}, mark *Element) *Element	在 mark 点之后插入元素，mark 点由其他插入函数提供
- InsertBefore(v interface {}, mark * Element) *Element	在 mark 点之前插入元素，mark 点由其他插入函数提供
- PushBackList(other *List)	添加 other 列表元素到尾部
- PushFrontList(other *List)	添加 other 列表元素到头部

```go
l := list.New()
l.PushBack("fist")
l.PushFront(67)
```

**从列表中删除元素**  
列表插入函数的返回值会提供一个 \*list.Element 结构，这个结构记录着列表元素的值以及与其他节点之间的关系等信息，从列表中删除元素时，需要用到这个结构进行快速删除。
```go
package main

import "container/list"

func main() {
    l := list.New()
    // 尾部添加
    l.PushBack("canon") // canon
    // 头部添加
    l.PushFront(67) //	67, canon
    // 尾部添加后保存元素句柄
    element := l.PushBack("fist") // 67, canon, fist
    // 在fist之后添加high
    l.InsertAfter("high", element) // 67, canon, fist, high
    // 在fist之前添加noon
    l.InsertBefore("noon", element) // 67, canon, noon, fist, high
    // 使用
    l.Remove(element) // 67, canon, noon, high
}
```

**遍历列表——访问列表的每一个元素**  
遍历双链表需要配合 Front() 函数获取头元素，遍历时只要元素不为空就可以继续进行，每一次遍历都会调用元素的 Next() 函数。  
```go
l := list.New()
// 尾部添加
l.PushBack("canon")
// 头部添加
l.PushFront(67)
for i := l.Front(); i != nil; i = i.Next() {
    fmt.Println(i.Value)
    // 67
    // canon
}
```
