package initialize

import (
	"ginchat/global"
	"ginchat/router"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// 暴露静态文件目录
	r.Static(global.MY_CONFIG.Local.Path, global.MY_CONFIG.Local.StorePath)

	// 调用 api
	router.Api(r)

	return r
}
