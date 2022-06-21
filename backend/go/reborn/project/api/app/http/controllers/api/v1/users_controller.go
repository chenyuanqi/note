package v1

import (
	// "api/app/models/user"
	// "api/app/policies"
	// "api/app/requests"
	"api/app/models/user"
	"api/app/requests"
	"api/pkg/auth"
	"api/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	// data := user.All()
	// response.Data(c, data)
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *UsersController) UpdateProfile(c *gin.Context) {

	request := requests.UserUpdateProfileRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateProfile); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Name = request.Name
	currentUser.City = request.City
	currentUser.Introduction = request.Introduction
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdateEmail(c *gin.Context) {

	request := requests.UserUpdateEmailRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateEmail); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Email = request.Email
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		// 失败，显示错误提示
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

// func (ctrl *UsersController) Index(c *gin.Context) {
// 	users := user.All()
// 	response.Data(c, users)
// }

// func (ctrl *UsersController) Show(c *gin.Context) {
// 	userModel := user.Get(c.Param("id"))
// 	if userModel.ID == 0 {
// 		response.Abort404(c)
// 		return
// 	}
// 	response.Data(c, userModel)
// }

// func (ctrl *UsersController) Store(c *gin.Context) {

// 	request := requests.UserRequest{}
// 	if ok := requests.Validate(c, &request, requests.UserSave); !ok {
// 		return
// 	}

// 	userModel := user.User{
// 		FieldName: request.FieldName,
// 	}
// 	userModel.Create()
// 	if userModel.ID > 0 {
// 		response.Created(c, userModel)
// 	} else {
// 		response.Abort500(c, "创建失败，请稍后尝试~")
// 	}
// }

// func (ctrl *UsersController) Update(c *gin.Context) {

// 	userModel := user.Get(c.Param("id"))
// 	if userModel.ID == 0 {
// 		response.Abort404(c)
// 		return
// 	}

// 	if ok := policies.CanModifyUser(c, userModel); !ok {
// 		response.Abort403(c)
// 		return
// 	}

// 	request := requests.UserRequest{}
// 	bindOk, errs := requests.Validate(c, &request, requests.UserSave)
// 	if !bindOk {
// 		return
// 	}
// 	if len(errs) > 0 {
// 		response.ValidationError(c, 20101, errs)
// 		return
// 	}

// 	userModel.FieldName = request.FieldName
// 	rowsAffected := userModel.Save()
// 	if rowsAffected > 0 {
// 		response.Data(c, userModel)
// 	} else {
// 		response.Abort500(c, "更新失败，请稍后尝试~")
// 	}
// }

// func (ctrl *UsersController) Delete(c *gin.Context) {

// 	userModel := user.Get(c.Param("id"))
// 	if userModel.ID == 0 {
// 		response.Abort404(c)
// 		return
// 	}

// 	if ok := policies.CanModifyUser(c, userModel); !ok {
// 		response.Abort403(c)
// 		return
// 	}

// 	rowsAffected := userModel.Delete()
// 	if rowsAffected > 0 {
// 		response.Success(c)
// 		return
// 	}

// 	response.Abort500(c, "删除失败，请稍后尝试~")
// }
