package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webApp/src/config"
	"webApp/src/models"
	"webApp/src/requisitions"
	"webApp/src/responses"
	"webApp/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplatesA(w, "login.html", nil)
}

func LoadUserRegistrationPage(w http.ResponseWriter, r *http.Request) {

	utils.ExecuteTemplatesA(w, "registration.html", nil)

}

func LoadMainPage(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/publications", config.APIURL)
	response, err := requisitions.MakeRequisitionWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorApi{ErrorAPi: err.Error()})
		return
	}
	if response.StatusCode >= 400 {
		responses.StatusCodeErrorTreatment(w, response)
		return
	}
	defer response.Body.Close()
	var publications []models.Publication
	if err = json.NewDecoder(response.Body).Decode(publications); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorApi{ErrorAPi: err.Error()})
		return
	}

	utils.ExecuteTemplatesA(w, "home.html", publications)
}
