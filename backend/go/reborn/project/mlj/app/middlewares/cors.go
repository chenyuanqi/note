package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

/**
 * 支持跨域
 */
func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	c.Header("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization,token,access-token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("x-container-pod", os.Getenv("HOSTNAME"))

	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
	}

	c.Next()
}
