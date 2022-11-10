package controllers

import (
	"api/src/dataBase"
	"api/src/modells"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user modells.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = user.PrepareUser(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := dataBase.ConnectDataBase()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repositories := repository.NewUserRepository(db)
	user.ID, err = repositories.CreateUser(user)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	db, err := dataBase.ConnectDataBase()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewUserRepository(db)
	user, err := repository.GetUser(nameOrNick)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, user)

}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting user by id"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User"))
}
