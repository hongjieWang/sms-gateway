package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	tdmq "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdmq/v20200217"
	"go-admin/app/admin/service/dto"
)

type Mq struct {
	service.Service
}

// DescribeClusters 获取集群列表
func (e *Mq) DescribeClusters() (model *tdmq.DescribeClustersResponse, err error) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewDescribeClustersRequest()
	return client.DescribeClusters(request)
}

// DescribeEnvironments 获取指定集群下命名空间
func (e *Mq) DescribeEnvironments(clusterId string) (model *tdmq.DescribeEnvironmentsResponse, err error) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewDescribeEnvironmentsRequest()
	request.ClusterId = common.StringPtr(clusterId)
	return client.DescribeEnvironments(request)
}

// CreateEnvironments 创建命名空间
func (e *Mq) CreateEnvironments(environment *dto.Environment) (model *tdmq.CreateEnvironmentResponse, err error) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewCreateEnvironmentRequest()
	request.ClusterId = environment.ClusterId
	request.EnvironmentId = environment.EnvironmentId
	request.Remark = environment.Remark
	request.MsgTTL = environment.MsgTTL
	return client.CreateEnvironment(request)
}

// DeleteEnvironments 删除命名空间
func (e *Mq) DeleteEnvironments(environment *dto.Environment) (model *tdmq.DeleteEnvironmentsResponse, err error) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewDeleteEnvironmentsRequest()
	request.ClusterId = environment.ClusterId
	request.EnvironmentIds = environment.EnvironmentIds
	return client.DeleteEnvironments(request)
}

// CreateTopic 创建Topic
func (e *Mq) CreateTopic(topic *dto.CreateTopic) (model *tdmq.CreateTopicResponse, err error) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewCreateTopicRequest()
	request.EnvironmentId = topic.EnvironmentId
	request.ClusterId = topic.ClusterId
	request.TopicName = topic.TopicName
	request.Remark = topic.Remark
	request.Partitions = topic.Partitions
	request.TopicType = topic.TopicType
	return client.CreateTopic(request)
}

// DescribeTopics 查询Topic
func (e *Mq) DescribeTopics(topic *dto.DescribeTopics) (response *tdmq.DescribeTopicsResponse, err error) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewDescribeTopicsRequest()
	request.EnvironmentId = topic.EnvironmentId
	request.ClusterId = topic.ClusterId
	request.TopicName = topic.TopicName
	request.Limit = topic.Limit
	request.Offset = topic.Offset
	return client.DescribeTopics(request)
}

// DeleteTopics 删除topic
func (e *Mq) DeleteTopics(topic *dto.DeleteTopics) (response *tdmq.DeleteTopicsResponse, err error) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewDeleteTopicsRequest()
	request.TopicSets = append(request.TopicSets, &tdmq.TopicRecord{
		EnvironmentId: topic.EnvironmentId,
		TopicName:     topic.TopicName,
	})
	request.ClusterId = topic.ClusterId
	request.EnvironmentId = topic.EnvironmentId
	return client.DeleteTopics(request)
}
