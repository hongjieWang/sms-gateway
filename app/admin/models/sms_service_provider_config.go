package models

import (
	"go-admin/common/models"
)

type SmsServiceProviderConfig struct {
	models.Model
	ProviderName    string `json:"providerName" gorm:"type:varchar(32);comment:服务商名称"`
	ProviderNo      string `json:"providerNo" gorm:"type:varchar(32);comment:服务商编号"`
	ChannelNo       string `json:"channelNo" gorm:"type:varchar(32);comment:短信渠道商"`
	AccessKeyId     string `json:"accessKeyId" gorm:"type:varchar(64);comment:身份标识"`
	AccessKeySecret string `json:"accessKeySecret" gorm:"type:varchar(64);comment:身份认证密钥"`
	Endpoint        string `json:"endpoint" gorm:"type:varchar(255);comment:调用域名"`
	SdkAppId        string `json:"sdkAppId" gorm:"type:varchar(32);comment:应用ID"`
	Region          string `json:"region" gorm:"type:varchar(32);comment:地域列表"`
	Remark          string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	ExtJson         string `json:"extJson" gorm:"type:json;comment:扩展字段"`
	Status          string `json:"status" gorm:"type:bit(1);comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (SmsServiceProviderConfig) TableName() string {
	return "sms_service_provider_config"
}

func (e *SmsServiceProviderConfig) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SmsServiceProviderConfig) GetId() interface{} {
	return e.Id
}
