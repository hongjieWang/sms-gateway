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

type TArticle struct {
	service.Service
}

// GetPage 获取TArticle列表
func (e *TArticle) GetPage(c *dto.TArticleGetPageReq, p *actions.DataPermission, list *[]models.TArticle, count *int64) error {
	var err error
	var data models.TArticle

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TArticleService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TArticle对象
func (e *TArticle) Get(d *dto.TArticleGetReq, p *actions.DataPermission, model *models.TArticle) error {
	var data models.TArticle

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTArticle error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TArticle对象
func (e *TArticle) Insert(c *dto.TArticleInsertReq) error {
    var err error
    var data models.TArticle
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TArticleService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TArticle对象
func (e *TArticle) Update(c *dto.TArticleUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.TArticle{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if db.Error != nil {
        e.Log.Errorf("TArticleService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除TArticle
func (e *TArticle) Remove(d *dto.TArticleDeleteReq, p *actions.DataPermission) error {
	var data models.TArticle

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveTArticle error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}