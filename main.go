package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//fmt.Println("hello world!")

	r := gin.Default()
	r.POST("/admin", func(c *gin.Context) {
		c.String(http.StatusOK,"hello world")
	})
	r.Run(":8888")
}
