package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ColumnRequest struct {
	Title       string `valid:"title" json:"title"`
	Description string `valid:"description" json:"description,omitempty"`
}

func ColumnSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:2", "max_cn:20"},
		"description": []string{"min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:名称为必填项",
			"min_cn:专栏长度需至少 2 个字",
			"max_cn:专栏长度不能超过 20 个字",
		},
		"description": []string{
			"min_cn:专栏描述长度需至少 3 个字",
			"max_cn:专栏描述长度不能超过 255 个字",
		},
	}
	return validate(data, rules, messages)
}
