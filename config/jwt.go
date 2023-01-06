package config

type JWT struct {
	SiginKey string `mapstructure:"sigin-key" json:"sigin-key" yaml:"sigin-key"` // 暴露的地址
	Issuer   string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`          // 暴露的地址
}
