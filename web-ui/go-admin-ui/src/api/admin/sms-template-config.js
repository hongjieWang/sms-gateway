import request from "@/utils/request";

// 查询SmsTemplateConfig列表
export function listSmsTemplateConfig(query) {
  return request({
    url: "/api/v1/sms-template-config",
    method: "get",
    params: query,
  });
}

// 查询SmsTemplateConfig详细
export function getSmsTemplateConfig(id) {
  return request({
    url: "/api/v1/sms-template-config/" + id,
    method: "get",
  });
}

//根据templateNo查询模版信息
export function getByTemplateNo(templateNo) {
  return request({
    url: "/api/v1/sms-template-config/getByTemplateNo",
    method: "get",
    params: { templateNo: templateNo },
  });
}

// 新增SmsTemplateConfig
export function addSmsTemplateConfig(data) {
  return request({
    url: "/api/v1/sms-template-config",
    method: "post",
    data: data,
  });
}

// 修改SmsTemplateConfig
export function updateSmsTemplateConfig(data) {
  return request({
    url: "/api/v1/sms-template-config/" + data.id,
    method: "put",
    data: data,
  });
}

// 删除SmsTemplateConfig
export function delSmsTemplateConfig(data) {
  return request({
    url: "/api/v1/sms-template-config",
    method: "delete",
    data: data,
  });
}
