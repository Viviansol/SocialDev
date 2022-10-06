package controllers

import (
	"api/src/dataBase"
	"api/src/modells"
	"api/src/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user modells.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}
	db, err := dataBase.ConnectDataBase()
	if err != nil {
		log.Fatal(err)
	}

	repositories := repository.NewUserRepository(db)
	repositories.CreateUser(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users"))
}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting user by id"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User"))
}
