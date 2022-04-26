
### 单元测试
在不使用 “单元测试” 的情况下，我们如何测试一个函数或方法的正确性？
```go
func  Add(num1, num2 int) int {
    return num1 + num2
}

func  main() {
    excepted := 5
    actual := Add(2, 3)
    if excepted == actual {
        fmt.Println("成功")
    } else {
        fmt.Println("失败")
    }
}
```
这样的测试方式，它有这些问题：测试代码和业务代码混乱、不分离；测试完后，测试代码必须删除；如果不删除，会参与编译。

**什么是单元测试**  
单元测试又称为模块测试，是针对程序模块（软件设计的最小单元）来进行正确性检验的测试工作。在 Go 语言中，测试的最小单元常常是函数和方法。

**测试文件**  
在很多语言中，常常把测试文件放在一个独立的目录下进行管理，而在 Go 语言中会和源文件放置在一块，即同一目录下。  
假如源文件的命名是 xxx.go, 那单元测试文件的命名则为 xxx_test.go。如果在编译阶段 xxx_test.go 文件会被忽略。
```go
// 对于上面的 Add 函数，所在文件是 add.go，那创建的测试文件也和它放在一块
// unitest 目录
//     add.go
//     add_test.go 单元测试
```

**单元测试文件-基本结构&内容**  
```go
// gobasic/unittest/add_test.go
package unittest
// 导入 testing 标准包
import  "testing"
// 创建一个 Test 开头的函数名 TestAdd，Test 是固定写法，后面的 Add 一般和你要测试的函数名对应，当然不对应也没有问题
// 参数类型 *tesing.T 用于打印测试结果，参数中也必须跟上
func  TestAdd(t *testing.T) {
    // excepted 函数期待的结果
    excepted := 4
    // actual 函数真实计算的结果
    actual := Add(2, 3)
    // 如果不相等，打印出错误
    if excepted != actual {
        t.Errorf("excepted：%d, actual:%d", excepted, actual)
    }
}
```
在 unittest 目录下运行 go test （或 go test ./）命令，表示运行 unittest 目录下的单元测试，不会再往下递归。如果想往下递归，即当前目录下还有目录，则运行 go test ./... 命令。
```bash
$ go test
--- FAIL: TestAdd (0.00s)
     add_test.go:11: excepted：4, actual:5
 FAIL
 FAIL    github.com/miaogaolin/gobasic/unittest  0.228s
 FAIL
```

**\*testing.T**  
参数类型 T 中的几个方法：
- Error 打印错误日志、标记为失败 FAIL，并继续往下执行。
- Errorf 格式化打印错误日志、标记为失败 FAIL，并继续往下执行。
- Fail 不打印日志，结果中只标记为失败 FAIL，并继续往下执行。
- FailNow 不打印日志，结果中只标记为失败 FAIL，但在当前测试函数中不继续往下执行。
- Fatal 打印日志、标记为失败，并且内部调用了 FaileNow 函数，也不往下执行。
- Fatalf 格式化打印错误日志、标记为失败，并且内部调用了 FaileNow 函数，也不往下执行。

`没有成功的方法，只要没有通知错误，那就说明是正确的。`

**测试资源**   
有时候在写单元测试时，可能需要读取文件，那这些相关的资源文件就放置在 testdata 目录下。
```go
// unitest 目录
//     add.go
//     add_test.go 单元测试
//     testdata 目录
```

**go test 和 go vet**  
在运行 go test 命令后，go vet 命令也会自动运行。go vet 命令用于代码的静态分析，检查编译器检查不出的错误。  
`在测试时无需单独运行 go vet 命令，一个 go test 命令就包含了。`
```go
package main

import  "fmt"

func  main() {
    // 占位符 % d 需要的是整数，但给的是字符串
    fmt.Printf("%d", "miao")
}
```
对于这种类似的错误，编译器是不会报错的，这时候就用到了 go vet 命令。
```bash
$ go vet
# github.com/miaogaolin/gobasic/vet
.\main.go:6:2: Printf format %d has arg "miao" of wrong type string
```

**表格驱动测试**  


