package model

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	Content string
	User    User
}
