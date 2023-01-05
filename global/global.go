package global

import (
	"ginchat/config"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	MY_DB     *gorm.DB
	MY_CONFIG config.Server
	MY_VIPER  *viper.Viper

	lock sync.RWMutex
)
