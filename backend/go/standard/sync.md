
### Go 同步
协调并发代码，我们来了解一下同步（sync）。

```golang
var (
  counter = 0
  // 默认的的 sync.Mutex 是未锁定状态
  lock sync.Mutex
)
func main() {
  for i := 0; i < 20; i++ {
    go incr()
  }
  time.Sleep(time.Millisecond * 10)
}
func incr() {
  // 粗糙的锁操作（覆盖着大量代码的锁操作）,另外一个问题是与死锁有关
  lock.Lock()
  defer lock.Unlock()
  counter++
  fmt.Println(counter)
}

func main() {
  go func() { lock.Lock() }()
  time.Sleep(time.Millisecond * 10)
  lock.Lock()
}
```
首先，有一个常见的锁叫读写互斥锁。它主要提供了两种锁功能：一个锁定读取和一个锁定写入。它的区别是允许多个同时读取，同时确保写入是独占的。  
在 Go 中， sync.RWMutex 就是这种锁。另外 sync.Mutex 结构不但提供了 Lock 和 Unlock 方法 ，也提供了 RLock 和 RUnlock 方法；其中 R 代表 Read。虽然读写锁很常用，它也给开发人员带来了额外的负担：我们不但要关注我们正在访问的数据，还要注意如何访问。此外，部分并发编程不只是通过为数不多的代码按顺序的访问变量； 它也需要协调多个协程。  

