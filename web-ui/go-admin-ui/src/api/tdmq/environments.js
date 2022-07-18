import request from "@/utils/request";

// 查询集群列表
export function listEnvironments(clustersId) {
  return request({
    url: "/api/v1/mq/describe-environments/" + clustersId,
    method: "get",
  });
}
//获取命名空间下拉
export function selectEnvironments(clustersId) {
  return request({
    url: "/api/v1/mq/describe-select-environments/" + clustersId,
    method: "get",
  });
}
