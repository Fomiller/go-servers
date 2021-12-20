package main

import (
	"net/http"

	"github.com/fomiller/go-servers/gin/jwtAuth1/controller"
	"github.com/fomiller/go-servers/gin/jwtAuth1/service"
	"github.com/gin-gonic/gin"
)

var (
	loginService    service.LoginService       = service.StaticLoginService()
	jwtService      service.JWTService         = service.JWTAuthService()
	loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)
)

func main() {

	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	r.Run(":8080")
}
