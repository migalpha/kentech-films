package redis

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func Test_TokenRepo_Save(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("could not initialize mini redis: %s", err.Error())
	}
	defer mr.Close()

	t.Run("Happy path", func(t *testing.T) {
		r := NewTokenRepository(redis.NewClient(&redis.Options{Addr: mr.Addr()}))

		err := r.Save(context.Background(), "token")

		assert.Nil(t, err)
	})
}

func Test_TokenRepo_IsTokenBlacklisted(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("could not initialize mini redis: %s", err.Error())
	}
	defer mr.Close()

	t.Run("Happy path - Not Blacklisted", func(t *testing.T) {
		r := NewTokenRepository(redis.NewClient(&redis.Options{Addr: mr.Addr()}))
		token := "test-token"
		ctx := context.Background()

		got, err := r.IsTokenBlacklisted(ctx, token)

		assert.Equal(t, false, got)
		assert.Nil(t, err)
	})
	t.Run("Happy path - Blacklisted", func(t *testing.T) {
		r := NewTokenRepository(redis.NewClient(&redis.Options{Addr: mr.Addr()}))
		token := "test-token"
		ctx := context.Background()

		r.Save(ctx, token)

		got, err := r.IsTokenBlacklisted(ctx, token)

		assert.Equal(t, true, got)
		assert.Nil(t, err)
	})
}
