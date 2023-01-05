package core

import (
	"ginchat/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	m := global.MY_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}

	db, err := gorm.Open(mysql.Open(m.Dsn()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
