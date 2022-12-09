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
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{userId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequireAuthentication: true,
	}, {
		Uri:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequireAuthentication: true,
	}, {
		Uri:                   "/users/{userId}/follow",
		Method:                http.MethodPost,
		Function:              controllers.FollowUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{userId}/unfollow",
		Method:                http.MethodPost,
		Function:              controllers.UnfollowUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{userId}/followers",
		Method:                http.MethodGet,
		Function:              controllers.SearchFollowers,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{userId}/following",
		Method:                http.MethodGet,
		Function:              controllers.SearchFollowing,
		RequireAuthentication: true,
	},
}
