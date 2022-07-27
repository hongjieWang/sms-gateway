import request from '@/utils/request'

// 获取验证码
export function getAllCount() {
  return request({
    url: '/api/v1/sms-dashboard/count',
    method: 'get'
  })
}

export function getSuccessCount() {
  return request({
    url: '/api/v1/sms-dashboard/get-success-count',
    method: 'get'
  })
}

export function getRequestCount() {
  return request({
    url: '/api/v1/sms-dashboard/get-request-count',
    method: 'get'
  })
}
