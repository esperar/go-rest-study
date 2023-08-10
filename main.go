package main

import (
	"gorest/database"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")
}

func initDB() {
	config := database.Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "Esperer123!",
		DB:         "gorest",
	}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)

	if err != nil {
		panic(err.Error())
	}
}
