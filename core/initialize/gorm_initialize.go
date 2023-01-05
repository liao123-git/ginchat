package initialize

import (
	"fmt"
	"ginchat/model"
	"os"

	"gorm.io/gorm"
)

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.UserBasic{},
	)
	if err != nil {
		fmt.Println("Auto Migrate Failed")
		os.Exit(0)
	}

	fmt.Println("Auto Migrate Success")
}
