package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
	"github.com/nEdAy/wx_attendance_api_server/config"
	"github.com/nEdAy/wx_attendance_api_server/router"
	"net/http"
	"os/signal"
	"log"
	"time"
	"context"
	"github.com/nEdAy/wx_attendance_api_server/internal/db"
	"github.com/google/logger"
	"fmt"
)

func main() {
	// 初始化配置文件
	cfg, err := config.NewConfig("")
	if err != nil {
		logger.Fatalln("配置文件读取失败:", err)
	}
	js, _ := json.Marshal(cfg)
	log.Println(string(js))
	// 初始化google/logger输出到文件
	err = initLogger("face_server", "debug" == cfg.Mode)
	if err != nil {
		logger.Fatalln("日志初始化失败:", err)
	}

	// 初始化mysql
	err = db.InitDB(cfg.MySQL)
	if err != nil {
		logger.Fatalln("mysql连接错误:", err)
	}

	gin.SetMode(config.Conf.Mode)

	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()
	// Logging to a file.
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	r := router.SetupRouter()
	r.Static("/assets", "./assets")

	// Listen and Server in 127.0.0.1:443
	address := fmt.Sprintf("%s:%d", config.Conf.Http.Address, config.Conf.Http.Port)
	server := &http.Server{
		Addr:           address,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		if err := server.ListenAndServeTLS("./data/ssl/face.cer", "./data/ssl/face.key"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
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

// 获取log文件对象
func initLogger(name string, verbose bool) error {
	logPath := fmt.Sprintf(".%sdata%slog%s%s_%d.log", string(os.PathSeparator), string(os.PathSeparator), string(os.PathSeparator), name, time.Now().Unix())

	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return err
	}
	logger.Init(name, verbose, false, lf)

	return nil
}
