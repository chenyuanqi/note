package main

import (
	"blog/pkg/route"

	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

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

// Article  对应一条文章数据
type Article struct {
	Title, Content string
	ID             int64
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	/* // 1. 获取 URL 参数
	vars := mux.Vars(r)
	// 获取路径参数
	id := vars["id"]
	// fmt.Fprint(w, "文章 ID："+id)

	// 2. 读取对应的文章数据
	article := Article{}
	query := "SELECT * FROM articles WHERE id = ?"
	// QueryOne() 读取单条数据。参数只有一个的情况下，我们称之为纯文本模式，多个参数的情况下称之为 Prepare 模式（封装 Prepare 方法的调用）
	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content) // Scan() 将查询结果赋值到我们的 article struct 中，传参应与数据表字段的顺序保持一致。sql.Row 是个指针变量，保存有 SQL 连接。当调用 Scan() 时，就会将连接释放。所以在每次 QueryRow 后使用 Scan 是必须的
	// 等同于
	// stmt, err := db.Prepare(query)
	// checkError(err)
	// defer stmt.Close()
	// err = stmt.QueryRow(id).Scan(&article.ID, &article.Title, &article.Body) */

	// 1. 获取 URL 参数
	id := getRouteVariable("id", r)

	// 2. 读取对应的文章数据
	article, err := getArticleByID(id)

	// 3. 如果出现错误
	if err != nil {
		// 当 Scan() 发现没有返回数据的话，会返回 sql.ErrNoRows 类型的错误
		if err == sql.ErrNoRows {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 读取成功
		// fmt.Fprint(w, "读取成功，文章标题 —— "+article.Title)

		// 4. 读取成功，显示文章
		// Funcs() 方法的传参是 template.FuncMap 类型的 Map 对象。键为模板里调用的函数名称，值为当前上下文的函数名称
		tmpl, err := template.New("show.gohtml").
			Funcs(template.FuncMap{
				"RouteName2URL": route.Name2URL,
				"Int64ToString": Int64ToString,
			}).ParseFiles("resources/views/articles/show.gohtml")
		checkError(err)

		err = tmpl.Execute(w, article)
		checkError(err)
	}
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "访问文章列表~")
	// 1. 执行查询语句，返回一个结果集。使用 Query() 从数据库中读取多条数据（调用方式与 QueryRow() 和 Exec() 一致，支持单一参数的纯文本模式，以及多个参数的 Prepare 模式；这里使用 纯文本模式）
	rows, err := db.Query("SELECT * from articles") // rows 包含从数据库里读取出来的数据和 SQL 连接
	checkError(err)
	defer rows.Close() // 需在检测 err 以后调用，否则会让运行时 panic

	var articles []Article
	// 2. 循环读取结果
	// 使用 rows.Next() 遍历数据，遍历到最后内部遇到 EOF 错误，会自动调用 rows.Close() 将 SQL 连接关闭；如遇错误，SQL 连接也会自动关闭（rows.Close() 可调用多次，使用 rows.Close() 可保证 SQL 连接永远是关闭的）
	for rows.Next() {
		var article Article
		// 2.1 扫描每一行的结果并赋值到一个 article 对象中
		err := rows.Scan(&article.ID, &article.Title, &article.Content)
		checkError(err) // 检测下是否有错误发生
		// 2.2 将 article 追加到 articles 的这个数组中
		articles = append(articles, article)
	}

	// 2.3 检测遍历时是否发生错误
	err = rows.Err()
	checkError(err)

	// 3. 加载模板
	tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
	checkError(err)

	// 4. 渲染模板，将所有文章的数据传输进去
	err = tmpl.Execute(w, articles)
	checkError(err)
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
	/* errors := make(map[string]string)
	// 验证标题
	// 注意：len() 由于 Go 语言的字符串都以 UTF-8 格式保存，每个中文占用 3 个字节
	// 如果希望按习惯上的字符个数来计算，就需要使用 Go 语言中 utf8 包提供的 RuneCountInString () 函数来计数比如 utf8.RuneCountInString(title)
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if len(title) < 3 || len(title) > 40 {
		errors["title"] = "标题长度需介于 3-40 个字符"
	}
	// 验证内容
	if content == "" {
		errors["content"] = "内容不能为空"
	} else if len(content) < 10 {
		errors["content"] = "内容长度需大于或等于 10 个字符"
	} */

	errors := validateArticleFormData(title, content)

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

func articlesEditHandler(w http.ResponseWriter, r *http.Request) {
	/* // 1. 获取 URL 参数
	vars := mux.Vars(r)
	id := vars["id"]

	// 2. 读取对应的文章数据
	article := Article{}
	query := "SELECT * FROM articles WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content) */

	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	article, err := getArticleByID(id)

	// 3. 如果出现错误
	if err != nil {
		if err == sql.ErrNoRows {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 读取成功，显示表单
		updateURL, _ := router.Get("articles.update").URL("id", id)
		data := ArticlesFormData{
			Title:   article.Title,
			Content: article.Content,
			URL:     updateURL,
			Errors:  nil,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
		checkError(err)

		err = tmpl.Execute(w, data)
		checkError(err)
	}
}

func articlesUpdateHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "更新文章")

	// 1. 获取 URL 参数
	id := getRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_, err := getArticleByID(id)

	// 3. 如果出现错误
	if err != nil {
		if err == sql.ErrNoRows {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 未出现错误
		// 4.1 表单验证
		title := r.PostFormValue("title")
		content := r.PostFormValue("content")

		/* errors := make(map[string]string)
		// 验证标题
		if title == "" {
			errors["title"] = "标题不能为空"
		} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
			errors["title"] = "标题长度需介于 3-40"
		}

		// 验证内容
		if content == "" {
			errors["content"] = "内容不能为空"
		} else if utf8.RuneCountInString(content) < 10 {
			errors["content"] = "内容长度需大于或等于 10 个字节"
		} */
		errors := validateArticleFormData(title, content)

		if len(errors) == 0 {
			// 4.2 表单验证通过，更新数据

			query := "UPDATE articles SET title = ?, content = ? WHERE id = ?"
			// Exec() 的用法与 QueryRow() 类似，支持单独参数的纯文本模式 与 多个参数的 Prepare 模式
			rs, err := db.Exec(query, title, content, id)

			if err != nil {
				checkError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
			}

			// √ 更新成功，跳转到文章详情页
			if n, _ := rs.RowsAffected(); n > 0 {
				showURL, _ := router.Get("articles.show").URL("id", id)
				http.Redirect(w, r, showURL.String(), http.StatusFound)
			} else {
				fmt.Fprint(w, "您没有做任何更改！")
			}
		} else {

			// 4.3 表单验证不通过，显示理由

			updateURL, _ := router.Get("articles.update").URL("id", id)
			data := ArticlesFormData{
				Title:   title,
				Content: content,
				URL:     updateURL,
				Errors:  errors,
			}
			tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
			checkError(err)

			err = tmpl.Execute(w, data)
			checkError(err)
		}
	}
}

func articlesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := getRouteVariable("id", r)

	// 2. 读取对应的文章数据
	article, err := getArticleByID(id)

	// 3. 如果出现错误
	if err != nil {
		if err == sql.ErrNoRows {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 未出现错误，执行删除操作
		rowsAffected, err := article.Delete()

		// 4.1 发生错误
		if err != nil {
			// 应该是 SQL 报错了
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		} else {
			// 4.2 未发生错误
			if rowsAffected > 0 {
				// 重定向到文章列表页
				indexURL, _ := router.Get("articles.index").URL()
				http.Redirect(w, r, indexURL.String(), http.StatusFound)
			} else {
				// Edge case
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "404 文章未找到")
			}
		}
	}
}

func getRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

func getArticleByID(id string) (Article, error) {
	article := Article{}
	query := "SELECT * FROM articles WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content)
	return article, err
}

func validateArticleFormData(title string, content string) map[string]string {
	errors := make(map[string]string)
	// 验证标题
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度需介于 3-40 个字节"
	}

	// 验证内容
	if content == "" {
		errors["content"] = "内容不能为空"
	} else if utf8.RuneCountInString(content) < 10 {
		errors["content"] = "内容长度需大于或等于 10 个字节"
	}

	return errors
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	showURL, err := router.Get("articles.show").URL("id", strconv.FormatInt(a.ID, 10))
	if err != nil {
		checkError(err)
		return ""
	}
	return showURL.String()
}

// Delete 方法用以从数据库中删除单条记录
func (a Article) Delete() (rowsAffected int64, err error) {
	// Exec() 使用的是纯文本模式的查询模式，因为 ID 我们是从数据库里拿出来的，是自增 ID ，无需担心 SQL 注入，这样可以少发送一次 SQL 请求
	rs, err := db.Exec("DELETE FROM articles WHERE id = " + strconv.FormatInt(a.ID, 10))
	if err != nil {
		return 0, err
	}

	// √ 删除成功，跳转到文章详情页
	if n, _ := rs.RowsAffected(); n > 0 {
		return n, nil
	}

	return 0, nil
}

// Int64ToString 将 int64 转换为 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
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

	route.Initialize()
	router = route.Router

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
	// 更新博文
	router.HandleFunc("/articles/{id:[0-9]+}/edit", articlesEditHandler).Methods("GET").Name("articles.edit")
	router.HandleFunc("/articles/{id:[0-9]+}", articlesUpdateHandler).Methods("POST").Name("articles.update")
	// 删除博文
	router.HandleFunc("/articles/{id:[0-9]+}/delete", articlesDeleteHandler).Methods("POST").Name("articles.delete")

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
