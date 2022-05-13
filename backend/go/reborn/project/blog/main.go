package main

import (
	"blog/app/http/middlewares"
	"blog/bootstrap"
	"blog/config"
	c "blog/pkg/config"

	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

// 语法错误：non-declaration statement outside function body 函数外无法使用变量赋值语句
// router := mux.NewRouter()
// 包级别的变量声明时不能使用 := 语法，修改为带关键词 var 的变量声明即可
var router = mux.NewRouter()

func main() {
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
