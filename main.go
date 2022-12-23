package main

import (
	"context"
	"github.com/htd/app/demo"
	"github.com/htd/framework/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建engine结构
	core := gin.New()
	// 绑定具体的服务
	core.Bind(&demo.DemoServiceProvider{})

	core.Use(gin.Recovery())

	core.GET("/demo", DemoHander)

	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	// 这个goroutine是启动服务的goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}

func DemoHander(c *gin.Context) {
	demoService := c.MustMake(demo.Key).(demo.Service)
	foo := demoService.GetHello()
	c.ISetOkStatus().IJsonp(foo)
}
