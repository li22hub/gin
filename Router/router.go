package Router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com!",
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/topgoer", helloHandler)
	return r
}