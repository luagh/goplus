// package auth处理逻辑
package auth

import (
	v1 "Goplus/app/http/controllers/api/v1"
	"Goplus/app/models/user"
	"Goplus/app/requests"
	"Goplus/pkg/response"
	"github.com/gin-gonic/gin"
)

// PasswordController 用户控制器
type PasswordController struct {
	v1.BaseAPIController
}

func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.ResetByPhoneRequest{}

	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}
	// 2. 更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}

func (pc PasswordController) ResetByEmail(c *gin.Context) {

	request := requests.ResetByEmailRequest{}

	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}
	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}
