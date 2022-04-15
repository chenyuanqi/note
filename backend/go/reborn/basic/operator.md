
### Go 运算符
Go 语言总共提供了 6 种常用的运算符，分别为算术运算符、关系运算符、逻辑运算符、位运算符、赋值运算符以及指针运算符。  

**算术运算符**  
算术运算符的意义和使用和数学上的概念很类似。需要注意的是，在做除法时，对于int类别结果只保留整数，即使无法整除，余数也会被丢弃；若要获取余数，需要用到取余（%）运算符。  
```go
var exampleNumA int = 10
var exampleNumB int = 20
var exampleNumC int = 30
var exampleNumD = exampleNumA + exampleNumB*exampleNumC
fmt.Println(exampleNumD) // 610

var exampleNumA int = 10
var exampleNumB int = 3
fmt.Println(exampleNumA / exampleNumB) // 3
fmt.Println(exampleNumA % exampleNumB) // 1
```
还有一种较为精简的自增（++）和自减（--）运算符，它们相当于加1和减1，然后再将计算结果赋值给自身变量。
```go
var exampleNumA int = 10
// exampleNumA = exampleNumA + 1
exampleNumA++
fmt.Println(exampleNumA)
```

**关系运算符**  
关系运算符则用来判断二者的关系，当结果与判断条件一致时，返回 true，反之则返回 false。  
```go
var exampleNumA int = 10
var exampleNumB int = 20
fmt.Println(exampleNumA <= exampleNumB) // true
```

**逻辑运算符**  
逻辑运算符通常用来将两个条件组合，获得组合后的关系，最终将输出布尔类型值。其组合方式包括与 “&&”、或 “||”、非 “!”。
```go
var exampleBoolA bool = true
var exampleBoolB bool = false
//逻辑与运算。当exampleBoolA和exampleBoolB均为true时，结果为true；其他情况均为false。
fmt.Println(exampleBoolA && exampleBoolB) // false
//逻辑或运算。当exampleBoolA或exampleBoolB有一个为true时，结果为true；当exampleBoolA和exampleBoolB都是false时，结果为false。
fmt.Println(exampleBoolA || exampleBoolB) // true
//逻辑非运算。将某个布尔类型的值取反。
fmt.Println(!exampleBoolB) // true
```

**位运算符**   
位运算符运用在整数型变量，在进行运算时，会首先将其它进制的数值转换为二进制的数值，然后使用二进制数值进行运算，最后以原始进制类型返回计算结果。
```go
//十进制7转二进制结果为0111
var exampleNumA int = 7
//十进制5转二进制结果为0101
var exampleNumB int = 5
fmt.Println(exampleNumA & exampleNumB) // 5
```

**赋值运算符**  
为某个变量赋初值时使用的“=”便是最为简单的赋值运算符了。此外，还有一些更为简便的经过运算的赋值运算符。  
```go
var exampleNumA int = 10
//exampleNumA = exampleNumA + 20
exampleNumA += 20
fmt.Println(exampleNumA) // 30
```

**指针运算符**   
指针运算符包含 & 和 * 两个运算符。  

### Go 运算符优先级
在实际开发中，通常会处理较为复杂的运算，通常会将多个变量与多种运算符一齐使用，如此便不可避免地出现先后次序地问题。

对于算术运算符，遵循数学上的运算顺序。 比如在既有乘法又有加法的情况下，会先进行乘法运算，再进行加法运算。当然，如果我们希望先进行加法运算，可以使用成对的小括号将加法部分包裹起来。这些做法在 Go 语言中也是相同的。  

下表中，优先级值越大，优先级越高。

| 优先级 | 运算符 |  
| ---: | :--- |  
| 1 | , |  
| 2 | =、+=、-=、*=、/=、 %=、 >=、 <<=、&=、^=、\|= |  
| 3 | \|\| |  
| 4 | && |  
| 5 | \| |  
| 6 | ^ |  
| 7 | & |  
| 8 | ==、!= |  
| 9 | <、<=、>、>= |  
| 10 | <<、>> |  
| 11 | +、- |  
| 12 | *（乘号）、/、% |  
| 13 | !、*（指针）、& 、++、--、+（正号）、-（负号） |  
| 14 | \( \)、\[ \]、-> |  
