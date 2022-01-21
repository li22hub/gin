package Models

import (
	"fmt"
	"package/Database"
)

// TreeList 菜单
type TreeList struct {
	Id        int
	Pid       int
	Shortname string
	Name      string
	Mergename string
	Level     int
	Pinyin    string
	Code      string
	Zip       string
	First     string
	Lng       string
	Lat       string
	Children  []TreeList
}

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
	Children  []ChinaAddress
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

// FormMenu 格式化菜单
func FormMenu(list []ChinaAddress, pid int) (formMenu []ChinaAddress) {
	for _, val := range list {
		if val.Pid == pid {
			if pid == 0 {
				// 顶层
				formMenu = append(formMenu, val)
			} else {
				var children []ChinaAddress
				child := val
				children = append(children, child)
			}
		}
	}
	return
}

// GetMenu 获取菜单
func GetMenu(menuList []*ChinaAddress, pid int) []TreeList {
	treeList := []TreeList{}
	for _, v := range menuList {
		if v.Pid == pid {
			child := GetMenu(menuList, v.Id)
			node := TreeList{
				Id:        v.Id,
				Pid:       v.Pid,
				Shortname: v.Shortname,
				Name:      v.Name,
				Mergename: v.Mergename,
				Level:     v.Level,
				Pinyin:    v.Pinyin,
				Code:      v.Code,
				Zip:       v.Zip,
				First:     v.First,
				Lng:       v.Lng,
				Lat:       v.Lat,
			}
			node.Children = child
			treeList = append(treeList, node)
		}
	}
	return treeList
}
