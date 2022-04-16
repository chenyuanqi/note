
### Go 系统常用命令
```bash
# 编译
go build
# 编译指定文件
go build main.go
#清除编译文件
go clean -i -n

# 编译源码，并且直接执行源码的 main() 函数，不会在当前目录留下可执行文件
go run main.go

# 格式化
go fm

# 将编译的中间文件放在 GOPATH 的 pkg 目录下，以及固定地将编译结果放在 GOPATH 的 bin 目录下
go install

# 一键获取代码、编译并安装
go get github.com/xxx

# 单元测试
go test
```
