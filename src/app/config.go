package app

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config :- default config for app package
type Config struct {
	// A secret string used for session cookies, passwords, etc.
	SecretKey []byte
}

// InitConfig :- default initialize function for config
func InitConfig() (*Config, error) {
	config := &Config{
		SecretKey: []byte(viper.GetString("secretkey")),
	}

	if len(config.SecretKey) == 0 {
		return nil, fmt.Errorf("SecretKey must be set")
	}

	return config, nil
}
