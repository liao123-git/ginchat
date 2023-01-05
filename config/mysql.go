package config

type Mysql struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址:端口
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             //:端口
	Dbname   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`       // 数据库名
	Account  string `mapstructure:"account" json:"account" yaml:"account"`    // 数据库用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	Config   string `mapstructure:"config" json:"config" yaml:"config"`       //全局表前缀，单独定义TableName则不生效
}

func (m *Mysql) Dsn() string {
	// "root:@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	return m.Account + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
