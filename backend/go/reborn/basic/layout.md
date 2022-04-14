
### Go 编译代码
简单的编译单个文件是这样的，
```bash
go build main.go
# 直接运行 Go 源码文件
go run main.go
```

复杂项目下 Go 程序的编译是怎样的呢？
```bash
# go mod init 初始化 go.mod 文件，一般建新项目时才会用这个命令
go mod init project
# 获取源码包
go get github.com/ethereum/go-ethereum
go get github.com/ethereum/go-ethereum@v1.10.1
# 自动添加依赖，go mod tidy 整理现有依赖，修改 go.mod 文件后执行会更新依赖
go mod tidy
# go mod graph 查看现有的依赖结构
# go mod vendor 导出项目所有依赖到 vendor 目录 （不建议使用）
go build main.go
```

### Go 布局
Go 可执行程序项目的典型结构布局是怎样的呢？
```
demo
├── cmd/
│   ├── app1/
│   │   └── main.go
│   └── app2/
│       └── main.go
├── go.mod
├── go.sum
├── internal/
│   ├── pkga/
│   │   └── pkg_a.go
│   └── pkgb/
│       └── pkg_b.go
├── pkg1/
│   └── pkg1.go
├── pkg2/
│   └── pkg2.go
└── vendor/
```
> cmd 目录就是存放项目要编译构建的可执行文件对应的 main 包的源文件  
> pkgN 目录，这是一个存放项目自身要使用、同样也是可执行文件对应 main 包所要依赖的库文件，同时这些目录下的包还可以被外部项目引用  
> go.mod 和 go.sum，它们是 Go 语言包依赖管理使用的配置文件  
> vendor 是 Go 1.5 版本引入的用于在项目本地缓存特定版本依赖包的机制，在 Go Modules 机制引入前，基于 vendor 可以实现可重现构建，保证基于同一源码构建出的可执行程序是等价的。  
> vendor 目录视为一个可选目录。原因在于，Go Module 本身就支持可再现构建，而无需使用 vendor。 当然 Go Module 机制也保留了 vendor 目录（通过 go mod vendor 可以生成 vendor 下的依赖包，通过 go build -mod=vendor 可以实现基于 vendor 的构建）。一般我们仅保留项目根目录下的 vendor 目录，否则会造成不必要的依赖选择的复杂性。  
> Go 库项目的初衷是为了对外部（开源或组织内部公开）暴露 API，对于仅限项目内部使用而不想暴露到外部的包，可以放在项目顶层的 internal 目录下面。当然 internal 也可以有多个并存在于项目结构中的任一目录层级中，关键是项目结构设计人员要明确各级 internal 包的应用层次和范围。  

公司的 Go 项目布局是怎样的呢？  
```
demo
├── README.md
├── bin
│   └── main.go
├── configs
│   ├── dev
│   ├── prod
│   └── test
├── controllers
│   ├── batch.go
├── docs
├── go.mod
├── go.sum
├── services
│   ├── batch
└── vendor
```

### Go 编码规范
[常用编码规范](https://github.com/kevinyan815/gocookbook/issues/61)  

**注释**  
单行注释也称为行注释，格式为以双斜杠（“//”）开头的一行，可以添加在代码的任何位置。  
多行注释也称为块注释，格式为以“/”开头，以“/”结束的一行或多行。  
```go
// 单行注释
/*
多行注释第一行

多行注释第三行
多行注释第四行
多行注释第五行
 */
```


