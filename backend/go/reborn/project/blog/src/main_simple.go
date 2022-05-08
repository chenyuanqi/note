package main

import (
	"fmt"
	"net/http"
	"strings"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 欢迎来到主页</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:vikey@example.com\">vikey@example.com</a>")
}

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprint(w, "<h1>Hello, 这里是主页</h1>")
// 	// fmt.Fprint(w, "当前请求路径为："+r.URL.Path)

// 	// 设置响应头
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	if r.URL.Path == "/" {
// 		fmt.Fprint(w, "<h1>Hello, 这里是主页</h1>")
// 	} else if r.URL.Path == "/about" {
// 		fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
// 			"<a href=\"mailto:vikey@example.com\">vikey@example.com</a>")
// 	} else {
// 		// 设置状态码
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
// 			"<p>如有疑惑，请联系我们。</p>")
// 	}
// }

func main() {
	// 项目自动重载，安装 cosmtrek/air：https://learnku.com/courses/go-basic/1.17/automatic-overloading/11490
	// http.HandleFunc("/", handlerFunc)
	// http.ListenAndServe(":8000", nil)

	// 使用路由
	router := http.NewServeMux()

	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about", aboutHandler)
	// 文章详情
	router.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.SplitN(r.URL.Path, "/", 3)[2]
		fmt.Fprint(w, "文章 ID："+id)
	})
	// 列表 or 创建
	router.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprint(w, "访问文章列表~")
		case "POST":
			fmt.Fprint(w, "创建新的文章")
		}
	})

	http.ListenAndServe(":8000", router)
}
