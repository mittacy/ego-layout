package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

// NewClientByName 直接通过配置名字获取新客户端
// @param name 配置名
// @return *gorm.DB
// @return error
func NewClientByName(name string) *gorm.DB {
	if conf, ok := dbConfig[name]; ok {
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	if conf.Params != "" {
		dsn = fmt.Sprintf("%s?%s", dsn, conf.Params)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 是否禁用表名复数形式
	})
	if err != nil {
		log.Panicf("连接数据库失败, 检查配置, err: %s, conf: %+v", err, conf)
	}

	return db
}
