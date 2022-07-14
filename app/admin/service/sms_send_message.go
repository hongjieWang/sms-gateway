package service

import (
	"encoding/json"
	api "github.com/qiniu/go-sdk/v7"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/service"
	"regexp"
	"strconv"
	"strings"
)

type SmsSendMessage struct {
	service.Service
}

func (s *SmsSendMessage) SendMessage(req *dto.SmsSendMessageReq) error {
	s.Log.Info(req)
	var business models.SmsBusinessConfig
	err := s.getBusinessNo(req.BusinessNo, &business)
	number := s.getAppAvailableNumber(business.AppNo)
	if number < 0 {
		return api.NewError("500", "应用套餐不足")
	}
	var template models.SmsTemplateConfig
	err = s.getTemplateNo(business.TemplateNo, &template)
	var provider models.SmsServiceProviderConfig
	err = s.getProviderNo(template.ProviderNo, &provider)
	tencent, err := s.sendTencent(&provider, &template, req)
	content := s.buildContent(template.TemplateContent, req)
	template.TemplateContent = content
	err = s.installSendLog(tencent, &business, &template)
	if err != nil {
		return err
	}
	s.Log.Infof("a", tencent)
	return nil
}

// 发送短信插入日志
func (s *SmsSendMessage) installSendLog(response *sms.SendSmsResponse, business *models.SmsBusinessConfig, template *models.SmsTemplateConfig) error {
	var data models.SmsSendLog

	for _, resp := range response.Response.SendStatusSet {
		b, _ := json.Marshal(*resp)
		model := models.SmsSendLog{
			AppNo:       business.AppNo,
			BusinessNo:  business.BusinessNo,
			Code:        *resp.Code,
			PhoneNumber: *resp.PhoneNumber,
			Fee:         strconv.FormatUint(*resp.Fee, 10),
			Message:     *resp.Message,
			Content:     template.TemplateContent,
			Remark:      *response.Response.RequestId,
			Status:      "1",
			ExtJson:     string(b),
		}
		go s.updateAppNumber(business.AppNo, *resp.Fee)
		err := s.Orm.Model(&data).Create(&model).Error
		if err != nil {
			s.Log.Errorf("SmsSendMessageService installSendLog error:%s \r\n", err)
			return err
		}
	}
	return nil
}

func (s *SmsSendMessage) buildContent(content string, req *dto.SmsSendMessageReq) string {
	r, _ := regexp.Compile("{(.)}")
	allString := r.FindAllString(content, -1)
	for i, s := range allString {
		content = strings.Replace(content, s, req.Params[i], -1)
	}
	return content
}

//获取当前应用剩余短信数量
func (s *SmsSendMessage) getAppAvailableNumber(appNo string) int64 {
	var data models.SmsAppConfig
	var app models.SmsAppConfig
	err := s.Orm.Model(&data).Where(models.SmsAppConfig{
		AppNo: appNo,
	}).First(&app).Error
	if err != nil {
		s.Log.Errorf("SmsSendMessageService updateAppNumber error:%s \r\n", err)
		return -1
	}
	return app.AvailableNumber
}

// 更新应用使用数量
func (s *SmsSendMessage) updateAppNumber(appNo string, fee uint64) {
	var data models.SmsAppConfig
	var app models.SmsAppConfig
	err := s.Orm.Model(&data).Where(models.SmsAppConfig{
		AppNo: appNo,
	}).First(&app).Error
	if err != nil {
		s.Log.Errorf("SmsSendMessageService updateAppNumber error:%s \r\n", err)
		return
	}
	app.UseNumber = int64(fee) + app.UseNumber
	app.AvailableNumber = app.AvailableNumber - int64(fee)
	s.Orm.Save(&app)
}

// 腾讯云发送短信
func (s *SmsSendMessage) sendTencent(provider *models.SmsServiceProviderConfig, template *models.SmsTemplateConfig, req *dto.SmsSendMessageReq) (response *sms.SendSmsResponse, err error) {
	credential := common.NewCredential(provider.AccessKeyId, provider.AccessKeySecret)
	client, _ := sms.NewClient(credential, provider.Region, profile.NewClientProfile())
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(provider.SdkAppId)
	request.SignName = common.StringPtr(template.SignName)
	request.TemplateId = common.StringPtr(template.ThirdPartyTemplateNo)
	request.TemplateParamSet = common.StringPtrs(req.Params)
	request.PhoneNumberSet = common.StringPtrs(req.Phones)
	response, err = client.SendSms(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return nil, err
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *SmsSendMessage) getProviderNo(providerNo string, provider *models.SmsServiceProviderConfig) error {
	var data models.SmsServiceProviderConfig
	err := s.Orm.Model(&data).Where(models.SmsServiceProviderConfig{
		ProviderNo: providerNo,
	}).First(provider).Error
	if err != nil {
		s.Log.Errorf("SmsSendMessageService getProviderNo error:%s \r\n", err)
		return err
	}
	return nil
}

// 根据模版编号获取短信运营商
func (s *SmsSendMessage) getTemplateNo(templateNo string, template *models.SmsTemplateConfig) error {
	var data models.SmsTemplateConfig
	err := s.Orm.Model(&data).Where(models.SmsTemplateConfig{
		TemplateNo: templateNo,
	}).First(template).Error
	if err != nil {
		s.Log.Errorf("SmsSendMessageService getTemplateNo error:%s \r\n", err)
		return err
	}
	return nil
}

//根据业务编号查询业务信息
func (s *SmsSendMessage) getBusinessNo(businessNo string, config *models.SmsBusinessConfig) error {
	var data models.SmsBusinessConfig
	err := s.Orm.Model(&data).Where(models.SmsBusinessConfig{
		BusinessNo: businessNo,
	}).First(config).Error
	if err != nil {
		s.Log.Errorf("SmsSendMessageService getBusinessNo error:%s \r\n", err)
		return err
	}
	return nil
}
