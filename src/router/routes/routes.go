package routes

import (
	"api/src/middleware"
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
	routes = append(routes, loginRoute)
	for _, route := range routes {
		if route.RequireAuthentication {
			r.HandleFunc(route.Uri,
				middleware.Logger(middleware.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, middleware.Logger(route.Function)).Methods(route.Method)
		}

	}
	return r
}
