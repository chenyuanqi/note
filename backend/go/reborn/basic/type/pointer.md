
### Go 指针
Go语言中的指针，它主要由两大核心概念构成：类型指针和切片指针。
- 类型指针：在传递数据时直接使用指针，可以避免创建数据的副本，节约内存开销。类型指针不能进行偏移和运算，可以避免非法修改为其它数据的风险，也更有利于垃圾回收机制及时找到并回收它们；
- 切片指针：切片由指向起始元素的指针、元素数量和总容量构成。当访问切片发生越界时，会发生宕机并输出堆栈信息。宕机是可以恢复的，而崩溃只能导致程序停止运行。

使用指针更有利于程序运行的性能和稳定性。另外，在某些操作中，如使用反射修改变量的值，必须使用可寻址的变量（通过指针）。  
在实际应用中，最为常用的便是获取变量的内存地址，以及获取某个地址对应的值。在Go语言中，前者使用“&”运算符，后者使用“*”运算符。它们互为反向操作，操作的对象也不同。
```go
//exampleNumberA变量（整数型变量）声明和赋值
var exampleNumberA int = 10
//获取exampleNumberA的地址，并赋值给exampleNumberAPtr变量（exampleNumberAPtr的类型是指针类型）
exampleNumberAPtr := &exampleNumberA
//输出exampleNumberAPtr变量的值（将输出内存地址）
fmt.Println(exampleNumberAPtr) // 0xc00001a088
//获取exampleNumberAPtr（指针变量）表示的实际数据值，并赋值给exampleNumberAPtrValue变量（整数型变量）
exampleNumberAPtrValue := *exampleNumberAPtr
//输出exampleNumberAPtrValue变量（整数型变量）的值
fmt.Println(exampleNumberAPtrValue) // 10
```

我们还可以使用 new() 函数直接创建指针变量，相当于在内存中创建了没有变量名的某种类型的变量。这样做无需产生新的数据“代号”，取值和赋值转而通过指针变量完成。常用在无需变量名或必须要传递指针变量值的场景中。  
```go
//使用new()函数创建名为exampleNumberAPtr指针类型变量，表示int64型值
exampleNumberAPtr := new(int64)
//修改exampleNumberAPtr表示的实际数据值
*exampleNumberAPtr = 100
//获取exampleNumberAPtr表示的实际数据值
fmt.Println(*exampleNumberAPtr) // 100
```
