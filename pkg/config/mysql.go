package config

const (
	MysqlDBPrefix = "mysql."
)

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Params   string `mapstructure:"params"`
}
