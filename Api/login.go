package Api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	models "package/Models"
	"time"
)

//测试接口
func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"message":"okok!",
	})
}

//登录接口
func Login(ctx *gin.Context) {
	fmt.Println("hello world!!")
	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"message":"小东西 还挺别致!",
	})
}

//创建用户
func CreateUser(ctx *gin.Context) {
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
