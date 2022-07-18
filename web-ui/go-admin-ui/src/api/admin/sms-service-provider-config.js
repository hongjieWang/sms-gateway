import request from '@/utils/request'

// 查询SmsServiceProviderConfig列表
export function listSmsServiceProviderConfig(query) {
    return request({
        url: '/api/v1/sms-service-provider-config',
        method: 'get',
        params: query
    })
}

// 查询SmsServiceProviderConfig详细
export function getSmsServiceProviderConfig (id) {
    return request({
        url: '/api/v1/sms-service-provider-config/' + id,
        method: 'get'
    })
}


// 新增SmsServiceProviderConfig
export function addSmsServiceProviderConfig(data) {
    return request({
        url: '/api/v1/sms-service-provider-config',
        method: 'post',
        data: data
    })
}

// 修改SmsServiceProviderConfig
export function updateSmsServiceProviderConfig(data) {
    return request({
        url: '/api/v1/sms-service-provider-config/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除SmsServiceProviderConfig
export function delSmsServiceProviderConfig(data) {
    return request({
        url: '/api/v1/sms-service-provider-config',
        method: 'delete',
        data: data
    })
}

