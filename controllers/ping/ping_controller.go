package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping return "pong" as a response when a request come to /ping
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
