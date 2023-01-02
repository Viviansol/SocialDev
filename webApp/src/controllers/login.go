package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
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
	response, err := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(user))
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
