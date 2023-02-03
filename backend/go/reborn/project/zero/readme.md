
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
```bash
go mod init zero
goctl api new greet
go mod tidy
# 编写逻辑
vim greet/internal/logic/greetlogic.go
```
greetlogic.go 修改 api 逻辑，
```go
func (l *GreetLogic) Greet(req *types.Request) (*types.Response, error) {
    return &types.Response{
        Message: "Hello go-zero",
    }, nil
}
```
启动服务，
```bash
cd greet
go run greet.go -f etc/greet-api.yaml
curl -i -X GET http://localhost:8888/from/you
```

**微服务**  
假设我们在开发一个商城项目，而开发者小明负责用户模块(user)和订单模块(order)的开发，我们姑且将这两个模块拆分成两个微服务。

创建 user rpc 服务。
```bash
mkdir -p mall/user/rpc
vim mall/user/rpc/user.proto
```
user.proto 代码如下：
```
syntax = "proto3";

package user;
  
// protoc-gen-go 版本大于1.4.0, proto文件需要加上go_package,否则无法生成
option go_package = "./user";

message IdRequest {
    string id = 1;
}
  
message UserResponse {
    // 用户id
    string id = 1;
    // 用户名称
    string name = 2;
    // 用户性别
    string gender = 3;
}
  
service User {
    rpc getUser(IdRequest) returns(UserResponse);
}
```
生成代码，
```bash
goctl env check -i -f
cd mall/user/rpc
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
# 填充业务逻辑
vim internal/logic/getuserlogic.go
```
getuserlogic.go 代码如下：
```go
package logic

import (
    "context"

    "zero/mall/user/rpc/internal/svc"
    "zero/mall/user/rpc/types/user"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
    return &GetUserLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
    return &user.UserResponse{
            Id:   "1",
            Name: "test",
    }, nil
}
```

创建 order api 服务。
```bash
# 回到 go-zero-demo/mall 目录
mkdir -p order/api && cd order/api
vim order.api
```
order.api 代码如下：
```
type(
    OrderReq {
        Id string `path:"id"`
    }
  
    OrderReply {
        Id string `json:"id"`
        Name string `json:"name"`
    }
)
service order {
    @handler getOrder
    get /api/order/get/:id (OrderReq) returns (OrderReply)
}
```
生成 order 服务，
```bash
goctl api go -api order.api -dir .
# 添加 user rpc 配置
vim internal/config/config.go
```
config.go 代码如下：
```go
package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
}
```
添加 yaml 配置，
```bash
vim etc/order.yaml 
```
order.yaml 代码如下：
```yaml
Name: order
Host: 0.0.0.0
Port: 8888
UserRpc:
  Etcd:
    Hosts:
    - 127.0.0.1:2379
    Key: user.rpc
```
完善服务依赖，
```bash
vim internal/svc/servicecontext.go
```
servicecontext.go 代码如下：
```go
package svc

import (
	"zero/mall/order/api/internal/config"
	"zero/mall/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
```
添加 order 演示逻辑，
```bash
vim internal/logic/getorderlogic.go
```
getorderlogic.go 代码如下：
```go
package logic

import (
    "context"
    "errors"

    "zero/mall/order/api/internal/svc"
    "zero/mall/order/api/internal/types"
    "zero/mall/user/rpc/types/user"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetOrderLogic {
    return GetOrderLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (*types.OrderReply, error) {
    user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
        Id: "1",
    })
    if err != nil {
        return nil, err
    }

    if user.Name != "test" {
        return nil, errors.New("用户不存在")
    }

    return &types.OrderReply{
        Id:   req.Id,
        Name: "test order",
    }, nil
}
```

启动服务并验证，
```bash
# 安装 etcd
wget https://github.com/etcd-io/etcd/releases/download/v3.5.7/etcd-v3.5.7-darwin-amd64.zip
unzip etcd-v3.5.7-darwin-amd64.zip
cd etcd-v3.5.7-darwin-amd64
sudo cp etcd /usr/local/bin
# 启动 etcd
etcd
# 在 go-zero-demo 目录下
go mod tidy

# 在 mall/user/rpc 目录启动 user rpc
go run user.go -f etc/user.yaml
# 在 mall/order/api 目录启动 order api
go run order.go -f etc/order.yaml
# 访问 order api
curl -i -X GET http://localhost:8888/api/order/get/1
```

### Go Protobuf
从 [Protobuf Release](https://github.com/protocolbuffers/protobuf/releases) 下载符合系统的安装包。
```bash
wget https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-osx-x86_64.zip
brew install p7zip
sudo 7z x protoc-21.12-osx-x86_64.zip -o/usr/local
sudo chmod 755 /usr/local/bin/protoc
protoc --version
```

在 Golang 中使用 protobuf，还需要安装 protoc-gen-go，这个工具用来将 .proto 文件转换为 Golang 代码。
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

在 google.golang.org/protobuf 中，protoc-gen-go 纯粹用来生成 pb 序列化相关的文件，不再承载 gRPC 代码生成功能。  
生成 gRPC 相关代码需要安装 grpc-go 相关的插件 protoc-gen-go-grpc。
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# 执行 code gen 命令
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    routeguide/route_guide.proto
```
