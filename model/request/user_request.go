package systemReq

import (
	"errors"
	"ginchat/global"
	"ginchat/model"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// 注册 User 相关的结构体
type Register struct {
	Name          string `json:"name" example:"name" binding:"required"`
	Password      string `json:"password" example:"password" binding:"required"`
	ConfirmedPass string `json:"confirmed password" example:"password" binding:"required,eqfield=Password"`
	Email         string `json:"email" example:"ldl@qq.com" binding:"required,email,aleadyRegister"`
}

type Login struct {
	Email    string `json:"email" example:"ldl@qq.com" binding:"required,email,isRegistered"`
	Password string `json:"password" example:"password" binding:"required"`
}

const (
	RegisterMethod = "register"
	LoginMethod    = "login"
)

// 注册自定义User验证器
func RegisterUserValidations(method string) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		switch method {
		case RegisterMethod:
			v.RegisterValidation("aleadyRegister", aleadyRegister)
			break
		case LoginMethod:
			v.RegisterValidation("isRegistered", isRegistered)
			break
		}
	}
}

// 实现验证器
var aleadyRegister validator.Func = func(fl validator.FieldLevel) bool {
	if email, ok := fl.Field().Interface().(string); ok {
		// 判断 email 是否注册
		if !errors.Is(global.MY_DB.Where("email = ?", email).First(&model.UserBasic{}).Error, gorm.ErrRecordNotFound) {
			return false
		}
	}

	return true
}

var isRegistered validator.Func = func(fl validator.FieldLevel) bool {
	// 未注册返回 false ，所以跟上面反一下就okl
	return !aleadyRegister(fl)
}
