package controller

import (
	"ginchat/utils/response"

	"github.com/gin-gonic/gin"
)

type ResourceController struct {
}

func (_ *ResourceController) Index(c *gin.Context) {

	response.OK(c, gin.H{
		"message": "index",
	})
}

func (_ *ResourceController) Show(c *gin.Context) {
	response.OK(c, gin.H{
		"message": "show",
		"id":      c.Param("id"),
	})
}

func (_ *ResourceController) Store(c *gin.Context) {
	response.OK(c, gin.H{
		"message": "store",
	})
}

func (_ *ResourceController) Edit(c *gin.Context) {
	response.OK(c, gin.H{
		"message": "edit",
		"id":      c.Param("id"),
	})
}

func (_ *ResourceController) Destroy(c *gin.Context) {
	response.OK(c, gin.H{
		"message": "delete",
		"id":      c.Param("id"),
	})
}
