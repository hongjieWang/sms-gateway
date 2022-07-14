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

type SmsSendLog struct {
	service.Service
}

// GetPage 获取SmsSendLog列表
func (e *SmsSendLog) GetPage(c *dto.SmsSendLogGetPageReq, p *actions.DataPermission, list *[]models.SmsSendLog, count *int64) error {
	var err error
	var data models.SmsSendLog

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).Order("updated_at desc").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SmsSendLogService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SmsSendLog对象
func (e *SmsSendLog) Get(d *dto.SmsSendLogGetReq, p *actions.DataPermission, model *models.SmsSendLog) error {
	var data models.SmsSendLog

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSmsSendLog error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SmsSendLog对象
func (e *SmsSendLog) Insert(c *dto.SmsSendLogInsertReq) error {
	var err error
	var data models.SmsSendLog
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SmsSendLogService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SmsSendLog对象
func (e *SmsSendLog) Update(c *dto.SmsSendLogUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SmsSendLog{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		e.Log.Errorf("SmsSendLogService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SmsSendLog
func (e *SmsSendLog) Remove(d *dto.SmsSendLogDeleteReq, p *actions.DataPermission) error {
	var data models.SmsSendLog

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSmsSendLog error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
