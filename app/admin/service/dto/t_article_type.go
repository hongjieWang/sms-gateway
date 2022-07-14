package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TArticleTypeGetPageReq struct {
	dto.Pagination     `search:"-"`
    TArticleTypeOrder
}

type TArticleTypeOrder struct {Id int `form:"idOrder"  search:"type:order;column:id;table:t_article_type"`
    TypeNo string `form:"typeNoOrder"  search:"type:order;column:type_no;table:t_article_type"`
    TypeName string `form:"typeNameOrder"  search:"type:order;column:type_name;table:t_article_type"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:t_article_type"`
    
}

func (m *TArticleTypeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TArticleTypeInsertReq struct {
    Id int `json:"-" comment:"主键"` // 主键
    TypeNo string `json:"typeNo" comment:"分类编号"`
    TypeName string `json:"typeName" comment:"分类名称"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *TArticleTypeInsertReq) Generate(model *models.TArticleType)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.TypeNo = s.TypeNo
    model.TypeName = s.TypeName
    model.Status = s.Status
}

func (s *TArticleTypeInsertReq) GetId() interface{} {
	return s.Id
}

type TArticleTypeUpdateReq struct {
    Id int `uri:"id" comment:"主键"` // 主键
    TypeNo string `json:"typeNo" comment:"分类编号"`
    TypeName string `json:"typeName" comment:"分类名称"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *TArticleTypeUpdateReq) Generate(model *models.TArticleType)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.TypeNo = s.TypeNo
    model.TypeName = s.TypeName
    model.Status = s.Status
}

func (s *TArticleTypeUpdateReq) GetId() interface{} {
	return s.Id
}

// TArticleTypeGetReq 功能获取请求参数
type TArticleTypeGetReq struct {
     Id int `uri:"id"`
}
func (s *TArticleTypeGetReq) GetId() interface{} {
	return s.Id
}

// TArticleTypeDeleteReq 功能删除请求参数
type TArticleTypeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TArticleTypeDeleteReq) GetId() interface{} {
	return s.Ids
}