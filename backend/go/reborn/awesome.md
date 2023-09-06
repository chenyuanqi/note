
### Go 官方
[官网](https://go.dev/)  
[Google Go 镜像下载](https://golang.google.cn/dl/)  
[Go 镜像下载](https://gomirrors.org/)  

Go 语言每半年发布一个主要版本，一般在每年的 2 月和 8 月。Go 语言的维护周期终止在落后两个主要版本时，如：Go 1.5 会在 Go 1.7 发布时停止支持；Go 1.6 会在 Go 1.8 发布时停止支持。  
[版本迭代](https://golang.org/doc/devel/release.html)  

### Go 之禅
> 每个 package 实现单一的目的  
> 显式处理错误  
> 尽早返回，而不是使用深嵌套  
> 让调用者处理并发（带来的问题）  
> 在启动一个 goroutine 时，需要知道何时它会停止  
> 避免 package 级别的状态  
> 简单很重要  
> 编写测试以锁定 package API 的行为  
> 如果你觉得慢，先编写 benchmark 来证明  
> 适度是一种美德  
> 可维护性  

**Go 箴言**  
> 不要通过共享内存进行通信，通过通信共享内存  
> 并发不是并行  
> 管道用于协调；互斥量（锁）用于同步  
> 接口越大，抽象就越弱  
> 利用好零值  
> 空接口 interface{} 没有任何类型约束  
> Gofmt 的风格不是人们最喜欢的，但 gofmt 是每个人的最爱  
> 允许一点点重复比引入一点点依赖更好  
> 系统调用必须始终使用构建标记进行保护  
> 必须始终使用构建标记保护 Cgo  
> Cgo 不是 Go  
> 使用标准库的 unsafe 包，不能保证能如期运行  
> 清晰比聪明更好  
> 反射永远不清晰  
> 错误是值  
> 不要只检查错误，还要优雅地处理它们  
> 设计架构，命名组件，（文档）记录细节  
> 文档是供用户使用的  
> 不要（在生产环境）使用 panic()  

### Go 入门
[Go awesome](https://github.com/avelino/awesome-go)  
[Go 中文文档](https://go-zh.org/doc/)  
[Go 模块代理](https://goproxy.cn/)  

[零基础通关 Go 语言](https://juejin.cn/book/7039174186522116131)  
[Go 实战：web 入门](https://learnku.com/courses/go-basic/1.17)  

[C 语言中文网](http://c.biancheng.net/)  
[Go 开发手记](https://github.com/kevinyan815/gocookbook)  
[Go 小抄 - cheat sheets](https://yourbasic.org/golang/#cheat-sheets)  

[Go 常见问题](https://learnku.com/go/wikis/38175)  
[50 个 Go 新手易犯的错误](https://learnku.com/go/wikis/49781)  

[Go 编码指南](https://learnku.com/go/wikis/38174)  
[Go 编码规范](https://learnku.com/go/wikis/38426)  
[Go 最佳实践](https://learnku.com/go/wikis/38430)  
[Go 语言圣经](https://yar999.gitbook.io/gopl-zh/)  
[Go 标准库](http://books.studygolang.com/The-Golang-Standard-Library-by-Example/)  
[地鼠文档](https://www.topgoer.cn/)  

### Go 进阶
[Go 语言编程之旅](https://golang2.eddycjy.com/)  

[Gopher 日报](https://github.com/bigwhite/gopherdaily)  

### Go 工具
[Aorm 数据库操作](https://github.com/tangpanqing/aorm)  
[Gocache 缓存管理](https://github.com/eko/gocache)  
[Goredis redis客户端](https://github.com/redis/go-redis)  

[Viper 配置管理](http://github.com/spf13/viper)  
[Zap 日志记录](https://github.com/uber-go/zap)  
[Imaging 图片处理](https://github.com/disintegration/imaging)  

[Lancet 柳叶刀工具][https://github.com/duke-git/lancet]  
[Cast 类型转换](https://github.com/spf13/cast)  
[Carbon 日期处理](https://github.com/golang-module/carbon)  
[Decimal 精度浮点数](https://github.com/shopspring/decimal)  
[ZeroWidth 去除零宽字符](https://github.com/trubitsyn/go-zero-width)  

[Govalidator 参数校验](https://github.com/asaskevich/govalidator)  
[Grequest http请求](https://github.com/levigross/grequests)  
[Fast http高性能服务](https://github.com/valyala/fasthttp)  

[Cobra 命令行交互](https://github.com/spf13/cobra)  
[Ansi 彩色字符](https://github.com/mgutz/ansi)  

[Goconvey 测试工具](https://github.com/smartystreets/goconvey)  
[Uuid uuid生成](https://github.com/pborman/uuid)  
[Faker 数据生产](https://github.com/go-faker/faker)  

[gzip gin中间件-gzip](https://github.com/gin-contrib/gzip)  
[cors gin中间件-cors](https://github.com/gin-contrib/cors)  
[jwt gin中间件-jwt](https://github.com/appleboy/gin-jwt)  
[session gin中间件-session](https://github.com/gin-contrib/sessions)  
[timeout gin中间件-timeout](https://github.com/gin-contrib/timeout)  
