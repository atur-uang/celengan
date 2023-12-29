package framework

import appConfig "github.com/atur-uang/celengan/config"

type Database struct {
	Name     string
	Username string
	Password *string
	Host     string
	Port     int
}

type Config struct {
}

func (config *Config) GetDatabaseConfiguration() Database {
	dbConfig := appConfig.DB()
	database := Database{}
	if dbName, ok := dbConfig["name"].(string); ok {
		database.Name = dbName
	}

	return database
}
