package controller

import (
	"github.com/gin-gonic/gin"
)

// 定义资源路由方法
type ResourceMethod interface {
	Index(c *gin.Context)   // GET /user
	Show(c *gin.Context)    // GET /user/id
	Store(c *gin.Context)   // POST /user
	Edit(c *gin.Context)    // PUT /user/id
	Destroy(c *gin.Context) // DELETE /user/id
}
