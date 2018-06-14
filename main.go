package main

import (
	"github.com/nEdAy/wx_attendance_api_server/router"
	"github.com/nEdAy/wx_attendance_api_server/config"
	"github.com/nEdAy/wx_attendance_api_server/model"
	"github.com/nEdAy/wx_attendance_api_server/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
	"context"
	"github.com/rs/zerolog/log"
)

func init() {
	// 初始化Logger
	logger.Setup()
	// 初始化Config
	config.Setup()
	// 初始化Database
	model.Setup()
	// 初始化Router
	router.Setup()
}

func main() {
	initGin()
	// 配置并启动Gin Server
	startGinServer()
}

func initGin() {
	// 配置Gin
	gin.SetMode(config.App.RunMode)
	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()
	// Logging to a file.
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
}

func startGinServer() {
	// Listen and Server in 127.0.0.1:8000
	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		Handler:        router.Router,
		ReadTimeout:    config.Server.ReadTimeout,
		WriteTimeout:   config.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		if config.Server.Protocol == "http" {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal().Msgf("http listen: %s\n", err)
			}
		} else {
			if err := server.ListenAndServeTLS(config.Path.CertFilePath, config.Path.KeyFilePath); err != nil && err != http.ErrServerClosed {
				log.Fatal().Msgf("https listen: %s\n", err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info().Msg("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Error().Msgf("Server Shutdown:", err)
	}
	log.Info().Msg("Server exiting")
}
