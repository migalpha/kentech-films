package config

import (
	"github.com/kelseyhightower/envconfig"
)

var postgresSettings PostgresSettings

type PostgresSettings struct {
	Host       string `envconfig:"POSTGRES_HOST" required:"true"`
	Port       string `envconfig:"POSTGRES_PORT" required:"true"`
	User       string `envconfig:"POSTGRES_USER" required:"true"`
	Password   string `envconfig:"POSTGRES_PASS" required:"true"`
	DBName     string `envconfig:"POSTGRES_DBNAME" required:"true"`
	DriverName string
	SSLMode    string
}

func Postgres() PostgresSettings {
	return postgresSettings
}

func InitializePostgres() error {
	if err := envconfig.Process("", &postgresSettings); err != nil {
		return err
	}

	postgresSettings.DriverName = "postgres"
	postgresSettings.SSLMode = "disable"
	return nil
}
