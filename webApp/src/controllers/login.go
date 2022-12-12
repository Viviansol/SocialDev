package controllers

import "net/http"

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(" login page"))
}
