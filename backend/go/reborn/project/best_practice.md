

### 最佳实践
Go 箴言
- 不要通过共享内存进行通信，通过通信共享内存
- 并发不是并行
- 管道用于协调；互斥量（锁）用于同步
- 接口越大，抽象就越弱
- 利用好零值
- 空接口 interface{} 没有任何类型约束
- Gofmt 的风格不是人们最喜欢的，但 gofmt 是每个人的最爱
- 允许一点点重复比引入一点点依赖更好
- 系统调用必须始终使用构建标记进行保护
- 必须始终使用构建标记保护 Cgo
- Cgo 不是 Go
- 使用标准库的 unsafe 包，不能保证能如期运行
- 清晰比聪明更好
- 反射永远不清晰
- 错误是值
- 不要只检查错误，还要优雅地处理它们
- 设计架构，命名组件，（文档）记录细节
- 文档是供用户使用的
- 不要（在生产环境）使用 panic()

Go 之禅
- 每个 package 实现单一的目的
- 显式处理错误
- 尽早返回，而不是使用深嵌套
- 让调用者处理并发（带来的问题）
- 在启动一个 goroutine 时，需要知道何时它会停止
- 避免 package 级别的状态
- 简单很重要
- 编写测试以锁定 package API 的行为
- 如果你觉得慢，先编写 benchmark 来证明
- 适度是一种美德
- 可维护性

一、代码规范
1. 命名规范
- 变量、函数、方法使用驼峰命名法。
- 全局常量使用全大写字母加下划线的方式命名。
- 缩写尽量避免使用，如果必须使用则按照大驼峰命名法。

2. 代码格式
- 使用 gofmt 工具格式化代码。
- 使用 4 个空格代替制表符缩进。
- 使用 UTF-8 编码。

3. 注释规范
- 使用 // 或 /* */ 注释。
- 注释应该清晰简洁，避免过多无用的注释。
- 函数、方法需要编写文档注释。

二、错误处理
1. 错误处理
- 函数、方法应该返回错误信息。
- 使用 errors 包或自定义错误类型处理错误。
- 在处理错误时，应该记录日志并返回错误信息。

2. 错误日志
- 使用 log 包记录错误日志。
- 日志应该包含错误信息、堆栈信息和错误发生的时间等关键信息。

三、并发处理
1. 并发模型
- 使用 goroutine 和 channel 实现并发模型。
- 避免使用共享变量，使用 channel 传递数据。
- 使用 sync 包提供的锁实现互斥访问。

2. 并发安全
- 避免出现竞态条件。
- 使用 sync 包提供的锁实现互斥访问。
- 使用 atomic 包提供的原子操作实现线程安全。

四、性能优化
1. 内存管理
- 避免过多的内存分配和释放。
- 使用 sync.Pool 重用对象池。
- 使用标准库提供的内存分配和释放函数。

2. 并发优化
- 避免过多的 goroutine 创建和销毁。
- 使用 sync.WaitGroup 等待 goroutine 完成任务。
- 使用 sync.Mutex 避免竞态条件。

3. 数据结构优化
- 使用 map 代替 slice 实现索引访问。
- 使用 bytes.Buffer 代替字符串拼接。
- 使用 sync.Map 提供并发安全的 map。

五、其他  
- 30 * time.Second 比 time.Duration(30) * time.Second 更好
- 按类型分组 const 声明，按逻辑和/或类型分组 var
```go
// BAD
const (
    foo = 1
    bar = 2
    message = "warn message"
)

// MOSTLY BAD
const foo = 1
const bar = 2
const message = "warn message"

// GOOD
const (
    foo = 1
    bar = 2
)

const message = "warn message"
```
- 多行字符串用反引号(`)  
- 用 _ 来跳过不用的参数
- 用 range 循环来进行数组或 slice 的迭代
- 如果你要比较时间戳，请使用 time.Before 或 time.After
- 在 Go 里面要小心使用 range: for i := range a and for i, v := range &a ，都不是 a 的副本；但是 for i, v := range a 里面的就是 a 的副本
- 要 marshal 任意的 JSON， 你可以 marshal 为 map[string]interface{}{}

