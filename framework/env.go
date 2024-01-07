package framework

import "github.com/spf13/viper"

type Env struct {
}

func (env Env) String(key string, initiate string) string {
	viper.SetDefault(key, initiate)
	return viper.GetString(key)
}

func (env Env) Integer(key string, initiate int) int {
	viper.SetDefault(key, initiate)
	return viper.GetInt(key)
}

func (env Env) Boolean(key string, initiate bool) bool {
	viper.SetDefault(key, initiate)
	return viper.GetBool(key)
}

func EnvString(key string, initiate string) string {
	return Env{}.String(key, initiate)
}

func EnvInt(key string, initiate int) int {
	return Env{}.Integer(key, initiate)
}
