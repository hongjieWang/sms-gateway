package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	tdmq "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdmq/v20200217"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"net/http"
)

type Mq struct {
	api.Api
}

// DescribeClusters
// @Summary 获取集群列表
// @Description 获取JSON
// @Tags TDMQ
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/mq/describe-clusters [get]
// @Security Bearer
func (e Mq) DescribeClusters(c *gin.Context) {
	s := service.Mq{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	response, err := s.DescribeClusters()
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(response.Response, "查询成功")
}

func (e Mq) SelectClusters(c *gin.Context) {
	s := service.Mq{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	response, err := s.DescribeClusters()
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	l := make([]dto.SelectClusters, 0)
	for _, i := range response.Response.ClusterSet {
		d := dto.SelectClusters{
			ClusterId:   *i.ClusterId,
			ClusterName: *i.ClusterName,
		}
		l = append(l, d)
	}
	e.OK(l, "查询成功")
}

func (e Mq) PageDescribeClusters(c *gin.Context) {
	s := service.Mq{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	list := make([]tdmq.Cluster, 0)
	var count int64
	response, err := s.DescribeClusters()
	count = *response.Response.TotalCount
	for _, cluster := range response.Response.ClusterSet {
		list = append(list, *cluster)
	}
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.PageOK(list, int(count), 1, 20, "查询成功")
}

// SelectEnvironments 命名空间下拉选择
func (e Mq) SelectEnvironments(c *gin.Context) {
	s := service.Mq{}
	cluster := dto.Clusters{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&cluster, nil).
		MakeService(&s.Service).
		Errors
	response, err := s.DescribeEnvironments(cluster.ClusterId)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	l := make([]dto.SelectEnvironments, 0)
	for _, i := range response.Response.EnvironmentSet {
		d := dto.SelectEnvironments{
			EnvironmentId: *i.EnvironmentId,
			Remark:        *i.EnvironmentId + "[" + *i.Remark + "]",
		}
		l = append(l, d)
	}
	e.OK(l, "查询成功")
}

func (e Mq) DescribeEnvironments(c *gin.Context) {
	s := service.Mq{}
	cluster := dto.Clusters{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&cluster, nil).
		MakeService(&s.Service).
		Errors
	response, err := s.DescribeEnvironments(cluster.ClusterId)
	list := make([]tdmq.Environment, 0)
	var count uint64
	count = *response.Response.TotalCount
	for _, environment := range response.Response.EnvironmentSet {
		list = append(list, *environment)
	}
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.PageOK(list, int(count), 1, 20, "查询成功")
}

// CreateEnvironments 创建命名空间
func (e Mq) CreateEnvironments(c *gin.Context) {
	s := service.Mq{}
	environment := dto.Environment{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&environment, binding.JSON).
		MakeService(&s.Service).
		Errors
	response, err := s.CreateEnvironments(&environment)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(response.Response, "创建命名空间成功")
}

// DeleteEnvironments 删除命名空间
func (e Mq) DeleteEnvironments(c *gin.Context) {
	s := service.Mq{}
	environment := dto.Environment{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&environment, binding.JSON).
		MakeService(&s.Service).
		Errors
	response, err := s.DeleteEnvironments(&environment)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(response.Response, "命名空间删除成功")
}

// CreateTopic 创建Topic
func (e Mq) CreateTopic(c *gin.Context) {
	s := service.Mq{}
	topic := dto.CreateTopic{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&topic, binding.JSON).
		MakeService(&s.Service).
		Errors
	response, err := s.CreateTopic(&topic)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "创建失败")
		return
	}
	e.OK(response.Response, "Topic创建成功")
}

// DescribeTopics 查询Topic
func (e Mq) DescribeTopics(c *gin.Context) {
	s := service.Mq{}
	describeTopics := dto.DescribeTopics{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&describeTopics, binding.JSON).
		MakeService(&s.Service).
		Errors
	response, err := s.DescribeTopics(&describeTopics)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(response.Response, "Topic查询成功")
}

// PageDescribeTopics Topic 分页查询
func (e Mq) PageDescribeTopics(c *gin.Context) {
	s := service.Mq{}
	describeTopics := dto.DescribeTopics{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&describeTopics, binding.JSON).
		MakeService(&s.Service).
		Errors
	response, err := s.DescribeTopics(&describeTopics)
	list := make([]tdmq.Topic, 0)
	var count uint64
	if err != nil {
		e.PageOK(list, 0, 1, 20, "查询成功")
		return
	}
	count = *response.Response.TotalCount
	for _, topic := range response.Response.TopicSets {
		list = append(list, *topic)
	}
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.PageOK(list, int(count), 1, 20, "查询成功")
}

// DeleteTopic 删除Topic
func (e Mq) DeleteTopic(c *gin.Context) {
	s := service.Mq{}
	deleteTopics := dto.DeleteTopics{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&deleteTopics, binding.JSON).
		MakeService(&s.Service).
		Errors
	response, err := s.DeleteTopics(&deleteTopics)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(response.Response, "Topic删除成功")
}
