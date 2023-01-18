package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func Fail(c *gin.Context, errmsg string) {
	JSON(c, gin.H{
		"result":  "error",
		"message": errmsg,
	})
}

func Success(c *gin.Context, data interface{}, message string) {
	JSON(c, gin.H{
		"result":  "success",
		"data":    data,
		"message": message,
	})
}
