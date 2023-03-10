package system

import (
	"ginchat/controller"

	"github.com/gin-gonic/gin"
)

type ResourceRouter struct {
}

func (_ *ResourceRouter) InitResourceRouter(r *gin.RouterGroup) {
	resourceController := new(controller.ResourceController)

	{ // Public
		resource(r, "/resource", resourceController)
	}

	{ // Prviate
		// resource(l, "/user", resourceController.UserController)
	}
}

func resource(r *gin.RouterGroup, path string, c controller.ResourceMethod) {
	// 获取全部
	r.GET(path, c.Index)

	// 获取单个
	r.GET(path+"/:id", c.Show)

	// 添加
	r.POST(path, c.Store)

	// 更新
	r.PUT(path+"/:id", c.Edit)

	// 删除
	r.DELETE(path+"/:id", c.Destroy)
}
