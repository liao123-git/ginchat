package middleware

import (
	"ginchat/utils"
	"ginchat/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 用户重新登录账户，上次登录生成的 token 还有效 (暂未解决)
		// 所以要实现注销功能需要借助缓存
		// 或者每个用户使用不同的jwtSecret（签名密钥），每次生成新的tokenString钱都更换jwtSecret，这样旧的token就会验证失败

		token := c.Request.Header.Get("x-jwt-token")
		if token == "" {
			response.Failed(c, http.StatusUnprocessableEntity, "Please login first")
			c.Abort()
			return
		}

		claims, err := utils.ParseJWT(token)
		if err != nil {
			response.Failed(c, http.StatusUnprocessableEntity, err.Error())
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}

}
