package controllers

import (
	"net/http"
	"webApp/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplatesA(w, "login.html", nil)
}
