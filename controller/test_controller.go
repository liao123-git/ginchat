package controller

import (
	"github.com/gin-gonic/gin"
)

type TestController struct {
}

// PingPong
// @Summary ping pong
// @Schemes
// @Description do ping
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ping [get]
func (_ *TestController) PingPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// resouce router controller
// func (_ *UserController) Index(c *gin.Context) {
// 	var users []*models.UserBasic

// 	global.MY_DB.Find(&users)

// 	response.Ok(c, gin.H{
// 		"message": "index",
// 		"users":   users,
// 	})
// }

// func (_ *UserController) Show(c *gin.Context) {
// 	response.Ok(c, gin.H{
// 		"message": "show",
// 		"id":      c.Param("id"),
// 	})
// }

// func (_ *UserController) Store(c *gin.Context) {
// 	response.Ok(c, gin.H{
// 		"message": "store",
// 	})
// }

// func (_ *UserController) Edit(c *gin.Context) {
// 	response.Ok(c, gin.H{
// 		"message": "edit",
// 		"id":      c.Param("id"),
// 	})
// }

// func (_ *UserController) Destroy(c *gin.Context) {
// 	response.Ok(c, gin.H{
// 		"message": "delete",
// 		"id":      c.Param("id"),
// 	})
// }
