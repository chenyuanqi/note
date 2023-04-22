
### 编码风格
[指南参考](https://tonybai.com/google-go-style/google-go-style-guide/)  
[决定参考](https://tonybai.com/google-go-style/google-go-style-decisions/)  
[最佳实践](https://tonybai.com/google-go-style/google-go-style-best-practices/)  

#### 关于命名
Go语言中的名称一般不应包含下划线，不过这一原则有三个例外：
- 仅由生成的代码导入的包名可以包含下划线。
- 在*_test.go文件中的Test、Benchmark和Example函数名称可以包含下划线。
- 与操作系统或cgo互操作的低级库可以重复使用标识符，就像在syscall中那样。但这在大多数代码库中是非常罕见的。

包名  
Go包的名称应该是短小的，并且只包含小写字母。由多个单词组成的包名的各个单词之间应该没有间断。例如，我们使用tabwriter作为包名，而不是tabWriter，TabWriter或者tab_writer。  
避免选择那些可能被常用的局部变量名称遮蔽的包名。例如，usercount是一个比count更好的包名，因为count是一个常用的变量名。

Receiver命名  
Receiver变量的名称必须满足下面要求：  
- 短（通常为一或两个字母的长度）。
- 类型本身的缩略语。
- 统一应用于该类型的每一个Receiver。

```go
// func (tray Tray)
func (t Tray)
// func (info *ResearchInfo)
func (ri *ResearchInfo)
// func (this *ReportWriter)
func (w *ReportWriter)
// func (self *Scanner)
func (s *Scanner)
```

常量命名
常量名称必须像Go中的其他名称一样使用驼峰命名法(MixedCaps)(导出的常量以大写字母开始，而未导出的常量以小写字母开始)。  
根据常量的作用来命名常量，而不是根据它们的值。如果一个常量除了它的值之外没有作用，那么就没有必要把它定义为一个常量。
```go
// Bad:
const MAX_PACKET_SIZE = 512
const kMaxBufferSize = 1024
const KMaxUsersPergroup = 500

// Good:
const MaxPacketSize = 512 // 译注：不应命名为FiveHundredTwelve
const (
    ExecuteBit = 1 << iota
    WriteBit
    ReadBit
)
```

首字母缩略词(Initialisms)  
名称中的单词如果是首字母缩略词或缩略语（例如，URL和NATO）应该使用相同的大小写命名。URL应该使用URL或url（如urlPony，或URLPony），而不是Url。这也适用于ID，当它是“标识符”的缩写时使用appID而不是appId。  
- 在有多个首字母缩略词的名字中（例如XMLAPI，它包含XML和API两个首字母缩略词），每个首字母缩略词中的字母都应该具有一致的大小写，但名字中的多个首字母缩略词不需要有相同的大小写。
- 在包含小写字母的首字母缩略词名称中（例如DDoS、iOS、gRPC），首字母应该保持其在缩略词中原有的样子，除非你需要为了导出该名称而改变第一个字母。在这些情况下，整个首字母缩略词中的字母应该采用相同的大小写（例如，ddos, IOS, GRPC）。

Getter命名  
函数和方法名称不应该使用Get或get前缀，除非底层概念使用“get”一词（例如HTTP GET）。我们倾向于直接用那个要Get的事物名词进行命名，例如使用Counts而不是GetCounts。  
如果函数涉及到执行复杂的计算或执行远程调用，可以使用不同的词，如Compute或Fetch来代替Get，以使读者清楚地知道函数调用可能需要时间，并可能阻塞或失败。  

变量命名  
一般的经验法则是，名字的长度应该与它使用的范围大小成正比，与它在该范围内使用的次数成反比。一个在文件范围内创建的变量，其名称可能需要由多个单词组成，而一个在单个内部代码块范围内的变量可能只需要用一个单词命名，甚至只有一两个字符，以保持代码的清晰和避免无关的信息。  
一个在小范围内可能非常清楚的名字（例如，c代表一个计数器）在大范围内可能是不足以胜任的，需要在代码中进一步澄清以提醒读者其目。当一个范围内有许多变量，或者有表示类似的值或概念的变量时，我们可能需要比范围建议的更长的变量名。  

单字母变量名  
单字母变量名可以是一个有用的工具，可以最大限度地减少重复，但这类变量名也可能使代码出现不必要地不透明。把它们的使用限制在其全词含义很明显的情况下，而且如果用全词来代替单字母变量，就会出现重复的情况。

变量名称 vs. 类型  
编译器总是知道变量的类型，而且在大多数情况下，读者也可以通过变量的使用方式清楚地知道它是什么类型。只有当一个变量的值在同一范围内出现两次时，才有必要澄清它的类型。
```go
// var numUsers int
var users int
// var nameString string
var name string
// var primaryProject *Project
var primary *Project
```

外部上下文 vs. 本地名称  
包含周围上下文信息的名字往往不仅没有带来好处，还会产生额外的噪音。包名、方法名、类型名、函数名、导入路径、甚至文件名都可以提供自动限定其中所有名称的上下文信息。
```go
// Bad:
// In package "ads/targeting/revenue/reporting"
type AdsTargetingRevenueReport struct{}
func (p *Project) ProjectName() string

// Good:
// In package "ads/targeting/revenue/reporting"
type Report struct{}
func (p *Project) Name() string
```

#### 关于注释
对注释的惯例（包括注释的内容、使用的风格、如何提供可运行的例子等）进行说明是为了更好地提升阅读公共API文档的体验。  

确保注释即使在狭窄的屏幕上也能从源码中读到。  
当一个注释变得太长时，建议将它拆成多个单行注释。在可能的情况下，争取使注释在80列宽的终端上也能很好阅读，但这并不是一个硬性规定；Go中的注释没有固定的行长限制。例如，标准库经常选择根据标点符号来中断注释，这有时会使个别行更接近60-70个字符。  
```go
// Bad:
// This is a comment paragraph. The length of individual lines doesn't matter in 
// Godoc;
// but the choice of wrapping causes jagged lines on narrow screens or in 
// Critique,
// which can be annoying, especially when in a comment block that will wrap 
// repeatedly.
//
// Don't worry too much about the long URL:
// https://supercalifragilisticexpialidocious.example.com:8080/Animalia/Chordata/Mammalia/Rodentia/Geomyoidea/Geomyidae/

// Good:
// This is a comment paragraph.
// The length of individual lines doesn't matter in Godoc;
// but the choice of wrapping makes it easy to read on narrow screens.
//
// Don't worry too much about the long URL:
// https://supercalifragilisticexpialidocious.example.com:8080/Animalia/Chordata/Mammalia/Rodentia/Geomyoidea/Geomyidae/
//
// Similarly, if you have other information that is made awkward
// by too many line breaks, use your judgment and include a long line
// if it helps rather than hinders.
```

文档注释
所有顶层导出的名字都必须有文档注释，具有不明显的行为或意义的未导出的类型或函数声明也应该如此。这些注释应该是以被描述对象的名称开始的完整句子。冠词（”a”、”an”、”the”）可以放在名字前面，使其读起来更自然。
```go
// Good:
// A Request represents a request to run a command.
type Request struct { ...

// Encode writes the JSON encoding of req to w.
func Encode(w io.Writer, req *Request) { ...

// Good:
// Options configure the group management service.
type Options struct {
    // General setup:
    Name  string
    Group *FooGroup

    // Dependencies:
    DB *sql.DB

    // Customization:
    LargeGroupThreshold int // optional; default: 10
    MinimumMembers      int // optional; default: 2
}

// Good:
// A Server handles serving quotes from the collected works of Shakespeare.
type Server struct {
    // BaseDir points to the base directory under which Shakespeare's works are stored.
    //
    // The directory structure is expected to be the following:
    //   {BaseDir}/manifest.json
    //   {BaseDir}/{name}/{name}-part{number}.txt
    BaseDir string

    WelcomeMessage  string // displayed when user logs in
    ProtocolVersion string // checked against incoming requests
    PageLength      int    // lines per page when printing (optional; default: 20)
}

// Bad:
// 当名称产生不必要的重复时，不要使用命名的返回值参数
func (n *Node) Parent1() (node *Node)
func (n *Node) Parent2() (node *Node, err error)
```

包注释  
包的注释必须紧挨着package子句出现，在注释和包名之间没有空行。  
```go
// Good:
// Package math provides basic constants and mathematical functions.
//
// This package does not guarantee bit-identical results across architectures.
package math
```

#### 关于包
包应分为两组导入：  
- 标准库包
- 其他包（项目和vendor包）
```go
// Good:
package main

import (
    "fmt"
    "hash/adler32"
    "os"

    "github.com/dsnet/compress/flate"
    "golang.org/x/text/encoding"
    "google.golang.org/protobuf/proto"

    foopb "myproj/foo/proto/proto"
    _ "myproj/rpc/protocols/dial"
    _ "myproj/security/auth/authhooks"
)
```

#### 关于错误
错误字符串  
错误字符串不应大写（除非是以导出名称、专有名词或首字母缩略词开始），也不应以标点符号结束。这是因为错误字符串在打印给用户之前，通常出现在其他环境中。  
另一方面，完整显示的消息（日志、测试失败、API响应或其他用户界面）的风格视情况而定，但通常应该以大写开头。
```go
// Bad:
err := fmt.Errorf("Something bad happened.")

// Good:
err := fmt.Errorf("something bad happened")

// Good:
log.Infof("Operation aborted: %v", err)
log.Errorf("Operation aborted: %v", err)
t.Errorf("Op(%q) failed unexpectedly; err=%v", args, err)
```

缩进错误流程  
在继续进行你的代码的其余部分之前，先处理错误。这可以提高代码的可读性，使读者能够迅速找到正常的路径。这个逻辑同样适用于任何测试一个条件是否为终止条件的代码块（例如，return、panic、log.Fatal）。  
如果终止条件没有得到满足，后续运行的代码应该出现在if块之后，而不应该放入缩进的else子句中。  
```go
// Bad:
if err != nil {
    // error handling
} else {
    // normal code that looks abnormal due to indentation
}

// Good:
if err != nil {
    // error handling
    return // or continue, etc.
}
// normal code
```

#### 关于指南
```go
// Bad:
// 使用=而不是:=可以完全改变这一行。
if user, err = db.UserByID(userID); err != nil {
    // ...
}

// Good:
u, err := db.UserByID(userID)
if err != nil {
    return fmt.Errorf("invalid origin user: %s", err)
}
user = u

// Bad:
// 这行代码中间的！很容易错过
leap := (year%4 == 0) && (!(year%100 == 0) || (year%400 == 0))

// Good:
// 格里高利闰年不能仅通过year%4==0来判定。
// 具体请参见https://en.wikipedia.org/wiki/Leap_year#Algorithm.
var (
    leap4   = year%4 == 0
    leap100 = year%100 == 0
    leap400 = year%400 == 0
)
leap := leap4 && (!leap100 || leap400)
```




