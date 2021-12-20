package controller

import (
	"github.com/fomiller/go-servers/gin/jwtAuth1/dto"
	"github.com/fomiller/go-servers/gin/jwtAuth1/service"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func LoginHandler(loginService service.LoginService,
	jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthorized := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthorized {
		return controller.jwtService.GenerateToken(credential.Email, true)
	}
	return ""
}
