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
// @Schemes 用户注册账号
// @Description 用户注册账号
// @Produce   application/json
// @Param    data      body      systemReq.Register          true  "name, password, confirmed password, email"
// @Success  200   {object}  response.Response{data=systemRes.UserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router /user/register [post]
func (_ *UserController) Register(c *gin.Context) {
	var params systemReq.Register
	var user model.UserBasic

	systemReq.RegisterUserValidations(systemReq.RegisterMethod)

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

// Login
// @Tags     User
// @Schemes 用户登陆账号
// @Description 用户登陆账号
// @Produce   application/json
// @Param    data      body      systemReq.Login          true  "email, password"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "用户登陆账号,返回包括用户信息"
// @Router /user/login [post]
func (_ *UserController) Login(c *gin.Context) {
	var params systemReq.Login
	var user model.UserBasic

	systemReq.RegisterUserValidations(systemReq.LoginMethod)

	// 验证参数
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.Failed(c, http.StatusUnprocessableEntity, request.GetErrMsg(params, err))
		return
	}

	// 验证密码
	global.MY_DB.Where("email = ?", params.Email).First(&user)
	if !utils.BcryptCheck(params.Password, user.Password) {
		response.Failed(c, http.StatusUnprocessableEntity, "Wrong email or password")
		return
	}

	// 获取token
	token, err := creatJWT(&user)
	if err != nil {
		response.Failed(c, http.StatusUnprocessableEntity, "Get token failed")
		return
	}

	response.OK(c, systemRes.LoginResponse{
		Name:  user.Name,
		Email: user.Email,
		UUID:  user.UUID,
		Token: token,
	})
}

func creatJWT(u *model.UserBasic) (string, error) {
	baseClaims := utils.BaseClaims{
		Name:  u.Name,
		Email: u.Email,
		UUID:  u.UUID,
	}

	myCustomClaims := utils.CreateClaims(baseClaims)

	return utils.CreateJWT(myCustomClaims)
}

// 测试 jwt
func (_ *UserController) TestJWT(c *gin.Context) {
	claims, has := c.Get("claims")
	if !has {
		response.Failed(c, http.StatusUnprocessableEntity, "Token failed")
		return
	}

	response.OK(c, gin.H{
		"claims": claims,
	})
}
