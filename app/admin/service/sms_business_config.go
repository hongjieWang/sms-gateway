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

type SmsBusinessConfig struct {
	service.Service
}

// GetPage 获取SmsBusinessConfig列表
func (e *SmsBusinessConfig) GetPage(c *dto.SmsBusinessConfigGetPageReq, p *actions.DataPermission, list *[]models.SmsBusinessConfig, count *int64) error {
	var err error
	var data models.SmsBusinessConfig

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).Order("updated_at desc").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SmsBusinessConfigService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SmsBusinessConfig对象
func (e *SmsBusinessConfig) Get(d *dto.SmsBusinessConfigGetReq, p *actions.DataPermission, model *models.SmsBusinessConfig) error {
	var data models.SmsBusinessConfig

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSmsBusinessConfig error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SmsBusinessConfig对象
func (e *SmsBusinessConfig) Insert(c *dto.SmsBusinessConfigInsertReq) error {
	var err error
	var data models.SmsBusinessConfig
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SmsBusinessConfigService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SmsBusinessConfig对象
func (e *SmsBusinessConfig) Update(c *dto.SmsBusinessConfigUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SmsBusinessConfig{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		e.Log.Errorf("SmsBusinessConfigService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SmsBusinessConfig
func (e *SmsBusinessConfig) Remove(d *dto.SmsBusinessConfigDeleteReq, p *actions.DataPermission) error {
	var data models.SmsBusinessConfig

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSmsBusinessConfig error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
