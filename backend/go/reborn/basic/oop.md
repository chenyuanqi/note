
### Go 面向对象
封装、继承、多态是面向对象的 3 个基本特征。  
Golang 实现面向对象的两个关键类型是 struct 和 interface。
> Go 支持面向对象(OOP)，并不是纯粹的面向对象语言；  
> Go 没有类的概念，结构体(struct)相当于其它编程语言的类(class)；  
> Go 面向对象编程非常简洁，通过接口(interface)关联，耦合性低，也非常灵活；  

**封装**   
封装就是把抽象出来的字段和操作方法封装在一起，数据被保护在内部，只有通过操作方法，才能对字段进行操作。  
```go
package main
​
import "fmt"
​
type Person struct { // 抽象出来的字段
    name string
}
​
func (person *Person) setName(name string) { // 封装方法
    person.name = name
}
​
func (person *Person) GetInfo() { // 封装方法
    fmt.Println(person.name)
}
​
func main() {
    p := Person{"go"}
    p.setName("golang")
    p.GetInfo() // 输出 新名字
}
```

**继承**  
继承顾名思义，可以解决代码复用。在 Go 中，只需要在结构体中嵌套一个匿名结构体。Go 没有显式的继承，而是通过组合实现继承。  
```go
package main
​
import "fmt"
​
type Person struct { // 抽象出的字段
    name string
}
​
func (p *Person) getName() { // 封装方法
  fmt.Println(p.name)
}
​
type Inherit struct { // 继承
    Person
}
​
func main() {
    i := Inherit{Person{"go"}}
    i.getName() // 输出 go
}
```

**多态**   
把它们共同的方法提炼出来定义一个抽象的接口，就是多态。  
```go
package main
​
import "fmt"
​
type Birds interface { // 将共同的方法提炼出来
    fly()
}
​
type Dove struct {
}
​
type Eagle struct {
}
​
func (d *Dove) fly() { // 封装方法
    fmt.Println("鸽子在飞")
}
​
func (e *Eagle) fly() { // 封装方法
    fmt.Println("鹰在飞")
}
​
func main() {
    var b Birds // 多态的抽象接口
​
    dove := Dove{}
    eagle := Eagle{}
​
    b = &dove
    b.fly() // 鸽子在飞
​
    b = &eagle
    b.fly() // 鹰在飞
}
```
