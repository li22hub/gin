package Router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"package/Api"
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

	v1 := r.Engine.Group("/v1")
	v1.GET("/login", Api.Login)


	v2 := r.Engine
	v2.POST("/GetUserOne",Api.GetUserOne)
	v2.POST("/GetUserList",Api.GetUserList)
}