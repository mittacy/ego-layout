package bootstrap

import "github.com/mittacy/ego-layout/config"

func InitTask(confPath, env string) {
	// conf
	InitViper(confPath, env, 0)

	// log
	InitLog()

	// configs
	config.InitGorm()
	config.InitRedis()
}
