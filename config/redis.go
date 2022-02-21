package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

var redisSettings RedisSettings

type RedisSettings struct {
	Host         string `envconfig:"REDIS_HOST" required:"true"`
	Port         string `envconfig:"REDIS_PORT" required:"true"`
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	PoolTimeout  time.Duration
}

func Redis() RedisSettings {
	return redisSettings
}

func InitializeRedis() error {
	if err := envconfig.Process("", &redisSettings); err != nil {
		return err
	}

	redisSettings.DialTimeout = 3 * time.Second
	redisSettings.ReadTimeout = 2 * time.Second
	redisSettings.WriteTimeout = 2 * time.Second
	redisSettings.PoolSize = 30
	redisSettings.PoolTimeout = 3 * time.Second
	return nil
}
