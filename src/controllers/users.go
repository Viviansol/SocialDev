package controllers

import (
	"api/src/dataBase"
	"api/src/modells"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
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
	if err = user.PrepareUser("registration"); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := dataBase.ConnectDataBase()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositories := repository.NewUserRepository(db)
	user.ID, err = repositories.CreateUser(user)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

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

	if err = user.PrepareUser("edition"); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := dataBase.ConnectDataBase()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	if err = repo.UpdateUser(userID, user); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)

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
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := dataBase.ConnectDataBase()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repository.NewUserRepository(db)
	user, err := repo.GetUserById(userId)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, user)

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := dataBase.ConnectDataBase()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repository.NewUserRepository(db)
	if err = repo.DeleteUser(userId); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)

}