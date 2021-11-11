package Models

import (
	"encoding/json"
	"package/Database"
)

type Result struct {
	Id               int    `gorm:"column:id" json:"id"`
	DesignResultId   string `gorm:"column:design_result_id" json:"design_result_id"`
	DesignId         string `gorm:"column:design_id" json:"design_id"`
	AlgorithmId      string `gorm:"column:algorithm_id" json:"algorithm_id"`
	DesignResultPath json.RawMessage `gorm:"column:design_result_path" json:"design_result_path"`
	Video            string `gorm:"column:video" json:"video"`
	RendererVersion  string `gorm:"column:renderer_version" json:"renderer_version"`
	DefaultShow      string `gorm:"column:default_show" json:"default_show"`
	ExecuteCount     int    `gorm:"column:execute_count" json:"execute_count"`
	IsDel            int    `gorm:"column:is_del" json:"is_del"`
	CreateUserId     string `gorm:"column:create_user_id" json:"create_user_id"`
	UpdateUserId     string `gorm:"column:update_user_id" json:"update_user_id"`
	CreateTime       string `gorm:"column:create_time" json:"create_time"`
	UpdateTime       string `gorm:"column:update_time" json:"update_time"`
}

//根据条件查询信息 并分页
func FindList(where Result, page int, pageSize int) ([]Result, int64, error) {
	result := Database.GetMysql()
	var data []Result
	var count int64
	db := result.DB.Table("user_result")
	if where.Id > 0 {
		db = db.Where("id = ?", where.Id)
	}
	if where.ExecuteCount > 0 {
		db = db.Where("execute_count = ?", where.ExecuteCount)
	}
	if where.IsDel > 0 {
		db = db.Where("is_del = ?", where.IsDel)
	}
	if where.Video != "" {
		db = db.Where("video = ?", where.Video)
	}
	if page > 0 && pageSize > 0 {
		db = db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	db = db.Find(&data).Count(&count)
	db.LogMode(true)
	err := db.Error
	if err != nil {
		defer result.DB.Close()
		return data, count, err
	}
	return data, count, nil
}
