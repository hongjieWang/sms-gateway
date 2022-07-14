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

type TArticleType struct {
	api.Api
}

// GetPage 获取文章分类表列表
// @Summary 获取文章分类表列表
// @Description 获取文章分类表列表
// @Tags 文章分类表
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.TArticleType}} "{"code": 200, "data": [...]}"
// @Router /api/v1/t-article-type [get]
// @Security Bearer
func (e TArticleType) GetPage(c *gin.Context) {
    req := dto.TArticleTypeGetPageReq{}
    s := service.TArticleType{}
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
	list := make([]models.TArticleType, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取文章分类表 失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取文章分类表
// @Summary 获取文章分类表
// @Description 获取文章分类表
// @Tags 文章分类表
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.TArticleType} "{"code": 200, "data": [...]}"
// @Router /api/v1/t-article-type/{id} [get]
// @Security Bearer
func (e TArticleType) Get(c *gin.Context) {
	req := dto.TArticleTypeGetReq{}
	s := service.TArticleType{}
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
	var object models.TArticleType

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取文章分类表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建文章分类表
// @Summary 创建文章分类表
// @Description 创建文章分类表
// @Tags 文章分类表
// @Accept application/json
// @Product application/json
// @Param data body dto.TArticleTypeInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/t-article-type [post]
// @Security Bearer
func (e TArticleType) Insert(c *gin.Context) {
    req := dto.TArticleTypeInsertReq{}
    s := service.TArticleType{}
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
		e.Error(500, err, fmt.Sprintf("创建文章分类表  失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改文章分类表
// @Summary 修改文章分类表
// @Description 修改文章分类表
// @Tags 文章分类表
// @Accept application/json
// @Product application/json
// @Param data body dto.TArticleTypeUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/t-article-type/{id} [put]
// @Security Bearer
func (e TArticleType) Update(c *gin.Context) {
    req := dto.TArticleTypeUpdateReq{}
    s := service.TArticleType{}
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
		e.Error(500, err, fmt.Sprintf("修改文章分类表 失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除文章分类表
// @Summary 删除文章分类表
// @Description 删除文章分类表
// @Tags 文章分类表
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/t-article-type [delete]
// @Security Bearer
func (e TArticleType) Delete(c *gin.Context) {
    s := service.TArticleType{}
    req := dto.TArticleTypeDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除文章分类表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}