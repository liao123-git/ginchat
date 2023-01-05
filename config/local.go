package config

type Local struct {
	Path      string `mapstructure:"path" json:"path" yaml:"path"`                   // 暴露的地址
	StorePath string `mapstructure:"store-path" json:"store-path" yaml:"store-path"` // 本地写入地址
}
