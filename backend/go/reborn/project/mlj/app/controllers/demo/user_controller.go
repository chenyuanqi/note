package demo

import (
	"mlj/app/controllers"
	"mlj/app/services/demo"
	"mlj/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	controllers.Controller
}

func (ctrl *UserController) Create(c *gin.Context) {
	u := &demo.UserCreate{}
	if err := c.ShouldBind(&u.Request); err != nil {
		response.Fail(c, err.Error())
		return
	}

	if err := u.Create(); err != nil {
		response.Fail(c, err.Error())
		return
	}

	response.Success(c, u.Response, "")
}

func (ctrl *UserController) Update(c *gin.Context) {
	u := &demo.UserUpdate{}
	if err := c.ShouldBind(&u.Request); err != nil {
		response.Fail(c, err.Error())
		return
	}

	if err := u.Update(); err != nil {
		response.Fail(c, err.Error())
		return
	}

	response.Success(c, u.Response, "")
}

func (ctrl *UserController) Delete(c *gin.Context) {
	u := &demo.UserDelete{}
	if err := c.ShouldBind(&u.Request); err != nil {
		response.Fail(c, err.Error())
		return
	}

	if err := u.Delete(); err != nil {
		response.Fail(c, err.Error())
		return
	}

	response.Success(c, u.Response, "")
}

func (ctrl *UserController) Detail(c *gin.Context) {
	u := &demo.UserDetail{}
	if err := c.ShouldBind(&u.Request); err != nil {
		response.Fail(c, err.Error())
		return
	}

	if err := u.Detail(); err != nil {
		response.Fail(c, err.Error())
		return
	}

	response.Success(c, u.Response, "")
}
