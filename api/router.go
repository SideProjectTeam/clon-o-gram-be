package api

import (
	"github.com/gin-gonic/gin"
	
	"github.com/SideProjectTeam/clon-o-gram-be/api/controllers/users"
)

func SetRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/register",users.RegisterUser)
	
	return r
}