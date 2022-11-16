

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

// defer
// 延迟执行；参数预计算；LIFO 执行顺序
func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

// 接口
// 声明和定义
type Shape interface {
  perimeter() float64
  area() float64
}
var s Shape
// 隐式地让一个类型实现接口
type Rectangle struct {
  a, b float64
}
func (r Rectangle) perimeter() float64 {
  return (r.a + r.b) * 2
}
func (r Rectangle) area() float64 {
  return r.a * r.b
}
// 接口的动态调用方式
var s Shape
s = Rectangle{3, 4}
s.perimeter()
s.area()
// 接口的嵌套
type ReadWriter interface {
  Reader
  Writer
}
type Reader interface {
  Read(p []byte) (n int, err error)
}
type Writer interface {
  Write(p []byte) (n int, err error)
}
// 接口的类型断言
func main(){
  var s Shape
  s = Rectangle{3, 4}
  rect := s.(Rectangle)
  fmt.Printf("长方形周长:%v, 面积:%v \\n",rect.perimeter(),rect.area())
}
// 根据空接口中动态类型的差异选择不同的处理方式
switch f := arg.(type) {
  case bool:
    p.fmtBool(f, verb)
  case float32:
    p.fmtFloat(float64(f), 32, verb)
  case float64:
    p.fmtFloat(f, 64, verb)
```

**并发编程**  
在 Go 语言中与并发编程紧密相关的就是协程与通道。  
- 进程、线程与协程。进程是操作系统资源分配的基本单位，线程是操作系统资源调度的基本单位。而协程位于用户态，是在线程基础上构建的轻量级调度单位。  
- 并发与并行。并行指的是同时做很多事情，并发是指同时管理很多事情。  
- 主协程与子协程。 main 函数是特殊的主协程，它退出之后整个程序都会退出。而其他的协程都是子协程，子协程退出之后，程序正常运行。

```golang
// 通道声明与初始化
chan int
chan <- float
<-chan string
// 通道写入数据
c <- 5
// 通道读取数据
data := <- c
// 通道关闭
close(c)
// 通道作为参数
func worker(id int, c chan int) {
  for n := range c {
    fmt.Printf("Worker %d received %c\\n",
      id, n)
  }
}
// 通道作为返回值（一般用于创建通道的阶段）
func createWorker(id int) chan int {
  c := make(chan int)
  go worker(id, c)
  return c
}
// 单方向的通道，用于只读和只写场景
func worker(id int, c <-chan int)
// select 监听多个通道实现多路复用。当 case 中多个通道状态准备就绪时，select 随机选择一个分支进行执行
select {
  case <-ch1:
    // ...
  case x := <-ch2:
    // ...use x...
  case ch3 <- y:
    // ...
  default:
    // ...
  }

// 用 context 来处理协程的优雅退出和级联退出

func Stream(ctx context.Context, out chan<- Value) error {
  for {
    v, err := DoSomething(ctx)
    if err != nil {
      return err
    }
    select {
    case <-ctx.Done():
      return ctx.Err()
    case out <- v:
    }
  }
// 传统的同步原语：原子锁。Go 提供了 atomic 包用于处理原子操作
func add() {
  for {
    if atomic.CompareAndSwapInt64(&flag, 0, 1) {
      count++
      atomic.StoreInt64(&flag, 0)
      return
    }
  }
}
// 传统的同步原语：互斥锁
var m sync.Mutex
func add() {
  m.Lock()
  count++
  m.Unlock()
}
// 传统的同步原语：读写锁。适合多读少写场景

type Stat struct {
  counters map[string]int64
  mutex sync.RWMutex
}
func (s *Stat) getCounter(name string) int64 {
  s.mutex.RLock()
  defer s.mutex.RUnlock()
  return s.counters[name]
}
func (s *Stat) SetCounter(name string){
  s.mutex.Lock()
  defer s.mutex.Unlock()
  s.counters[name]++
}
```




