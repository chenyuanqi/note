package bi

import (
	"mlj/app/controllers"
	bs "mlj/app/services/bussiness"
	"mlj/pkg/response"

	"github.com/gin-gonic/gin"
)

type BussinessController struct {
	controllers.Controller
}

func (ctrl *BussinessController) Query(c *gin.Context) {
	b := &bs.Business{}
	if err := c.ShouldBind(&b.Request); err != nil {
		response.Fail(c, err.Error())
	}

	if err := b.Query(); err != nil {
		response.Fail(c, err.Error())
	}

	response.Success(c, b.Response, "")
}
