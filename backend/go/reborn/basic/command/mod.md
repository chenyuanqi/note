
### Go mod
一些环境变量的说明。  

```bash
go mod init	# 生成 go.mod 文件
go mod download	# 下载 go.mod 文件中指明的所有依赖
go mod tidy	# 整理现有的依赖
go mod graph # 查看现有的依赖结构
go mod edit	# 编辑 go.mod 文件
go mod vendor # 导出项目所有的依赖到vendor目录
go mod verify # 校验一个模块是否被篡改过
go mod why # 查看为什么需要依赖某模块
```
GO111MODULE  
Go 语言提供了 GO111MODULE 这个环境变量来作为 Go modules 的开关，其允许设置以下参数：  
● auto：只要项目包含了 go.mod 文件的话启用 Go modules，目前在 Go1.11 至 Go1.14 中仍然是默认值。  
● on：启用 Go modules，推荐设置，将会是未来版本中的默认值。  
● off：禁用 Go modules，不推荐设置。  
```bash
go env -w GO111MODULE=on
```

GOPROXY  
这个环境变量主要是用于设置 Go 模块代理（Go module proxy）,其作用是用于使 Go 在后续拉取模块版本时直接通过镜像站点来快速拉取。  
GOPROXY 的默认值是：https://proxy.golang.org,direct  
```bash
go env -w GOPROXY=https://goproxy.cn,direct
```
