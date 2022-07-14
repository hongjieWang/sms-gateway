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

type SmsSendLog struct {
	api.Api
}

// GetPage 获取短信发送记录列表
// @Summary 获取短信发送记录列表
// @Description 获取短信发送记录列表
// @Tags 短信发送记录
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SmsSendLog}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sms-send-log [get]
// @Security Bearer
func (e SmsSendLog) GetPage(c *gin.Context) {
	req := dto.SmsSendLogGetPageReq{}
	s := service.SmsSendLog{}
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
	list := make([]models.SmsSendLog, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取短信发送记录 失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取短信发送记录
// @Summary 获取短信发送记录
// @Description 获取短信发送记录
// @Tags 短信发送记录
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.SmsSendLog} "{"code": 200, "data": [...]}"
// @Router /api/v1/sms-send-log/{id} [get]
// @Security Bearer
func (e SmsSendLog) Get(c *gin.Context) {
	req := dto.SmsSendLogGetReq{}
	s := service.SmsSendLog{}
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
	var object models.SmsSendLog

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取短信发送记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建短信发送记录
// @Summary 创建短信发送记录
// @Description 创建短信发送记录
// @Tags 短信发送记录
// @Accept application/json
// @Product application/json
// @Param data body dto.SmsSendLogInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sms-send-log [post]
// @Security Bearer
func (e SmsSendLog) Insert(c *gin.Context) {
	req := dto.SmsSendLogInsertReq{}
	s := service.SmsSendLog{}
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
		e.Error(500, err, fmt.Sprintf("创建短信发送记录  失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改短信发送记录
// @Summary 修改短信发送记录
// @Description 修改短信发送记录
// @Tags 短信发送记录
// @Accept application/json
// @Product application/json
// @Param data body dto.SmsSendLogUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sms-send-log/{id} [put]
// @Security Bearer
func (e SmsSendLog) Update(c *gin.Context) {
	req := dto.SmsSendLogUpdateReq{}
	s := service.SmsSendLog{}
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
		e.Error(500, err, fmt.Sprintf("修改短信发送记录 失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除短信发送记录
// @Summary 删除短信发送记录
// @Description 删除短信发送记录
// @Tags 短信发送记录
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sms-send-log [delete]
// @Security Bearer
func (e SmsSendLog) Delete(c *gin.Context) {
	s := service.SmsSendLog{}
	req := dto.SmsSendLogDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除短信发送记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
