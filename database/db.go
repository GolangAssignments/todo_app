package database

import (
	"fmt"
	"log"
	"todo_app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect Sqlite Database:", err)
	}
	err = DB.AutoMigrate(&models.Todo{})

	if err != nil {
		log.Fatal("Failed to migrate the table:", err)
	}

	fmt.Println("Sqlite database connected successfully!")
}
