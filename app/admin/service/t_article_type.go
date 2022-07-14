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

type TArticleType struct {
	service.Service
}

// GetPage 获取TArticleType列表
func (e *TArticleType) GetPage(c *dto.TArticleTypeGetPageReq, p *actions.DataPermission, list *[]models.TArticleType, count *int64) error {
	var err error
	var data models.TArticleType

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TArticleTypeService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TArticleType对象
func (e *TArticleType) Get(d *dto.TArticleTypeGetReq, p *actions.DataPermission, model *models.TArticleType) error {
	var data models.TArticleType

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTArticleType error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TArticleType对象
func (e *TArticleType) Insert(c *dto.TArticleTypeInsertReq) error {
    var err error
    var data models.TArticleType
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TArticleTypeService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TArticleType对象
func (e *TArticleType) Update(c *dto.TArticleTypeUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.TArticleType{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if db.Error != nil {
        e.Log.Errorf("TArticleTypeService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除TArticleType
func (e *TArticleType) Remove(d *dto.TArticleTypeDeleteReq, p *actions.DataPermission) error {
	var data models.TArticleType

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveTArticleType error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}