package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

// 中间件（拦截器），功能：预处理，登录授权、验证、分页、耗时统计...
// func myHandler() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// 通过自定义中间件，设置的值，在后续处理只要调用了这个中间件的都可以拿到这里的参数
// 		ctx.Set("usersesion", "userid-1")
// 		ctx.Next()  // 放行
// 		ctx.Abort() // 阻止
// 	}
// }

func main() {
	// 创建一个服务
	ginServer := gin.Default()
	ginServer.Use(favicon.New("./Arctime.ico")) // 这里如果添加了东西然后再运行没有变化，请重启浏览器，浏览器有缓存

	// 加载静态页面
	ginServer.LoadHTMLGlob("templates/*") // 一种是全局加载，一种是加载指定的文件

	// 加载资源文件
	ginServer.Static("/static", "./static")

	// 相应一个页面给前端

	ginServer.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "This data is come from Go background.",
		})
	})

	// 能加载静态页面也可以加载测试文件

	// 获取请求中的参数

	// 传统方式：usl?userid=xxx&username=conqueror712
	// Rustful方式：/user/info/1/conqueror712

	// 下面是传统方式的例子
	ginServer.GET("/user/info", func(context *gin.Context) { // 这个格式是固定的
		userid := context.Query("userid")
		username := context.Query("username")
		// 拿到之后返回给前端
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})
	// 此时执行代码之后，在浏览器中可以输入http://localhost:8081/user/info?userid=111&username=666
	// 就可以看到返回了JSON格式的数据

	// 下面是Rustful方式的例子
	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		// 还是一样，返回给前端
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})
	// 指定代码后，只需要在浏览器中http://localhost:8081/user/info/111/555
	// 就可以看到返回了JSON数据了，非常方便简洁

	// 序列化
	// 前端给后端传递JSON
	ginServer.POST("/json", func(ctx *gin.Context) {
		// request.body
		data, _ := ctx.GetRawData()
		var m map[string]interface{} // Go语言中object一般用空接口来表示，可以接收anything
		// 顺带一提，1.18以上，interface可以直接改成any
		_ = json.Unmarshal(data, &m)
		ctx.JSON(http.StatusOK, m)
	})
	// 用apipost或者postman写一段json传到localhost:8081/json里就可以了
	/*
		json示例：
		{
			"name": "Conqueror712",
			"age": 666,
			"address": "Mars"
		}
	*/
	// 看到后端的实时响应里面接收到数据就可以了

	// 处理表单请求 这些都是支持函数式编程，Go语言特性，可以把函数作为参数传进来
	ginServer.POST("/user/add", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		ctx.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	// 路由
	ginServer.GET("/test", func(ctx *gin.Context) {
		// 重定向 -> 301
		ctx.Redirect(301, "https://conqueror712.gitee.io/conqueror712.gitee.io/")
	})
	// http://localhost:8081/test

	// 404
	ginServer.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(404, "404.html", nil)
	})

	// 路由组暂略

	// 服务器端口，用服务器端口来访问地址
	ginServer.Run(":8081") // 不写的话默认是8080，也可以更改
}
