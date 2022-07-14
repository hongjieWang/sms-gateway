package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SmsServiceProviderConfig struct {
	service.Service
}

func (e *SmsServiceProviderConfig) GetByProviderNo(c *dto.SmsServiceProviderConfigGetPageReq, model *models.SmsServiceProviderConfig) error {
	var data models.SmsServiceProviderConfig
	err := e.Orm.Model(&data).Where(models.SmsServiceProviderConfig{ProviderNo: c.ProviderNo}).
		First(model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSmsServiceProviderConfig error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetPage 获取SmsServiceProviderConfig列表
func (e *SmsServiceProviderConfig) GetPage(c *dto.SmsServiceProviderConfigGetPageReq, p *actions.DataPermission, list *[]models.SmsServiceProviderConfig, count *int64) error {
	var err error
	var data models.SmsServiceProviderConfig

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).Order("updated_at desc").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SmsServiceProviderConfigService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SmsServiceProviderConfig对象
func (e *SmsServiceProviderConfig) Get(d *dto.SmsServiceProviderConfigGetReq, p *actions.DataPermission, model *models.SmsServiceProviderConfig) error {
	var data models.SmsServiceProviderConfig

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSmsServiceProviderConfig error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SmsServiceProviderConfig对象
func (e *SmsServiceProviderConfig) Insert(c *dto.SmsServiceProviderConfigInsertReq) error {
	var err error
	var data models.SmsServiceProviderConfig
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SmsServiceProviderConfigService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SmsServiceProviderConfig对象
func (e *SmsServiceProviderConfig) Update(c *dto.SmsServiceProviderConfigUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SmsServiceProviderConfig{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		e.Log.Errorf("SmsServiceProviderConfigService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SmsServiceProviderConfig
func (e *SmsServiceProviderConfig) Remove(d *dto.SmsServiceProviderConfigDeleteReq, p *actions.DataPermission) error {
	var data models.SmsServiceProviderConfig

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSmsServiceProviderConfig error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
