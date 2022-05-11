package controllers

import (
	"blog/app/models/article"
	"blog/pkg/logger"
	"blog/pkg/route"
	"blog/pkg/types"
	"strconv"
	"unicode/utf8"

	"fmt"
	"html/template"
	"net/http"

	"gorm.io/gorm"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
}

// Article  对应一条文章数据
type Article struct {
	Title, Content string
	ID             uint64
}

// Index 文章列表页
func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	// 1. 获取结果集
	articles, err := article.GetAll()

	if err != nil {
		// 数据库错误
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	} else {
		// 2. 加载模板
		tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
		logger.LogError(err)

		// 3. 渲染模板，将所有文章的数据传输进去
		err = tmpl.Execute(w, articles)
		logger.LogError(err)
	}
}

// Show 文章详情页面
func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	article, err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		// 当 Scan() 发现没有返回数据的话，会返回 sql.ErrNoRows 类型的错误
		if err == gorm.ErrRecordNotFound {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 读取成功
		fmt.Fprint(w, "读取成功，文章标题 —— "+article.Title)

		// 4. 读取成功，显示文章
		// Funcs() 方法的传参是 template.FuncMap 类型的 Map 对象。键为模板里调用的函数名称，值为当前上下文的函数名称

		tmpl, err := template.New("show.gohtml").
			Funcs(template.FuncMap{
				"RouteName2URL":  route.Name2URL,
				"Uint64ToString": types.Uint64ToString,
			}).ParseFiles("resources/views/articles/show.gohtml")
		logger.LogError(err)

		err = tmpl.Execute(w, article)
		logger.LogError(err)
	}
}

// ArticlesFormData 创建博文表单数据
type ArticlesFormData struct {
	Title, Content string
	URL            string
	Errors         map[string]string
}

// Create 文章创建页面
func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	storeURL := route.Name2URL("articles.store")
	data := ArticlesFormData{
		Title:   "",
		Content: "",
		URL:     storeURL,
		Errors:  nil,
	}
	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

// Store 文章创建页面
func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")

	// 检查是否有错误
	errors := validateArticleFormData(title, content)
	if len(errors) == 0 {
		_article := article.Article{
			Title:   title,
			Content: content,
		}
		_article.Create()
		if _article.ID > 0 {
			fmt.Fprint(w, "插入成功，ID 为"+strconv.FormatUint(_article.ID, 10))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章失败，请联系管理员")
		}
	} else {

		storeURL := route.Name2URL("articles.store")

		data := ArticlesFormData{
			Title:   title,
			Content: content,
			URL:     storeURL,
			Errors:  errors,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")

		logger.LogError(err)

		err = tmpl.Execute(w, data)
		logger.LogError(err)
	}
}

func validateArticleFormData(title string, content string) map[string]string {
	errors := make(map[string]string)
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
	}

	return errors
}
