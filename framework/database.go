package framework

import (
	"fmt"
	"github.com/atur-uang/celengan/app/models"
	"github.com/atur-uang/celengan/framework/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"strconv"
)

var DB *gorm.DB

type Database struct {
}

func (d Database) SetDatabaseConnection(config config.Config) {
	var database *gorm.DB
	var dbErr error

	db := config.GetDatabaseConfiguration()
	host := db.Host
	username := db.Username
	password := db.Password
	port := strconv.Itoa(db.Port)
	dbName := db.Name

	switch db.Driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
		database, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "sqlserver":
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", username, password, host, port, dbName)
		database, dbErr = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, username, password, dbName, port)
		database, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "sqlite":
		dsn := dbName
		database, dbErr = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	}

	DB = database
	if dbErr != nil {
		panic("failed to connect database")
	}

	err := DB.AutoMigrate(&models.User{}, &models.Vehicle{})
	if err != nil {
		return
	}
}

func GetDB() *gorm.DB {
	return DB
}
