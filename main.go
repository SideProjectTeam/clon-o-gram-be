package main

import (
	"github.com/SideProjectTeam/clon-o-gram-be/api"
	"github.com/gin-gonic/gin"
	"github.com/SideProjectTeam/clon-o-gram-be/api/mngdb"
)

func main() {
	//start db
	mngdb.InitDB()
	
	r := gin.Default()
	api.SetRoutes(r)
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
