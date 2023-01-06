package systemRes

import "ginchat/model"

type UserResponse struct {
	User model.UserBasic `json:"user"`
}

// 定义 tag 对应的报错信息
var tagToMsg = map[string]string{
	"eqfield":        "The two passwords must be consistent",
	"aleadyRegister": "The email aleady registered",
	"isRegistered":   "This email has not been registered yet",
}

func GetUserErrMsg(tag string) string {
	msg, ok := tagToMsg[tag]

	if ok {
		return msg
	}

	return ""
}
