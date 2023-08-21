package sqlite

import (
	"github.com/willpang/data-entry/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {

	database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate User database!")
	}

	return database
}
