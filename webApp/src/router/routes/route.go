package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"webApp/middleware"
)

type Route struct {
	URI                  string
	Method               string
	Function             func(w http.ResponseWriter, r *http.Request)
	RequireAuthenticaion bool
}

func Configure(router *mux.Router) *mux.Router {

	routes := routesLogin
	routes = append(routes, usersRoutes...)
	routes = append(routes, mainPageRoute)

	for _, route := range routes {
		if route.RequireAuthenticaion == true {
			router.HandleFunc(route.URI, middleware.Logger(middleware.Authenticate(route.Function))).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, middleware.Logger(route.Function)).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
