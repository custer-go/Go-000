package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	stop := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		return serveApp(stop)
	})

	g.Go(func() error {
		return serveSignal(stop)
	})

	if err := g.Wait(); err != nil {
		fmt.Println("g.Wait err: ", err)
	}
}

func serveApp(stop chan struct{}) error {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"result": "pong"})
		})
		v1.GET("/close", func(c *gin.Context) {
			close(stop)
			c.JSON(200, gin.H{"result": "close"})
		})
	}

	server := &http.Server{Addr: ":9090", Handler: router}
	done := make(chan error)
	go (func() { // 启动 HTTP
		if err := server.ListenAndServe(); err != nil {
			<-done
			logrus.Fatalf("服务启动失败: %s\n", err)
		}
	})()

	// 超时控制
	go (func() {
		<-stop

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		done <- server.Shutdown(ctx)
	})()

	return nil
}

func serveSignal(stop chan struct{}) error {
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stop:
		close(quit)
	case <-quit:
		close(stop)
	}

	signal.Reset(syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("正在平滑关闭服务器...")
	return nil
}
