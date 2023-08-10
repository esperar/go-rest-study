package main

import (
	"github.com/gorilla/mux"
	"gorest/database"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	log.Fatal(http.ListenAndServe(":8090", router))

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
