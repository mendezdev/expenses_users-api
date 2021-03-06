package app

import (
	"github.com/mendezdev/expenses_users-api/controllers/ping"
	"github.com/mendezdev/expenses_users-api/controllers/users"
)

func mapUrls() {
	// Ping controller
	router.GET("/ping", ping.Ping)

	// Users controller
	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.POST("/users/login", users.Login)
}
