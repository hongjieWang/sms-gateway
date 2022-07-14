package models

import (
	"go-admin/common/models"
)

type SmsAppConfig struct {
	models.Model

	AppNo           string `json:"appNo" gorm:"type:varchar(32);comment:应用编号"`
	AppName         string `json:"appName" gorm:"type:varchar(64);comment:应用名称"`
	AvailableNumber int64  `json:"availableNumber" gorm:"type:int;comment:可用数量"`
	CurrentLimiting int64  `json:"currentLimiting" gorm:"type:int;comment:限流数量"`
	UseNumber       int64  `json:"useNumber" gorm:"type:int;comment:已用数量"`
	Remark          string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	ExtJson         string `json:"extJson" gorm:"type:json;comment:扩展信息"`
	Status          string `json:"status" gorm:"type:bit(1);comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (SmsAppConfig) TableName() string {
	return "sms_app_config"
}

func (e *SmsAppConfig) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SmsAppConfig) GetId() interface{} {
	return e.Id
}
