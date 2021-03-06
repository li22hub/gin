package Models

import (
	"encoding/json"
	"fmt"
	"package/Database"
)

//定义表结构结构体
type TestDetail struct {
	Id       int             `gorm:"column:id" db:"id" form:"id"`
	TestId   int             `gorm:"column:test_id" db:"test_id" form:"test_id" `
	Name     string          `gorm:"column:name" db:"name" form:"name" `
	Data     json.RawMessage `gorm:"column:data" db:"data" form:"data" `
	Time     int64           `gorm:"column:time" db:"time" form:"time" `
	DataTime string          `gorm:"column:datatime" db:"datatime" form:"datatime" `
}

//获取detail表数据
func List(t []*TestDetail) ([]*TestDetail, int, error) {
	userTestDetail := Database.GetMysql()
	var count int
	db := userTestDetail.DB.Table("user_test_detail")
	db = db.Limit(20).Find(&t).Count(&count)
	err := db.Error
	fmt.Println(err)
	if err != nil {
		defer userTestDetail.DB.Close()
		return t, count, err
	}
	return t, count, nil
}

