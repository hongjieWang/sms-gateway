package tencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	smsDto "go-admin/app/sms/dto"
	"regexp"
	"strings"
)

type Tencent struct {
}

func (s Tencent) BuildContent(content string, req *dto.SmsSendMessageReq) string {
	r, _ := regexp.Compile("{(.)}")
	allString := r.FindAllString(content, -1)
	for i, s := range allString {
		content = strings.Replace(content, s, req.Params[i], -1)
	}
	return content
}

func (s Tencent) SendTencent(provider *models.SmsServiceProviderConfig, template *models.SmsTemplateConfig, req *dto.SmsSendMessageReq) (resp smsDto.SmsResponse, err error) {
	credential := common.NewCredential(provider.AccessKeyId, provider.AccessKeySecret)
	client, _ := sms.NewClient(credential, provider.Region, profile.NewClientProfile())
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(provider.SdkAppId)
	request.SignName = common.StringPtr(template.SignName)
	request.TemplateId = common.StringPtr(template.ThirdPartyTemplateNo)
	request.TemplateParamSet = common.StringPtrs(req.Params)
	request.PhoneNumberSet = common.StringPtrs(req.Phones)
	response, err := client.SendSms(request)
	requestId := response.Response.RequestId
	var statuses []smsDto.SmsSendStatus
	for _, status := range response.Response.SendStatusSet {
		statuses = append(statuses, smsDto.SmsSendStatus{
			Fee:     *status.Fee,
			Message: *status.Message,
			Code:    *status.Code,
			Phone:   *status.PhoneNumber,
		})
	}
	resp = smsDto.SmsResponse{
		RequestId:     *requestId,
		SmsSendStatus: statuses,
	}
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return resp, err
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		return resp, err
	}
	return resp, nil
}
