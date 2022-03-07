package model

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	Content string `json:"content"`
	User    User
}
