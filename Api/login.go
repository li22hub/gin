package Api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	common "package/Common"
	"package/Models"
)


//登录接口
func Login(ctx *gin.Context) {
	fmt.Println("hello world!!")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "小东西 还挺别致!",
	})
}

//一条用户信息
func GetUserOne(ctx *gin.Context) {
	var resData common.Response
	data := Models.User{}
	data, err := Models.UserListOne(data)
	resData.Data = data
	if err != nil {
		resData.Code = 500
		resData.Msg = fmt.Sprintf("%v", err)
		ctx.JSON(http.StatusInternalServerError, resData)
		return
	}
	resData.Code = 200
	resData.Msg = "操作成功"
	ctx.JSON(http.StatusOK, resData)
	return
}


//用户列表
func GetUserList(ctx *gin.Context) {
	var resData common.Response
	data := []*Models.User{}
	data, err := Models.UserList(data)
	resData.Data = data
	if err != nil {
		resData.Code = 500
		resData.Msg = fmt.Sprintf("%v", err)
		ctx.JSON(http.StatusInternalServerError, resData)
		return
	}
	resData.Code = 200
	resData.Msg = "操作成功"
	ctx.JSON(http.StatusOK, resData)
	return
}