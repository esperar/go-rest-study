package main

import (
	"github.com/gorilla/mux"
	"gorest/controller"
	"gorest/database"
	"gorest/entity"
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

func initalizeHandlers(router *mux.Router) {
	router.HandleFunc("/create", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controller.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controller.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controller.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controller.DeletePersonByID).Methods("DELETE")
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
	database.Migrate(^entity.Person{})
}
