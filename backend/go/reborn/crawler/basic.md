

### Golang 
[Go开挂入门](https://book.golang-dream.com/)  

### Golang 基础
基础部分主要包括：开发环境；基础语法；语法特性；并发编程；项目组织；工具与库。  

Go 的特性：  
- 良好的编译器和依赖设计  
- 面向组合而不是继承  
- 并发原语  
- 简单与健壮性  
- 强大丰富的标准库与工具集  

```golang
// 基础语法

// 内置类型
// int  int8  int16  int32  int64
// uint  uint8  uint16  uint32  uint64  uintptr
// float32  float64  complex128  complex64
// bool  byte  rune  string

// 表达式与运算符
// 优先级(由高到低)              操作符
//   5                *  /  %  <<  >>  &  &^
//   4                +  -  |  ^
//   3                ==  !=  <  <=  >  >=
//   2                &&
//   1                ||

// 基本控制结构
// 顺序结构
if{

} else if {

} else {

}
switch var1 {
    case val1:
        ...
    case val2,val3:
        ...
    default:
        ...
}
// 循环结构
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
i := 1
for i < 100 {
    fmt.Println(i)
    i = i * 2
}
for {
    fmt.Println("Hello")
}
evenVals := []int{2, 4, 6, 8, 10, 12}
for i, v := range evenVals {
    fmt.Println(i, v)
}

// 函数
// 基本声明
func name(parameter-list) (result-list) {
    body
}
// 多返回值
func div (a,b int) (int,error){
    if b == 0 {
     return 0, errors.New("b cat't be 0")
    }
    return a/b,nil
}
// 可变参数
func Println(a ...interface{}) (n int, err error)
// 递归
func f(n int) int {
  if n == 1 {
    return 1
  }
  return n * f(n-1)
}
// 函数作为参数
// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {
    for _, v := range list {
        f(v)
    }
}

func main() {
    // 使用匿名函数打印切片内容
    visit([]int{1, 2, 3, 4}, func(v int) {
        fmt.Println(v)
    })
}
// 函数作为返回值
func logging(f http.HandlerFunc) http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL.Path)
    f(w,r)
  }
}
// 函数作为值
var opMap = map[string]func(int, int) int{
    "+": add,
    "-": sub,
    "*": mul,
    "/": div,
}
f := opMap[op]
f()

// 复合类型
// 切片
// 声明与赋值
var slice1 []int
numbers:= []int{1,2,3,4,5,6,7,8}
var x = []int{1, 5: 4, 6, 10: 100, 15}
// 添加元素
y := []int{20, 30, 40}
x = append(x, y...)
// 截取
numbers:= []int{1,2,3,4,5,6,7,8}
// 从下标2 一直到下标4，但是不包括下标4
numbers1 :=numbers[2:4]
// 从下标0 一直到下标3，但是不包括下标3
numbers2 :=numbers[:3]
// 从下标3 一直到结尾
numbers3 :=numbers[3:]

// map
// 声明和初始化
var hash map[T]T
var hash = make(map[T]T,NUMBER)
var country = map[string]string{
"China": "Beijing",
"Japan": "Tokyo",
"India": "New Delhi",
"France": "Paris",
"Italy": "Rome",
}
// 访问
v := hash[key]
v,ok := hash[key]
// 赋值和初始化
m := map[string]int{
    "hello": 5,
    "world": 10,
}
delete(m, "hello")

// 结构体
// 声明和赋值
type Nat struct {
    n  int
    d  int
}
var nat Nat
nat := Nat{
    2,
    3
}
nat.n = 4
natq := Nat{
    d:  3,
    n:  2,
}
// 匿名结构体
var person struct {
    name string
    age  int
    pet  string
}

pet := struct {
    name string
    kind string
}{
    name: "Fido",
    kind: "dog",
}
```







