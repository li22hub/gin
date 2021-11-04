package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"package/Config"
	"package/Router"
	"syscall"
	"time"
)

type HttpTransport struct {
	Engine   *gin.Engine
	StopHttp chan bool
}

//开启http服务
func Start() {
	httpTransport := NewHttpTransport()
	httpServer := httpTransport.HttpServer()
	handleSignals(httpServer)
}

//信号捕捉 优雅退出
func handleSignals(httpServer *http.Server) {
	signCh := make(chan os.Signal, 2)
	signal.Notify(signCh, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGKILL)
	exitCh := make(chan int)

	go func() {
		select {
		case s := <-signCh:
			stopHttpServer(httpServer)
			fmt.Println("捕捉到信号:", s)
			goto ExitProcess
		}

	ExitProcess:
		log.Println("Exit Service")
		exitCh <- 0
	}()

	code := <-exitCh
	os.Exit(code)
}

//停止http服务
func stopHttpServer(httpServer *http.Server) {
	fmt.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error("HttpServer Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Info("timeout of 2 seconds.")
	}
	log.Info("Server exiting")
}

//结构体赋值 类似PHP中实例化
func NewHttpTransport() *HttpTransport {
	return &HttpTransport{}
}

//http服务主方法
func (h *HttpTransport) HttpServer() *http.Server {

	cfg, err := Config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	h.Engine = gin.Default()
	r := Router.NewRoute(h.Engine)
	r.ApiRoutes()

	srv := &http.Server{
		Addr:    cfg.HttpAddr,
		Handler: h.Engine,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.WithFields(log.Fields{
		"address": cfg.HttpAddr,
	}).Info("api: Running HTTP server")
	return srv
}