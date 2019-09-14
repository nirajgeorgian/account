package api

import (
	"github.com/spf13/viper"
)

// Config :- base config structure for api package
type Config struct {
	// The port to bind the web application server to
	Port int

	// The number of proxies positioned in front of the API. This is used to interpret
	// X-Forwarded-For headers.
	ProxyCount int
}

// InitConfig :- default initialize function for config
func InitConfig() (*Config, error) {
	config := &Config{
		Port:       viper.GetInt("port"),
		ProxyCount: viper.GetInt("proxycount"),
	}

	if config.Port == 0 {
		config.Port = 3030
	}
	return config, nil
}
