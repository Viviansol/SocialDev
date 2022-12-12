package routes

import (
	"net/http"
	"webApp/src/controllers"
)

var routesLogin = []Route{

	{
		URI:                  "/",
		Method:               http.MethodGet,
		Function:             controllers.LoadLoginPage,
		RequireAuthenticaion: false,
	},
	{
		URI:                  "/login",
		Method:               http.MethodGet,
		Function:             controllers.LoadLoginPage,
		RequireAuthenticaion: false,
	},
}
