import request from "@/utils/request";

// 查询集群列表
export function listClusters(query) {
  return request({
    url: "/api/v1/mq/describe-page-clusters",
    method: "get",
    params: query,
  });
}
//查询集群列表-Select
export function describeSelectClusters() {
  return request({
    url: "/api/v1/mq/describe-select-clusters",
    method: "get",
  });
}
