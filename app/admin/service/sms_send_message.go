package service

import (
	"encoding/json"
	api "github.com/qiniu/go-sdk/v7"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/app/sms/aliyun"
	smsDto "go-admin/app/sms/dto"
	"go-admin/app/sms/tencent"
	"go-admin/common/service"
	"strconv"
)

type SmsSendMessage struct {
	service.Service
}

func (s *SmsSendMessage) SendMessage(req *dto.SmsSendMessageReq) (string, error) {
	s.Log.Info(req)
	var business models.SmsBusinessConfig
	err := s.getBusinessNo(req.BusinessNo, &business)
	number := s.getAppAvailableNumber(business.AppNo)
	if number < 0 {
		return "应用套餐不足", api.NewError("500", "应用套餐不足")
	}
	var template models.SmsTemplateConfig
	err = s.getTemplateNo(business.TemplateNo, &template)
	var provider models.SmsServiceProviderConfig
	err = s.getProviderNo(template.ProviderNo, &provider)
	var sendStatus smsDto.SmsResponse
	if provider.ChannelNo == smsDto.Tencent {
		t := tencent.Tencent{}
		sendStatus, err = t.SendTencent(&provider, &template, req)
		content := t.BuildContent(template.TemplateContent, req)
		template.TemplateContent = content
	} else if provider.ChannelNo == smsDto.AliYun {
		yun := aliyun.AliYun{}
		sendStatus, err = yun.SendSms(&provider, &template, req)
		if err != nil {
			return "", err
		}
		content := yun.BuildContent(template.TemplateContent, req)
		template.TemplateContent = content
	}
	err = s.installSendLog(sendStatus, &business, &template)
	if err != nil {
		return err.Error(), err
	}
	marshal, _ := json.Marshal(sendStatus)
	return string(marshal), nil
}

// 发送短信插入日志
func (s *SmsSendMessage) installSendLog(response smsDto.SmsResponse, business *models.SmsBusinessConfig, template *models.SmsTemplateConfig) error {
	var data models.SmsSendLog
	for _, resp := range response.SmsSendStatus {
		b, _ := json.Marshal(resp)
		model := models.SmsSendLog{
			AppNo:       business.AppNo,
			BusinessNo:  business.BusinessNo,
			Code:        resp.Code,
			PhoneNumber: resp.Phone,
			Fee:         strconv.FormatUint(resp.Fee, 10),
			Message:     resp.Message,
			Content:     template.TemplateContent,
			Remark:      response.RequestId,
			Status:      "1",
			ExtJson:     string(b),
		}
		s.updateAppNumber(business.AppNo, resp.Fee)
		err := s.Orm.Model(&data).Create(&model).Error
		if err != nil {
			s.Log.Errorf("SmsSendMessageService installSendLog error:%s \r\n", err)
			return err
		}
	}
	return nil
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
