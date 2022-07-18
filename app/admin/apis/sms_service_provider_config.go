package apis

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SmsServiceProviderConfig struct {
	api.Api
}

// GetPage 获取服务商配置列表
// @Summary 获取服务商配置列表
// @Description 获取服务商配置列表
// @Tags 服务商配置
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SmsServiceProviderConfig}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sms-service-provider-config [get]
// @Security Bearer
func (e SmsServiceProviderConfig) GetPage(c *gin.Context) {
	req := dto.SmsServiceProviderConfigGetPageReq{}
	s := service.SmsServiceProviderConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.SmsServiceProviderConfig, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取服务商配置 失败，\r\n失败信息 %s", err.Error()))
		return
	}
	for i := range list {
		list[i].AccessKeySecret = "**********************************"
	}
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取服务商配置
// @Summary 获取服务商配置
// @Description 获取服务商配置
// @Tags 服务商配置
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.SmsServiceProviderConfig} "{"code": 200, "data": [...]}"
// @Router /api/v1/sms-service-provider-config/{id} [get]
// @Security Bearer
func (e SmsServiceProviderConfig) Get(c *gin.Context) {
	req := dto.SmsServiceProviderConfigGetReq{}
	s := service.SmsServiceProviderConfig{}
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
	var object models.SmsServiceProviderConfig

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取服务商配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建服务商配置
// @Summary 创建服务商配置
// @Description 创建服务商配置
// @Tags 服务商配置
// @Accept application/json
// @Product application/json
// @Param data body dto.SmsServiceProviderConfigInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sms-service-provider-config [post]
// @Security Bearer
func (e SmsServiceProviderConfig) Insert(c *gin.Context) {
	req := dto.SmsServiceProviderConfigInsertReq{}
	s := service.SmsServiceProviderConfig{}
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
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建服务商配置  失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改服务商配置
// @Summary 修改服务商配置
// @Description 修改服务商配置
// @Tags 服务商配置
// @Accept application/json
// @Product application/json
// @Param data body dto.SmsServiceProviderConfigUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sms-service-provider-config/{id} [put]
// @Security Bearer
func (e SmsServiceProviderConfig) Update(c *gin.Context) {
	req := dto.SmsServiceProviderConfigUpdateReq{}
	s := service.SmsServiceProviderConfig{}
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
		e.Error(500, err, fmt.Sprintf("修改服务商配置 失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除服务商配置
// @Summary 删除服务商配置
// @Description 删除服务商配置
// @Tags 服务商配置
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sms-service-provider-config [delete]
// @Security Bearer
func (e SmsServiceProviderConfig) Delete(c *gin.Context) {
	s := service.SmsServiceProviderConfig{}
	req := dto.SmsServiceProviderConfigDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除服务商配置失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
