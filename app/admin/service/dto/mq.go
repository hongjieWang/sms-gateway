package dto

type Clusters struct {
	ClusterId string `uri:"clusterId"`
}

type SelectClusters struct {
	ClusterId   string `json:"value"`
	ClusterName string `json:"label"`
}

type SelectEnvironments struct {
	EnvironmentId string `json:"value"`
	Remark        string `json:"label"`
}

type Environment struct {
	// 环境（命名空间）名称，不支持中字以及除了短线和下划线外的特殊字符且不超过16个字符。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 未消费消息过期时间，单位：秒，最小60，最大1296000，（15天）。
	MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

	// 说明，128个字符以内。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 消息保留策略
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitempty" name:"RetentionPolicy"`

	// 删除命名空间 最多20个
	EnvironmentIds []*string `json:"EnvironmentIds,omitempty" name:"EnvironmentIds"`
}

type RetentionPolicy struct {

	// 消息保留时长
	TimeInMinutes *int64 `json:"TimeInMinutes,omitempty" name:"TimeInMinutes"`

	// 消息保留大小
	SizeInMB *int64 `json:"SizeInMB,omitempty" name:"SizeInMB"`
}

type CreateTopic struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名，不支持中字以及除了短线和下划线外的特殊字符且不超过64个字符。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 0：非分区topic，无分区；非0：具体分区topic的分区数，最大不允许超过128。
	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

	// 0： 普通消息；
	// 1 ：全局顺序消息；
	// 2 ：局部顺序消息；
	// 3 ：重试队列；
	// 4 ：死信队列。
	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`

	// 备注，128字符以内。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeTopics struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名模糊匹配。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// topic类型描述：
	// 0：普通消息；
	// 1：全局顺序消息；
	// 2：局部顺序消息；
	// 3：重试队列；
	// 4：死信队列；
	// 5：事务消息。
	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DeleteTopics struct {

	// 主题集合，每次最多删除20个。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// pulsar集群Id。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 是否强制删除，默认为false
	Force *bool `json:"Force,omitempty" name:"Force"`
}
