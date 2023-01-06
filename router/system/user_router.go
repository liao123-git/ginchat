package system

import (
	"ginchat/controller"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (_ *UserRouter) InitUserRouter(p *gin.RouterGroup, l *gin.RouterGroup) {
	pUser := p.Group("/user")
	userController := new(controller.UserController)

	{ // Public
		pUser.POST("/register", userController.Register)
		pUser.POST("/login", userController.Login)
	}

	{ // Prviate

	}
}
