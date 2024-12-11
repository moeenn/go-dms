package config

import (
	"errors"
	"os"
)

var (
	ErrEnvDBConnectionNotSet = errors.New("env variable not set: DB_CONNECTION")
)

type Config struct {
	Database *DatabaseConfig
}

func NewConfig() (*Config, error) {
	dbConfig, err := NewDatabaseConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{
		Database: dbConfig,
	}

	return c, nil
}

type DatabaseConfig struct {
	ConnectionURI string
}

func NewDatabaseConfig() (*DatabaseConfig, error) {
	uri := os.Getenv("DB_CONNECTION")
	if uri == "" {
		return nil, ErrEnvDBConnectionNotSet
	}

	c := &DatabaseConfig{
		ConnectionURI: uri,
	}

	return c, nil
}
