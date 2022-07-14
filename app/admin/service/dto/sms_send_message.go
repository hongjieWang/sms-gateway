package dto

// SmsSendMessageReq 发送短信请求对象
type SmsSendMessageReq struct {
	BusinessNo string   `json:"businessNo" form:"businessNo" comment:"服务编号"`
	TemplateNo string   `json:"templateNo" form:"templateNo" comment:"模版编号"`
	Phones     []string `json:"phones" form:"phones" comment:"手机号集合"`
	Params     []string `json:"params" form:"params" comment:"参数列表"`
}
