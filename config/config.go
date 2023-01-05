package config

type Server struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`    // gorm mysql
	Local  Local  `mapstructure:"local" json:"local" yaml:"local"`    // local
	System System `mapstructure:"system" json:"system" yaml:"system"` // local
}
