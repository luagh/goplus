package auth

import (
	v1 "Goplus/app/http/controllers/api/v1"
	"Goplus/app/requests"
	"Goplus/pkg/auth"
	"Goplus/pkg/jwt"
	"Goplus/pkg/response"
	"github.com/gin-gonic/gin"
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseAPIController
}

func (lc LoginController) LoginByPhone(c *gin.Context) {

	request := requests.LoginByPhoneRequest{}

	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}
	// 2. 尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在")
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
