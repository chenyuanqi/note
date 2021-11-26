
### Go Module
Go 源码需要先编译，再分发和运行。如果是单 Go 源文件的情况，我们可以直接使用 go build 命令 +Go 源文件名的方式编译。不过，对于复杂的 Go 项目，我们需要在 Go Module 的帮助下完成项目的构建。  

```bash
# 配置国内镜像代理（使用阿里云镜像）
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
# 验证
go env | grep GOPROXY

# 单文件
go build main.go
# 开发时
go run main.go

# 复杂项目
go mod init
# 自动添加依赖相关版本
go mod tidy
go build main.go
```
