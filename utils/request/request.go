package request

import (
	"reflect"

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

		field, _ := reflect.TypeOf(u).FieldByName(fieldName) // 通过反射获取filed
		errMsg := field.Tag.Get(tag + "_err_msg")            // 定义自定义报错信息 例子：required_err_msg

		if errMsg != "" {
			return errMsg
		} else {
			return fieldName + " must be " + tag
		}
	}
	return ""
}
