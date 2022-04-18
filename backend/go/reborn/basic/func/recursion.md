
### 递归函数
谓递归函数指的是在函数内部调用函数自身的函数，从数学解题思路来说，递归就是把一个大问题拆分成多个小问题，再各个击破，在实际开发过程中，递归函数可以解决许多数学问题，如计算给定数字阶乘、产生斐波系列等。  

构成递归需要具备以下条件：
- 一个问题可以被拆分成多个子问题；
- 拆分前的原问题与拆分后的子问题除了数据规模不同，但处理问题的思路是一样的；
- 不能无限制的调用本身，子问题需要有退出递归状态的条件。
`注意：编写递归函数时，一定要有终止条件，否则就会无限调用下去，直到内存溢出。`

**斐波那契数列**  
斐波那契数列的形式（num3 = num1 + num2）如下所示：
1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, …
```go
package main
import "fmt"
func main() {
    result := 0
    for i := 1; i <= 10; i++ {
        result = fibonacci(i)
        fmt.Printf("fibonacci(%d) is: %d\n", i, result)
    }
}
func fibonacci(n int) (res int) {
    if n <= 2 {
        res = 1
    } else {
        res = fibonacci(n-1) + fibonacci(n-2)
    }
    return
}
```

**数字阶乘**  
一个正整数的阶乘（factorial）是所有小于及等于该数的正整数的积，并且 0 的阶乘为 1，自然数 n 的阶乘写作 n!，“基斯顿·卡曼”在 1808 年发明了 n! 这个运算符号。  
例如，n!=1×2×3×…×n，阶乘亦可以递归方式定义：0!=1，n!=(n-1)!×n。  
```go
package main
import "fmt"
func Factorial(n uint64) (result uint64) {
    if n > 0 {
        result = n * Factorial(n-1)
        return result
    }
    return 1
}
func main() {
    var i int = 5
    fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}
// 递归的过程
// 5 * Factorial(4)
// 5 * (4 * Factorial(3))
// 5 * (4 * (3 * Factorial(2)))
// 5 * (4 * (3 * (2 * Factorial(1))))
// 5 * (4 * (3 * (2 * (1 * Factorial(0)))))
// 5 * (4 * (3 * (2 * (1 * 1))))
// 5 * (4 * (3 * (2 * 1)))
// 5 * (4 * (3 * 2))
// 5 * (4 * 6)
// 5 * 24
// 120
```

**多个函数组成递归**  
Go 语言中也可以使用相互调用的递归函数，多个函数之间相互调用形成闭环，因为 Go 语言编译器的特殊性，这些函数的声明顺序可以是任意的。  
```go
package main
import (
    "fmt"
)
func main() {
    fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
    fmt.Printf("%d is odd: is %t\n", 17, odd(17)) // 17 is odd: is true
    fmt.Printf("%d is odd: is %t\n", 18, odd(18)) // 18 is odd: is false
}
func even(nr int) bool {
    if nr == 0 {
        return true
    }
    return odd(RevSign(nr) - 1)
}
func odd(nr int) bool {
    if nr == 0 {
        return false
    }
    return even(RevSign(nr) - 1)
}
func RevSign(nr int) int {
    if nr < 0 {
        return -nr
    }
    return nr
}
```
