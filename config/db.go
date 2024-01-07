package config

import (
	"github.com/atur-uang/celengan/framework/env"
)

func DB() map[string]any {

	return map[string]any{
		"driver":   env.Str("DB_DRIVER", "mysql"),
		"host":     env.Str("DB_HOST", "172.0.0.1"),
		"port":     env.Int("DB_PORT", 5432),
		"username": env.Str("DB_USERNAME", "john"),
		"password": env.Str("DB_PASSWORD", "password"),
		"name":     env.Str("DB_NAME", "celengan"),
	}
}
