package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorest/database"
	"gorest/entity"
	"io/ioutil"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)

	database.Connector.Create(person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	database.Connector.First(&person, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}
