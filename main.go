package main

import (
	"os"
	"net/http"
	"os/signal"
	"time"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/nEdAy/wx_attendance_api_server/router"
	"github.com/nEdAy/wx_attendance_api_server/config"
	_ "github.com/nEdAy/wx_attendance_api_server/model"
)

func main() {
	// 配置Zap
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	// 配置Gin
	gin.SetMode(config.App.RunMode)
	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()
	// Logging to a file.
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// Listen and Server in 127.0.0.1:8000
	address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	server := &http.Server{
		Addr:           address,
		Handler:        router.Router,
		ReadTimeout:    config.Server.ReadTimeout,
		WriteTimeout:   config.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		if config.Server.Protocol == "http" {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				sugar.Fatalf("http listen: %s\n", err)
			}
		} else {
			if err := server.ListenAndServeTLS(config.Path.CertFilePath, config.Path.KeyFilePath); err != nil && err != http.ErrServerClosed {
				sugar.Fatalf("https listen: %s\n", err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	sugar.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		sugar.Fatal("Server Shutdown:", err)
	}
	sugar.Info("Server exiting")
}
