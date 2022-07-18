import request from '@/utils/request'

// 查询Segments列表
export function listSegments(query) {
    return request({
        url: '/api/v1/segments',
        method: 'get',
        params: query
    })
}

// 查询Segments详细
export function getSegments (bizTag) {
    return request({
        url: '/api/v1/segments/' + bizTag,
        method: 'get'
    })
}


// 新增Segments
export function addSegments(data) {
    return request({
        url: '/api/v1/segments',
        method: 'post',
        data: data
    })
}

// 修改Segments
export function updateSegments(data) {
    return request({
        url: '/api/v1/segments/'+data.bizTag,
        method: 'put',
        data: data
    })
}

// 删除Segments
export function delSegments(data) {
    return request({
        url: '/api/v1/segments',
        method: 'delete',
        data: data
    })
}

