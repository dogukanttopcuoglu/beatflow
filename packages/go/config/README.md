# BeatFlow Shared Config Package

This package provides shared configuration loading utilities for BeatFlow Go services.

It uses Viper to support default values, environment variable overrides, service-specific config structs, and environment variable prefixes.

## Purpose

Each BeatFlow service should own its own configuration struct.

This shared package does not define service-specific configuration fields. Instead, it provides a reusable `Load` function that can populate any service config struct from defaults and environment variables.

The goal is to keep configuration loading consistent across services while allowing each service to define its own config shape.

## Service-Specific Config Pattern

Each service should define its own config struct inside the owning service.

Example config struct for `auth-service`:

```go
type Config struct {
	Service ServiceConfig `mapstructure:"service"`
	HTTP    HTTPConfig    `mapstructure:"http"`
}

type ServiceConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
}

type HTTPConfig struct {
	Port int `mapstructure:"port"`
}
```

## Example Usage

```go
package main

import (
	"log"

	"github.com/dogukanttopcuoglu/beatflow/packages/go/config"
)

type Config struct {
	Service ServiceConfig `mapstructure:"service"`
	HTTP    HTTPConfig    `mapstructure:"http"`
}

type ServiceConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
}

type HTTPConfig struct {
	Port int `mapstructure:"port"`
}

func main() {
	var cfg Config

	err := config.Load(&cfg, config.Options{
		EnvPrefix: "AUTH",
		Defaults: map[string]any{
			"service.name": "auth-service",
			"service.env":  "development",
			"http.port":    8080,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting %s on port %d", cfg.Service.Name, cfg.HTTP.Port)
}
```

## Environment Variable Mapping

With this option:

```go
EnvPrefix: "AUTH"
```

Config keys map to environment variables like this:

```txt
service.name -> AUTH_SERVICE_NAME
service.env  -> AUTH_SERVICE_ENV
http.port    -> AUTH_HTTP_PORT
```

Environment variables override default values.

Example:

```bash
AUTH_SERVICE_NAME=auth-api
AUTH_SERVICE_ENV=local
AUTH_HTTP_PORT=9090
```

With these environment variables, the loaded config becomes:

```txt
cfg.Service.Name = "auth-api"
cfg.Service.Env  = "local"
cfg.HTTP.Port    = 9090
```

## Design Rule

This package should only contain technical configuration loading behavior.

Service-specific configuration structs should live inside the owning service.

For example:

```txt
apps/auth-service/internal/config
```

should define the Auth Service config struct, while:

```txt
packages/go/config
```

should only provide the shared loading utility.

## Current Scope

This initial version supports:

- isolated Viper instance per load call
- default values
- environment variable overrides
- environment variable prefixes
- nested config keys
- struct unmarshalling with `mapstructure`
- error handling for nil config targets

Future improvements may include:

- required config key validation
- config file support
- config value validation
- secret handling rules