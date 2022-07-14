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

type Segments struct {
	api.Api
}

// GetPage 获取Segments列表
// @Summary 获取Segments列表
// @Description 获取Segments列表
// @Tags Segments
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Segments}} "{"code": 200, "data": [...]}"
// @Router /api/v1/segments [get]
// @Security Bearer
func (e Segments) GetPage(c *gin.Context) {
	req := dto.SegmentsGetPageReq{}
	s := service.Segments{}
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
	list := make([]models.Segments, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Segments 失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

func (e Segments) GetId(c *gin.Context) {
	req := dto.SegmentsGetReq{}
	s := service.Segments{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	number := s.GetId("test")
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Segments失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(number, "查询成功")
}

// Get 获取Segments
// @Summary 获取Segments
// @Description 获取Segments
// @Tags Segments
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.Segments} "{"code": 200, "data": [...]}"
// @Router /api/v1/segments/{id} [get]
// @Security Bearer
func (e Segments) Get(c *gin.Context) {
	req := dto.SegmentsGetReq{}
	s := service.Segments{}
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
	var object models.Segments

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Segments失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Segments
// @Summary 创建Segments
// @Description 创建Segments
// @Tags Segments
// @Accept application/json
// @Product application/json
// @Param data body dto.SegmentsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/segments [post]
// @Security Bearer
func (e Segments) Insert(c *gin.Context) {
	req := dto.SegmentsInsertReq{}
	s := service.Segments{}
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
		e.Error(500, err, fmt.Sprintf("创建Segments  失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Segments
// @Summary 修改Segments
// @Description 修改Segments
// @Tags Segments
// @Accept application/json
// @Product application/json
// @Param data body dto.SegmentsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/segments/{id} [put]
// @Security Bearer
func (e Segments) Update(c *gin.Context) {
	req := dto.SegmentsUpdateReq{}
	s := service.Segments{}
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
		e.Error(500, err, fmt.Sprintf("修改Segments 失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Segments
// @Summary 删除Segments
// @Description 删除Segments
// @Tags Segments
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/segments [delete]
// @Security Bearer
func (e Segments) Delete(c *gin.Context) {
	s := service.Segments{}
	req := dto.SegmentsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Segments失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
