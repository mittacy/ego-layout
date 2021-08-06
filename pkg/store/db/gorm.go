package db

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	dbPool map[string]*gorm.DB
)

func init() {
	dbPool = make(map[string]*gorm.DB, 0)
}

// ConnectGorm 连接Mysql，获取gorm连接句柄
// @param name 数据库配置名
// @return *gorm.DB
func ConnectGorm(name string) *gorm.DB {
	key := fmt.Sprintf("%s.%s", MysqlDBPrefix, name)

	if db, ok := dbPool[key]; ok {
			return db
		}

	var dbConfig MysqlConf
	if err := viper.UnmarshalKey(key, &dbConfig); err != nil {
		zap.S().Panicf("连接数据库失败, 检查%s的配置, err: %s", key, err)
	}

	db, err := ConnectGormByConf(dbConfig)
	if err != nil {
		zap.S().Panicf("连接数据库失败, 检查%s的配置, err: %s", key, err)
	}

	dbPool[key] = db

	return db
}

// ConnectGormByConf 连接数据库
// @param conf 连接配置信息
// @return *gorm.DB
// @return error
func ConnectGormByConf(conf MysqlConf) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	if conf.Params != "" {
		dsn = fmt.Sprintf("%s?%s", dsn, conf.Params)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 是否禁用表名复数形式
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
