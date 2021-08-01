package config

import "time"

const (
	RedisConnPrefix = "redis."
)

type Redis struct {
	Expire    int64 `mapstructure:"expire"`
	Deviation int64 `mapstructure:"deviation"`
}

type RedisConfig struct {
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
