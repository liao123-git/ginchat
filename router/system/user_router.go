package system

import (
	"ginchat/controller"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (_ *UserRouter) InitUserRouter(p *gin.RouterGroup, l *gin.RouterGroup) {
	pUser := p.Group("/user")
	ctl := new(controller.Controller).UserController

	{ // Public
		pUser.POST("/register", ctl.Register)
	}

	{ // Prviate

	}
}
