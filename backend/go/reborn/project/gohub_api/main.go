package main

import (
    "fmt"
    "flag"

	"github.com/gin-gonic/gin"

	"gohub_api/bootstrap"
	btsConfig "gohub_api/config"
    "gohub_api/pkg/config"
)

func init() {
    btsConfig.Initialize()
}

func main() {
    var env string
    flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
    flag.Parse()
    config.InitConfig(env)

	r := gin.New()

	bootstrap.SetupRoute(r)

	if err := r.Run(":" + config.Get("app.port")); err != nil {
		fmt.Println(err.Error())
	}
}
