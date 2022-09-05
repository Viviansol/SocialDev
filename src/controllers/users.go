package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users"))
}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting user by id"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User"))
}
