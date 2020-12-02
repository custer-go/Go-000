package main

import (
	"Week02/src/dbs"
	"Week02/src/handlers"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	dbs.InitDB() // 连接数据库

	router := gin.Default()
	v1 := router.Group("/v1/user")
	{
		v1.GET("", new(handlers.UserController).GetUserInfo)
	}

	server := &http.Server{Addr: ":9090", Handler: router}
	go (func() { // 启动 HTTP
		if err := server.ListenAndServe(); err != nil {
			logrus.Fatalf("服务启动失败: %s\n", err)
		}
	})()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Println("正在平滑关闭服务器...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatal("强制关闭服务器: ", err)
	}
	logrus.Println("服务器平滑关闭成功")
}
