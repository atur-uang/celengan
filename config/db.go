package config

import "github.com/spf13/viper"

func DB() map[string]any {
	viper.SetDefault("DB_NAME", "celengan")

	return map[string]any{
		"name": viper.GetString("DB_NAME"),
	}
}
