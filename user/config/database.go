package config

import (
	"../models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// OpenConnection Open a connection to mysql server
func OpenConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/gin_docker_train?parseTime=true")
	if err != nil {
		panic(err.Error)
	}

	db.AutoMigrate(&models.User{})

	return db
}
