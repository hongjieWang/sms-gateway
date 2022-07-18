import request from "@/utils/request";

// 查询SmsAppConfig列表
export function listSmsAppConfig(query) {
  return request({
    url: "/api/v1/sms-app-config",
    method: "get",
    params: query,
  });
}

// 查询SmsAppConfig详细
export function getSmsAppConfig(id) {
  return request({
    url: "/api/v1/sms-app-config/" + id,
    method: "get",
  });
}

// 新增SmsAppConfig
export function addSmsAppConfig(data) {
  return request({
    url: "/api/v1/sms-app-config",
    method: "post",
    data: {
      appName: data.appName,
      availableNumber: parseInt(data.availableNumber),
      currentLimiting: parseInt(data.currentLimiting),
      extJson: "{}",
      status: data.status,
      appNo: data.appNo,
      remark: data.remark,
    },
  });
}

// 修改SmsAppConfig
export function updateSmsAppConfig(data) {
  return request({
    url: "/api/v1/sms-app-config/" + data.id,
    method: "put",
    data: {
      appName: data.appName,
      availableNumber: parseInt(data.availableNumber),
      currentLimiting: parseInt(data.currentLimiting),
      extJson: data.extJson,
      status: data.status,
      appNo: data.appNo,
      remark: data.remark,
    },
  });
}

// 删除SmsAppConfig
export function delSmsAppConfig(data) {
  return request({
    url: "/api/v1/sms-app-config",
    method: "delete",
    data: data,
  });
}
