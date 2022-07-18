import request from '@/utils/request'

// 查询SmsSendLog列表
export function listSmsSendLog(query) {
    return request({
        url: '/api/v1/sms-send-log',
        method: 'get',
        params: query
    })
}

// 查询SmsSendLog详细
export function getSmsSendLog (id) {
    return request({
        url: '/api/v1/sms-send-log/' + id,
        method: 'get'
    })
}


// 新增SmsSendLog
export function addSmsSendLog(data) {
    return request({
        url: '/api/v1/sms-send-log',
        method: 'post',
        data: data
    })
}

// 修改SmsSendLog
export function updateSmsSendLog(data) {
    return request({
        url: '/api/v1/sms-send-log/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除SmsSendLog
export function delSmsSendLog(data) {
    return request({
        url: '/api/v1/sms-send-log',
        method: 'delete',
        data: data
    })
}

