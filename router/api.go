package router

import (
	"ginchat/router/system"

	"github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) {
	// 定义基本路由
	r := router.Group("/api")
	routerGroup := new(system.RouterGroup)

	// 调用路由
	{
		routerGroup.InitSwagRouter(router) // swagger 路由
		routerGroup.InitResourceRouter(r)  // 资源路由

		routerGroup.InitUserRouter(r)
	}
}
