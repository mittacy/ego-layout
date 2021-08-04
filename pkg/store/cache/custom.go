package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/mittacy/ego-layout/pkg/config"
	"go.uber.org/zap"
	"math/rand"
)

type CustomRedis struct {
	apiName        string      // api名
	pool           *redis.Pool // redis 缓冲池
	cachePrefixKey string      // redis缓存前缀
	expireRange    ExpireRange // 缓存有效期范围
}

// GetCustomRedisByConf 获取 CustomRedis 连接
// @param name redis配置名
// @param apiName
// @return CustomRedigo
func GetCustomRedisByConf(name, apiName string) CustomRedis {
	pool := GetRedisPool(name)

	return GetCustomRedisByPool(pool, apiName)
}

// GetCustomRedisByPool 获取 CustomRedis 连接
// @param conn redis.Conn连接句柄
// @param apiName api名字，用户区分各个api类型的缓存，防止缓存键冲突
// @return CustomRedigo
func GetCustomRedisByPool(pool *redis.Pool, apiName string) CustomRedis {
	if _, err := pool.Get().Do("ping"); err != nil {
		zap.S().Panicf("连接redis失败, 检查redis配置, err: %s", err)
	}

	return CustomRedis{
		apiName:        apiName,
		pool:           pool,
		cachePrefixKey: fmt.Sprintf("%s:%s", config.Global.Server.Name, apiName),
		expireRange:    NewExpireRange(GlobalRedisConf.Expire, GlobalRedisConf.Deviation),
	}
}

type ExpireRange struct {
	MinSecond int64
	MaxSecond int64
}

// NewExpireRange 生成随机缓存的范围
// @param expire 缓存有效期, 单位: 小时
// @param deviation 随机偏移范围
// @return ExpireRange
func NewExpireRange(expire, deviation int64) ExpireRange {
	expireRange := ExpireRange{}

	// 初始化随机过期时间范围
	deviation = deviation * 3600
	expire = expire * 3600
	expireRange.MinSecond = expire - deviation
	expireRange.MaxSecond = expire + deviation

	return expireRange
}

// CachePrefixKey 获取缓存前缀
// @return string
func (c *CustomRedis) CachePrefixKey() string {
	return c.cachePrefixKey
}

// ExpireRange 获取缓存有效期范围
// @return ExpireRange
func (c *CustomRedis) ExpireRange() ExpireRange {
	return c.expireRange
}

// Do sends a command to the server and returns the received reply.
func (c *CustomRedis) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn, err := c.getConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	reply, err = conn.Do(commandName, args...)
	return
}

// Send writes the command to the client's output buffer.
func (c *CustomRedis) Send(commandName string, args ...interface{}) error {
	conn, err := c.getConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	return conn.Send(commandName, args...)
}

// Flush flushes the output buffer to the Redis server.
func (c *CustomRedis) Flush() error {
	conn, err := c.getConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	return conn.Flush()
}

// Receive receives a single reply from the Redis server
func (c *CustomRedis) Receive() (reply interface{}, err error) {
	conn, err := c.getConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return conn.Receive()
}

// Del 删除键
// @param key 缓存字符串键
func (c *CustomRedis) Del(key ...interface{}) error {
	conn, err := c.getConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("del", key...)

	return err
}

// CacheString 缓存字符串
// @param key 字符串键
// @param value 字符串值
func (c *CustomRedis) CacheString(key string, value string) error {
	conn, err := c.getConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	expire := c.RandomExpire()
	_, err = conn.Do("setex", key, expire, value)

	return err
}

// RandomExpire 生成随机有效期时间，单位秒
// @return int64 有效期时间
func (c *CustomRedis) RandomExpire() int64 {
	return rand.Int63n(c.expireRange.MaxSecond-c.expireRange.MinSecond) + c.expireRange.MinSecond
}

// getConn 获取 redis 连接
// @return redis.Conn
// @return error
func (c *CustomRedis) getConn() (redis.Conn, error) {
	conn := c.pool.Get()
	if err := conn.Err(); err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *CustomRedis) GetConn() (redis.Conn, error) {
	conn := c.pool.Get()
	if err := conn.Err(); err != nil {
		return nil, err
	}

	return conn, nil
}
