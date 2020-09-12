package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication set all the url mappings and start the server
func StartApplication() {
	mapUrls()

	fmt.Println("about to start the application...")
	router.Run(":8080")
}
