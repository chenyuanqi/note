package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql" // 此驱动会自行初始化（利用 init() 函数）并注册自己到 Golang 的 database/sql 上下文中
	"github.com/gorilla/mux"
)

// 语法错误：non-declaration statement outside function body 函数外无法使用变量赋值语句
// router := mux.NewRouter()
// 包级别的变量声明时不能使用 := 语法，修改为带关键词 var 的变量声明即可
var router = mux.NewRouter()

// 连接池对象：sql.DB 结构体是 database/sql 包封装的一个数据库操作对象，包含了操作数据库的基本方法。声明为包级别的变量，方便各个函数中访问
var db *sql.DB

func initDB() {

	var err error
	// 准备生成 DNS 信息（DSN 全称为 Data Source Name，表示 数据源信息，用于定义如何连接数据库）
	config := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "demo",
		AllowNativePasswords: true,
	}

	// 准备数据库连接池，config.FormatDSN() 用来生成 DSN 信息，返回一个 *sql.DB 结构体实例
	db, err = sql.Open("mysql", config.FormatDSN()) // root:root@tcp(127.0.0.1:3306)/demo?checkConnLiveness=false&maxAllowedPacket=0
	checkError(err)

	// 设置最大连接数（参考数据库 show variables like 'max_connections';）
	db.SetMaxOpenConns(100)
	// 设置最大空闲连接数（<= 0 表示不设置空闲连接数，默认为 2）
	db.SetMaxIdleConns(25)
	// 设置每个链接的过期时间（过期会自动关闭链接），设置的值不应该超过 MySQL 的 wait_timeout 设置项（默认情况下是 8 个小时）
	db.SetConnMaxLifetime(5 * time.Minute)

	// 尝试连接，失败会报错
	err = db.Ping() // 检测连接状态
	checkError(err)
}

func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    content longtext COLLATE utf8mb4_unicode_ci
); `
	// Exec() 用于执行没有返回结果集的 SQL 语句，如 CREATE TABLE, INSERT, UPDATE, DELETE 等
	_, err := db.Exec(createArticlesSQL)
	// Exec() 方法的第一个返回值为一个实现了 sql.Result 接口的类型
	// type Result interface {
	// LastInsertId() (int64, error)    // 使用 INSERT 向数据插入记录，数据表有自增 id 时，该函数有返回值
	// RowsAffected() (int64, error)    // 表示影响的数据表行数
	// }
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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

// ArticlesFormData 创建博文表单数据，给模板文件传输变量
type ArticlesFormData struct {
	Title, Content string
	URL            *url.URL
	Errors         map[string]string
}

func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "创建博文表单")

	// 多行字符串使用 ``
	/* html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	    <title>创建文章 —— 我的技术博客</title>
	</head>
	<body>
	    <form action="%s" method="post">
	        <p><input type="text" name="title"></p>
	        <p><textarea name="content" cols="30" rows="10"></textarea></p>
	        <p><button type="submit">提交</button></p>
	    </form>
	</body>
	</html>
	` */
	// 获取创建博文的链接，使用命名路由的好处是为 URL 修改提供了灵活性
	storeURL, _ := router.Get("articles.store").URL()
	// fmt.Fprintf(w, html, storeURL)

	data := ArticlesFormData{
		Title:   "",
		Content: "",
		URL:     storeURL,
		Errors:  nil,
	}
	// 加载模板文件
	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(w, data); err != nil {
		panic(err)
	}
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
	// POST PostForm: map[content:[content ] title:[title]]
	// POST Form: map[content:[content ] title:[title]]
	// title 的值为: titler.PostForm 中 title 的值为: title
	// r.Form 中 title 的值为: title

	// 表单验证
	content := r.PostFormValue("content")
	errors := make(map[string]string)
	// 验证标题
	// 注意：len() 由于 Go 语言的字符串都以 UTF-8 格式保存，每个中文占用 3 个字节
	// 如果希望按习惯上的字符个数来计算，就需要使用 Go 语言中 utf8 包提供的 RuneCountInString () 函数来计数比如 utf8.RuneCountInString(title)
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if len(title) < 3 || len(title) > 40 {
		errors["title"] = "标题长度需介于 3-40"
	}
	// 验证内容
	if content == "" {
		errors["content"] = "内容不能为空"
	} else if len(content) < 10 {
		errors["content"] = "内容长度需大于或等于 10 个字节"
	}
	// 检查是否有错误
	if len(errors) == 0 {
		fmt.Fprint(w, "验证通过!<br>")
		fmt.Fprintf(w, "title 的值为: %v <br>", title)
		fmt.Fprintf(w, "title 的长度为: %v <br>", len(title))
		fmt.Fprintf(w, "content 的值为: %v <br>", content)
		fmt.Fprintf(w, "content 的长度为: %v <br>", len(content))

		// 插入数据操作
		lastInsertID, err := saveArticleToDB(title, content)
		if lastInsertID > 0 {
			// strconv.FormatInt() 方法来将类型为 int64 的 lastInsertID 转换为字符串，10 是十进制
			fmt.Fprint(w, "插入成功，ID 为"+strconv.FormatInt(lastInsertID, 10))
		} else {
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// fmt.Fprintf(w, "有错误发生，errors 的值为: %v <br>", errors)
		/* html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
		    <title>创建文章 —— 我的技术博客</title>
		    <style type="text/css">.error {color: red;}</style>
		</head>
		<body>
		    <form action="{{ .URL }}" method="post">
		        <p><input type="text" name="title" value="{{ .Title }}"></p>
		        {{ with .Errors.title }}
		        <p class="error">{{ . }}</p>
		        {{ end }}
		        <p><textarea name="content" cols="30" rows="10">{{ .Content }}</textarea></p>
		        {{ with .Errors.content }}
		        <p class="error">{{ . }}</p>
		        {{ end }}
		        <p><button type="submit">提交</button></p>
		    </form>
		</body>
		</html>
		` */
		// 通过路由参数生成 URL 路径
		storeURL, _ := router.Get("articles.store").URL()
		data := ArticlesFormData{
			Title:   title,
			Content: content,
			URL:     storeURL,
			Errors:  errors,
		}
		// template.New() 包的初始化。html 变量里是包含模板语法的内容，模板语法以双层大括号 {{ }} 包起来
		// tmpl, err := template.New("create-form").Parse(html)

		// 使用模板文件
		// 模板后缀名 .gohtml ，可以使用任意后缀名，这不会影响代码的运行；常见的 Go 模板后缀名有 .tmpl、.tpl、 .gohtml 等（推荐 .gohtml）
		tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
		if err != nil {
			panic(err)
		}
		if err = tmpl.Execute(w, data); err != nil {
			panic(err)
		}
	}
}

func saveArticleToDB(title string, content string) (int64, error) {
	// 变量初始化
	var (
		id   int64
		err  error
		rs   sql.Result
		stmt *sql.Stmt
	)

	// 1. 获取一个 prepare 声明语句
	// stmt 是 statement 的简写，是声明、陈述的意思。可以理解为将包含变量占位符 ? 的语句先告知 MySQL 服务器端
	stmt, err = db.Prepare("INSERT INTO articles (title, content) VALUES(?,?)")
	// 例行的错误检测
	if err != nil {
		return 0, err
	}

	// 2. 在此函数运行结束后关闭此语句，防止占用 SQL 连接
	// stmt 是一个指针变量，会占用 SQL 连接
	defer stmt.Close()

	// 3. 执行请求，传参进入绑定的内容，stmt.Exec() 的参数依次对应 db.Prepare() 参数中 SQL 变量占位符 ?
	rs, err = stmt.Exec(title, content)
	// 返回值是一个 sql.Result 对象
	// type Result interface {
	// 	// 使用 INSERT 向数据插入记录，数据表有自增 ID 时，该函数有返回值
	// 	LastInsertId() (int64, error)
	// 	// 表示影响的数据表行数，常用于 UPDATE/DELETE 等 SQL 语句中
	// 	RowsAffected() (int64, error)
	// }
	if err != nil {
		return 0, err
	}

	// 4. 插入成功的话，会返回自增 ID
	if id, err = rs.LastInsertId(); id > 0 {
		return id, nil
	}

	return 0, err
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
	initDB()
	createTables()

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
