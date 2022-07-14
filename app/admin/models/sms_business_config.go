package models

import (
	"go-admin/common/models"
)

type SmsBusinessConfig struct {
	models.Model

	AppNo        string `json:"appNo" gorm:"type:varchar(32);comment:应用ID"`
	BusinessName string `json:"businessName" gorm:"type:varchar(255);comment:业务名称"`
	BusinessNo   string `json:"businessNo" gorm:"type:varchar(32);comment:业务编号"`
	TemplateNo   string `json:"templateNo" gorm:"type:varchar(32);comment:模版编号"`
	BusinessDesc string `json:"businessDesc" gorm:"type:varchar(255);comment:业务说明"`
	Remark       string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	ExtJson      string `json:"extJson" gorm:"type:json;comment:扩展信息"`
	Status       string `json:"status" gorm:"type:bit(1);comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (SmsBusinessConfig) TableName() string {
	return "sms_business_config"
}

func (e *SmsBusinessConfig) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SmsBusinessConfig) GetId() interface{} {
	return e.Id
}
