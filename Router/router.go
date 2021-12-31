package Router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"package/Api"
	"package/MiddleWare"
)

type route struct {
	Engine *gin.Engine
}

func NewRoute(Engine *gin.Engine) *route {
	return &route{
		Engine: Engine,
	}
}

//设置路由
func (r *route) ApiRoutes() {
	r.Engine.GET("/ping", func(cxt *gin.Context) {
		cxt.JSON(http.StatusOK, gin.H{
			"code": 200,
			"info": "清风寨朱子明",
		})
	})

	//加载中间件
	r.Engine.Use(MiddleWare.MiddleWare())
	{
		r.Engine.GET("/ce", func(c *gin.Context) {
			//获取中间件的变量值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			//接口返回值
			c.JSON(200, gin.H{"request": req})
		})
	}

	v1 := r.Engine.Group("/v1")
	v1.GET("/login", Api.Login)

	v0 := r.Engine
	v0.GET("/test", Api.Test)
	v0.GET("/tryRedis", Api.TryRedis)
	v0.GET("/hello", Api.Hello)

	v2 := r.Engine
	v2.POST("/GetUserList", Api.GetUserList)
	v2.POST("/GetUserOne", Api.GetUserOne)
	v2.POST("/UpdateUserOne", Api.UpdateUserOne)
	v2.POST("/DelUserOne", Api.DelUserOne)
	v2.POST("/AddUserOne", Api.AddUserOne)
	v2.POST("/AddUserList", Api.AddUserList)
	v2.POST("/UpFile", Api.UpFile)
	v2.POST("/FindList", Api.FindList)
	v2.POST("/FindAll", Api.FindAll)
	v2.POST("/PingMQsend", Api.PingMQsend)
	v2.POST("/PingMQreceive", Api.PingMQreceive)
	v2.POST("/ChannelT", Api.ChannelT)
}
