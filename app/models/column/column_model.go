// Package column 模型
package column

import (
	"Goplus/app/models"
	"Goplus/pkg/database"
)

type Column struct {
	models.BaseModel
	Author      uint64 `json:"author,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`

	models.CommonTimestampsField
}

func (column *Column) Create() {
	database.DB.Create(&column)
}

func (column *Column) Save() (rowsAffected int64) {
	result := database.DB.Save(&column)
	return result.RowsAffected
}

func (column *Column) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&column)
	return result.RowsAffected
}
