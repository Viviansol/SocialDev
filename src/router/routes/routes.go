package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Uri                   string
	Method                string
	Function              func(w http.ResponseWriter, r *http.Request)
	RequireAuthentication bool
}

func ConfigRouter(r *mux.Router) *mux.Router {
	routes := usersRoutes
	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}
	return r
}
