package main

import (
	"autoplay-hub/dao/mysql"
	"autoplay-hub/dao/redis"
	"autoplay-hub/logger"
	"autoplay-hub/pkg/snowflake"
	"autoplay-hub/router"
	"autoplay-hub/settings"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// 1. 加载配置文件
	if err := settings.Init(); err != nil {
		zap.L().Error("init settings failed", zap.Error(err))
		return
	}
	// 2.初始化日志
	if err := logger.Init(settings.Conf.LoggerConfig); err != nil {
		zap.L().Error("init logger failed", zap.Error(err))
		return
	}
	zap.L().Info("init logger success")
	defer zap.L().Sync()
	// 3.初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		zap.L().Error("init mysql failed", zap.Error(err))
		return
	}
	zap.L().Info("init mysql success")
	defer mysql.Close()
	// 4.初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		zap.L().Error("init redis failed", zap.Error(err))
		return
	}
	zap.L().Info("init redis success")
	defer redis.Close()
	// 初始化雪花算法
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		zap.L().Error("init snowflake failed", zap.Error(err))
		return
	}
	zap.L().Info("init snowflake success")
	// 5. 初始化路由
	r := router.SetupRouter()
	// 6.启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}
	// 优雅关机的实现
	// 启动一个goroutine 防止阻塞后续的优雅关机逻辑
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Fatal("Listen failed", zap.Error(err))
		}
	}()
	fmt.Printf("Server running on 127.0.0.1:%d\n", settings.Conf.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}
}
