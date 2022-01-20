package Models

import (
	"fmt"
	"package/Database"
)

type ChinaAddress struct {
	Id        int    `gorm:"column:id" db:"id" form:"id"`
	Pid       int    `gorm:"column:pid" db:"pid" form:"pid"`
	Shortname string `gorm:"column:shortname" db:"shortname" form:"shortname"`
	Name      string `gorm:"column:name" db:"name" form:"name"`
	Mergename string `gorm:"column:mergename" db:"mergename" form:"mergename"`
	Level     int    `gorm:"column:level" db:"level" form:"level"`
	Pinyin    string `gorm:"column:pinyin" db:"pinyin" form:"pinyin"`
	Code      string `gorm:"column:code" db:"code" form:"code"`
	Zip       string `gorm:"column:zip" db:"zip" form:"zip"`
	First     string `gorm:"column:first" db:"first" form:"first"`
	Lng       string `gorm:"column:lng" db:"id" form:"lng"`
	Lat       string `gorm:"column:lat" db:"lat" form:"lat"`
}

func ListAddress(c []*ChinaAddress) ([]*ChinaAddress, int, error) {
	userChinaAddress := Database.GetMysql()
	var count int
	db := userChinaAddress.DB.Table("user_china_address")
	db = db.Find(&c).Count(&count)
	err := db.Error
	fmt.Println(err)
	if err != nil {
		defer userChinaAddress.DB.Close()
		return c, count, err
	}
	return c, count, nil
}
