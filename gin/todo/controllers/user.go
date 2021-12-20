package controller

import (
	"net/http"

	"github.com/fomiller/go-servers/gin/todo/models"
	"github.com/fomiller/go-servers/gin/todo/store"
	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func FindUser(c *gin.Context) {
	var user models.User
	err := store.DB.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil {
		handleError(c, err)
		return
	}
}

func FindUsers() {

}

func CreateUser() {

}

func UpdateUser() {

}

func DeleteUser() {

}

func LoginUser() {

}

func LogoutUser() {

}
