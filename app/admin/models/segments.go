package models

import (
	"go-admin/common/models"
)

type Segments struct {
	BizTag string `json:"biz_tag" gorm:"primaryKey;comment:主键编码"`
	MaxId  int64  `json:"maxId" gorm:"type:bigint;comment:MaxId"`
	Step   string `json:"step" gorm:"type:int;comment:Step"`
	Remark string `json:"remark" gorm:"type:varchar(200);comment:Remark"`
	models.ModelTime
	models.ControlBy
}

func (Segments) TableName() string {
	return "segments"
}

func (e *Segments) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Segments) GetId() interface{} {
	return e.BizTag
}
