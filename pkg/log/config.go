package log

type Config struct {
	ServerName   string // 服务名
	Path         string // 日志目录地址
	LowLevel     Level  // 记录的最小级别：-1:debug、0:info、1:warn、2:error
	LogInConsole bool   // 打印到控制台
}

var logConf = Config{}
