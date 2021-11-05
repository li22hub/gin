package Models

import (
	"fmt"
	"package/Database"
)

type User struct {
	Id       int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	Username string `gorm:"column:username" db:"username" json:"username" form:"username"`
	Age      int    `gorm:"column:age" db:"age" json:"age" form:"age"`
	Address  string `gorm:"column:address" db:"address" json:"address" form:"address"`
	Time     int64  `gorm:"column:time" db:"time" json:"time" form:"time"`
	Status   int    `gorm:"column:status" db:"status" json:"status" form:"status"`
	IsDel    int    `gorm:"column:is_del" db:"is_del" json:"is_del" form:"is_del"`
}


//根据id查询一条用户信息
func UserListOne(u User) (User, error) {
	userTest := Database.GetMysql()
	err := userTest.DB.Table("user_test").First(&u).Error
	fmt.Println(err)
	if err != nil {
		defer userTest.DB.Close()
		return u, err
	}
	return u, nil
}


//查询所有用户信息
func UserList(u []*User) ([]*User, error) {
	userTest := Database.GetMysql()
	err := userTest.DB.Table("user_test").Find(&u).Error
	fmt.Println(err)
	if err != nil {
		defer userTest.DB.Close()
		return u, err
	}
	return u, nil
}