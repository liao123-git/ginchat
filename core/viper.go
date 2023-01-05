package core

import (
	"fmt"
	"ginchat/global"

	"github.com/spf13/viper"
)

func InitViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")
	err := v.ReadInConfig() // 读取配置文件
	if err != nil {         // 报错处理
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.WatchConfig()

	if err = v.Unmarshal(&global.MY_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
