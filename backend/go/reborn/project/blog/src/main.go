package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// 语法错误：non-declaration statement outside function body 函数外无法使用变量赋值语句
// router := mux.NewRouter()
// 包级别的变量声明时不能使用 := 语法，修改为带关键词 var 的变量声明即可
var router = mux.NewRouter()

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hello, 欢迎来到主页</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:vikey@example.com\">vikey@example.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// 获取路径参数
	id := vars["id"]
	fmt.Fprint(w, "文章 ID："+id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "访问文章列表~")
}

func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "创建博文表单")

	// 多行字符串使用 ``
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <title>创建文章 —— 我的技术博客</title>
</head>
<body>
    <form action="%s" method="post">
        <p><input type="text" name="title"></p>
        <p><textarea name="body" cols="30" rows="10"></textarea></p>
        <p><button type="submit">提交</button></p>
    </form>
</body>
</html>
`
	// 获取创建博文的链接，使用命名路由的好处是为 URL 修改提供了灵活性
	storeURL, _ := router.Get("articles.store").URL()
	fmt.Fprintf(w, html, storeURL)
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "创建新的文章~")

	// 从请求中解析请求参数 r.ParseForm()
	if err := r.ParseForm(); err != nil {
		// 解析错误，这里应该有错误处理
		fmt.Fprint(w, "请提供正确的数据！")
		return
	}

	title := r.PostForm.Get("title")

	// PostForm：存储了 post、put 参数，在使用之前需要调用 ParseForm 方法
	fmt.Fprintf(w, "POST PostForm: %v <br>", r.PostForm)
	// Form：存储了 post、put 和 get 参数，在使用之前需要调用 ParseForm 方法
	fmt.Fprintf(w, "POST Form: %v <br>", r.Form)
	// 从 PostForm 中读取 title
	fmt.Fprintf(w, "title 的值为: %v", title)
	fmt.Fprintf(w, "r.PostForm 中 title 的值为: %v <br>", r.PostFormValue("title"))
	// 从 Form 中读取 title
	fmt.Fprintf(w, "r.Form 中 title 的值为: %v <br>", r.FormValue("title"))
	// POST PostForm: map[body:[content ] title:[title]]
	// POST Form: map[body:[content ] title:[title]]
	// title 的值为: titler.PostForm 中 title 的值为: title
	// r.Form 中 title 的值为: title
}

func forceHTMLMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 2. 继续处理请求
		h.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Remove trailing slash
		// 1. 除首页以外，移除所有请求路径后面的斜杆
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}

		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}

func main() {
	// router := mux.NewRouter()
	// .StrictSlash(true) 处理 url 尾部 /
	// router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")

	// 查看博文
	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	// 创建博文
	router.HandleFunc("/articles/create", articlesCreateHandler).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")

	// 自定义 404 页面
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// 中间件：强制内容类型为 HTML
	router.Use(forceHTMLMiddleware)

	// 通过命名路由获取 URL 示例
	// homeURL, _ := router.Get("home").URL()
	// fmt.Println("homeURL: ", homeURL)
	// articleURL, _ := router.Get("articles.show").URL("id", "23")
	// fmt.Println("articleURL: ", articleURL)

	http.ListenAndServe(":8000", removeTrailingSlash(router))
}
