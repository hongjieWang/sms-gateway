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

type SmsBusinessConfig struct {
	api.Api
}

// GetPage 获取业务配置列表
// @Summary 获取业务配置列表
// @Description 获取业务配置列表
// @Tags 业务配置
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SmsBusinessConfig}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sms-business-config [get]
// @Security Bearer
func (e SmsBusinessConfig) GetPage(c *gin.Context) {
	req := dto.SmsBusinessConfigGetPageReq{}
	s := service.SmsBusinessConfig{}
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
	list := make([]models.SmsBusinessConfig, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取业务配置 失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取业务配置
// @Summary 获取业务配置
// @Description 获取业务配置
// @Tags 业务配置
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.SmsBusinessConfig} "{"code": 200, "data": [...]}"
// @Router /api/v1/sms-business-config/{id} [get]
// @Security Bearer
func (e SmsBusinessConfig) Get(c *gin.Context) {
	req := dto.SmsBusinessConfigGetReq{}
	s := service.SmsBusinessConfig{}
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
	var object models.SmsBusinessConfig

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取业务配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建业务配置
// @Summary 创建业务配置
// @Description 创建业务配置
// @Tags 业务配置
// @Accept application/json
// @Product application/json
// @Param data body dto.SmsBusinessConfigInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sms-business-config [post]
// @Security Bearer
func (e SmsBusinessConfig) Insert(c *gin.Context) {
	req := dto.SmsBusinessConfigInsertReq{}
	s := service.SmsBusinessConfig{}
	seq := service.Segments{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).MakeService(&seq.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	if req.BusinessNo == "" {
		// 设置业务编号
		id := seq.GetId("business")
		req.BusinessNo = fmt.Sprintf("B-%06d", id)
	}
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建业务配置  失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改业务配置
// @Summary 修改业务配置
// @Description 修改业务配置
// @Tags 业务配置
// @Accept application/json
// @Product application/json
// @Param data body dto.SmsBusinessConfigUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sms-business-config/{id} [put]
// @Security Bearer
func (e SmsBusinessConfig) Update(c *gin.Context) {
	req := dto.SmsBusinessConfigUpdateReq{}
	s := service.SmsBusinessConfig{}
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
		e.Error(500, err, fmt.Sprintf("修改业务配置 失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除业务配置
// @Summary 删除业务配置
// @Description 删除业务配置
// @Tags 业务配置
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sms-business-config [delete]
// @Security Bearer
func (e SmsBusinessConfig) Delete(c *gin.Context) {
	s := service.SmsBusinessConfig{}
	req := dto.SmsBusinessConfigDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除业务配置失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
