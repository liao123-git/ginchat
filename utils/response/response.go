package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

var defaultMsg = map[int]string{
	http.StatusOK: "Success",

	http.StatusUnauthorized:        "User not logged in",
	http.StatusNotFound:            "Not Found",
	http.StatusUnprocessableEntity: "Data was wrong",
}

func Result(status int, data interface{}, c *gin.Context, msg []string) {
	returnMsg := defaultMsg[status]
	if len(msg) == 1 {
		returnMsg = msg[0]
	} else if returnMsg == "" {
		returnMsg = "something was wrong"
	}

	c.JSON(status, Response{
		data,
		returnMsg,
	})
}

func OK(c *gin.Context, data interface{}, msg ...string) {
	Result(http.StatusOK, data, c, msg)
}

func Failed(c *gin.Context, status int, msg ...string) {
	Result(status, map[string]interface{}{}, c, msg)
}
