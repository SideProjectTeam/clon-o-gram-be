package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SideProjectTeam/clon-o-gram-be/api/mngdb"
	"github.com/SideProjectTeam/clon-o-gram-be/api/router"
)

//Run is starter func for server
func Run() {
	mngdb.InitDB()
	fmt.Println("Listening [::]:3000")
	r := router.New()
	log.Fatal(http.ListenAndServe(":3000", r))
}
