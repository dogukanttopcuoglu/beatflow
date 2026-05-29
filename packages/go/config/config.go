package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Options struct {
	EnvPrefix string
	Defaults  map[string]any
}

func Load(target any, opts Options) error {
	if target == nil {
		return fmt.Errorf("config target cannot be nil")
	}

	v := viper.New()

	if opts.EnvPrefix != "" {
		v.SetEnvPrefix(opts.EnvPrefix)
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	v.AutomaticEnv()

	for key, val := range opts.Defaults {
		v.SetDefault(key, val)

		if err := v.BindEnv(key); err != nil {
			return fmt.Errorf("bind evn for key %q: %w", key, err)
		}
	}

	if err := v.Unmarshal(target); err != nil {
		return fmt.Errorf("unmarshal config: %w", err)
	}
	return nil
}
