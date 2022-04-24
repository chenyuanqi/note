
### big package：对整数的高精度计算
实际开发中，对于超出 int64 或者 uint64 类型的大数进行计算时，如果对精度没有要求，使用 float32 或者 float64 就可以胜任，但如果对精度有严格要求的时候，我们就不能使用浮点数了，因为浮点数在内存中只能被近似的表示。  
Go 语言中 math/big 包实现了大数字的多精度计算，支持 Int（有符号整数）、Rat（有理数）和 Float（浮点数）等数字类型。  
这些类型可以实现任意位数的数字，只要内存足够大，但缺点是需要更大的内存和处理开销，这使得它们使用起来要比内置的数字类型慢很多。  
```go
// 在 math/big 包中，Int 类型定义
// An Int represents a signed multi-precision integer.
// The zero value for an Int represents the value 0.
type Int struct {
    neg bool // sign
    abs nat  // absolute value of the integer
}

// 生成 Int 类型的方法为 NewInt ()
// NewInt allocates and returns a new Int set to x.
func NewInt(x int64) *Int {
    return new(Int).SetInt64(x)
}
```
`注意：NewInt () 函数只对 int64 有效，其他类型必须先转成 int64 才行。`

Go 语言中还提供了许多 Set 函数，可以方便的把其他类型的整形存入 Int ，因此，我们可以先 new (int) 然后再调用 Set 函数。  
```go
// SetInt64 函数将 z 转换为 x 并返回 z。
func (z *Int) SetInt64(x int64) *Int {
    neg := false
    if x < 0 {
        neg = true
        x = -x
    }
    z.abs = z.abs.setUint64(uint64(x))
    z.neg = neg
    return z
}
​
// SetUint64 函数将 z 转换为 x 并返回 z。
func (z *Int) SetUint64(x uint64) *Int {
    z.abs = z.abs.setUint64(x)
    z.neg = false
    return z
}
​
// Set 函数将 z 转换为 x 并返回 z。
func (z *Int) Set(x *Int) *Int {
    if z != x {
        z.abs = z.abs.set(x.abs)
        z.neg = x.neg
    }
    return z
}

package main
import (
    "fmt"
    "math/big"
)
func main() {
    big1 := new(big.Int).SetUint64(uint64(1000))
    fmt.Println("big1 is: ", big1) // big1 is:  1000
    big2 := big1.Uint64()
    fmt.Println("big2 is: ", big2) // big2 is:  1000
}
```

除了上述的 Set 函数，math/big 包中还提供了一个 SetString () 函数，可以指定进制数，比如二进制、十进制或者十六进制等。  
```go
// SetString sets z to the value of s, interpreted in the given base,
// and returns z and a boolean indicating success. The entire string
// (not just a prefix) must be valid for success. If SetString fails,
// the value of z is undefined but the returned value is nil.
//
// The base argument must be 0 or a value between 2 and MaxBase. If the base
// is 0, the string prefix determines the actual conversion base. A prefix of
// ``0x'' or ``0X'' selects base 16; the ``0'' prefix selects base 8, and a
// ``0b'' or ``0B'' prefix selects base 2. Otherwise the selected base is 10.
//
func (z *Int) SetString(s string, base int) (*Int, bool) {
    r := strings.NewReader(s)
    if _, _, err := z.scan(r, base); err != nil {
        return nil, false
    }
    // entire string must have been consumed
    if _, err := r.ReadByte(); err != io.EOF {
        return nil, false
    }
    return z, true // err == io.EOF => scan consumed all of s
}

package main
import (
    "fmt"
    "math/big"
)
func main() {
    big1, _ := new(big.Int).SetString("1000", 10)
    fmt.Println("big1 is: ", big1) // big1 is:  1000
    big2 := big1.Uint64()
    fmt.Println("big2 is: ", big2) // big2 is:  1000
}
```

因为 Go 语言不支持运算符重载，所以所有大数字类型都有像是 Add () 和 Mul () 这样的方法。
```go
// Add 方法的定义，将 z 转换为 x + y 并返回 z
func (z *Int) Add(x, y *Int) *Int
```

**计算第 1000 位的斐波那契数列**  
```go
package main
import (
    "fmt"
    "math/big"
    "time"
)
const LIM = 1000 //求第1000位的斐波那契数列
var fibs [LIM]*big.Int //使用数组保存计算出来的数列的指针
func main() {
    result := big.NewInt(0)
    start := time.Now()
    for i := 0; i < LIM; i++ {
        result = fibonacci(i)
        fmt.Printf("数列第 %d 位: %d\n", i+1, result)
    }
    end := time.Now()
    delta := end.Sub(start)
    fmt.Printf("执行完成，所耗时间为: %s\n", delta)
}
func fibonacci(n int) (res *big.Int) {
    if n <= 1 {
        res = big.NewInt(1)
    } else {
        temp := new(big.Int)
        res = temp.Add(fibs[n-1], fibs[n-2])
    }
    fibs[n] = res
    return
}
```

