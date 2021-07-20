
### Go 指针
传值还是传指针？  
许多开发者认为应该传递 b 或者返回它在一个函数中会更加高效。然而，传递 / 返回的是切片的副本，但是切片本身就是一个引用。所以传递返回切片本身，没有什么区别。  
决定使用指针数组还是值数组归结为你如何使用单个值，而不是你用数组还是映射。
```golang
a := make([]Saiyan, 10)
// 或者
b := make([]*Saiyan, 10)
```