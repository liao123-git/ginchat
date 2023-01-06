package model

import (
	"errors"
	"ginchat/global"

	"gorm.io/gorm"
)

func IsExistInDB(query interface{}, args interface{}, m interface{}) bool {
	return !errors.Is(global.MY_DB.Where(query, args).First(m).Error, gorm.ErrRecordNotFound)
}
