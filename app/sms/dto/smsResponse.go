package dto

type SmsResponse struct {
	SmsSendStatus []SmsSendStatus `json:"sms_send_status" common:"响应集合"`
	RequestId     string          `json:"request_id" common:"请求ID"`
}

type SmsSendStatus struct {
	Fee     uint64 `json:"fee" common:"计费条数"`
	Message string `json:"message" common:"响应信息"`
	Code    string `json:"code" common:"响应消息"`
	Phone   string `json:"phone" common:"手机号码"`
}

const (
	Tencent = "Tencent"
	AliYun  = "Aliyun"
)
