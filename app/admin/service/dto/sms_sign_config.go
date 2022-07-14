package dto

import ("time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SmsSignConfigGetPageReq struct {
	dto.Pagination     `search:"-"`
    SmsSignConfigOrder
}

type SmsSignConfigOrder struct {Id int `form:"idOrder"  search:"type:order;column:id;table:sms_sign_config"`
    SignName string `form:"signNameOrder"  search:"type:order;column:sign_name;table:sms_sign_config"`
    ProviderNo string `form:"providerNoOrder"  search:"type:order;column:provider_no;table:sms_sign_config"`
    ProviderName string `form:"providerNameOrder"  search:"type:order;column:provider_name;table:sms_sign_config"`
    Remark string `form:"remarkOrder"  search:"type:order;column:remark;table:sms_sign_config"`
    ExtJson string `form:"extJsonOrder"  search:"type:order;column:ext_json;table:sms_sign_config"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:sms_sign_config"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:sms_sign_config"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:sms_sign_config"`
    CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:sms_sign_config"`
    UpdatedAt time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sms_sign_config"`
    DeletedAt time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sms_sign_config"`
    
}

func (m *SmsSignConfigGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SmsSignConfigInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    SignName string `json:"signName" comment:"签名名称"`
    ProviderNo string `json:"providerNo" comment:"服务商编号"`
    ProviderName string `json:"providerName" comment:"服务商名称"`
    Remark string `json:"remark" comment:"备注"`
    ExtJson string `json:"extJson" comment:"扩展信息"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *SmsSignConfigInsertReq) Generate(model *models.SmsSignConfig)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.SignName = s.SignName
    model.ProviderNo = s.ProviderNo
    model.ProviderName = s.ProviderName
    model.Remark = s.Remark
    model.ExtJson = s.ExtJson
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *SmsSignConfigInsertReq) GetId() interface{} {
	return s.Id
}

type SmsSignConfigUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    SignName string `json:"signName" comment:"签名名称"`
    ProviderNo string `json:"providerNo" comment:"服务商编号"`
    ProviderName string `json:"providerName" comment:"服务商名称"`
    Remark string `json:"remark" comment:"备注"`
    ExtJson string `json:"extJson" comment:"扩展信息"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *SmsSignConfigUpdateReq) Generate(model *models.SmsSignConfig)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.SignName = s.SignName
    model.ProviderNo = s.ProviderNo
    model.ProviderName = s.ProviderName
    model.Remark = s.Remark
    model.ExtJson = s.ExtJson
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SmsSignConfigUpdateReq) GetId() interface{} {
	return s.Id
}

// SmsSignConfigGetReq 功能获取请求参数
type SmsSignConfigGetReq struct {
     Id int `uri:"id"`
}
func (s *SmsSignConfigGetReq) GetId() interface{} {
	return s.Id
}

// SmsSignConfigDeleteReq 功能删除请求参数
type SmsSignConfigDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SmsSignConfigDeleteReq) GetId() interface{} {
	return s.Ids
}