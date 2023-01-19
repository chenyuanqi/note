package main

import (
	"log"
	"net/http"

	"mlj/app/controllers/bi"
	"mlj/app/controllers/demo"
	"mlj/app/middlewares"
	"mlj/pkg/common/consts"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	gin.SetMode(gin.DebugMode)
	if consts.EnvModeIsRelease {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middlewares.Cors)
	r.Use(middlewares.Recover)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result":  "success",
			"message": "ping success",
		})
	})

	biGroup := r.Group("bi")
	{
		bb := new(bi.BiController)
		biGroup.GET("index", bb.Query)

		business := new(bi.BusinessController)
		biGroup.GET("business", business.Query)

		weibo := new(bi.WeiboController)
		biGroup.GET("weibo", weibo.Query)
	}

	userGroup := r.Group("user")
	{
		u := new(demo.UserController)
		userGroup.GET("detail", u.Detail)
		userGroup.POST("create", u.Create)
		userGroup.PUT("update", u.Update)
		userGroup.DELETE("delete", u.Delete)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result":  "error",
			"message": "no server",
		})
	})

	if err := http.ListenAndServe(":9999", r); err != nil {
		log.Fatalf("ListenAndServe err: %s", err)
	}
}
