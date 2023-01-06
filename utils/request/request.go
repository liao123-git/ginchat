package request

import (
	systemRes "ginchat/model/response"

	"github.com/go-playground/validator/v10"
)

// 获取自定义报错
func GetErrMsg(u interface{}, err error) string {

	invalid, ok := err.(*validator.InvalidValidationError) // 如果是输入参数无效，则直接返回输入参数错误
	if ok {
		return "input param wrong: " + invalid.Error()
	}

	validationErrs := err.(validator.ValidationErrors) // 断言是ValidationErrors
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field() // 获取是哪个字段不符合格式
		tag := validationErr.Tag()         // 获取是哪个tag不符合格式

		errMsg := systemRes.GetUserErrMsg(tag) // tag 转报错信息

		if errMsg != "" {
			return errMsg
		} else {
			return fieldName + " must be " + tag
		}
	}
	return ""
}
