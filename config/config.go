package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var c commons

type commons struct {
	AppEnv     string `envconfig:"APP_ENV" default:"development"`
	AppName    string `envconfig:"APP_NAME" default:"films"`
	Port       string `envconfig:"PORT" default:"8080"`
	Host       string `envconfig:"HOST" default:"0.0.0.0"`
	JWTSecret  string `envconfig:"JWT_SECRET" required:"true"`
	DateFormat string
	BCryptCost int
}

func Commons() commons {
	return c
}

func Initialize() {

	if err := envconfig.Process("", &c); err != nil {
		log.Fatal(err.Error())
	}
	c.DateFormat = time.RFC3339
	c.BCryptCost = 14

	if err := InitializeHttp(); err != nil {
		log.Fatal(err.Error())
	}

	if err := InitializePostgres(); err != nil {
		log.Fatal(err.Error())
	}

	if err := InitializeRedis(); err != nil {
		log.Fatal(err.Error())
	}
}
