package batch

import (
	"go-message-api/config"
	"go-message-api/model"
)

/*
マイグレーション
*/
func Migration() {
	config.SetUp()

	config.Db.AutoMigrate(&model.User{})
	config.Db.AutoMigrate(&model.Memo{})
}
