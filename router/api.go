package router

import (
	"ginchat/middleware"
	"ginchat/router/system"

	"github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) {
	// 定义基本路由
	r := router.Group("/api")
	routerGroup := new(system.RouterGroup)

	// 公开组
	publicGroup := r.Group("")

	// 私密组
	prviateGroup := r.Group("")
	prviateGroup.Use(middleware.PrivateGroup())

	// 调用路由
	{
		routerGroup.InitSwagRouter(router)                        // swagger 路由
		routerGroup.InitResourceRouter(publicGroup, prviateGroup) // 资源路由

		routerGroup.InitUserRouter(publicGroup, prviateGroup)
	}
}
