package bi

import (
	"mlj/app/controllers"
	bs "mlj/app/services/business"
	"mlj/pkg/response"

	"github.com/gin-gonic/gin"
)

type BusinessController struct {
	controllers.Controller
}

func (ctrl *BusinessController) Query(c *gin.Context) {
	b := &bs.Business{}
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
