package models

import (
	"go-admin/common/models"
)

type SmsSendLog struct {
	models.Model

	AppNo       string `json:"appNo" gorm:"type:varchar(32);comment:应用ID"`
	BusinessNo  string `json:"businessNo" gorm:"type:varchar(32);comment:业务编号"`
	Status      string `json:"status" gorm:"type:int;comment:状态"`
	Fee         string `json:"fee" gorm:"type:int;comment:计价条数"`
	PhoneNumber string `json:"phoneNumber" gorm:"type:varchar(255);comment:发送手机号"`
	Message     string `json:"message" gorm:"type:varchar(255);comment:接口响应消息"`
	Code        string `json:"code" gorm:"type:varchar(32);comment:接口响应状态码"`
	Content     string `json:"content" gorm:"type:varchar(255);comment:发送内容"`
	Remark      string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	ExtJson     string `json:"extJson" gorm:"type:json;comment:扩展信息"`
	models.ModelTime
	models.ControlBy
}

func (SmsSendLog) TableName() string {
	return "sms_send_log"
}

func (e *SmsSendLog) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SmsSendLog) GetId() interface{} {
	return e.Id
}
