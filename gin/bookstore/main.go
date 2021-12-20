package main

import (
	"github.com/fomiller/go-servers/gin/bookstore/controllers"
	"github.com/fomiller/go-servers/gin/bookstore/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("author/:name", controllers.FindBooksByAuthor)

	r.Run()
}
