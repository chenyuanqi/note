package bi

import (
	"mlj/app/controllers"
	bb "mlj/app/services/bi"
	"mlj/pkg/response"

	"github.com/gin-gonic/gin"
)

type BiController struct {
	controllers.Controller
}

func (ctrl *BiController) Query(c *gin.Context) {
	b := &bb.Bi{}
	if err := c.ShouldBind(&b.Request); err != nil {
		response.Fail(c, err.Error())
		return
	}

	if err := b.Query(); err != nil {
		response.Fail(c, err.Error())
		return
	}

	response.Success(c, b.Response, "")
}
