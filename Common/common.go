package Common

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseList struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	TotalCount int         `json:"totalCount"`
}

func ResponseData(ctx *gin.Context, httpCode int, code int, msg string, data interface{}) {
	var resData Response
	resData.Code = code
	resData.Msg = msg
	resData.Data = data
	ctx.JSON(httpCode, resData)
	return
}
