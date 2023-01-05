package core

import (
	"fmt"

	"ginchat/core/initialize"
	"ginchat/global"
)

func RunServer() {
	r := initialize.Router()

	info()

	// 运行
	r.Run(global.MY_CONFIG.System.Port)
}

func info() {
	welcome := "\n欢迎使用 ginchat\n"
	version := "当前版本: v0.9\n"
	auth := "作者: LDL\n"
	wechat := "微信号: LDLkukukaki\n"
	swagger := "默认自动化文档地址:http://127.0.0.1" + global.MY_CONFIG.System.Port + "/swagger/index.html\n"
	defaultAddr := "默认前端文件运行地址: http://127.0.0.1" + global.MY_CONFIG.System.Port + "\n"

	fmt.Println(welcome + version + auth + wechat + swagger + defaultAddr)
}
