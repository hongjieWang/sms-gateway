import request from "@/utils/request";

// 查询SmsBusinessConfig列表
export function listSmsBusinessConfig(query) {
  return request({
    url: "/api/v1/sms-business-config",
    method: "get",
    params: query,
  });
}

// 查询SmsBusinessConfig详细
export function getSmsBusinessConfig(id) {
  return request({
    url: "/api/v1/sms-business-config/" + id,
    method: "get",
  });
}

// 查询SmsBusinessConfig详细
export function getByBusinessNo(businessNo) {
  return request({
    url: "/api/v1/sms-business-config/getByBusinessNo",
    method: "get",
    params: { businessNo: businessNo },
  });
}

// 新增SmsBusinessConfig
export function addSmsBusinessConfig(data) {
  return request({
    url: "/api/v1/sms-business-config",
    method: "post",
    data: data,
  });
}

// 修改SmsBusinessConfig
export function updateSmsBusinessConfig(data) {
  return request({
    url: "/api/v1/sms-business-config/" + data.id,
    method: "put",
    data: data,
  });
}

// 删除SmsBusinessConfig
export function delSmsBusinessConfig(data) {
  return request({
    url: "/api/v1/sms-business-config",
    method: "delete",
    data: data,
  });
}
