package aliyun

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	smsDto "go-admin/app/sms/dto"
	"regexp"
	"strings"
)

type AliYun struct {
}

// CreateClient /**
func (a *AliYun) CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func (a *AliYun) SendSms(provider *models.SmsServiceProviderConfig, template *models.SmsTemplateConfig, req *dto.SmsSendMessageReq) (smsDto.SmsResponse, error) {
	client, _err := a.CreateClient(tea.String(provider.AccessKeyId), tea.String(provider.AccessKeySecret))
	if _err != nil {
		return smsDto.SmsResponse{}, _err
	}
	//处理手机号
	phones := ""
	for _, phone := range req.Phones {
		phones = phones + phone + ","
	}
	phones = phones[:len(phones)-1]
	//处理请求参数
	r, _ := regexp.Compile("\\$\\{\\w+\\}")
	allString := r.FindAllString(template.TemplateContent, -1)
	var mapJson = make(map[string]string)
	for i, s := range allString {
		r1, _ := regexp.Compile("[a-z]")
		findAllString := r1.FindAllString(s, -1)
		key := ""
		for _, s := range findAllString {
			key = key + s
		}
		mapJson[key] = req.Params[i]
	}
	marshal, _ := json.Marshal(mapJson)
	param := string(marshal)
	sendSmsReq := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &phones,
		SignName:      &template.SignName,
		TemplateCode:  &template.ThirdPartyTemplateNo,
		TemplateParam: &param,
	}
	sms, _err := client.SendSms(sendSmsReq)
	var sendStatus []smsDto.SmsSendStatus
	for _, phone := range req.Phones {
		sendStatus = append(sendStatus, smsDto.SmsSendStatus{
			Phone:   phone,
			Code:    *sms.Body.Code,
			Message: *sms.Body.Message,
			Fee:     1,
		})
	}
	return smsDto.SmsResponse{
		SmsSendStatus: sendStatus,
		RequestId:     *sms.Body.RequestId,
	}, _err
}

func (a *AliYun) BuildContent(content string, req *dto.SmsSendMessageReq) string {
	r, _ := regexp.Compile("\\$\\{\\w+\\}")
	allString := r.FindAllString(content, -1)
	for i, s := range allString {
		fmt.Println(s)
		r1, _ := regexp.Compile("[a-z]")
		findAllString := r1.FindAllString(s, -1)
		key := ""
		for _, s := range findAllString {
			key = key + s
		}
		content = strings.Replace(content, s, req.Params[i], -1)
	}
	return content
}
