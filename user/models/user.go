package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:100"`
	Password string `json:"password" gorm:"type:text"`
	Name     string `json:"name" gorm:"size:50"`
}
