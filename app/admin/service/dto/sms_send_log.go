package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SmsSendLogGetPageReq struct {
	dto.Pagination `search:"-"`
	BusinessNo     string `form:"businessNo"  search:"type:contains;column:business_no;table:sms_send_log"`
	AppNo          string `form:"appNo"  search:"type:contains;column:app_no;table:sms_send_log"`
	SmsSendLogOrder
}

type SmsSendLogOrder struct {
	Id          int       `form:"idOrder"  search:"type:order;column:id;table:sms_send_log"`
	AppNo       string    `form:"appNoOrder"  search:"type:order;column:app_no;table:sms_send_log"`
	BusinessNo  string    `form:"businessNoOrder"  search:"type:order;column:business_no;table:sms_send_log"`
	Status      string    `form:"statusOrder"  search:"type:order;column:status;table:sms_send_log"`
	Fee         string    `form:"feeOrder"  search:"type:order;column:fee;table:sms_send_log"`
	PhoneNumber string    `form:"phoneNumberOrder"  search:"type:order;column:phone_number;table:sms_send_log"`
	Message     string    `form:"messageOrder"  search:"type:order;column:message;table:sms_send_log"`
	Code        string    `form:"codeOrder"  search:"type:order;column:code;table:sms_send_log"`
	Content     string    `form:"contentOrder"  search:"type:order;column:content;table:sms_send_log"`
	Remark      string    `form:"remarkOrder"  search:"type:order;column:remark;table:sms_send_log"`
	ExtJson     string    `form:"extJsonOrder"  search:"type:order;column:ext_json;table:sms_send_log"`
	CreateBy    string    `form:"createByOrder"  search:"type:order;column:create_by;table:sms_send_log"`
	UpdateBy    string    `form:"updateByOrder"  search:"type:order;column:update_by;table:sms_send_log"`
	CreatedAt   time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:sms_send_log"`
	UpdatedAt   time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sms_send_log"`
	DeletedAt   time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sms_send_log"`
}

func (m *SmsSendLogGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SmsSendLogInsertReq struct {
	Id          int    `json:"-" comment:"??????"` // ??????
	AppNo       string `json:"appNo" comment:"??????ID"`
	BusinessNo  string `json:"businessNo" comment:"????????????"`
	Status      string `json:"status" comment:"??????"`
	Fee         string `json:"fee" comment:"????????????"`
	PhoneNumber string `json:"phoneNumber" comment:"???????????????"`
	Message     string `json:"message" comment:"??????????????????"`
	Code        string `json:"code" comment:"?????????????????????"`
	Content     string `json:"content" comment:"????????????"`
	Remark      string `json:"remark" comment:"??????"`
	ExtJson     string `json:"extJson" comment:"????????????"`
	common.ControlBy
}

func (s *SmsSendLogInsertReq) Generate(model *models.SmsSendLog) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AppNo = s.AppNo
	model.BusinessNo = s.BusinessNo
	model.Status = s.Status
	model.Fee = s.Fee
	model.PhoneNumber = s.PhoneNumber
	model.Message = s.Message
	model.Code = s.Code
	model.Content = s.Content
	model.Remark = s.Remark
	model.ExtJson = s.ExtJson
	model.CreateBy = s.CreateBy // ?????????????????????????????????????????????
}

func (s *SmsSendLogInsertReq) GetId() interface{} {
	return s.Id
}

type SmsSendLogUpdateReq struct {
	Id          int    `uri:"id" comment:"??????"` // ??????
	AppNo       string `json:"appNo" comment:"??????ID"`
	BusinessNo  string `json:"businessNo" comment:"????????????"`
	Status      string `json:"status" comment:"??????"`
	Fee         string `json:"fee" comment:"????????????"`
	PhoneNumber string `json:"phoneNumber" comment:"???????????????"`
	Message     string `json:"message" comment:"??????????????????"`
	Code        string `json:"code" comment:"?????????????????????"`
	Content     string `json:"content" comment:"????????????"`
	Remark      string `json:"remark" comment:"??????"`
	ExtJson     string `json:"extJson" comment:"????????????"`
	common.ControlBy
}

func (s *SmsSendLogUpdateReq) Generate(model *models.SmsSendLog) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AppNo = s.AppNo
	model.BusinessNo = s.BusinessNo
	model.Status = s.Status
	model.Fee = s.Fee
	model.PhoneNumber = s.PhoneNumber
	model.Message = s.Message
	model.Code = s.Code
	model.Content = s.Content
	model.Remark = s.Remark
	model.ExtJson = s.ExtJson
	model.UpdateBy = s.UpdateBy // ?????????????????????????????????????????????
}

func (s *SmsSendLogUpdateReq) GetId() interface{} {
	return s.Id
}

// SmsSendLogGetReq ????????????????????????
type SmsSendLogGetReq struct {
	Id int `uri:"id"`
}

func (s *SmsSendLogGetReq) GetId() interface{} {
	return s.Id
}

// SmsSendLogDeleteReq ????????????????????????
type SmsSendLogDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SmsSendLogDeleteReq) GetId() interface{} {
	return s.Ids
}
