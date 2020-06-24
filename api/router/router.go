package router

import (
	"github.com/SideProjectTeam/clon-o-gram-be/api/router/routes"
	"github.com/gorilla/mux"
)


func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}