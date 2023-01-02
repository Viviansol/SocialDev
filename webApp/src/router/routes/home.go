package routes

import (
	"net/http"
	"webApp/src/controllers"
)

var mainPageRoute = Route{
	URI:                  "/home",
	Method:               http.MethodGet,
	Function:             controllers.LoadMainPage,
	RequireAuthenticaion: true,
}
