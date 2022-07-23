package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/admin/service"
)

type SmsDashboard struct {
	api.Api
}

func (e SmsDashboard) GetAllCount(c *gin.Context) {
	s := service.SmsDashboard{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	count := s.GetAllCount()
	e.OK(count, "查询成功")
}

func (e SmsDashboard) GetSuccessCount(c *gin.Context) {
	s := service.SmsDashboard{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	count := s.GetSuccessCount()
	e.OK(count, "查询成功")
}

func (e SmsDashboard) GetRequestCount(c *gin.Context) {
	s := service.SmsDashboard{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	count := s.GetRequestCount()
	e.OK(count, "查询成功")
}
