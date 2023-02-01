package main

import (
	"log"
	"net/http"

	"mlj/app/controllers/bi"
	"mlj/app/controllers/demo"
	"mlj/app/middlewares"
	btConfig "mlj/config"
	"mlj/pkg/common/consts"
	"mlj/pkg/config"
	"mlj/pkg/facade"
	"mlj/pkg/response"

	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化配置
	btConfig.Initialize()
	config.InitConfig(config.Get("app.env"))
}

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

	requestGroup := r.Group("request")
	{
		requestGroup.GET("merchant-coupon-list", func(c *gin.Context) {
			res, err := facade.GetMerchantCouponList(1, 10)
			if err != nil {
				response.Fail(c, err.Error())
				return
			}

			response.Success(c, res, "")
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result":  "error",
			"message": "no server",
		})
	})

	if err := http.ListenAndServe(":"+config.Get("app.port"), r); err != nil {
		log.Fatalf("ListenAndServe err: %s", err)
	}
}
