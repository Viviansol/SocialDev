package routes

import (
	"api/src/controllers"
	"net/http"
)

var publicationsRoutes = []Route{
	{
		Uri:                   "/publications",
		Method:                http.MethodPost,
		Function:              controllers.CreatePublication,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/publications",
		Method:                http.MethodGet,
		Function:              controllers.SearchPublications,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/publications/{publicationId}",
		Method:                http.MethodGet,
		Function:              controllers.SearchPublicationById,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/publications/{publicationId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePublication,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/publications/{publicationId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePublication,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{userId}/publications",
		Method:                http.MethodGet,
		Function:              controllers.SearchPublicationsByUSer,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/publications/{publicationId}/like",
		Method:                http.MethodPost,
		Function:              controllers.LikePublication,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/publications/{publicationId}/Unlike",
		Method:                http.MethodPost,
		Function:              controllers.UnlikePublication,
		RequireAuthentication: true,
	},
}
