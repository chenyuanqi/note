
### 常见问题
- make 和 new 的区别
> make 会进行初始化，new 会返回一个零值的指针。  
> 
> new (T) 和 make (T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。
> new (T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T 的值。换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。
> make (T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make () 只适用于 slice、map 和 channel.

- 在多行 array、slice、map 语句中为什么最后一项需要，号
> 
```go
x := []int {
    1,
    2    // syntax error: unexpected newline, expecting comma or }
}
```

- Go 的 Slice 如何扩容？
> 在使用 append 向 slice 追加元素时，若 slice 空间不足则会发生扩容，扩容会重新分配一块更大的内存，将原 slice 拷贝到新 slice ，然后返回新 slice。扩容后再将数据追加进去。  
> 扩容操作只对容量，扩容后的 slice 长度不变，容量变化规则如下：  
> 若 slice 容量小于 1024 个元素，那么扩容的时候 slice 的 cap 就翻番，乘以 2；一旦元素个数超过 1024 个元素，增长因子就变成 1.25，即每次增加原来容量的四分之一。  
> 若 slice 容量够用，则将新元素追加进去，slice.len++，返回原 slice。  
> 若 slice 容量不够用，将 slice 先扩容，扩容得到新 slice，将新元素追加进新 slice，slice.len++，返回新 slice。  

