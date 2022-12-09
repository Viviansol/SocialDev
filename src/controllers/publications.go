package controllers

import (
	"api/src/authentication"
	"api/src/dataBase"
	"api/src/modells"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtracUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication modells.Publication
	if err = json.Unmarshal(requestBody, &publication); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
	publication.AuthorId = userId
	db, err := dataBase.ConnectDataBase()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repository.NewPublicationRepository(db)
	publication.Id, err = repo.CreatePublication(publication)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, publication)

}

func SearchPublications(w http.ResponseWriter, r *http.Request) {

}

func SearchPublicationById(w http.ResponseWriter, r *http.Request) {

}
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}
func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
