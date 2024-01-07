package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/atur-uang/celengan/app"
	"github.com/atur-uang/celengan/app/models"
	"github.com/atur-uang/celengan/framework"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gitlab.com/go-box/pongo2gin/v6"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
)

var DB *gorm.DB

func main() {
	application := gin.Default()
	router := app.Routes(application)

	setupLogger()
	loadEnvironmentVariable()
	loadConfiguration()

	loadViews(application)
	loadStaticFiles(application)
	runServer(router)
}

func loadEnvironmentVariable() {
	// Load environment variable from the .env file
	dotenvErr := godotenv.Load(".env")
	if dotenvErr != nil {
		// todo: log the error
	}

	viper.AutomaticEnv()

}

func setupLogger() {
	// Set the log to the file
	gin.ForceConsoleColor()
	f, _ := os.Create("storage/logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func loadStaticFiles(route *gin.Engine) {
	// Serving static file
	route.Static("/assets", "./resources/assets")
}

func loadViews(route *gin.Engine) {
	// Load the go html view
	//route.LoadHTMLGlob("resources/views/**/*")
	route.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: "resources/views",
		TemplateSet: nil,
		ContentType: "text/html; charset=utf-8",
	})
}

func runServer(router *gin.Engine) {
	server := &http.Server{Addr: ":8080", Handler: router}

	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

func loadConfiguration() {

	config := framework.Config{}

	setDatabaseConnection(config)
}

func setDatabaseConnection(config framework.Config) {
	var database *gorm.DB
	var dbErr error

	db := config.GetDatabaseConfiguration()
	host := db.Host
	username := db.Username
	password := db.Password
	port := string(rune(db.Port))
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

	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

}
