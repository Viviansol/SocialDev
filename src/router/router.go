package router

import (
	"api/src/router/routes"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRouter(r)
}
