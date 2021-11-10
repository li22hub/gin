package Api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	common "package/Common"
	"package/Models"
	"strconv"
	"time"
)

//登录接口
func Login(ctx *gin.Context) {
	fmt.Println("hello world!!")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "小东西 还挺别致!",
	})
}

//用户列表
func GetUserList(ctx *gin.Context) {
	var resData common.ResponseList
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
	//fmt.Println(reflect.TypeOf(resData))
	//resData.TotalCount = len(resData)
	ctx.JSON(http.StatusOK, resData)
	return
}

//根据id查询一条用户信息
func GetUserOne(ctx *gin.Context) {
	var where Models.User
	var err error
	where.Id, err = strconv.Atoi(ctx.Query("id"))
	if err != nil || where.Id <= 0 {
		common.ResponseData(ctx, http.StatusOK, 1, "id参数错误", nil)
		fmt.Println(err) //打日志
		return
	}
	data, err := Models.UserListOne(where)
	if err != nil {
		fmt.Println(err)
		if err == gorm.ErrRecordNotFound {
			common.ResponseData(ctx, http.StatusOK, 1, "没有相关数据", nil) //查不到
			return
		}
		common.ResponseData(ctx, http.StatusOK, 1, "查询失败", nil)
		return
	}
	common.ResponseData(ctx, http.StatusOK, 0, "查询成功", data)
	return
}

//根据id更新一条数据
func UpdateUserOne(ctx *gin.Context) {
	var where Models.User
	var err error
	var data map[string]interface{}
	where.Id, err = strconv.Atoi(ctx.Query("id"))
	if err != nil || where.Id <= 0 {
		common.ResponseData(ctx, http.StatusOK, 1, "id参数缺少", nil)
		return
	}
	data = make(map[string]interface{})
	data["username"] = ctx.PostForm("username")
	data["age"], _ = strconv.Atoi(ctx.PostForm("age"))
	data["address"] = ctx.PostForm("address")
	data["time"] = time.Now().Unix()
	data["status"], _ = strconv.Atoi(ctx.PostForm("status"))
	err = Models.UpdateUserOne(where, data)
	if err != nil {
		common.ResponseData(ctx, http.StatusOK, 1, "更新失败", err)
		return
	}
	common.ResponseData(ctx, http.StatusOK, 0, "更新成功", struct{}{})
	return
}

//根据id删除一条数据
func DelUserOne(ctx *gin.Context) {
	var where Models.User
	var err error
	where.Id, err = strconv.Atoi(ctx.Query("id"))
	if err != nil || where.Id <= 0 {
		common.ResponseData(ctx, http.StatusOK, 1, "id参数缺少", nil)
		return
	}
	err = Models.DelUserOne(where)
	if err != nil {
		common.ResponseData(ctx, http.StatusOK, 1, "删除失败", err)
		return
	}
	common.ResponseData(ctx, http.StatusOK, 0, "删除成功", struct{}{})
	return
}

//插入一条数据
func AddUserOne(ctx *gin.Context) {
	var err error
	var data Models.User
	data.Username = ctx.PostForm("username")
	if data.Username == "" {
		common.ResponseData(ctx, http.StatusOK, 1, "参数username为空", err)
		return
	}
	data.Age, _ = strconv.Atoi(ctx.PostForm("age"))
	if data.Age <= 0 {
		common.ResponseData(ctx, http.StatusOK, 1, "参数age为空", err)
		return
	}
	data.Address = ctx.PostForm("address")
	if data.Address == "" {
		common.ResponseData(ctx, http.StatusOK, 1, "参数address为空", err)
		return
	}
	data.Time = time.Now().Unix()
	err = Models.AddUserOne(data)
	if err != nil {
		common.ResponseData(ctx, http.StatusOK, 1, "插入失败", err)
		return
	}
	common.ResponseData(ctx, http.StatusOK, 0, "插入成功", data)
	return
}

//批量插入数组数据
func AddUserList(ctx *gin.Context) {
	array := [...][3]string{
		{"雷欧", "26", "北京市南京西路步行街"},
		{"高斯", "27", "深圳市南京西路步行街"},
		{"戴拿", "25", "南京市南京西路步行街"},
	}
	var data Models.User
	var err error
	for _, v := range array {
		data.Username = v[0]
		data.Age, _ = strconv.Atoi(v[1])
		data.Address = v[2]
		data.Time = time.Now().Unix()
		data.Status = 0
		err = Models.AddUserList(data)
	}
	if err != nil {
		common.ResponseData(ctx, http.StatusOK, 1, "批量插入失败", err)
		return
	}
	common.ResponseData(ctx, http.StatusOK, 0, "批量插入成功", struct{}{})
	return
}
