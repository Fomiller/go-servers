package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"data": "Hello"})
	})

	r.Run(":8080")
}
