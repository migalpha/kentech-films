package main

import (
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_setupServer(t *testing.T) {
	app := application{
		redis:    &redis.Client{},
		postgres: &gorm.DB{},
	}

	got := setupServer(app)

	assert.NotNil(t, got)
}
