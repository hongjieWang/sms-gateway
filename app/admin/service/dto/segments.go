package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SegmentsGetPageReq struct {
	dto.Pagination `search:"-"`
	SegmentsOrder
}

type SegmentsOrder struct {
	BizTag     string `form:"bizTagOrder"  search:"type:order;column:biz_tag;table:segments"`
	MaxId      string `form:"maxIdOrder"  search:"type:order;column:max_id;table:segments"`
	Step       string `form:"stepOrder"  search:"type:order;column:step;table:segments"`
	Remark     string `form:"remarkOrder"  search:"type:order;column:remark;table:segments"`
	CreateTime string `form:"createTimeOrder"  search:"type:order;column:create_time;table:segments"`
	UpdateTime string `form:"updateTimeOrder"  search:"type:order;column:update_time;table:segments"`
}

func (m *SegmentsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SegmentsInsertReq struct {
	BizTag     string `json:"bizTag" comment:""`
	MaxId      int64  `json:"maxId" comment:""`
	Step       string `json:"step" comment:""`
	Remark     string `json:"remark" comment:""`
	CreateTime string `json:"createTime" comment:""`
	UpdateTime string `json:"updateTime" comment:""`
	common.ControlBy
}

func (s *SegmentsInsertReq) Generate(model *models.Segments) {
	model.BizTag = s.BizTag
	model.MaxId = s.MaxId
	model.Step = s.Step
	model.Remark = s.Remark
}

func (s *SegmentsInsertReq) GetId() interface{} {
	return s.BizTag
}

type SegmentsUpdateReq struct {
	BizTag     string `uri:"bizTag" comment:""` //
	MaxId      int64  `json:"maxId" comment:""`
	Step       string `json:"step" comment:""`
	Remark     string `json:"remark" comment:""`
	CreateTime string `json:"createTime" comment:""`
	UpdateTime string `json:"updateTime" comment:""`
	common.ControlBy
}

func (s *SegmentsUpdateReq) Generate(model *models.Segments) {
	model.BizTag = s.BizTag
	model.MaxId = s.MaxId
	model.Step = s.Step
	model.Remark = s.Remark
}

func (s *SegmentsUpdateReq) GetId() interface{} {
	return s.BizTag
}

// SegmentsGetReq 功能获取请求参数
type SegmentsGetReq struct {
	BizTag string `uri:"bizTag"`
}

func (s *SegmentsGetReq) GetId() interface{} {
	return s.BizTag
}

// SegmentsDeleteReq 功能删除请求参数
type SegmentsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SegmentsDeleteReq) GetId() interface{} {
	return s.Ids
}
