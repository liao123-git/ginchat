package main

import (
	"ginchat/core"
	"ginchat/core/initialize"
	"ginchat/global"
)

func main() {
	global.MY_VIPER = core.InitViper()
	global.MY_DB = core.InitGorm()

	if global.MY_DB != nil {
		initialize.RegisterTables(global.MY_DB) // 初始化表

		// 程序结束前关闭数据库链接
		db, _ := global.MY_DB.DB()
		defer db.Close()
	}

	core.RunServer()
}
