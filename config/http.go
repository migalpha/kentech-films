package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

var httpSettings HTTPSettings

type HTTPSettings struct {
	Address           string `envconfig:"SERVER_ADDRESS" default:":8080"`
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
}

func HTTP() HTTPSettings {
	return httpSettings
}

func InitializeHttp() error {
	if err := envconfig.Process("", &httpSettings); err != nil {
		return err
	}

	httpSettings.ReadTimeout = 2 * time.Second
	httpSettings.ReadHeaderTimeout = 2 * time.Second
	httpSettings.WriteTimeout = 2 * time.Second

	return nil
}
