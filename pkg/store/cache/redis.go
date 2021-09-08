package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/spf13/viper"
	"time"
)

var (
	cachePool map[string]*Redis
)

func init() {
	cachePool = make(map[string]*Redis, 0)
}

// NewRedis 获取redis连接
// @param confName redis配置名, eg: redis.local
// @return Redis
func NewRedis(confName, apiName string) *Redis {
	if conn, ok := cachePool[confName]; ok {
		return conn
	}

	var f Config
	if err := viper.UnmarshalKey(confName, &f); err != nil {
		log.Panicf("解析redis配置失败, 检查%s的配置, err: %s", confName, err)
	}

	pool, err := connectRedis(f)
	if err != nil {
		log.Panicf("连接redis失败, 检查%s的配置, err: %s", confName, err)
	}
	
	pre := viper.GetString("redis.pre")
	if pre == "" {
		panic("获取redis配置失败，检查redis.pre配置")
	}
	
	r := Redis{
		apiName:        apiName,
		pool:           pool,
		cachePrefixKey: fmt.Sprintf("%s:%s", pre, apiName),
	}

	cachePool[confName] = &r
	return &r
}

// connectRedis 连接Redis
// @param conf redis缓冲池配置
// @return *redis.Pool redis连接
// @return error
func connectRedis(conf Config) (*redis.Pool, error) {
	pool := redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(conf.Network, fmt.Sprintf("%s:%d", conf.Host, conf.Port))
			if err != nil {
				return nil, err
			}

			if conf.Password != "" {
				if _, err := c.Do("AUTH", conf.Password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", conf.DB); err != nil {
				c.Close()
				return nil, err
			}

			return c, nil
		},
		MaxIdle:         conf.MaxIdle,                       // 最大空闲连接数
		MaxActive:       conf.MaxActive,                     // 连接池最大数目,为0则不限制
		IdleTimeout:     conf.IdleTimeout * time.Second,     // 空闲连接超时时间，超过时间的空闲连接会被关闭,为0将不会被关闭,应该设置一个比redis服务端超时时间更短的时间
		Wait:            conf.Wait,                          // 如果为true且已经达到MaxActive的限制，则等待连接池
		MaxConnLifetime: conf.MaxConnLifeTime * time.Second, // 超过该时间关闭连接，如果为0则不根据时间来关闭连接
	}

	return &pool, nil
}

type Redis struct {
	apiName        string      // api名
	pool           *redis.Pool // redis 缓冲池
	cachePrefixKey string      // redis缓存前缀
}

// Do sends a command to the server and returns the received reply.
func (c *Redis) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn, err := c.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	reply, err = conn.Do(commandName, args...)
	return
}

// Send writes the command to the client's output buffer.
func (c *Redis) Send(commandName string, args ...interface{}) error {
	conn, err := c.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	return conn.Send(commandName, args...)
}

// Flush flushes the output buffer to the Redis server.
func (c *Redis) Flush() error {
	conn, err := c.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	return conn.Flush()
}

// Receive receives a single reply from the Redis server
func (c *Redis) Receive() (reply interface{}, err error) {
	conn, err := c.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return conn.Receive()
}

// utils

// GetCachePrefixKey 获取缓存前缀
// @return string
func (c *Redis) GetCachePrefixKey() string {
	return c.cachePrefixKey
}

// GetConn 获取 redis 连接
// @return redis.Conn
// @return error
func (c *Redis) GetConn() (redis.Conn, error) {
	conn := c.pool.Get()
	if err := conn.Err(); err != nil {
		return nil, err
	}

	return conn, nil
}

// 常用方法封装

// GetRandomExpire 获取随机缓存时间，单位: 秒
// @param re 过期时间结构体, 传nil则使用配置文件信息
// @return int64
func (c *Redis) GetRandomExpire(re *RandomExpire) int64 {
	if re == nil {
		re = &RandomExpire{
			Expire:    viper.GetDuration("redis.expire") * time.Hour,
			Deviation: viper.GetDuration("redis.deviation") * time.Hour,
		}
	}

	return re.RandomExpire()
}

// SetStringRandomExpire 缓存字符串，设置随机过期时间
// @param key 键名
// @param value 键值
// @param rx 过期时间结构体, 传nil则使用配置文件信息
// @return error
func (c *Redis) SetStringRandomExpire(key string, value string, re *RandomExpire) error  {
	conn, err := c.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	expire := c.GetRandomExpire(re)

	_, err = conn.Do("setex", key, expire, value)
	return err
}

// Del 批量删除键
// @param key 缓存字符串键
func (c *Redis) DelKeys(key ...interface{}) error {
	conn, err := c.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("del", key...)

	return err
}


