package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	defaultExpiration = 24 * time.Hour
)

type TokenRepo struct {
	DB *redis.Client
}

func (repo TokenRepo) IsTokenBlacklisted(ctx context.Context, key string) (bool, error) {
	_, err := repo.DB.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return true, err
	}
	return true, nil
}

func (repo TokenRepo) Save(ctx context.Context, key string) error {
	err := repo.DB.Set(ctx, key, 1, defaultExpiration).Err()
	if err != nil {
		return fmt.Errorf("[TokenRepo:Save][err:%w]", err)
	}
	return nil

}
