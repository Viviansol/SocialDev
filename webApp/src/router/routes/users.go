package routes

import (
	"net/http"
	"webApp/src/controllers"
)

var usersRoutes = []Route{

	{
		URI:                  "/create-user",
		Method:               http.MethodGet,
		Function:             controllers.LoadUserRegistrationPage,
		RequireAuthenticaion: false,
	},
	{
		URI:                  "/users",
		Method:               http.MethodPost,
		Function:             controllers.CreateUser,
		RequireAuthenticaion: false,
	},
}
