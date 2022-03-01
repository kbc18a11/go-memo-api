package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUp() error {
	dsn := "root:password@tcp(127.0.0.1:3306)/onemessage?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		return err
	}

	return nil
}
