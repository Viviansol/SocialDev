package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		Uri:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.GetUser,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.GetUserById,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users/{userId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequireAuthentication: false,
	}, {
		Uri:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequireAuthentication: false,
	},
}