package Api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"package/Config"
)

type HttpTransport struct {
	Engine   *gin.Engine
	StopHttp chan bool
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
	h.ApiRoutes()
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

//设置路由
func (h *HttpTransport) ApiRoutes() {
	h.Engine.GET("/ping", func(cxt *gin.Context) {
		cxt.JSON(http.StatusOK, gin.H{
			"code": 200,
			"info": "清风寨朱子明",
		})
	})

	v1 := h.Engine.Group("/v1")
	v1.GET("/login", h.Login)
	v1.POST("/create", h.createUser)
}
