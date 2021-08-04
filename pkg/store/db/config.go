package db

const (
	MysqlDBPrefix = "mysql"
)

type MysqlConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Params   string `mapstructure:"params"`
}
