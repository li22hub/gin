package Api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"package/Config"
	models "package/Models"
	"time"
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

//登录接口
func (h *HttpTransport) Login(ctx *gin.Context) {
	fmt.Println("hello world!!")
	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"message":"小东西 还挺别致!",
	})
}

//创建用户
func (h *HttpTransport) createUser(ctx *gin.Context) {
	userModel := models.NewUser()
	userModel.Username = "admin"
	userModel.Age = 26
	userModel.Address = ctx.PostForm("address")
	userModel.Time = time.Now().Unix()
	userModel.Status = 0
	userModel.IsDel = 0

	data, err := userModel.Create()
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"info": "create user failed!!",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
		"info": "create user success!!",
	})
}