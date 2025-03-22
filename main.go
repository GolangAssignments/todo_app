package main

import (
	"log"
	"todo_app/database"
	"todo_app/routes"
)

func init() {
	database.ConnectDB()
}

func main() {
	routes.SetupRoutes().Run(":8080")

	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Fatal("Failed to get DB connection object:", err)
	}
	defer sqlDB.Close()
}
