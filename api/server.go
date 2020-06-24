package api

import (
	"github.com/SideProjectTeam/clon-o-gram-be/api/router"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	fmt.Println("Listening [::]:3000")
	r := router.New()
	log.Fatal(http.ListenAndServe(":3000", r))
}
