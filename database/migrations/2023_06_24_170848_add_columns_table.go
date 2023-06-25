package migrations

import (
	"Goplus/app/models"
	"Goplus/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Column struct {
		models.BaseModel

		Author      uint64 `gorm:"type:bigint;"`
		Title       string `gorm:"type:varchar(255);index;default:null"`
		Description string `gorm:"type:varchar(255);index;default:null"`
		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Column{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Column{})
	}

	migrate.Add("2023_06_24_170848_add_columns_table", up, down)
}
