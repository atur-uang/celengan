package config

import (
	"github.com/atur-uang/celengan/framework"
)

func DB() map[string]any {

	return map[string]any{
		"driver":   framework.EnvString("DB_DRIVER", "mysql"),
		"host":     framework.EnvString("DB_HOST", "172.0.0.1"),
		"port":     framework.EnvInt("DB_PORT", 5432),
		"username": framework.EnvString("DB_USERNAME", "john"),
		"password": framework.EnvString("DB_PASSWORD", "password"),
		"name":     framework.EnvString("DB_NAME", "celengan"),
	}
}
