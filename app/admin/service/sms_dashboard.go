package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/models"
)

type SmsDashboard struct {
	service.Service
}

func (d *SmsDashboard) GetAllCount() int64 {
	var smsSendLog models.SmsSendLog
	count := int64(0)
	err := d.Orm.Model(smsSendLog).Count(&count).Error
	if err != nil {
		d.Log.Errorf("SmsDashboard GetAllCount error:%s \r\n", err)
		return 0
	}
	return count
}

func (d *SmsDashboard) GetSuccessCount() int64 {
	var smsSendLog models.SmsSendLog
	count := int64(0)
	err := d.Orm.Model(smsSendLog).Where(models.SmsSendLog{
		Code: "OK",
	}).Count(&count).Error
	if err != nil {
		d.Log.Errorf("SmsDashboard GetSuccessCount error:%s \r\n", err)
		return 0
	}
	return count
}

func (d *SmsDashboard) GetRequestCount() int64 {
	var sysOperaLog models.SysOperaLog
	count := int64(0)
	err := d.Orm.Model(sysOperaLog).Count(&count).Error
	if err != nil {
		d.Log.Errorf("SmsDashboard GetRequestCount error:%s \r\n", err)
		return 0
	}
	return count
}
