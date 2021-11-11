package Api

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math"
	"net/http"
	common "package/Common"
	"package/Database"
	"package/Models"
	"reflect"
	"strconv"
)

//测试redis
func TryRedis(ctx *gin.Context) {
	rs, _ := Database.RedisConnect()
	value, err := redis.String(rs.Do("get", "newK"))
	fmt.Println(reflect.TypeOf(value))
	if value != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "获取成功",
			"value":   json.RawMessage(value),
		})
	} else {
		//存key value值
		rs.Do("set", "golang", "学go就图一乐,真功夫还得看php")
		//存json
		k := "newK"
		vMap := map[string]string{
			"name":     "admin",
			"password": "girl-kill",
		}
		value, _ := json.Marshal(vMap)
		rs.Do("setnx", k, value)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "key不存在 插入redis成功",
			"err":     err,
		})
	}
}

//不定条件查询 加分页
func FindList(ctx *gin.Context) {
	var where Models.Result
	var err error
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	if page == 0 || pageSize == 0 {
		common.ResponseData(ctx,http.StatusInternalServerError,1,"缺少必填参数",nil)
	}
	where.Id, _ = strconv.Atoi(ctx.Query("id"))
	where.ExecuteCount, _ = strconv.Atoi(ctx.Query("execute_count"))
	where.IsDel, _ = strconv.Atoi(ctx.Query("is_del"))
	where.Video = ctx.Query("video")
	data, count, err := Models.FindList(where, page, pageSize)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			common.ResponseData(ctx, http.StatusOK, 1, "没有相关数据", nil) //查不到
			return
		}
		common.ResponseData(ctx, http.StatusOK, 1, "查询失败", nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":       http.StatusOK,
		"msg":        "查询成功",
		"data":       data,
		"page":       page,
		"pageSize":   pageSize,
		"pageCount":  math.Ceil(float64(count) / float64(pageSize)),
		"totalCount": count,
	})
	return
}
