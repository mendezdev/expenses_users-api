package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mendezdev/expenses_users-api/utils"
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

func routes(httpServices *httpServices) *gin.Engine {
	router := gin.Default()

	// PING HANDLERS
	router.GET("/ping", utils.Ping)

	// USERS HANDLERS
	router.POST("/users", httpServices.users.CreateUserHandler)
	router.GET("/users/:user_id", httpServices.users.GetUserByIDHandler)
	router.DELETE("/users/:user_id", httpServices.users.GetUserByIDHandler)

	return router
}
