package system

import (
	"ginchat/controller"
	"ginchat/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (_ *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	router := r.Group("/user")
	userController := new(controller.UserController)

	{
		router.POST("/register", userController.Register)
		router.POST("/login", userController.Login)
	}

	prviateRouter := router.Use(middleware.JWTAuth())
	{ // Prviate
		prviateRouter.GET("/testJWT", userController.TestJWT)
	}
}
