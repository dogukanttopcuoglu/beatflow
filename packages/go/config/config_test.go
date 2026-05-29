package config

import "testing"

type sampleConfig struct {
	Service sampleServiceConfig `mapstructure:"service"`
	HTTP    sampleHTTPConfig    `mapstructure:"http"`
}

type sampleServiceConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
}

type sampleHTTPConfig struct {
	Port int `mapstructure:"port"`
}

func TestLoadUsesDefaults(t *testing.T) {
	var cfg sampleConfig

	err := Load(&cfg, Options{
		EnvPrefix: "AUTH",
		Defaults: map[string]any{
			"service.name": "auth-service",
			"service.env":  "development",
			"http.port":    8080,
		},
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.Service.Name != "auth-service" {
		t.Fatalf("expected service name auth-service, got %q", cfg.Service.Name)
	}

	if cfg.Service.Env != "development" {
		t.Fatalf("expected service env development, got %q", cfg.Service.Env)
	}

	if cfg.HTTP.Port != 8080 {
		t.Fatalf("expected HTTP port 8080, got %d", cfg.HTTP.Port)
	}

}

func TestLoadOverridesDefaultsWithEnvironmentVariables(t *testing.T) {
	t.Setenv("AUTH_SERVICE_NAME", "auth-api")
	t.Setenv("AUTH_HTTP_PORT", "9090")

	var cfg sampleConfig

	err := Load(&cfg, Options{
		EnvPrefix: "AUTH",
		Defaults: map[string]any{
			"service.name": "auth-service",
			"service.env":  "development",
			"http.port":    8080,
		},
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.Service.Name != "auth-api" {
		t.Fatalf("expected service name auth-api, got %q", cfg.Service.Name)
	}

	if cfg.Service.Env != "development" {
		t.Fatalf("expected service env development, got %q", cfg.Service.Env)
	}

	if cfg.HTTP.Port != 9090 {
		t.Fatalf("expected HTTP port 9090, got %d", cfg.HTTP.Port)
	}
}

func TestLoadReturnsErrorForNilTarget(t *testing.T) {
	err := Load(nil, Options{})

	if err == nil {
		t.Fatal("expected error for nil target")
	}
}
