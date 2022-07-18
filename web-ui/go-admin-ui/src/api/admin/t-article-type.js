import request from '@/utils/request'

// 查询TArticleType列表
export function listTArticleType(query) {
    return request({
        url: '/api/v1/t-article-type',
        method: 'get',
        params: query
    })
}

// 查询TArticleType详细
export function getTArticleType (id) {
    return request({
        url: '/api/v1/t-article-type/' + id,
        method: 'get'
    })
}


// 新增TArticleType
export function addTArticleType(data) {
    return request({
        url: '/api/v1/t-article-type',
        method: 'post',
        data: data
    })
}

// 修改TArticleType
export function updateTArticleType(data) {
    return request({
        url: '/api/v1/t-article-type/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除TArticleType
export function delTArticleType(data) {
    return request({
        url: '/api/v1/t-article-type',
        method: 'delete',
        data: data
    })
}

