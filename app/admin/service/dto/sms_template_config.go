package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SmsTemplateConfigGetPageReq struct {
	dto.Pagination `search:"-"`
	SmsTemplateConfigOrder
}

type SmsTemplateConfigOrder struct {
	Id                   int       `form:"idOrder"  search:"type:order;column:id;table:sms_template_config"`
	SignName             string    `form:"signNameOrder"  search:"type:order;column:sign_name;table:sms_template_config"`
	TemplateNo           string    `form:"templateNoOrder"  search:"type:order;column:template_no;table:sms_template_config"`
	TemplateName         string    `form:"templateNameOrder"  search:"type:order;column:template_name;table:sms_template_config"`
	TemplateContent      string    `form:"templateContentOrder"  search:"type:order;column:template_content;table:sms_template_config"`
	ThirdPartyTemplateNo string    `form:"thirdPartyTemplateNoOrder"  search:"type:order;column:third_party_template_no;table:sms_template_config"`
	ProviderNo           string    `form:"providerNoOrder"  search:"type:order;column:provider_no;table:sms_template_config"`
	Remark               string    `form:"remarkOrder"  search:"type:order;column:remark;table:sms_template_config"`
	ExtJson              string    `form:"extJsonOrder"  search:"type:order;column:ext_json;table:sms_template_config"`
	Status               string    `form:"statusOrder"  search:"type:order;column:status;table:sms_template_config"`
	CreateBy             string    `form:"createByOrder"  search:"type:order;column:create_by;table:sms_template_config"`
	UpdateBy             string    `form:"updateByOrder"  search:"type:order;column:update_by;table:sms_template_config"`
	CreatedAt            time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:sms_template_config"`
	UpdatedAt            time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sms_template_config"`
	DeletedAt            time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sms_template_config"`
}

func (m *SmsTemplateConfigGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SmsTemplateConfigGetByTemplateNoReq struct {
	TemplateNo string `json:"templateNo"  form:"templateNo" comment:"????????????"`
	common.ControlBy
}

type SmsTemplateConfigInsertReq struct {
	Id                   int    `json:"-" comment:"??????"` // ??????
	SignName             string `json:"signName" comment:"????????????"`
	TemplateNo           string `json:"templateNo" comment:"????????????"`
	TemplateName         string `json:"templateName" comment:"??????????????????"`
	TemplateContent      string `json:"templateContent" comment:"????????????"`
	ThirdPartyTemplateNo string `json:"thirdPartyTemplateNo" comment:"?????????????????????"`
	ProviderNo           string `json:"providerNo" comment:"???????????????"`
	Remark               string `json:"remark" comment:"??????"`
	ExtJson              string `json:"extJson" comment:"????????????"`
	Status               string `json:"status" comment:"??????"`
	common.ControlBy
}

func (s *SmsTemplateConfigInsertReq) Generate(model *models.SmsTemplateConfig) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.SignName = s.SignName
	model.TemplateNo = s.TemplateNo
	model.TemplateName = s.TemplateName
	model.TemplateContent = s.TemplateContent
	model.ThirdPartyTemplateNo = s.ThirdPartyTemplateNo
	model.ProviderNo = s.ProviderNo
	model.Remark = s.Remark
	model.ExtJson = s.ExtJson
	model.Status = s.Status
	model.CreateBy = s.CreateBy // ?????????????????????????????????????????????
}

func (s *SmsTemplateConfigInsertReq) GetId() interface{} {
	return s.Id
}

type SmsTemplateConfigUpdateReq struct {
	Id                   int    `uri:"id" comment:"??????"` // ??????
	SignName             string `json:"signName" comment:"????????????"`
	TemplateNo           string `json:"templateNo" comment:"????????????"`
	TemplateName         string `json:"templateName" comment:"????????????"`
	TemplateContent      string `json:"templateContent" comment:"????????????"`
	ThirdPartyTemplateNo string `json:"thirdPartyTemplateNo" comment:"?????????????????????"`
	ProviderNo           string `json:"providerNo" comment:"???????????????"`
	Remark               string `json:"remark" comment:"??????"`
	ExtJson              string `json:"extJson" comment:"????????????"`
	Status               string `json:"status" comment:"??????"`
	common.ControlBy
}

func (s *SmsTemplateConfigUpdateReq) Generate(model *models.SmsTemplateConfig) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.SignName = s.SignName
	model.TemplateNo = s.TemplateNo
	model.TemplateName = s.TemplateName
	model.TemplateContent = s.TemplateContent
	model.ThirdPartyTemplateNo = s.ThirdPartyTemplateNo
	model.ProviderNo = s.ProviderNo
	model.Remark = s.Remark
	model.ExtJson = s.ExtJson
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // ?????????????????????????????????????????????
}

func (s *SmsTemplateConfigUpdateReq) GetId() interface{} {
	return s.Id
}

// SmsTemplateConfigGetReq ????????????????????????
type SmsTemplateConfigGetReq struct {
	Id int `uri:"id"`
}

func (s *SmsTemplateConfigGetReq) GetId() interface{} {
	return s.Id
}

// SmsTemplateConfigDeleteReq ????????????????????????
type SmsTemplateConfigDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SmsTemplateConfigDeleteReq) GetId() interface{} {
	return s.Ids
}
