### ProtoBuf

protobuf 即 Protocol Buffers，是一种轻便高效的结构化数据存储格式，与语言、平台无关，可扩展可序列化。protobuf 性能和效率大幅度优于 JSON、XML 等其他的结构化数据格式。protobuf 是以二进制方式存储的，占用空间小，但也带来了可读性差的缺点。protobuf 在通信协议和数据存储等领域应用广泛。例如著名的分布式缓存工具 [Memcached](https://memcached.org/) 的 Go 语言版本[groupcache](https://github.com/golang/groupcache) 就使用了 protobuf 作为其 RPC 数据格式。

Protobuf 在 `.proto` 定义需要处理的结构化数据，可以通过 `protoc` 工具，将 `.proto` 文件转换为 C、C++、Golang、Java、Python 等多种语言的代码，兼容性好，易于使用。

**Protobuf 安装**  

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

**使用步骤**  
1. 定义服务：使用 Protocol Buffer 编写服务定义文件，定义服务的接口，包括服务的名称、接口的名称、参数和返回值等。
2. 生成代码：使用 Protocol Buffer 编译器，根据服务定义文件生成服务端和客户端的代码，这些代码可以用于实现服务的客户端和服务端。
3. 实现服务：使用生成的代码，实现服务的客户端和服务端，实现服务的功能。
4. 部署服务：将服务端部署到服务器上，客户端可以通过网络访问服务端，实现服务的调用

**定义消息类型**  
student.proto 定义如下：
```proto
syntax = "proto3";

package demo;

// 输出路径
option go_package=".";

// this is a comment
message Student {
  string name = 1;
  bool male = 2;
  repeated int32 scores = 3;
}
```

