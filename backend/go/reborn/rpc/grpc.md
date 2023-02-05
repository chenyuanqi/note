
### rpc
rpc 框架中最核心的几个要点：
- 序列化
- 通信协议
- IDL（接口描述语言）

这些在 gRPC 中分别对应的是：
- 基于 Protocol Buffer 序列化协议，性能高效。
- 基于 HTTP/2 标准协议开发，自带 stream、多路复用等特性；同时由于是标准协议，第三方工具的兼容性会更好（比如负载均衡、监控等）
- 编写一份 .proto 接口文件，便可生成常用语言代码。

**HTTP/2**
由于 HTTP/1.1 是一个文本协议，对人类非常友好，相反的对机器性能就比较低。
需要反复对文本进行解析，效率自然就低了；要对机器更友好就得采用二进制，HTTP/2 自然做到了。
除此之外还有其他优点：
- 多路复用：可以并行的收发消息，互不影响
- HPACK 节省 header 空间，避免 HTTP1.1 对相同的 header 反复发送。

**Protocol**  
gRPC 采用的是 Protocol 序列化，发布时间比 gRPC 早一些，所以也不仅只用于 gRPC，任何需要序列化 IO 操作的场景都可以使用它。它会更加的省空间、高性能。  
```proto
package order.v1;

service OrderService{

  rpc Create(OrderApiCreate) returns (Order) {}

  rpc Close(CloseApiCreate) returns (Order) {}

  // 服务端推送
  rpc ServerStream(OrderApiCreate) returns (stream Order) {}

  // 客户端推送
  rpc ClientStream(stream OrderApiCreate) returns (Order) {}
  
  // 双向推送
  rpc BdStream(stream OrderApiCreate) returns (stream Order) {}
}

message OrderApiCreate{
  int64 order_id = 1;
  repeated int64 user_id = 2;
  string remark = 3;
  repeated int32 reason_id = 4;
}
```

使用起来也是非常简单的，只需要定义自己的 .proto 文件，便可用命令行工具生成对应语言的 SDK。[更多参考](https://grpc.io/docs/languages/go/generated-code/)  

```bash
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
test.proto
```

生成代码之后编写服务端就非常简单了，只需要实现生成的接口即可。
```go
func (o *Order) Create(ctx context.Context, in *v1.OrderApiCreate) (*v1.Order, error) {
	// 获取 metadata
	md, ok := metadata.FromIncomin gContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "failed to get metadata")
	}
	fmt.Println(md)
	fmt.Println(in.OrderId)
	return &v1.Order{
		OrderId: in.OrderId,
		Reason:  nil,
	}, nil
}
```

客户端也非常简单，只需要依赖服务端代码，创建一个 connection 然后就和调用本地方法一样了。这是经典的 unary(一元)调用，类似于 http 的请求响应模式，一个请求对应一次响应。  
