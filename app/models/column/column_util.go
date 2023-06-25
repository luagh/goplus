package column

import (
	"Goplus/pkg/app"
	"Goplus/pkg/database"
	"Goplus/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (column Column) {
	database.DB.Where("id", idstr).First(&column)
	return
}

func GetBy(field, value string) (column Column) {
	database.DB.Where("? = ?", field, value).First(&column)
	return
}

func All() (columns []Column) {
	database.DB.Find(&columns)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Column{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (columns []Column, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Column{}),
		&columns,
		app.V1URL(database.TableName(&Column{})),
		perPage,
	)
	return
}
