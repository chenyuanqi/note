
### go-zero 环境准备
```bash
# 开启 GO111MODULE
go env GO111MODULE
go env -w GO111MODULE="on"
# 设置 GOPROXY
go env -w GOPROXY=https://goproxy.cn
# 设置 GOMODCACHE（如果目录不为空或者/dev/null，请跳过）
go env GOMODCACHE
go env -w GOMODCACHE=$GOPATH/pkg/mod

# 安装 goctl
go install github.com/zeromicro/go-zero/tools/goctl@latest
goctl -v

# protoc & protoc-gen-go 安装
goctl env check -i -f --verbose
# brew install protobuf protoc-gen-go protoc-gen-go-grpc
protoc --version
```

### 快速入门
＊*单体服务**  
