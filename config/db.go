package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

/*
DB接続の初期化
*/
func SetUp() error {
	dsn := "root:password@tcp(127.0.0.1:3306)/memo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		return err
	}

	Db = db

	return nil
}
