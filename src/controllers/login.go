package controllers

import (
	"api/src/authentication"
	"api/src/dataBase"
	"api/src/modells"
	"api/src/repository"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user modells.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.JSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := dataBase.ConnectDataBase()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repository.NewUserRepository(db)

	savedUser, err := repo.GetUserByEmail(user.Email)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(savedUser.Password, user.Password); err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(savedUser.ID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
	}

	userId := strconv.FormatUint(savedUser.ID, 10)

	response.JSON(w, http.StatusOK, modells.AuthData{ID: userId, Token: token})

}
