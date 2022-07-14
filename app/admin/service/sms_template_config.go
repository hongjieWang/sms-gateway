package service

import (
	"errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SmsTemplateConfig struct {
	service.Service
}

// GetPage 获取SmsTemplateConfig列表
func (e *SmsTemplateConfig) GetPage(c *dto.SmsTemplateConfigGetPageReq, p *actions.DataPermission, list *[]models.SmsTemplateConfig, count *int64) error {
	var err error
	var data models.SmsTemplateConfig

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).Order("updated_at desc").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SmsTemplateConfigService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SmsTemplateConfig对象
func (e *SmsTemplateConfig) Get(d *dto.SmsTemplateConfigGetReq, p *actions.DataPermission, model *models.SmsTemplateConfig) error {
	var data models.SmsTemplateConfig

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSmsTemplateConfig error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetByTemplateNo 根据模版编号查询模版信息
func (e *SmsTemplateConfig) GetByTemplateNo(templateNo string, model *models.SmsTemplateConfig) error {
	var data models.SmsTemplateConfig
	err := e.Orm.Model(&data).
		Where(models.SmsTemplateConfig{
			TemplateNo: templateNo,
		}).
		First(model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSmsTemplateConfig error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SmsTemplateConfig对象
func (e *SmsTemplateConfig) Insert(c *dto.SmsTemplateConfigInsertReq) error {
	var err error
	var data models.SmsTemplateConfig
	c.Generate(&data)

	templateId, err := e.createTemplateTencent(c)
	if err != nil {
		e.Log.Errorf("SmsTemplateConfigService createTemplateTencent error:%s \r\n", err)
		return err
	}

	data.ThirdPartyTemplateNo = templateId

	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SmsTemplateConfigService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

//腾讯云创建短信模版
func (e *SmsTemplateConfig) createTemplateTencent(c *dto.SmsTemplateConfigInsertReq) (string, error) {
	var provider models.SmsServiceProviderConfig
	var providerData models.SmsServiceProviderConfig
	err := e.Orm.Model(&provider).Where(models.SmsServiceProviderConfig{
		ProviderNo: c.ProviderNo,
	}).First(&providerData).Error
	if err != nil {
		e.Log.Errorf("SmsTemplateConfigService createTemplateTencent error:%s \r\n", err)
		return "", err
	}
	credential := common.NewCredential(
		providerData.AccessKeyId,
		providerData.AccessKeySecret,
	)
	client, _ := sms.NewClient(credential, providerData.Region, profile.NewClientProfile())
	request := sms.NewAddSmsTemplateRequest()

	request.TemplateName = &c.TemplateName
	request.TemplateContent = &c.TemplateContent
	request.SmsType = common.Uint64Ptr(0)
	request.International = common.Uint64Ptr(0)
	request.Remark = &c.Remark

	response, err := client.AddSmsTemplate(request)
	return *response.Response.AddTemplateStatus.TemplateId, nil
}

// Update 修改SmsTemplateConfig对象
func (e *SmsTemplateConfig) Update(c *dto.SmsTemplateConfigUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SmsTemplateConfig{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		e.Log.Errorf("SmsTemplateConfigService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SmsTemplateConfig
func (e *SmsTemplateConfig) Remove(d *dto.SmsTemplateConfigDeleteReq, p *actions.DataPermission) error {
	var data models.SmsTemplateConfig

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSmsTemplateConfig error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
