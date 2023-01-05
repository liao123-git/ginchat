package systemRes

import "ginchat/model"

type UserResponse struct {
	User model.UserBasic `json:"user"`
}
