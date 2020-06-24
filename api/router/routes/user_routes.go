package routes

import (
	"github.com/SideProjectTeam/clon-o-gram-be/api/controllers"
	"net/http"
)
var usersRoutes = [] Route{
	Route{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
}
