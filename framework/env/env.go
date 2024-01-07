package env

import "github.com/spf13/viper"

func Bool(key string, initiate bool) bool {
	viper.SetDefault(key, initiate)
	return viper.GetBool(key)
}

func Str(key string, initiate string) string {
	viper.SetDefault(key, initiate)
	return viper.GetString(key)
}

func Int(key string, initiate int) int {
	viper.SetDefault(key, initiate)
	return viper.GetInt(key)
}
