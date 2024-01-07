package framework

import appConfig "github.com/atur-uang/celengan/config"

type Database struct {
	Driver   string
	Name     string
	Username string
	Password string
	Host     string
	Port     int
}

type Config struct {
}

func (config *Config) GetDatabaseConfiguration() Database {
	dbConfig := appConfig.DB()
	database := Database{}
	if dbDriver, ok := dbConfig["driver"].(string); ok {
		database.Driver = dbDriver
	}
	if dbHost, ok := dbConfig["host"].(string); ok {
		database.Host = dbHost
	}
	if dbPort, ok := dbConfig["port"].(int); ok {
		database.Port = dbPort
	}
	if dbUsername, ok := dbConfig["username"].(string); ok {
		database.Username = dbUsername
	}
	if dbPassword, ok := dbConfig["password"].(string); ok {
		database.Password = dbPassword
	}
	if dbName, ok := dbConfig["name"].(string); ok {
		database.Name = dbName
	}

	return database
}
