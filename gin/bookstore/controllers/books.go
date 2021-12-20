package controllers

import (
	"net/http"

	"github.com/fomiller/go-servers/gin/bookstore/models"
	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.IndentedJSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		handleError(c, err)
		return
	}

	// Create Book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.IndentedJSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	var book models.Book
	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		handleError(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": book})
}

func FindBooksByAuthor(c *gin.Context) {
	var books []models.Book
	err := models.DB.Where("author = ?", c.Param("name")).Find(&books).Error
	if err != nil {
		handleError(c, err)
		return
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{"data": books})
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		handleError(c, err)
		return
	}

	var input models.UpdateBookInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		handleError(c, err)
		return
	}

	models.DB.Model(&book).Updates(input)
	c.IndentedJSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		handleError(c, err)
		return
	}

	models.DB.Delete(&book)
	c.IndentedJSON(http.StatusOK, gin.H{"deleted": book})
}
