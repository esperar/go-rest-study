package main

import "gorest/database"

func main() {
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
