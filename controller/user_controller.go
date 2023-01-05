package controller

import (
	"errors"
	"ginchat/global"
	"ginchat/model"
	systemReq "ginchat/model/request"
	systemRes "ginchat/model/response"
	"ginchat/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserController struct {
}

// Register
// @Tags     User
// @Schemes
// @Description 用户注册账号
// @Produce   application/json
// @Param    data      body      systemReq.Register          true  "name, password, email"
// @Success  200   {object}  response.Response{data=systemRes.UserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router /user/register [post]
func (_ *UserController) Register(c *gin.Context) {
	var params systemReq.Register
	var user model.UserBasic

	reqIP := c.ClientIP()
	err := c.ShouldBindJSON(&params)

	if err != nil {
		response.Failed(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if !errors.Is(global.MY_DB.Where("email = ?", params.Email).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		response.Failed(c, http.StatusUnauthorized, "该用户名已注册")
		return
	}

	// 生成uuid
	u1, err := uuid.NewRandom()
	for err != nil {
		u1, err = uuid.NewRandom()
	}

	user = model.UserBasic{
		Name:     params.Name,
		Password: params.Password,
		Email:    params.Email,
		ClientIP: reqIP,
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
