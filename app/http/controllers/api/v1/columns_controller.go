package v1

import (
	"Goplus/app/models/column"
	"Goplus/app/requests"
	"Goplus/pkg/auth"
	"Goplus/pkg/response"
	"Goplus/pkg/snowflake"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ColumnsController struct {
	BaseAPIController
}

func (ctrl *ColumnsController) Store(c *gin.Context) {

	request := requests.ColumnRequest{}
	if ok := requests.Validate(c, &request, requests.ColumnSave); !ok {
		return
	}
	Cid, _ := snowflake.GetID()
	author := auth.CurrentUID(c)
	Aid, _ := strconv.ParseUint(author, 10, 64)

	columnModel := column.Column{
		Author:      Aid,
		Title:       request.Title,
		Description: request.Description,
	}
	columnModel.ID = Cid
	columnModel.Create()
	if columnModel.ID > 0 {
		response.Created(c, columnModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}
