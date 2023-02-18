
### regexp package：正则表达式
正则表达式是一种进行模式匹配和文本操纵的复杂而又强大的工具。虽然正则表达式比纯粹的文本匹配效率低，但是它却更灵活，按照它的语法规则，根据需求构造出的正则表达式能够从原始文本中筛选出几乎任何你想要得到的字符组合。

Go语言通过 regexp 包为正则表达式提供了官方支持，其采用 RE2 语法，除了\c、\C外，Go语言和 Perl、Python 等语言的正则基本一致。

**正则表达式语法规则**  
正则表达式是由普通字符（例如字符 a 到 z）以及特殊字符（称为"元字符"）构成的文字序列，可以是单个的字符、字符集合、字符范围、字符间的选择或者所有这些组件的任意组合。

1) 字符  
[![LTl1yD.md.png](https://s1.ax1x.com/2022/04/25/LTl1yD.md.png)](https://imgtu.com/i/LTl1yD)

2) 数量词（用在字符或 (...) 之后）  
[![LTlafP.md.png](https://s1.ax1x.com/2022/04/25/LTlafP.md.png)](https://imgtu.com/i/LTlafP)  

3) 边界匹配  
[![LTl6Ts.md.png](https://s1.ax1x.com/2022/04/25/LTl6Ts.md.png)](https://imgtu.com/i/LTl6Ts)

4) 逻辑、分组  
[![LT1OEj.md.png](https://s1.ax1x.com/2022/04/25/LT1OEj.md.png)](https://imgtu.com/i/LT1OEj)

5) 特殊构造（不作为分组）  
[![LT1z80.md.png](https://s1.ax1x.com/2022/04/25/LT1z80.md.png)](https://imgtu.com/i/LT1z80)

**Regexp 包的使用**   
【示例 1】匹配指定类型的字符串
```go
package main
import (
    "fmt"
    "regexp"
)
func main() {
    buf := "abc azc a7c aac 888 a9c  tac"
    //解析正则表达式，如果成功返回解释器
    reg1 := regexp.MustCompile(`a.c`)
    if reg1 == nil {
        fmt.Println("regexp err")
        return
    }
    //根据规则提取关键信息
    result1 := reg1.FindAllStringSubmatch(buf, -1)
    fmt.Println("result1 = ", result1)
}
// result1 =  [[abc] [azc] [a7c] [aac] [a9c]]　
```

【示例 2】匹配 a 和 c 中间包含一个数字的字符串。
```go
package main
import (
    "fmt"
    "regexp"
)
func main() {
    buf := "abc azc a7c aac 888 a9c  tac"
    //解析正则表达式，如果成功返回解释器
    reg1 := regexp.MustCompile(`a[0-9]c`)
    // 或
    // reg1 := regexp.MustCompile(`a\dc`)
    if reg1 == nil { //解释失败，返回nil
        fmt.Println("regexp err")
        return
    }
    //根据规则提取关键信息
    result1 := reg1.FindAllStringSubmatch(buf, -1)
    fmt.Println("result1 = ", result1)
}
// result1 =  [[a7c] [a9c]]
```

【示例 3】匹配字符串中的小数。
```go
package main
import (
    "fmt"
    "regexp"
)
func main() {
    buf := "43.14 567 agsdg 1.23 7. 8.9 1sdljgl 6.66 7.8   "
    //解释正则表达式
    reg := regexp.MustCompile(`\d+\.\d+`)
    if reg == nil {
        fmt.Println("MustCompile err")
        return
    }
    //提取关键信息
    //result := reg.FindAllString(buf, -1)
    result := reg.FindAllStringSubmatch(buf, -1)
    fmt.Println("result = ", result)
}
// result =  [[43.14] [1.23] [8.9] [6.66] [7.8]]
```

【示例 4】匹配 div 标签中的内容。
```go
package main
import (
    "fmt"
    "regexp"
)
func main() {
    // 原生字符串
    buf := `
    
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>C语言中文网 | Go语言入门教程</title>
</head>
<body>
    <div>Go语言简介</div>
    <div>Go语言基本语法
    Go语言变量的声明
    Go语言教程简明版
    </div>
    <div>Go语言容器</div>
    <div>Go语言函数</div>
</body>
</html>
    `
    //解释正则表达式
    reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
    if reg == nil {
        fmt.Println("MustCompile err")
        return
    }
    //提取关键信息
    result := reg.FindAllStringSubmatch(buf, -1)
    //过滤<></>
    for _, text := range result {
        fmt.Println("text[1] = ", text[1])
    }
}
// text[1] =  Go语言简介
// text[1] =  Go语言基本语法
//     Go语言变量的声明
//     Go语言教程简明版
//   
// text[1] =  Go语言容器
// text[1] =  Go语言函数
```

【示例 5】通过 Compile 方法返回一个 Regexp 对象，实现匹配，查找，替换相关的功能。
```go
package main
import (
    "fmt"
    "regexp"
    "strconv"
)
func main() {
    //目标字符串
    searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
    pat := "[0-9]+.[0-9]+"          //正则
    f := func(s string) string{
        v, _ := strconv.ParseFloat(s, 32)
        return strconv.FormatFloat(v * 2, 'f', 2, 32)
    }
    if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
        fmt.Println("Match Found!")
    }
    // Compile 方法可以解析并返回一个正则表达式，如果成功返回，则说明该正则表达式正确可用于匹配文本
    // 也可以使用 MustCompile 方法，它也可以像 Compile 方法一样检验正则的有效性，但是当正则不合法时程序将 panic 
    re, _ := regexp.Compile(pat)
    //将匹配到的部分替换为 "##.#"
    str := re.ReplaceAllString(searchIn, "##.#")
    fmt.Println(str)
    //参数为函数时
    str2 := re.ReplaceAllStringFunc(searchIn, f)
    fmt.Println(str2)
}
// Match Found!
// John: ##.# William: ##.# Steve: ##.#
// John: 5156.68 William: 9134.46 Steve: 11264.36
```

**验证字符串匹配模式**  
```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`^\d+$`)
	fmt.Println(re.MatchString("13877474"))
}
```

**子模式匹配**  
```go
package main

import (
	"fmt";
	"regexp"
)

func main() {
re := regexp.MustCompile("\\$\\{(.*?)\\}")
match := re.FindStringSubmatch("git commit -m '${abc}'")
	fmt.Println(match)
}
```
SKU有两类一类是定值的比如"sku_goods1_9000_0"，这种就需要验证下单参数的中的金额与SKU中的金额是否是一样， 另外一种充值订单对应的SKU不是定值的就不用验证，比如"sku_recharge_other_0"，写一个函数完成这个功能。
```go
// 检查创建订单时参数里的amount与skuSymbol中的amount是否一致
func checkSkuAmount(skuSymbol string, amount int64) bool {
	if strings.Contains(skuSymbol, "other") {
		return true
	}
	re := regexp.MustCompile(`[A-Za-z_]+?(\d+?)_\d+$`)
	match := re.FindStringSubmatch(skuSymbol)
	if match == nil {
		return false
	}
	skuAmount, err := strconv.ParseInt(match[1], 10, 64)
	if err != nil || skuAmount != amount {
		return false
	}

	return true
}
```

**正则替换**  
```go
func HideUserName(userName string) string {
	if userName == "" {
		return ""
	}
	re := regexp.MustCompile(`^(.).+?(.)?$`)
	result := re.ReplaceAllString(userName, `$1*$2`)
	return result
}
```

# 例子

## 是否匹配

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, err := regexp.MatchString("h[a-z]+.*d$", "hello world")
	if err != nil {
		panic(err)
	}
	fmt.Println(match)

	match, err = regexp.MatchString("h[a-z]+.*d$", "ello world")
	if err != nil {
		panic(err)
	}
	fmt.Println(match)
}

// $ go run main.go
// 输出如下
/**
  true
  false
*/
```

## 匹配所有子字符串

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	c, err := regexp.Compile("h[a-z]")
	if err != nil {
		panic(err)
	}

	res := c.FindAllString("hello world", -1)
	fmt.Printf("res = %v\n", res)

	res2 := c.FindAllString("hello world hi ha h1", -1)
	fmt.Printf("res2 = %v\n", res2)
}

// $ go run main.go
// 输出如下
/**
  res = [he]
  res2 = [he hi ha]
*/
```

## 替换所有子字符串

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	c, err := regexp.Compile("h[a-z]")
	if err != nil {
		panic(err)
	}

	res := c.ReplaceAll([]byte("hello world"), []byte("?"))
	fmt.Printf("res = %s\n", res)

	res2 := c.ReplaceAll([]byte("hello world hi ha h1"), []byte("?"))
	fmt.Printf("res2 = %s\n", res2)
}

// $ go run main.go
// 输出如下
/**
  res = ?llo world
  res2 = ?llo world ? ? h1
*/
```

## 匹配中文

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, err := regexp.MatchString("\\x{4e00}-\\x{9fa5}", "hello world")
	if err != nil {
		panic(err)
	}
	fmt.Println(match)

	match, err = regexp.MatchString("\\p{Han}+", "hello 世界")
	if err != nil {
		panic(err)
	}
	fmt.Println(match)
}

// $ go run main.go
// 输出如下
/**
  false
  true
*/
```

