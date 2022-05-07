
### select 语句
select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。

select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。 select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。  
在一个 select 语句中，Go 会按顺序从头到尾评估每一个发送和接收的语句。  
如果其中的任意一个语句可以继续执行（即没有被阻塞），那么就从那些可以执行的语句中任意选择一条来使用。 如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有两种可能的情况：   
①如果给出了 default 语句，那么就会执行 default 的流程，同时程序的执行会从 select 语句后的语句中恢复。   
②如果没有 default 语句，那么 select 语句将被阻塞，直到至少有一个 case 可以进行下去。
```go
// 每个case都必须是一个通信
// 所有channel表达式都会被求值
// 所有被发送的表达式都会被求值
// 如果任意某个通信可以进行，它就执行；其他被忽略。
// 如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
// 否则：
// 如果有default子句，则执行该语句。
// 如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值
select {
    case communication clause:
       statement(s);      
    case communication clause:
       statement(s);
    /* 你可以定义任意数量的 case */
    default : /* 可选 */
       statement(s);
}

select { //不停的在这里检测
    case <-chanl : //检测有没有数据可以读
    //如果chanl成功读取到数据，则进行该case处理语句
    case chan2 <- 1 : //检测有没有可以写
    //如果成功向chan2写入数据，则进行该case处理语句


    //假如没有default，那么在以上两个条件都不成立的情况下，就会在此阻塞//一般default会不写在里面，select中的default子句总是可运行的，因为会很消耗CPU资源
    default:
    //如果以上都没有符合条件，那么则进行default处理流程
}
```

select 可以监听 channel 的数据流动。  
select 的用法与 switch 语法非常类似，由 select 开始的一个新的选择块，每个选择条件由 case 语句来描述。  
与 switch 语句可以选择任何使用相等比较的条件相比，select 由比较多的限制，其中最大的一条限制就是每个 case 语句里必须是一个 IO 操作。  

**select 的使用及典型用法**  
select中的 default 子句总是可运行的。  
如果有多个 case 都可以运行，select 会随机公平地选出一个执行，其他不会执行。  
如果没有可运行的 case 语句，且有 default 语句，那么就会执行 default 的动作。  
如果没有可运行的 case 语句，且没有 default 语句，select 将阻塞，直到某个 case 通信可以运行。  
```go
// 示例
var c1, c2, c3 chan int
var i1, i2 int
select {
    case i1 = <-c1:
        fmt.Printf("received ", i1, " from c1\n")
    case c2 <- i2:
        fmt.Printf("sent ", i2, " to c2\n")
    case i3, ok := (<-c3):  // same as: i3, ok := <-c3
        if ok {
        fmt.Printf("received ", i3, " from c3\n")
        } else {
        fmt.Printf("c3 is closed\n")
        }
    default:
        fmt.Printf("no communication\n")
} 
// no communication
```

超时判断  
比如在下面的场景中，使用全局 resChan 来接受 response，如果时间超过 3S,resChan 中还没有数据返回，则第二条 case 将执行。
```go
var resChan = make(chan int)
// do request
func test() {
    select {
    case data := <-resChan:
        doData(data)
    case <-time.After(time.Second * 3):
        fmt.Println("request time out")
    }
}

func doData(data int) {
    //...
}
```

退出
```go
//主线程（协程）中如下：
var shouldQuit=make(chan struct{})
fun main(){
    {
        //loop
    }
    //...out of the loop
    select {
        case <-c.shouldQuit:
            cleanUp()
            return
        default:
        }
    //...
}

// 在另外一个协程中，如果运行遇到非法操作或不可处理的错误，就向 shouldQuit 发送数据通知程序停止运行
close(shouldQuit)
```

判断 channel 是否阻塞
```go
//在某些情况下是存在不希望channel缓存满了的需求的，可以用如下方法判断
ch := make (chan int, 5)
//...
data：=0
select {
case ch <- data:
default:
    //做相应操作，比如丢弃 data。视需求而定
}
```
