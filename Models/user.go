package Models

import (
	"fmt"
	"package/Database"
)

//对应数据库
type User struct {
	Id       int    `gorm:"column:id" db:"id" form:"id"`
	Username string `gorm:"column:username" db:"username" form:"username"`
	Age      int    `gorm:"column:age" db:"age" form:"age"`
	Address  string `gorm:"column:address" db:"address" form:"address"`
	Time     int64  `gorm:"column:time" db:"time" form:"time"`
	Status   int    `gorm:"column:status" db:"status" form:"status"`
	IsDel    int    `gorm:"column:is_del" db:"is_del" form:"is_del"`
}

//对应response
type Users struct {
	Id       int    `json:"id" db:"id" form:"id"`
	Username string `json:"username" db:"username" form:"username"`
	Age      int    `json:"age" db:"age" form:"age"`
	Address  string `json:"address" db:"address" form:"address"`
	Time     string `json:"time" db:"time" form:"time"`
	Status   int    `json:"status" db:"status" form:"status"`
	IsDel    int    `json:"is_del" db:"is_del" form:"is_del"`
}

//查询所有用户信息
func UserList(u []*User) ([]*User, int, error) {
	userTest := Database.GetMysql()
	var count int
	db := userTest.DB.Table("user_test")
	db = db.Limit(20).Find(&u).Count(&count)
	err := db.Error
	fmt.Println(err)
	if err != nil {
		defer userTest.DB.Close()
		return u, count, err
	}
	return u, count, nil
}

//根据id查询一条用户信息
func UserListOne(where User) (User, error) {
	userTest := Database.GetMysql()
	var data User
	db := userTest.DB.Table("user_test")
	if where.Id > 0 {
		db = db.Where("id = ?", where.Id)
	}
	db = db.First(&data)
	err := db.Error
	if err != nil {
		fmt.Println(err)
		defer userTest.DB.Close()
		return data, err
	}
	return data, nil
}

//根据id更新一条用户信息
func UpdateUserOne(where User, data map[string]interface{}) (err error) {
	userTest := Database.GetMysql()
	db := userTest.DB.Table("user_test")
	if where.Id > 0 {
		db = db.Where("id = ?", where.Id).Update(data)
	}
	err = db.Error
	if err != nil {
		fmt.Println(err)
		defer userTest.DB.Close()
		return err
	}
	return nil
}

//根据id删除一条数据
func DelUserOne(where User) (err error) {
	usetTest := Database.GetMysql()
	db := usetTest.DB.Table("user_test")
	if where.Id > 0 {
		db = db.Where("id = ?", where.Id).Delete(&where)
	}
	err = db.Error
	if err != nil {
		defer usetTest.DB.Close()
		return err
	}
	return nil
}

//插入一条数据
func AddUserOne(u User) (err error) {
	userTest := Database.GetMysql()
	db := userTest.DB.Table("user_test")
	db = db.Create(&u)
	db.LogMode(true)
	err = db.Error
	if err != nil {
		defer userTest.DB.Close()
		return err
	}
	return nil
}

//批量插入数组数据
func AddUserList(u User) (err error) {
	userTest := Database.GetMysql()
	db := userTest.DB.Table("user_test")
	db = db.Create(&u)
	db.LogMode(true)
	err = db.Error
	if err != nil {
		defer userTest.DB.Close()
		return err
	}
	return nil
}
