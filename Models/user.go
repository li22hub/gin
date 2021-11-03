package models

import (
	"github.com/jinzhu/gorm"
	"package/Database"
)

type User struct {
	Id       int64 `gorm:"primary_key"`
	Username string
	Age      int64
	Address	 string
	Time	 int64
	Status   int
	IsDel    int
}

var (
	userTest *Database.Mysql
)

//获取表名
func (User) TableName() string {
	return "user_test"
}

func NewUser() *User {
	return &User{}
}

//创建用户
func (u *User) Create() (*gorm.DB, error) {
	userTest, _ = Database.GetMysql()
	defer userTest.DB.Close()
	//fmt.Println(u.Username)
	data := userTest.DB.Create(&u)
	return data, nil
}
