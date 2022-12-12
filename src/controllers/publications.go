package controllers

import (
	"api/src/authentication"
	"api/src/dataBase"
	"api/src/modells"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
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

	if err = publication.PreparePublication(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

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

	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["publicationId"], 10, 64)
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

	repo := repository.NewPublicationRepository(db)

	publication, err := repo.SearchPublicationById(publicationID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publication)

}
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}
func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
