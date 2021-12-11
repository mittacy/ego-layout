package mysql

import (
	"fmt"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"sync"
	"time"
)

var (
	dbPool     map[string]*gorm.DB // 因为gorm内部维护了连接池，所以根据mysql dsn判断，如果dsn相同直接返回同一个连接即可
	logInit    sync.Once           // 初始一次日志文件
	gormLogger zapgorm2.Logger     // gorm日志句柄
)

func init() {
	dbPool = make(map[string]*gorm.DB, 0)
}

// NewClientByName 直接通过配置名字获取新客户端
// @param name 配置名
// @return *gorm.DB
// @return error
func NewClientByName(name string) *gorm.DB {
	if conf, ok := GetConfig(name); ok {
		return NewClient(conf)
	}

	log.Panicf("%s 配置不存在, 请在 dbConfig 中配置", name)
	return nil
}

// NewClient 获取新客户端
// @param conf 配置信息
// @return *gorm.DB gorm连接
// @return error
func NewClient(conf Conf) *gorm.DB {
	// 初始化日志句柄
	logInit.Do(func() {
		slowThreshold := viper.GetDuration("GORM_SLOW_LOG_THRESHOLD") * time.Millisecond
		if slowThreshold == 0 {
			slowThreshold = time.Millisecond * 100
		}

		l := log.New("gorm")
		gormLogger = zapgorm2.New(l.GetZap())
		gormLogger.SlowThreshold = slowThreshold
		gormLogger.SetAsDefault()
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	if conf.Params != "" {
		dsn = fmt.Sprintf("%s?%s", dsn, conf.Params)
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
