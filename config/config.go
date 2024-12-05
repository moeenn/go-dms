package config

import (
	"errors"
	"os"
)

type DatabaseConfig struct {
	ConnectionURI string
}

func NewDatabaseConfig() (*DatabaseConfig, error) {
	c := &DatabaseConfig{}

	c.ConnectionURI = os.Getenv("DB_CONNECTION")
	if c.ConnectionURI == "" {
		return c, errors.New("env variable not set: DB_CONNECTION")
	}

	return c, nil
}
