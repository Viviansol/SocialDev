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
}
