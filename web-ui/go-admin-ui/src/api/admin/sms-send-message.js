import request from "@/utils/request";
// 新增smsSendMessage
export function smsSendMessage(data) {
  return request({
    url: "/api/v1/send-message",
    method: "post",
    data: data,
  });
}
