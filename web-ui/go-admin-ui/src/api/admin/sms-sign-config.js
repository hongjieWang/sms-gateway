import request from '@/utils/request'

// 查询SmsSignConfig列表
export function listSmsSignConfig(query) {
    return request({
        url: '/api/v1/sms-sign-config',
        method: 'get',
        params: query
    })
}

// 查询SmsSignConfig详细
export function getSmsSignConfig (id) {
    return request({
        url: '/api/v1/sms-sign-config/' + id,
        method: 'get'
    })
}


// 新增SmsSignConfig
export function addSmsSignConfig(data) {
    return request({
        url: '/api/v1/sms-sign-config',
        method: 'post',
        data: data
    })
}

// 修改SmsSignConfig
export function updateSmsSignConfig(data) {
    return request({
        url: '/api/v1/sms-sign-config/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除SmsSignConfig
export function delSmsSignConfig(data) {
    return request({
        url: '/api/v1/sms-sign-config',
        method: 'delete',
        data: data
    })
}

