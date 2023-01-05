package system

import (
	"ginchat/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwagRouter struct {
}

func (_ *SwagRouter) InitSwagRouter(p *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api"

	p.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
