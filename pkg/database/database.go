package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

// Connect 链接数据库
func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	var err error
	//使用gorm.open链接数据库
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}

}
