package cache

import (
	"math/rand"
	"time"
)

type Config struct {
	Network         string        `mapstructure:"network"`
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Password        string        `mapstructure:"password"`
	DB              string        `mapstructure:"db"`
	MaxIdle         int           `mapstructure:"maxIdle"`
	MaxActive       int           `mapstructure:"maxActive"`
	IdleTimeout     time.Duration `mapstructure:"idleTimeout"`
	Wait            bool          `mapstructure:"wait"`
	MaxConnLifeTime time.Duration `mapstructure:"maxConnLifeTime"`
}

type RandomExpire struct {
	Expire    time.Duration // 过期时间
	Deviation time.Duration // 随机偏离范围
}

// RandomExpire 生成随机有效期时间，单位:秒
// @return {}
func (r *RandomExpire) RandomExpire() int64 {
	max := int64((r.Expire + r.Deviation) * time.Second)
	min := int64((r.Expire - r.Deviation)  * time.Second)

	return rand.Int63n(max-min) + min
}
