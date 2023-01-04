package controllers

import (
	"fmt"
	"net/http"
	"webApp/src/config"
	"webApp/src/requisitions"
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
	fmt.Println(response.StatusCode, err)
	utils.ExecuteTemplatesA(w, "home.html", nil)
}
