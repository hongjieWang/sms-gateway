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

type SmsSignConfig struct {
	service.Service
}

// GetPage 获取SmsSignConfig列表
func (e *SmsSignConfig) GetPage(c *dto.SmsSignConfigGetPageReq, p *actions.DataPermission, list *[]models.SmsSignConfig, count *int64) error {
	var err error
	var data models.SmsSignConfig

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).Order("updated_at desc").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SmsSignConfigService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SmsSignConfig对象
func (e *SmsSignConfig) Get(d *dto.SmsSignConfigGetReq, p *actions.DataPermission, model *models.SmsSignConfig) error {
	var data models.SmsSignConfig

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSmsSignConfig error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetBySignName 根据签名名称查找签名数据
func (e *SmsSignConfig) GetBySignName(signName string) models.SmsSignConfig {
	var data models.SmsSignConfig
	var model models.SmsSignConfig
	e.Orm.Model(&data).Where(models.SmsSignConfig{SignName: signName}).
		First(&model)
	return model
}

// Insert 创建SmsSignConfig对象
func (e *SmsSignConfig) Insert(c *dto.SmsSignConfigInsertReq) error {
	var err error
	var data models.SmsSignConfig
	c.Generate(&data)

	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SmsSignConfigService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SmsSignConfig对象
func (e *SmsSignConfig) Update(c *dto.SmsSignConfigUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SmsSignConfig{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		e.Log.Errorf("SmsSignConfigService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SmsSignConfig
func (e *SmsSignConfig) Remove(d *dto.SmsSignConfigDeleteReq, p *actions.DataPermission) error {
	var data models.SmsSignConfig

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSmsSignConfig error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
