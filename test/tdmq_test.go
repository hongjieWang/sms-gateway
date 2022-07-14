package test

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	tdmq "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdmq/v20200217"
	"testing"
)

func TestGoDescribeClusters(t *testing.T) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewDescribeClustersRequest()
	_, err = client.DescribeClusters(request)
	if err != nil {
		panic(err)
	}

}

//获取租户下命名空间列表
func TestDescribeEnvironments(t *testing.T) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewDescribeEnvironmentsRequest()
	request.ClusterId = common.StringPtr("pulsar-jm95pr3kmme7")
	_, err = client.DescribeEnvironments(request)
	if err != nil {
		panic(err)
	}
}

// 创建命名空间
func TestCreateEnvironments(t *testing.T) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewCreateEnvironmentRequest()
	request.ClusterId = common.StringPtr("pulsar-jm95pr3kmme7")
	request.EnvironmentId = common.StringPtr("go-test")
	request.Remark = common.StringPtr("go-api测试")
	request.MsgTTL = common.Uint64Ptr(200)
	environment, err := client.CreateEnvironment(request)
	if err != nil {
		panic(err)
	}
	println(environment)
}

//删除命名空间
func TestDeleteEnvironments(t *testing.T) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewDeleteEnvironmentsRequest()
	request.ClusterId = common.StringPtr("pulsar-jm95pr3kmme7")
	request.EnvironmentIds = common.StringPtrs([]string{"go-test"})
	environment, err := client.DeleteEnvironments(request)
	if err != nil {
		panic(err)
	}
	println(environment)
}

//创建Topic
func TestCreateTopic(t *testing.T) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewCreateTopicRequest()
	request.EnvironmentId = common.StringPtr("go-test")
	request.ClusterId = common.StringPtr("pulsar-jm95pr3kmme7")
	request.TopicName = common.StringPtr("go-test-topic")
	request.Remark = common.StringPtr("go-test-topic")
	request.Partitions = common.Uint64Ptr(1)
	request.TopicType = common.Uint64Ptr(0)
	topic, err := client.CreateTopic(request)
	if err != nil {
		panic(err)
	}
	println(topic)
}

// 查询主题列表
func TestDescribeTopics(t *testing.T) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewDescribeTopicsRequest()
	request.EnvironmentId = common.StringPtr("go-test")
	request.ClusterId = common.StringPtr("pulsar-jm95pr3kmme7")
	request.TopicType = common.Uint64Ptr(0)
	topic, err := client.DescribeTopics(request)
	if err != nil {
		panic(err)
	}
	println(topic)
}

//删除主题
func TestDeleteTopics(t *testing.T) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewDeleteTopicsRequest()
	request.EnvironmentId = common.StringPtr("go-test")
	request.ClusterId = common.StringPtr("pulsar-jm95pr3kmme7")
	sets := request.TopicSets
	sets = append(sets, &tdmq.TopicRecord{
		EnvironmentId: common.StringPtr("go-test"),
		TopicName:     common.StringPtr("go-test-topic"),
	})
	request.TopicSets = sets
	topic, err := client.DeleteTopics(request)
	if err != nil {
		panic(err)
	}
	println(topic)
}

//发送消息测试
func TestSendMessages(t *testing.T) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewSendMessagesRequest()
	request.Topic = common.StringPtr("pulsar-jm95pr3kmme7/go-test/go-test-whj")
	request.Payload = common.StringPtr("Hello TDMQ BY GOLANG")
	topic, err := client.SendMessages(request)
	if err != nil {
		panic(err)
	}
	println(topic)
}

//获取消息列表
func TestReceiveMessage(t *testing.T) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewReceiveMessageRequest()
	request.Topic = common.StringPtr("pulsar-jm95pr3kmme7/go-test/go-test-whj")
	request.SubscriptionName = common.StringPtr("subscriptionName")
	topic, err := client.ReceiveMessage(request)
	if err != nil {
		panic(err)
	}
	println(topic)
}

//确认消息
func TestAcknowledgeMessage(t *testing.T) {
	credential := common.NewCredential("AKIDf8KFuIQHq89tLWQ3OTO2PUJr88sKVO2P", "8zb3OcgcNxMadXPBk47tOcZFenfGady9")
	client, err := tdmq.NewClient(credential, regions.Nanjing, profile.NewClientProfile())
	if err != nil {
		panic(err)
	}
	request := tdmq.NewAcknowledgeMessageRequest()
	request.MessageId = common.StringPtr("393040:0:0")
	request.AckTopic = common.StringPtr("pulsar-jm95pr3kmme7/go-test/go-test-topic")
	request.SubName = common.StringPtr("subscriptionName")
	topic, err := client.AcknowledgeMessage(request)
	if err != nil {
		panic(err)
	}
	println(topic)
}
