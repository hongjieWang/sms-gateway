package models

import (

	"go-admin/common/models"

)

type TArticleType struct {
    models.Model
    
    TypeNo string `json:"typeNo" gorm:"type:varchar(32);comment:分类编号"` 
    TypeName string `json:"typeName" gorm:"type:varchar(50);comment:分类名称"` 
    Status string `json:"status" gorm:"type:bit(1);comment:状态"`
    models.ModelTime
    models.ControlBy
}

func (TArticleType) TableName() string {
    return "t_article_type"
}

func (e *TArticleType) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TArticleType) GetId() interface{} {
	return e.Id
}