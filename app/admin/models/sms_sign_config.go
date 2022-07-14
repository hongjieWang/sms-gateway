package models

import (
	_ "time"

	"go-admin/common/models"
)

type SmsSignConfig struct {
	models.Model

	SignName     string `json:"signName" gorm:"type:varchar(32);comment:签名名称"`
	ProviderNo   string `json:"providerNo" gorm:"type:varchar(32);comment:服务商编号"`
	ProviderName string `json:"providerName" gorm:"type:varchar(64);comment:服务商名称"`
	Remark       string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	ExtJson      string `json:"extJson" gorm:"type:json;comment:扩展信息"`
	Status       string `json:"status" gorm:"type:bit(1);comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (SmsSignConfig) TableName() string {
	return "sms_sign_config"
}

func (e *SmsSignConfig) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SmsSignConfig) GetId() interface{} {
	return e.Id
}
