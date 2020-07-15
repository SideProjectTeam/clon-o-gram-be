package api

import (
	"github.com/gin-gonic/gin"
	
	"github.com/SideProjectTeam/clon-o-gram-be/api/controllers/users"
)

//SetRoutes is where you add all controllers to gin router
func SetRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/register",users.Register)
	r.POST("/login",users.Login)
	r.POST("/update_user",users.Update)
	return r
}