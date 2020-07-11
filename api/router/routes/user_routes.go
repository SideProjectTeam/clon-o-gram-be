package routes

import (
	"net/http"

	"github.com/SideProjectTeam/clon-o-gram-be/api/controllers"
)

var usersRoutes = []Route{
	Route{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
	Route{
		Uri:     "/register",
		Method:  http.MethodPost,
		Handler: controllers.RegisterUser,
	},
}
