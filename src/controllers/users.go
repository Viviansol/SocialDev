package controllers

import (
	"api/src/authentication"
	"api/src/dataBase"
	"api/src/modells"
	"api/src/repository"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"errors"
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
	userIDToken, err := authentication.ExtracUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
	}

	if userID != userIDToken {
		response.Erro(w, http.StatusForbidden, errors.New("It's not possible to update a different user"))
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
	userIDToken, err := authentication.ExtracUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
	}

	if userId != userIDToken {
		response.Erro(w, http.StatusForbidden, errors.New("It's not possible to delete a different user"))
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

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := authentication.ExtracUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
	}

	if followerId == userId {
		response.Erro(w, http.StatusForbidden, errors.New("You can´t follow yourself"))
		return
	}

	db, err := dataBase.ConnectDataBase()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repository.NewUserRepository(db)
	if err = repo.FollowUser(userId, followerId); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := authentication.ExtracUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
	}

	if followerId == userId {
		response.Erro(w, http.StatusForbidden, errors.New("You can´t unffollow yourself"))
		return
	}

	db, err := dataBase.ConnectDataBase()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repository.NewUserRepository(db)
	if err = repo.UnfollowUser(userId, followerId); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)

}

func SearchFollowers(w http.ResponseWriter, r *http.Request) {
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
	followers, err := repo.SearchFollowers(userId)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, followers)
}

func SearchFollowing(w http.ResponseWriter, r *http.Request) {
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
	users, err := repo.SearchFollowing(userId)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
	}
	response.JSON(w, http.StatusOK, users)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIdToken, err := authentication.ExtracUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if userIdToken != userId {
		response.Erro(w, http.StatusForbidden, errors.New("You can't update an user different of yours "))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	var password modells.Password
	if err = json.Unmarshal(requestBody, &password); err != nil {
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

	passwordInTheDb, err := repo.SearchPasswordById(userId)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(passwordInTheDb, password.Actual); err != nil {
		response.Erro(w, http.StatusInternalServerError, errors.New("Actual password is not the same as the password stored in the db"))
		return
	}

	passwordWithHash, err := security.Hash(password.New)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = repo.UpdatePassword(userId, string(passwordWithHash)); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}
