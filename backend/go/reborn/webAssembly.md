
### WebAssembly
WebAssembly 是一种新的编码方式，可以在现代的网络浏览器中运行 － 它是一种低级的类汇编语言，具有紧凑的二进制格式，可以接近原生的性能运行，并为诸如 C / C++ 等语言提供一个编译目标，以便它们可以在 Web 上运行。它也被设计为可以与 JavaScript 共存，允许两者一起工作。

Go 语言在 1.11 版本(2018年8月) 加入了对 WebAssembly (Wasm) 的原生支持，使用 Go 语言开发 WebAssembly 相关的应用变得更加地简单。Go 语言的内建支持是 Go 语言进军前端的一个重要的里程碑。在这之前，如果想使用 Go 语言开发前端，需要使用 GopherJS，GopherJS 是一个编译器，可以将 Go 语言转换成可以在浏览器中运行的 JavaScript 代码。新版本的 Go 则直接将 Go 代码编译为 wasm 二进制文件，而不再需要转为 JavaScript 代码。更巧的是，实现 GopherJS 和在 Go 语言中内建支持 WebAssembly 的是同一拨人。

Go 语言实现的函数可以直接导出供 JavaScript 代码调用，同时，Go 语言内置了 syscall/js 包，可以在 Go 语言中直接调用 JavaScript 函数，包括对 DOM 树的操作。

1、Hello World  
```go
package main

import "syscall/js"

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("Hello World!")
}
```

2、将 main.go 编译为 static/main.wasm  
```bash
GOOS=js GOARCH=wasm go build -o static/main.wasm
```

3、拷贝 wasm_exec.js (JavaScript 支持文件，加载 wasm 文件时需要) 到 static 文件夹  
```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" static
```

4、创建 index.html，引用 static/main.wasm 和 static/wasm_exec.js  
```html
<html>
<script src="static/wasm_exec.js"></script>
<script>
	const go = new Go();
	WebAssembly.instantiateStreaming(fetch("static/main.wasm"), go.importObject)
		.then((result) => go.run(result.instance));
</script>

</html>
```

5、使用 goexec 启动 Web 服务  
> 如果没有安装 goexec，可用 go get -u github.com/shurcooL/goexec 安装，需要将 $GOBIN 或 $GOPATH/bin 加入环境变量  

当前的目录结构如下：  
```
   |--static/  
      |--wasm_exec.js  
      |--main.wasm  
   |--main.go  
   |--index.html  
```

```bash
goexec 'http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))'
```

为了避免每次编译都需要输入繁琐的命令，可将这个过程写在 Makefile 中。
```
all: static/main.wasm static/wasm_exec.js
	goexec 'http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))'

static/wasm_exec.js:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" static

static/main.wasm : main.go
	GO111MODULE=auto GOOS=js GOARCH=wasm go build -o static/main.wasm .
```
这样一个敲一下 make 就够了。