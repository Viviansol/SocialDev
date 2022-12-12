package router

import (
	"github.com/gorilla/mux"
	"webApp/src/router/routes"
)

func Generate() *mux.Router {

	r := mux.NewRouter()
	return routes.Configure(r)
}
