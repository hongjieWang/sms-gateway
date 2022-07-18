import request from "@/utils/request";

// 查询集群列表
export function listTopic(data) {
  return request({
    url: "/api/v1/mq/page-describe-topics",
    method: "post",
    data: data,
  });
}
// 创建topic
export function addTopic(data) {
  return request({
    url: "/api/v1/mq/create-topic",
    method: "post",
    data: data,
  });
}
//删除topic
export function delTopic(data) {
  return request({
    url: "/api/v1/mq/delete-topic",
    method: "post",
    data: data,
  });
}
