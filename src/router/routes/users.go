package routes

import (
	usersController "devbook/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    usersController.Create,
		RequestAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    usersController.Index,
		RequestAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodGet,
		Function:    usersController.Show,
		RequestAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodPut,
		Function:    usersController.Update,
		RequestAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodDelete,
		Function:    usersController.Delete,
		RequestAuth: false,
	},
}
