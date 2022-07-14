package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SmsTemplateConfig struct {
	api.Api
}

// GetPage 获取短信模版配置列表
// @Summary 获取短信模版配置列表
// @Description 获取短信模版配置列表
// @Tags 短信模版配置
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SmsTemplateConfig}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sms-template-config [get]
// @Security Bearer
func (e SmsTemplateConfig) GetPage(c *gin.Context) {
	req := dto.SmsTemplateConfigGetPageReq{}
	s := service.SmsTemplateConfig{}
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
	list := make([]models.SmsTemplateConfig, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取短信模版配置 失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取短信模版配置
// @Summary 获取短信模版配置
// @Description 获取短信模版配置
// @Tags 短信模版配置
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.SmsTemplateConfig} "{"code": 200, "data": [...]}"
// @Router /api/v1/sms-template-config/{id} [get]
// @Security Bearer
func (e SmsTemplateConfig) Get(c *gin.Context) {
	req := dto.SmsTemplateConfigGetReq{}
	s := service.SmsTemplateConfig{}
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
	var object models.SmsTemplateConfig

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取短信模版配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// GetByTemplateNo 根据短信模版查询数据
func (e SmsTemplateConfig) GetByTemplateNo(c *gin.Context) {
	req := dto.SmsTemplateConfigGetByTemplateNoReq{}
	s := service.SmsTemplateConfig{}
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
	var object models.SmsTemplateConfig
	err = s.GetByTemplateNo(req.TemplateNo, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取短信模版配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建短信模版配置
// @Summary 创建短信模版配置
// @Description 创建短信模版配置
// @Tags 短信模版配置
// @Accept application/json
// @Product application/json
// @Param data body dto.SmsTemplateConfigInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sms-template-config [post]
// @Security Bearer
func (e SmsTemplateConfig) Insert(c *gin.Context) {
	req := dto.SmsTemplateConfigInsertReq{}
	s := service.SmsTemplateConfig{}
	seg := service.Segments{}
	sign := service.SmsSignConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).MakeService(&seg.Service).MakeService(&sign.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	//生成模版编号
	id := seg.GetId("smsTemplate")
	req.TemplateNo = fmt.Sprintf("T-%06d", id)
	//根据签名名称查询相关信息
	name := sign.GetBySignName(req.SignName)
	req.ProviderNo = name.ProviderNo
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建短信模版配置  失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Update 修改短信模版配置
// @Summary 修改短信模版配置
// @Description 修改短信模版配置
// @Tags 短信模版配置
// @Accept application/json
// @Product application/json
// @Param data body dto.SmsTemplateConfigUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sms-template-config/{id} [put]
// @Security Bearer
func (e SmsTemplateConfig) Update(c *gin.Context) {
	req := dto.SmsTemplateConfigUpdateReq{}
	s := service.SmsTemplateConfig{}
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
		e.Error(500, err, fmt.Sprintf("修改短信模版配置 失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除短信模版配置
// @Summary 删除短信模版配置
// @Description 删除短信模版配置
// @Tags 短信模版配置
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sms-template-config [delete]
// @Security Bearer
func (e SmsTemplateConfig) Delete(c *gin.Context) {
	s := service.SmsTemplateConfig{}
	req := dto.SmsTemplateConfigDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除短信模版配置失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
