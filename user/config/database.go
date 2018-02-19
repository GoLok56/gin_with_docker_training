package config

import (
	"github.com/golok56/gin_with_docker_training/user/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// OpenConnection Open a connection to mysql server
func OpenConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@db-training:5656/gin_docker_train?parseTime=true")
	if err != nil {
		panic(err.Error)
	}

	db.AutoMigrate(&models.User{})

	return db
}
