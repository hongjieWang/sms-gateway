import request from '@/utils/request'

// 查询TArticle列表
export function listTArticle(query) {
    return request({
        url: '/api/v1/t-article',
        method: 'get',
        params: query
    })
}

// 查询TArticle详细
export function getTArticle (id) {
    return request({
        url: '/api/v1/t-article/' + id,
        method: 'get'
    })
}


// 新增TArticle
export function addTArticle(data) {
    return request({
        url: '/api/v1/t-article',
        method: 'post',
        data: data
    })
}

// 修改TArticle
export function updateTArticle(data) {
    return request({
        url: '/api/v1/t-article/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除TArticle
export function delTArticle(data) {
    return request({
        url: '/api/v1/t-article',
        method: 'delete',
        data: data
    })
}

