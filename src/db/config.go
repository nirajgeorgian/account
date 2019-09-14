package db

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config :- default config for db package
type Config struct {
	DatabaseURI string
}

// InitConfig :- default initialize function for config
func InitConfig() (*Config, error) {
	config := &Config{
		DatabaseURI: viper.GetString("databaseuri"),
	}

	if config.DatabaseURI == "" {
		return nil, fmt.Errorf("DatabaseURI must be set")
	}

	return config, nil
}
