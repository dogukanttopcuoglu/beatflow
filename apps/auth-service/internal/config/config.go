package config

import (
	sharedconfig "github.com/dogukanttopcuoglu/beatflow/packages/go/config"
)

type Config struct {
	Service ServiceConfig `mapstructure:"service"`
	Log     LogConfig     `mapstructure:"log"`
	HTTP    HTTPConfig    `mapstructure:"http"`
}

type ServiceConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}

type HTTPConfig struct {
	Port int `mapstructure:"port"`
}

func Load() (Config, error) {
	var cfg Config
	err := sharedconfig.Load(&cfg, sharedconfig.Options{
		EnvPrefix: "AUTH",
		Defaults: map[string]any{
			"service.name": "auth-service",
			"service.env":  "development",
			"log.level":    "debug",
			"http.port":    8080,
		},
	})

	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
