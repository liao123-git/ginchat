package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func PublicGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before public group")
		c.Next()
		log.Println("after public group")
	}
}

func PrivateGroup() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before private group")
		c.Next()
		log.Println("after private group")
	}
}
