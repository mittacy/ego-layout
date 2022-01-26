package mysql

import (
	"fmt"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"time"
)

var (
	dbPool     map[string]*gorm.DB
	gormLogger zapgorm2.Logger // gorm日志句柄
)

func init() {
	dbPool = make(map[string]*gorm.DB, 0)

	l := log.New("gorm")
	gormLogger = zapgorm2.New(l.GetZap())

	if viper.GetDuration("GORM_SLOW_LOG_THRESHOLD") == 0 {
		gormLogger.SlowThreshold = time.Millisecond * 100
	} else {
		gormLogger.SlowThreshold = viper.GetDuration("GORM_SLOW_LOG_THRESHOLD") * time.Millisecond
	}
	gormLogger.LogLevel = logger.Info
	gormLogger.IgnoreRecordNotFoundError = true
	gormLogger.SetAsDefault()
}

// NewClientByName 直接通过配置名字获取新客户端
// @param name 配置名
// @return *gorm.DB
// @return error
func NewClientByName(name string) *gorm.DB {
	if conf, ok := GetConfig(name); ok {
		return NewClient(conf)
	}

	log.Panicf("%s 配置不存在, 请在 pkg/mysql/config.go GetConfig() 中配置", name)
	return nil
}

// NewClient 获取新客户端
// @param conf 配置信息
// @return *gorm.DB gorm连接
// @return error
func NewClient(conf Conf) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	if conf.Params != "" {
		dsn = fmt.Sprintf("%s&%s", dsn, conf.Params)
	}

	if db, ok := dbPool[dsn]; ok {
		return db
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 是否禁用表名复数形式
		Logger:         gormLogger,
	})
	if err != nil {
		log.Panicf("连接数据库失败, 检查配置, err: %s, conf: %+v", err, conf)
	}

	dbPool[dsn] = db

	return db
}
