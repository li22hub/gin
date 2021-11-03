/**
所有服务开启或关闭再此调用
*/
package service

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"package/Api"
	"syscall"
	"time"
)

//开启http服务
func Start() {
	httpTransport := Api.NewHttpTransport()
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