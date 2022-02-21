package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/migalpha/kentech-films/config"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type application struct {
	// tracer metrics.Tracer

	redis    *redis.Client
	postgres *gorm.DB
}

func (app *application) setupDependencies() (err error) {
	config.Initialize()

	// err = app.setupTracer()
	// if err != nil {
	// 	return fmt.Errorf("can't setup tracer: %w", err)
	// }

	err = app.setupPostgres()
	if err != nil {
		return fmt.Errorf("can't setup postgres: %w", err)
	}

	err = app.setupRedis()
	if err != nil {
		return fmt.Errorf("can't setup redis: %w", err)
	}

	return nil
}

func (app *application) setupPostgres() error {
	sqlDB, err := sql.Open(config.Postgres().DriverName, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Postgres().Host,
		config.Postgres().Port,
		config.Postgres().User,
		config.Postgres().DBName,
		config.Postgres().Password,
		config.Postgres().SSLMode,
	))
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	app.postgres = gormDB

	return nil
}

func (app *application) setupRedis() error {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.Redis().Host, config.Redis().Port),
		DialTimeout:  config.Redis().DialTimeout,
		PoolSize:     config.Redis().PoolSize,
		PoolTimeout:  config.Redis().PoolTimeout,
		ReadTimeout:  config.Redis().ReadTimeout,
		WriteTimeout: config.Redis().WriteTimeout,
	})
	if redisClient == nil {
		return fmt.Errorf("can't create redis client")
	}

	if redisClient.Ping(context.Background()).Err() != nil {
		return fmt.Errorf("can't connect to redis")
	}

	app.redis = redisClient
	return nil
}
