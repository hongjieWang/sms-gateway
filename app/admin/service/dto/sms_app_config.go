package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SmsAppConfigGetPageReq struct {
	dto.Pagination `search:"-"`
	SmsAppConfigOrder
}

type SmsAppConfigOrder struct {
	Id              int       `form:"idOrder"  search:"type:order;column:id;table:sms_app_config"`
	AppNo           string    `form:"appNoOrder"  search:"type:order;column:app_no;table:sms_app_config"`
	AppName         string    `form:"appNameOrder"  search:"type:order;column:app_name;table:sms_app_config"`
	AvailableNumber string    `form:"availableNumberOrder"  search:"type:order;column:available_number;table:sms_app_config"`
	CurrentLimiting string    `form:"currentLimitingOrder"  search:"type:order;column:current_limiting;table:sms_app_config"`
	UseNumber       string    `form:"useNumberOrder"  search:"type:order;column:use_number;table:sms_app_config"`
	Remark          string    `form:"remarkOrder"  search:"type:order;column:remark;table:sms_app_config"`
	ExtJson         string    `form:"extJsonOrder"  search:"type:order;column:ext_json;table:sms_app_config"`
	Status          string    `form:"statusOrder"  search:"type:order;column:status;table:sms_app_config"`
	CreateBy        string    `form:"createByOrder"  search:"type:order;column:create_by;table:sms_app_config"`
	UpdateBy        string    `form:"updateByOrder"  search:"type:order;column:update_by;table:sms_app_config"`
	CreatedAt       time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:sms_app_config"`
	UpdatedAt       time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sms_app_config"`
	DeletedAt       time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sms_app_config"`
}

func (m *SmsAppConfigGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SmsAppConfigInsertReq struct {
	Id              int    `json:"-" comment:"主键ID"` // 主键ID
	AppNo           string `json:"appNo" comment:"应用编号"`
	AppName         string `json:"appName" comment:"应用名称"`
	AvailableNumber int64  `json:"availableNumber" comment:"可用数量"`
	CurrentLimiting int64  `json:"currentLimiting" comment:"限流数量"`
	UseNumber       int64  `json:"useNumber" comment:"已用数量"`
	Remark          string `json:"remark" comment:"备注"`
	ExtJson         string `json:"extJson" comment:"扩展信息"`
	Status          string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *SmsAppConfigInsertReq) Generate(model *models.SmsAppConfig) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AppNo = s.AppNo
	model.AppName = s.AppName
	model.AvailableNumber = s.AvailableNumber
	model.CurrentLimiting = s.CurrentLimiting
	model.UseNumber = s.UseNumber
	model.Remark = s.Remark
	model.ExtJson = s.ExtJson
	model.Status = s.Status
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *SmsAppConfigInsertReq) GetId() interface{} {
	return s.Id
}

type SmsAppConfigUpdateReq struct {
	Id              int    `uri:"id" comment:"主键ID"` // 主键ID
	AppNo           string `json:"appNo" comment:"应用编号"`
	AppName         string `json:"appName" comment:"应用名称"`
	AvailableNumber int64  `json:"availableNumber" comment:"可用数量"`
	CurrentLimiting int64  `json:"currentLimiting" comment:"限流数量"`
	UseNumber       int64  `json:"useNumber" comment:"已用数量"`
	Remark          string `json:"remark" comment:"备注"`
	ExtJson         string `json:"extJson" comment:"扩展信息"`
	Status          string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *SmsAppConfigUpdateReq) Generate(model *models.SmsAppConfig) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AppNo = s.AppNo
	model.AppName = s.AppName
	model.AvailableNumber = s.AvailableNumber
	model.CurrentLimiting = s.CurrentLimiting
	model.UseNumber = s.UseNumber
	model.Remark = s.Remark
	model.ExtJson = s.ExtJson
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SmsAppConfigUpdateReq) GetId() interface{} {
	return s.Id
}

// SmsAppConfigGetReq 功能获取请求参数
type SmsAppConfigGetReq struct {
	Id int `uri:"id"`
}

func (s *SmsAppConfigGetReq) GetId() interface{} {
	return s.Id
}

// SmsAppConfigDeleteReq 功能删除请求参数
type SmsAppConfigDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SmsAppConfigDeleteReq) GetId() interface{} {
	return s.Ids
}
