package controller

import (
	"ginchat/global"
	"ginchat/model"
	systemReq "ginchat/model/request"
	systemRes "ginchat/model/response"
	"ginchat/utils"
	"ginchat/utils/request"
	"ginchat/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
}

// Register
// @Tags     User
// @Schemes
// @Description 用户注册账号
// @Produce   application/json
// @Param    data      body      systemReq.Register          true  "name, password, confirmed password, email"
// @Success  200   {object}  response.Response{data=systemRes.UserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router /user/register [post]
func (_ *UserController) Register(c *gin.Context) {
	var params systemReq.Register
	var user model.UserBasic

	params.InitValidations()

	// 验证参数
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.Failed(c, http.StatusUnprocessableEntity, request.GetErrMsg(params, err))
		return
	}

	// 生成uuid
	u1, err := uuid.NewRandom()
	for err != nil {
		u1, err = uuid.NewRandom()
	}

	user = model.UserBasic{
		Name:     params.Name,
		Password: utils.BcryptHash(params.Password),
		Email:    params.Email,
		ClientIP: c.ClientIP(),
		UUID:     u1.String(),
	}

	// 添加到数据库
	err = global.MY_DB.Create(&user).Error
	if err != nil {
		response.Failed(c, http.StatusUnprocessableEntity, "注册失败")
		return
	}

	response.OK(c, systemRes.UserResponse{User: user})
}
