package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mendezdev/expenses_users-api/controllers/ping"
)

func mapUrls() {
	// Ping controller

	// Users controller
	/*router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.POST("/users/login", users.Login)*/
}

func routes(userService *httpServices) *gin.Engine {
	router := gin.Default()

	// PING HANDLERS
	router.GET("/ping", ping.Ping)

	// USERS HANDLERS
	router.POST("/users", userService.userHTTPService.CreateUserHandler)
	router.GET("/users/:user_id", userService.userHTTPService.GetUserByIDHandler)
	router.DELETE("/users/:user_id", userService.userHTTPService.GetUserByIDHandler)

	return router
}
