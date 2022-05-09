
### Go 泛型
泛型程序设计（generic programming）是程序设计语言的一种风格或范式。泛型允许程序员在强类型语言中编写代码时，使用一些以后才确定的类型，其在真正实例化时才会为这些参数指确定类型。另外各语言和其编译器、运行环境对泛型的支持均不一样，因此需要针对来辩证。  
泛型本质上并不是绝对的必需品，更不是 Go 语言的早期目标，因此在过往的发展阶段没有过多重视这一点，而是把精力放在了其他 feature 上。

简单来讲，泛型就是参数化多态。其可根据实参类型生成不同的版本，支持任意数量的调用。
```go
func F(a, b T) T{ return a+b }

// T 为 int
F(1, 2)

// T 为 string
F("1", "2")
```
在编译时期编译器便确定其 T 的入参类型。这也是 Go 泛型实现的要求之一 “编译时类型安全”。

