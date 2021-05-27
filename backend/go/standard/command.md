
### Go 命令行
Go 国内加速镜像的使用。  
```bash
# 启用 Go Modules 功能
go env -w GO111MODULE=on

# 配置 GOPROXY 环境变量，以下三选一
# 1. 七牛 CDN
go env -w  GOPROXY=https://goproxy.cn,direct
# 2. 阿里云
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
# 3. 官方
go env -w  GOPROXY=https://goproxy.io,direct

# 检查配置是否生效
go env | grep GOPROXY


# Windows
# 启用 Go Modules 功能
$env:GO111MODULE="on"
# 配置 GOPROXY 环境变量，以下三选一
# 1. 七牛 CDN
$env:GOPROXY="https://goproxy.cn,direct"
# 2. 阿里云
$env:GOPROXY="https://mirrors.aliyun.com/goproxy/,direct"
# 3. 官方
$env:GOPROXY="https://goproxy.io,direct"
# 测试一下
time go get golang.org/x/tour
```

常用命令行。  
```bash
# go get 可以下载指定的软件包及其依赖，同时会像 go install 命令一样安装下载好的软件包
go get github.com/golang/lint/golint
# -u 强制软件包版本为最新版本
# -d 可以跳过编译和安装的步骤，只将存储库克隆到 GOPATH 工作区

# 清空模块缓存
go clean --modcache

# go build 命令允许您在自己的平台上为 Go 支持的任何目标平台构建可执行文件
GOOS=windows GOARCH=amd64 go build github.com/mholt/caddy/caddy
# -ldflags 选项：优化 Go 二进制大小，在构建过程中设置变量值
go build -ldflags="-X main.Version 1.0.0"

# go test 测试
go test
# -race 运行 Go race detector
# -run 来过滤要由 regex 和 -run 标志运行的测试: go test -run=FunctionName
# -bench 去运行基准测试。 - -cpuprofile cpu.out 退出前将 CPU 配置文件写入指定的文件
# -memprofile mem.out 在所有测试通过后，将内存配置文件写入文件
# -v 打印测试名称、状态 (失败或通过)、运行测试需要多少时间、测试用例中的任何日志等等
# -cover 度量在运行一组测试时执行的代码行的百分比

# go list 列出了由导入路径命名的包，每行一个
go list

# go env 打印 Go 环境变量信息
go env

# go fmt 根据 Go 的标准重新格式化你的代码
go fmt

# go vet 检查 Go 源代码并报告可疑的构造，如参数与格式字符串不一致的 Printf 调用
go vet

# go generate 命令是在 Go 1.4 版本加入的，是「在编译前自动运行生成源代码的工具」
# Go 工具会扫描与当前包相关的文件，寻找带有表单 //go:generate command arguments「魔力注释」的行
go generate

# go doc/godoc  Go 文档没有任何额外的格式化规则。所有内容都是纯文本
go doc json.Encoder
```

### 非标准 Go 工具

**golint**  
```bash
go get -u github.com/golang/lint/golint
```

*errcheck*  
```bash
go get github.com/kisielk/errcheck
```
