package routes

import (
	"api/src/router/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		NeedAuth: false,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controllers.GetUsers,
		NeedAuth: false,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodGet,
		Function: controllers.GetUserById,
		NeedAuth: false,
	},
	{
		URI:      "/users/{id}/update",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		NeedAuth: false,
	},
	{
		URI:      "/users/{id}/delete",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: false,
	},
}
