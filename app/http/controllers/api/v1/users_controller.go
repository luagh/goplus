package v1

import (
	"Goplus/pkg/auth"
	"Goplus/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// 从auth中获取当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userMOdel := auth.CurrentUser(c)
	response.Data(c, userMOdel)
}
