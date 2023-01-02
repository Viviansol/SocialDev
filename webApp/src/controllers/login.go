package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webApp/src/config"
	"webApp/src/models"
	"webApp/src/responses"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorApi{ErrorAPi: err.Error()})
		return
	}
	url := fmt.Sprintf("%s/login", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorApi{ErrorAPi: err.Error()})
		return
	}

	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.StatusCodeErrorTreatment(w, response)
		return
	}

	var authData models.AuthData
	if err := json.NewDecoder(response.Body).Decode(&authData); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorApi{ErrorAPi: err.Error()})
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
