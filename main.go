package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/routes"
	"web_app/settings"
)

// Web开发脚手架
func main() {
	if len(os.Args) < 2 { // 命令行输入
		fmt.Println("请输入配置文件路径！")
		return
	}
	// 1.加载配置文件
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("init settings faild,err:%v\n", err)
		return
	}
	// 2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger faild,err:%v\n", err)
		return
	}
	defer zap.L().Sync()

	// 3.初始化mysql连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql faild,err:%v\n", err)
		return
	}
	defer mysql.Close()

	// 4.初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis faild,err:%v\n", err)
		return
	}
	defer redis.Close()

	// 5.注册路由
	r := routes.SetUp()
	// 6.启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()

	// 等待信号中断来优雅关闭服务器,5s超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill默认发送syscall.SIGTERM信号
	// ...
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内关闭（将未处理完的请求处理完再关闭），超过5秒就退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
