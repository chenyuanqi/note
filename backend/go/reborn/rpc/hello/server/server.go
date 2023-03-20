package main

import (
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {
	_ = rpc.RegisterName("HelloService", &HelloService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic("监听端口失败")
	}
	conn, err := listener.Accept()
	if err != nil {
		panic("建立链接失败")
	}
	rpc.ServeConn(conn)
}
