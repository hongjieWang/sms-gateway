package dto

import ("time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SmsServiceProviderConfigGetPageReq struct {
	dto.Pagination     `search:"-"`
    SmsServiceProviderConfigOrder
}

type SmsServiceProviderConfigOrder struct {Id int `form:"idOrder"  search:"type:order;column:id;table:sms_service_provider_config"`
    ProviderName string `form:"providerNameOrder"  search:"type:order;column:provider_name;table:sms_service_provider_config"`
    ProviderNo string `form:"providerNoOrder"  search:"type:order;column:provider_no;table:sms_service_provider_config"`
    AccessKeyId string `form:"accessKeyIdOrder"  search:"type:order;column:access_key_id;table:sms_service_provider_config"`
    AccessKeySecret string `form:"accessKeySecretOrder"  search:"type:order;column:access_key_secret;table:sms_service_provider_config"`
    Endpoint string `form:"endpointOrder"  search:"type:order;column:endpoint;table:sms_service_provider_config"`
    SdkAppId string `form:"sdkAppIdOrder"  search:"type:order;column:sdk_app_id;table:sms_service_provider_config"`
    Region string `form:"regionOrder"  search:"type:order;column:region;table:sms_service_provider_config"`
    Remark string `form:"remarkOrder"  search:"type:order;column:remark;table:sms_service_provider_config"`
    ExtJson string `form:"extJsonOrder"  search:"type:order;column:ext_json;table:sms_service_provider_config"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:sms_service_provider_config"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:sms_service_provider_config"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:sms_service_provider_config"`
    CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:sms_service_provider_config"`
    UpdatedAt time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sms_service_provider_config"`
    DeletedAt time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sms_service_provider_config"`
    
}

func (m *SmsServiceProviderConfigGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SmsServiceProviderConfigInsertReq struct {
    Id int `json:"-" comment:"主键ID"` // 主键ID
    ProviderName string `json:"providerName" comment:"服务商名称"`
    ProviderNo string `json:"providerNo" comment:"服务商编号"`
    AccessKeyId string `json:"accessKeyId" comment:"身份标识"`
    AccessKeySecret string `json:"accessKeySecret" comment:"身份认证密钥"`
    Endpoint string `json:"endpoint" comment:"调用域名"`
    SdkAppId string `json:"sdkAppId" comment:"应用ID"`
    Region string `json:"region" comment:"地域列表"`
    Remark string `json:"remark" comment:"备注"`
    ExtJson string `json:"extJson" comment:"扩展字段"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *SmsServiceProviderConfigInsertReq) Generate(model *models.SmsServiceProviderConfig)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.ProviderName = s.ProviderName
    model.ProviderNo = s.ProviderNo
    model.AccessKeyId = s.AccessKeyId
    model.AccessKeySecret = s.AccessKeySecret
    model.Endpoint = s.Endpoint
    model.SdkAppId = s.SdkAppId
    model.Region = s.Region
    model.Remark = s.Remark
    model.ExtJson = s.ExtJson
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *SmsServiceProviderConfigInsertReq) GetId() interface{} {
	return s.Id
}

type SmsServiceProviderConfigUpdateReq struct {
    Id int `uri:"id" comment:"主键ID"` // 主键ID
    ProviderName string `json:"providerName" comment:"服务商名称"`
    ProviderNo string `json:"providerNo" comment:"服务商编号"`
    AccessKeyId string `json:"accessKeyId" comment:"身份标识"`
    AccessKeySecret string `json:"accessKeySecret" comment:"身份认证密钥"`
    Endpoint string `json:"endpoint" comment:"调用域名"`
    SdkAppId string `json:"sdkAppId" comment:"应用ID"`
    Region string `json:"region" comment:"地域列表"`
    Remark string `json:"remark" comment:"备注"`
    ExtJson string `json:"extJson" comment:"扩展字段"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *SmsServiceProviderConfigUpdateReq) Generate(model *models.SmsServiceProviderConfig)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.ProviderName = s.ProviderName
    model.ProviderNo = s.ProviderNo
    model.AccessKeyId = s.AccessKeyId
    model.AccessKeySecret = s.AccessKeySecret
    model.Endpoint = s.Endpoint
    model.SdkAppId = s.SdkAppId
    model.Region = s.Region
    model.Remark = s.Remark
    model.ExtJson = s.ExtJson
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SmsServiceProviderConfigUpdateReq) GetId() interface{} {
	return s.Id
}

// SmsServiceProviderConfigGetReq 功能获取请求参数
type SmsServiceProviderConfigGetReq struct {
     Id int `uri:"id"`
}
func (s *SmsServiceProviderConfigGetReq) GetId() interface{} {
	return s.Id
}

// SmsServiceProviderConfigDeleteReq 功能删除请求参数
type SmsServiceProviderConfigDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SmsServiceProviderConfigDeleteReq) GetId() interface{} {
	return s.Ids
}