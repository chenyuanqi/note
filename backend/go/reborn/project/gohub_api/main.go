package main

import (
    "fmt"

	"github.com/gin-gonic/gin"

	"gohub_api/bootstrap"
)

func main() {
	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	bootstrap.SetupRoute(r)

	if err := r.Run(":9999"); err != nil {
		fmt.Println(err.Error())
	}
}
