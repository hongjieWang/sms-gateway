package models

import (
	"go-admin/common/models"
)

type SmsTemplateConfig struct {
	models.Model

	SignName             string `json:"signName" gorm:"type:varchar(32);comment:签名名称"`
	TemplateNo           string `json:"templateNo" gorm:"type:varchar(32);comment:模版编号"`
	TemplateName         string `json:"templateName" gorm:"type:varchar(32);comment:模版名称"`
	TemplateContent      string `json:"templateContent" gorm:"type:varchar(255);comment:模版内容"`
	ThirdPartyTemplateNo string `json:"thirdPartyTemplateNo" gorm:"type:varchar(32);comment:第三方模版编号"`
	ProviderNo           string `json:"providerNo" gorm:"type:varchar(32);comment:服务商编号"`
	Remark               string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	ExtJson              string `json:"extJson" gorm:"type:json;comment:扩展信息"`
	Status               string `json:"status" gorm:"type:bit(1);comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (SmsTemplateConfig) TableName() string {
	return "sms_template_config"
}

func (e *SmsTemplateConfig) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SmsTemplateConfig) GetId() interface{} {
	return e.Id
}
