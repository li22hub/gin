package Api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"package/Database"
)

//写入redis
func SetRedis(ctx *gin.Context){
	redis,err := Database.RedisConnect()
	redis.Do("set","go","学go就图一乐,真功夫还得看php")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"code":http.StatusInternalServerError,
			"message":"插入redis失败",
		})
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"message":"插入redis成功",
	})
}
