package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SmsServiceProviderConfigGetPageReq struct {
	dto.Pagination `search:"-"`
	SmsServiceProviderConfigOrder
}

type SmsServiceProviderConfigOrder struct {
	Id              int       `form:"idOrder"  search:"type:order;column:id;table:sms_service_provider_config"`
	ProviderName    string    `form:"providerNameOrder"  search:"type:order;column:provider_name;table:sms_service_provider_config"`
	ProviderNo      string    `form:"providerNoOrder"  search:"type:order;column:provider_no;table:sms_service_provider_config"`
	AccessKeyId     string    `form:"accessKeyIdOrder"  search:"type:order;column:access_key_id;table:sms_service_provider_config"`
	AccessKeySecret string    `form:"accessKeySecretOrder"  search:"type:order;column:access_key_secret;table:sms_service_provider_config"`
	Endpoint        string    `form:"endpointOrder"  search:"type:order;column:endpoint;table:sms_service_provider_config"`
	SdkAppId        string    `form:"sdkAppIdOrder"  search:"type:order;column:sdk_app_id;table:sms_service_provider_config"`
	Region          string    `form:"regionOrder"  search:"type:order;column:region;table:sms_service_provider_config"`
	Remark          string    `form:"remarkOrder"  search:"type:order;column:remark;table:sms_service_provider_config"`
	ExtJson         string    `form:"extJsonOrder"  search:"type:order;column:ext_json;table:sms_service_provider_config"`
	Status          string    `form:"statusOrder"  search:"type:order;column:status;table:sms_service_provider_config"`
	CreateBy        string    `form:"createByOrder"  search:"type:order;column:create_by;table:sms_service_provider_config"`
	UpdateBy        string    `form:"updateByOrder"  search:"type:order;column:update_by;table:sms_service_provider_config"`
	CreatedAt       time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:sms_service_provider_config"`
	UpdatedAt       time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sms_service_provider_config"`
	DeletedAt       time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sms_service_provider_config"`
}

func (m *SmsServiceProviderConfigGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SmsServiceProviderConfigInsertReq struct {
	Id              int    `json:"-" comment:"??????ID"` // ??????ID
	ProviderName    string `json:"providerName" comment:"???????????????"`
	ProviderNo      string `json:"providerNo" comment:"???????????????"`
	ChannelNo       string `json:"channelNo" comment:"???????????????"`
	AccessKeyId     string `json:"accessKeyId" comment:"????????????"`
	AccessKeySecret string `json:"accessKeySecret" comment:"??????????????????"`
	Endpoint        string `json:"endpoint" comment:"????????????"`
	SdkAppId        string `json:"sdkAppId" comment:"??????ID"`
	Region          string `json:"region" comment:"????????????"`
	Remark          string `json:"remark" comment:"??????"`
	ExtJson         string `json:"extJson" comment:"????????????"`
	Status          string `json:"status" comment:"??????"`
	common.ControlBy
}

func (s *SmsServiceProviderConfigInsertReq) Generate(model *models.SmsServiceProviderConfig) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ProviderName = s.ProviderName
	model.ChannelNo = s.ChannelNo
	model.ProviderNo = s.ProviderNo
	model.AccessKeyId = s.AccessKeyId
	model.AccessKeySecret = s.AccessKeySecret
	model.Endpoint = s.Endpoint
	model.SdkAppId = s.SdkAppId
	model.Region = s.Region
	model.Remark = s.Remark
	model.ExtJson = s.ExtJson
	model.Status = s.Status
	model.CreateBy = s.CreateBy // ?????????????????????????????????????????????
}

func (s *SmsServiceProviderConfigInsertReq) GetId() interface{} {
	return s.Id
}

type SmsServiceProviderConfigUpdateReq struct {
	Id              int    `uri:"id" comment:"??????ID"` // ??????ID
	ProviderName    string `json:"providerName" comment:"???????????????"`
	ProviderNo      string `json:"providerNo" comment:"???????????????"`
	ChannelNo       string `json:"channelNo" comment:"?????????????????????"`
	AccessKeyId     string `json:"accessKeyId" comment:"????????????"`
	AccessKeySecret string `json:"accessKeySecret" comment:"??????????????????"`
	Endpoint        string `json:"endpoint" comment:"????????????"`
	SdkAppId        string `json:"sdkAppId" comment:"??????ID"`
	Region          string `json:"region" comment:"????????????"`
	Remark          string `json:"remark" comment:"??????"`
	ExtJson         string `json:"extJson" comment:"????????????"`
	Status          string `json:"status" comment:"??????"`
	common.ControlBy
}

func (s *SmsServiceProviderConfigUpdateReq) Generate(model *models.SmsServiceProviderConfig) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ProviderName = s.ProviderName
	model.ProviderNo = s.ProviderNo
	model.ChannelNo = s.ChannelNo
	model.AccessKeyId = s.AccessKeyId
	model.AccessKeySecret = s.AccessKeySecret
	model.Endpoint = s.Endpoint
	model.SdkAppId = s.SdkAppId
	model.Region = s.Region
	model.Remark = s.Remark
	model.ExtJson = s.ExtJson
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // ?????????????????????????????????????????????
}

func (s *SmsServiceProviderConfigUpdateReq) GetId() interface{} {
	return s.Id
}

// SmsServiceProviderConfigGetReq ????????????????????????
type SmsServiceProviderConfigGetReq struct {
	Id int `uri:"id"`
}

func (s *SmsServiceProviderConfigGetReq) GetId() interface{} {
	return s.Id
}

// SmsServiceProviderConfigDeleteReq ????????????????????????
type SmsServiceProviderConfigDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SmsServiceProviderConfigDeleteReq) GetId() interface{} {
	return s.Ids
}
