package database

import (
	"fmt"
	"notes_server/config"
	"notes_server/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	if config.DB.Type == "sqlite" {
		db, err := gorm.Open(sqlite.Open(config.DB.DatabaseFileName), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		DB = db

	} else if config.DB.Type == "mysql" {
		// Not tested
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB.User, config.DB.Password, config.DB.DBHost, config.DB.DBPort, config.DB.DBName)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		DB = db
	} else {
		panic("invalid database config")
	}

	DB.AutoMigrate(&models.User{}, &models.Note{})
}
