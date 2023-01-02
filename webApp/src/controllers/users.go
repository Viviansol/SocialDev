package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webApp/src/responses"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nickName": r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorApi{ErrorAPi: err.Error()})
	}
	response, err := http.Post("http://localhost:5300/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorApi{ErrorAPi: err.Error()})
	}
	defer response.Body.Close()
	fmt.Println(response.Body)
	if response.StatusCode >= 400 {
		responses.StatusCodeErrorTreatment(w, response)
		return
	}
	responses.JSON(w, response.StatusCode, nil)
}
