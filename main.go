package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/atur-uang/celengan/app"
	"github.com/atur-uang/celengan/framework"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var DB *gorm.DB

func main() {
	application := gin.Default()
	router := app.Routes(application)

	setupLogger()
	loadEnvironmentVariable()
	setDatabaseConnection()

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
	route.LoadHTMLGlob("resources/views/**/*")
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

func setDatabaseConnection() {

	config := framework.Config{}
	db := config.GetDatabaseConfiguration()
	name := db.Name
	fmt.Println(name)

	//database, dbErr := gorm.Open(mysql.Open(""), &gorm.Config{})
	//DB = database
	//if dbErr != nil {
	//	panic("failed to connect database")
	//}
	//
	//err := DB.AutoMigrate(&models.User{})
	//if err != nil {
	//	return
	//}

}