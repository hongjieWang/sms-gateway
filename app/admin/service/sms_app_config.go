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

type SmsAppConfig struct {
	service.Service
}

// GetPage 获取SmsAppConfig列表
func (e *SmsAppConfig) GetPage(c *dto.SmsAppConfigGetPageReq, p *actions.DataPermission, list *[]models.SmsAppConfig, count *int64) error {
	var err error
	var data models.SmsAppConfig

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).Order("updated_at desc").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SmsAppConfigService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SmsAppConfig对象
func (e *SmsAppConfig) Get(d *dto.SmsAppConfigGetReq, p *actions.DataPermission, model *models.SmsAppConfig) error {
	var data models.SmsAppConfig

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSmsAppConfig error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SmsAppConfig对象
func (e *SmsAppConfig) Insert(c *dto.SmsAppConfigInsertReq) error {
	var err error
	var data models.SmsAppConfig
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SmsAppConfigService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SmsAppConfig对象
func (e *SmsAppConfig) Update(c *dto.SmsAppConfigUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SmsAppConfig{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		e.Log.Errorf("SmsAppConfigService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SmsAppConfig
func (e *SmsAppConfig) Remove(d *dto.SmsAppConfigDeleteReq, p *actions.DataPermission) error {
	var data models.SmsAppConfig

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSmsAppConfig error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
