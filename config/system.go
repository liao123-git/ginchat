package config

type System struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"` // 服务运行的端口
}
