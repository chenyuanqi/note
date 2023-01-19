package bi

import (
	"mlj/app/controllers"
	"mlj/app/services/weibo"
	"mlj/pkg/response"

	"github.com/gin-gonic/gin"
)

type WeiboController struct {
	controllers.Controller
}

func (ctrl *WeiboController) Query(c *gin.Context) {
	w := &weibo.Weibo{}
	if err := c.ShouldBind(&w.Request); err != nil {
		response.Fail(c, err.Error())
	}

	if err := w.Query(); err != nil {
		response.Fail(c, err.Error())
	}

	response.Success(c, w.Response, "")
}
