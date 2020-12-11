package app

import (
	"fmt"

	"github.com/mendezdev/expenses_users-api/db/mongodb"
	users "github.com/mendezdev/expenses_users-api/users/web"
)

type httpServices struct {
	users users.UserHTTPService
}

//StartApplication set all the url mappings and start the server
func StartApplication() {
	db := mongodb.ConnectToDB("mongodb://localhost:27017")
	userService := users.NewUserHTTPService(db)
	services := &httpServices{
		users: userService,
	}

	router := routes(services)

	fmt.Println("about to start the application...")
	router.Run(":8080")
}
