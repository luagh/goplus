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

// 所有专栏
func (ctrl *ColumnsController) Index(c *gin.Context) {
	data, pager := column.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"paper": pager,
	})
}

// 更新专栏
func (ctrl *ColumnsController) Update(c *gin.Context) {
	// 验证 url 参数 id 是否正确
	columnModel := column.Get(c.Param("id"))
	if columnModel.ID == 0 {
		response.Abort404(c)
		return
	}
	// 表单验证
	request := requests.ColumnRequest{}
	if ok := requests.Validate(c, &request, requests.ColumnSave); !ok {
		return
	}

	// 保存数据
	columnModel.Title = request.Title
	columnModel.Description = request.Description
	rowsAffected := columnModel.Save()

	if rowsAffected > 0 {
		response.Data(c, columnModel)
	} else {
		response.Abort500(c)
	}
}
