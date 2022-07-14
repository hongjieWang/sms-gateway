package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SmsBusinessConfigGetPageReq struct {
	dto.Pagination `search:"-"`
	SmsBusinessConfigOrder
}

type SmsBusinessConfigOrder struct {
	Id           int       `form:"idOrder"  search:"type:order;column:id;table:sms_business_config"`
	AppNo        string    `form:"appNoOrder"  search:"type:order;column:app_no;table:sms_business_config"`
	BusinessName string    `form:"businessNameOrder"  search:"type:order;column:business_name;table:sms_business_config"`
	BusinessNo   string    `form:"businessNoOrder"  search:"type:order;column:business_no;table:sms_business_config"`
	BusinessDesc string    `form:"businessDescOrder"  search:"type:order;column:business_desc;table:sms_business_config"`
	Remark       string    `form:"remarkOrder"  search:"type:order;column:remark;table:sms_business_config"`
	ExtJson      string    `form:"extJsonOrder"  search:"type:order;column:ext_json;table:sms_business_config"`
	Status       string    `form:"statusOrder"  search:"type:order;column:status;table:sms_business_config"`
	CreateBy     string    `form:"createByOrder"  search:"type:order;column:create_by;table:sms_business_config"`
	UpdateBy     string    `form:"updateByOrder"  search:"type:order;column:update_by;table:sms_business_config"`
	CreatedAt    time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:sms_business_config"`
	UpdatedAt    time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sms_business_config"`
	DeletedAt    time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sms_business_config"`
}

func (m *SmsBusinessConfigGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SmsBusinessConfigInsertReq struct {
	Id           int    `json:"-" comment:"主键"` // 主键
	AppNo        string `json:"appNo" comment:"应用ID"`
	BusinessName string `json:"businessName" comment:"业务名称"`
	BusinessNo   string `json:"businessNo" comment:"业务编号"`
	TemplateNo   string `json:"templateNo" comment:"模版编号"`
	BusinessDesc string `json:"businessDesc" comment:"业务说明"`
	Remark       string `json:"remark" comment:"备注"`
	ExtJson      string `json:"extJson" comment:"扩展信息"`
	Status       string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *SmsBusinessConfigInsertReq) Generate(model *models.SmsBusinessConfig) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AppNo = s.AppNo
	model.BusinessName = s.BusinessName
	model.BusinessNo = s.BusinessNo
	model.BusinessDesc = s.BusinessDesc
	model.Remark = s.Remark
	model.TemplateNo = s.TemplateNo
	model.ExtJson = s.ExtJson
	model.Status = s.Status
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *SmsBusinessConfigInsertReq) GetId() interface{} {
	return s.Id
}

type SmsBusinessConfigUpdateReq struct {
	Id           int    `uri:"id" comment:"主键"` // 主键
	AppNo        string `json:"appNo" comment:"应用ID"`
	BusinessName string `json:"businessName" comment:"业务名称"`
	BusinessNo   string `json:"businessNo" comment:"业务编号"`
	TemplateNo   string `json:"templateNo" comment:"模版编号"`
	BusinessDesc string `json:"businessDesc" comment:"业务说明"`
	Remark       string `json:"remark" comment:"备注"`
	ExtJson      string `json:"extJson" comment:"扩展信息"`
	Status       string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *SmsBusinessConfigUpdateReq) Generate(model *models.SmsBusinessConfig) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AppNo = s.AppNo
	model.BusinessName = s.BusinessName
	model.BusinessNo = s.BusinessNo
	model.BusinessDesc = s.BusinessDesc
	model.TemplateNo = s.TemplateNo
	model.Remark = s.Remark
	model.ExtJson = s.ExtJson
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SmsBusinessConfigUpdateReq) GetId() interface{} {
	return s.Id
}

// SmsBusinessConfigGetReq 功能获取请求参数
type SmsBusinessConfigGetReq struct {
	Id int `uri:"id"`
}

func (s *SmsBusinessConfigGetReq) GetId() interface{} {
	return s.Id
}

// SmsBusinessConfigDeleteReq 功能删除请求参数
type SmsBusinessConfigDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SmsBusinessConfigDeleteReq) GetId() interface{} {
	return s.Ids
}
