package systemReq

import (
	"errors"
	"ginchat/global"
	"ginchat/model"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Register User register structure
type Register struct {
	Name          string `json:"name" example:"name" binding:"required"`
	Password      string `json:"password" example:"password" binding:"required"`
	ConfirmedPass string `json:"confirmed password" example:"password" binding:"required,eqfield=Password" eqfield_err_msg:"The two passwords must be consistent"`
	Email         string `json:"email" example:"ldl@qq.com" binding:"required,email,aleadyRegister" aleadyRegister_err_msg:"The email aleady registered"`
}

// 自定义验证器
func (_ *Register) InitValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("aleadyRegister", aleadyRegister)
	}
}

var aleadyRegister validator.Func = func(fl validator.FieldLevel) bool {
	email, ok := fl.Field().Interface().(string)

	if ok {
		// 判断用户名是否注册
		if !errors.Is(global.MY_DB.Where("email = ?", email).First(&model.UserBasic{}).Error, gorm.ErrRecordNotFound) {
			return false
		}
	}

	return true
}
