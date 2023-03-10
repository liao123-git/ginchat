package model

import (
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name     string `json:"name" gorm:"required"`
	Password string `json:"-" gorm:"required"`
	Email    string `json:"email" gorm:"required,unique"`
	UUID     string `json:"uuid" gorm:"required,unique"`
	ClientIP string `json:"clientIP" gorm:"required"`
	// ClientPort    string `json:"clientPort" gorm:"required"`
	LoginTime     time.Time `json:"loginTime"`
	HeartbeatTime time.Time `json:"heartbeatTime"` // 心跳检测时间
	LogoutTime    time.Time `json:"logoutTime"`
	IsLogout      bool      `json:"isLogout" gorm:"default:true"`
	DeviceInfo    string    `json:"-"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
