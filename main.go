package main

import (
	"context"
	"errors"
	"github.com/atur-uang/celengan/app"
	"github.com/atur-uang/celengan/framework"
	"github.com/atur-uang/celengan/framework/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gitlab.com/go-box/pongo2gin/v6"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	application := gin.Default()

	// Middleware
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "https://celengan.online", "https://www.celengan.online"}
	corsConfig.AllowCredentials = true
	application.Use(cors.New(corsConfig))
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
	logPath := "/var/log/celengan/gin.log"
	var f *os.File

	if _, err := os.Stat(logPath); err == nil {
		f, _ = os.Open(logPath)
	} else {
		f, _ = os.Create(logPath)
	}

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

	c := config.Config{}
	setDatabaseConnection(c)
}

func setDatabaseConnection(config config.Config) {
	db := framework.Database{}
	db.SetDatabaseConnection(config)
}
