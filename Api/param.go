package Api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	common "package/Common"
	"package/Database"
	"strconv"
	"time"
)

type UserData struct {
	Id       int    `gorm:"column:id" db:"id" form:"id"`
	Username string `gorm:"column:username" db:"username" form:"username"`
	Age      int    `gorm:"column:age" db:"age" form:"age"`
	Address  string `gorm:"column:address" db:"address" form:"address"`
	Time     int64  `gorm:"column:time" db:"time" form:"time"`
	Status   int    `gorm:"column:status" db:"status" form:"status"`
	IsDel    int    `gorm:"column:is_del" db:"is_del" form:"is_del"`
}

func TestStudy(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.DefaultQuery("id", ""))
	//NAME := ctx.Param("name")
	NAME := ctx.Query("name")
	ADD := ctx.PostForm("add")
	if ID == 0 || NAME == "" || ADD == "" {
		common.ResponseData(
			ctx, http.StatusOK, 404, "缺少参数!", "",
		)
		return
	}
	var where UserData
	where.Id = ID
	where.Username = NAME
	where.Address = ADD
	res, _ := FirstData(where)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "你很别致!",
		"data": res,
	})
}

func FirstData(where UserData) (UserData, error) {
	userTest := Database.GetMysql()
	var data UserData
	db := userTest.DB.Table("user_test")
	if where.Id > 0 {
		db = db.Where("id = ?", where.Id)
	}
	//if where.Username != ""{
	//	db = db.Where("user_name LIKE ?","%where.Username%")
	//	db = db.Where("name <> ?", "name")
	//	db = db.Where("name in (?)", []string{"name", "name 2"})
	//	db = db.Where("name = ? AND age >= ?", "name", "22")
	//	db = db.Where("updated_at > ?", lastWeek)
	//	db = db.Where("created_at BETWEEN ? AND ?", lastWeek, today)
	//}
	db = db.First(&data)
	err := db.Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func ConvertTime(utime uint64) string {
	timeLayout := "2006-01-02 15:04:05"
	format := time.Unix(int64(utime), 0).Format(timeLayout)
	return format
}
