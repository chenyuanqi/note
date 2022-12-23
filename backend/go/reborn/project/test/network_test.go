package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

func TestConn(t *testing.T) {
	// 监听一个未被占用的端口，并返回 Listener
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler)
	// 启动 http 服务
	go http.Serve(ln, nil)

	// 发起一个 Get 请求，检查返回值是否正确
	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handleError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "hello world" {
		t.Fatal("expected hello world, but got", string(body))
	}
}

// 针对 http 开发的场景，使用标准库 net/http/httptest 进行测试更为高效
func TestConn2(t *testing.T) {
	req := httptest.NewRequest("GET", "http://127.0.0.1/hello", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	bytes, _ := ioutil.ReadAll(w.Result().Body)

	if string(bytes) != "hello world" {
		t.Fatal("expected hello world, but got", string(bytes))
	}
}
