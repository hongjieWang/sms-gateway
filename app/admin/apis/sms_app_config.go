package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SmsAppConfig struct {
	api.Api
}

// GetPage 获取应用管理配置列表
// @Summary 获取应用管理配置列表
// @Description 获取应用管理配置列表
// @Tags 应用管理配置
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SmsAppConfig}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sms-app-config [get]
// @Security Bearer
func (e SmsAppConfig) GetPage(c *gin.Context) {
	req := dto.SmsAppConfigGetPageReq{}
	s := service.SmsAppConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.SmsAppConfig, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取应用管理配置 失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取应用管理配置
// @Summary 获取应用管理配置
// @Description 获取应用管理配置
// @Tags 应用管理配置
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.SmsAppConfig} "{"code": 200, "data": [...]}"
// @Router /api/v1/sms-app-config/{id} [get]
// @Security Bearer
func (e SmsAppConfig) Get(c *gin.Context) {
	req := dto.SmsAppConfigGetReq{}
	s := service.SmsAppConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SmsAppConfig

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取应用管理配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建应用管理配置
// @Summary 创建应用管理配置
// @Description 创建应用管理配置
// @Tags 应用管理配置
// @Accept application/json
// @Product application/json
// @Param data body dto.SmsAppConfigInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sms-app-config [post]
// @Security Bearer
func (e SmsAppConfig) Insert(c *gin.Context) {
	req := dto.SmsAppConfigInsertReq{}
	s := service.SmsAppConfig{}
	seg := service.Segments{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).MakeService(&seg.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	//设置应用编号
	if req.AppNo == "" {
		appNo := seg.GetId("app")
		req.AppNo = fmt.Sprintf("A-%06d", appNo)
	}
	req.UseNumber = 0
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建应用管理配置  失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改应用管理配置
// @Summary 修改应用管理配置
// @Description 修改应用管理配置
// @Tags 应用管理配置
// @Accept application/json
// @Product application/json
// @Param data body dto.SmsAppConfigUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sms-app-config/{id} [put]
// @Security Bearer
func (e SmsAppConfig) Update(c *gin.Context) {
	req := dto.SmsAppConfigUpdateReq{}
	s := service.SmsAppConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改应用管理配置 失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除应用管理配置
// @Summary 删除应用管理配置
// @Description 删除应用管理配置
// @Tags 应用管理配置
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sms-app-config [delete]
// @Security Bearer
func (e SmsAppConfig) Delete(c *gin.Context) {
	s := service.SmsAppConfig{}
	req := dto.SmsAppConfigDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除应用管理配置失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
