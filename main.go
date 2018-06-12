package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/nEdAy/wx_attendance_api_server/router"
	"net/http"
	"os/signal"
	"log"
	"time"
	"context"
	"fmt"
	"github.com/nEdAy/wx_attendance_api_server/config"
	"github.com/nEdAy/wx_attendance_api_server/model"
)

func main() {
	// 初始化Config
	config.Setup()
	// 初始化Database
	model.Setup()
	// 配置Gin
	gin.SetMode(config.App.RunMode)

	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()
	// Logging to a file.
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 配置Router
	r := router.SetupRouter()

	// Listen and Server in 127.0.0.1:8000
	address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	server := &http.Server{
		Addr:           address,
		Handler:        r,
		ReadTimeout:    config.Server.ReadTimeout,
		WriteTimeout:   config.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		if config.Server.Protocol == "http" {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("http listen: %s\n", err)
			}
		} else {
			if err := server.ListenAndServeTLS(config.Path.CertFilePath, config.Path.KeyFilePath); err != nil && err != http.ErrServerClosed {
				log.Fatalf("https listen: %s\n", err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
