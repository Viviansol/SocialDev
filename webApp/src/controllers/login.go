package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
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

}
