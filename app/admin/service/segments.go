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

type Segments struct {
	service.Service
}

// GetPage 获取Segments列表
func (e *Segments) GetPage(c *dto.SegmentsGetPageReq, p *actions.DataPermission, list *[]models.Segments, count *int64) error {
	var err error
	var data models.Segments

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SegmentsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Segments对象
func (e *Segments) Get(d *dto.SegmentsGetReq, p *actions.DataPermission, model *models.Segments) error {
	var data models.Segments

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSegments error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

func (e *Segments) GetId(bizTag string) int64 {
	var data models.Segments
	var model models.Segments
	err := e.Orm.Model(&data).Where(models.Segments{BizTag: bizTag}).
		First(&model).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return 0
	}
	model.MaxId = model.MaxId + 1
	e.Orm.Save(&model)
	return model.MaxId
}

// Insert 创建Segments对象
func (e *Segments) Insert(c *dto.SegmentsInsertReq) error {
	var err error
	var data models.Segments
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SegmentsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Segments对象
func (e *Segments) Update(c *dto.SegmentsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Segments{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		e.Log.Errorf("SegmentsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Segments
func (e *Segments) Remove(d *dto.SegmentsDeleteReq, p *actions.DataPermission) error {
	var data models.Segments

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSegments error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
