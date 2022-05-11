package main

import (
	"blog/app/http/middlewares"
	"blog/bootstrap"
	"blog/pkg/logger"

	"net/http"

	"github.com/gorilla/mux"
)

// 语法错误：non-declaration statement outside function body 函数外无法使用变量赋值语句
// router := mux.NewRouter()
// 包级别的变量声明时不能使用 := 语法，修改为带关键词 var 的变量声明即可
var router = mux.NewRouter()

func main() {
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	err := http.ListenAndServe(":8000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
