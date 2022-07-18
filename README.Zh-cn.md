# 短信网关总体设计

在消息系统中，短信网关和短信渠道的对接是最核心的功能。其中短信网关是对外提供服务的接口，所有需要发送短信的操作都需要通过短信网关分发到对应的渠道上。一旦定型，后续就很少，也很难调整。而短信渠道是接收网关的请求，调用渠道接口执行真正的发送短信操作。每个渠道的接口，传输方式都不尽相同，所以在这里，短信网关相对短信渠道模块的作用，类似设计模式模式中的wrapper，封装各个渠道的差异，对网关呈现统一的接口。而网关的功能就是为业务提供通用接口，一些和渠道交互的公共操作，也会放置到网关中。

## 一、功能概述

消息系统对其他系统提供的服务包括渠道商管理、签名管理、短信模版管理、应用管理、业务类型管理、短信记录管理等。

- 渠道商管理：短息网关维护多个渠道商，配置渠道商的身份ID、身份key等其他信息，支持的渠道商：腾讯云、阿里云等
- 签名管理：配置渠道商下不同的短信签名，为后续发送短信提供数据基础。
- 短信模版管理：配置不同短信模版信息，可在也业务配置中选择不同的模版信息，注：不同渠道模版略有不同，详情参见说明文档。
- 应用管理：配置短信网关对接的应用信息，可在应用管理中配置应用可用短信条数、每分钟限制短信条数等信息，通过应用可用条数控制不同应用使用短信数量。
- 业务管理：配置发送短信业务场景，如注册业务、登录业务、下单通知、支付成功通知等不同的业务信息。添加业务信息后，生成业务编号，业务系统根据业务编号发送短信模版。

## 二、整体架构

![smsGolang](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/smsGolang.png)

## 三、管理平台

![杰子学编程-渠道商配置](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/sms_service_provider_config.png)

![杰子学编程-添加渠道商](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/sms_service_provider_config_add.png)

![杰子学编程-短信签名模版](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/sms_sign_config.png)

![杰子学编程-添加短信签名](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/sms_sign_config_add.png)

![杰子学编程-模版列表](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/sms_template_config.png)

![杰子学编程-添加短信模版](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/sms_template_config_add.png)

![杰子学编程-应用管理](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/sms_app_config.png)

![杰子学编程-添加应用](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/sms_app_config_add.png)

![杰子学编程-业务模版配置](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/sms_business_config.png)

![杰子学编程-添加业务模版](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/sms_business_config_add.png)

![杰子学编程-发送记录](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/sms_send_log.png)

![杰子学编程-发送短信测试页面](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/sms_send_test_page.png)

## 四、服务API

请求地址：/api/v1/send-message

请求方式：POST

请求参数：

| 字段名称   | 字段说明     | 字段类型 | 是否必填 | 备注              |
| ---------- | ------------ | -------- | -------- | ----------------- |
| businessNo | 业务编号     | String   | 是       |                   |
| phones     | 手机号集合   | List     | 是       | 多个手机号","分割 |
| params     | 模版填充参数 | List     | 否       | 多个参数","分割   |

请求参数示例：

```json
{
    "businessNo": "B-000003",
    "phones": [
        "+86182XXXXXX68"
    ],
    "params": ["123456"]
}
```

响应参数：

| 字段名称  | 字段说明     | 字段类型 | 是否必填 | 备注 |
| --------- | ------------ | -------- | -------- | ---- |
| requestId | requestId    | String   | 是       |      |
| code      | 状态码       | Int      | 是       |      |
| msg       | 渠道返回消息 | String   | 是       |      |
| data      | 响应数据     | Int      | 否       |      |

响应参数示例：

```json
{
  "requestId": "b9c0fd5e-3a43-4039-a224-2a5ae99385e1",
  "code": 200,
  "msg": "{\"sms_send_status\":[{\"fee\":1,\"message\":\"OK\",\"code\":\"OK\",\"phone\":\"+8618232533068\"}],\"request_id\":\"771D1C9C-74AD-582E-B5EB-5FB12C038497\"}",
  "data": 200
}
```



## 无、数据库设计

### 5.1 短信渠道配置表

配置不同渠道商信息，如腾讯云、阿里云、七牛云等

sms_service_provider_config (服务商配置表)

| 字段名称          | 字段说明     | 字段类型 | 是否必填 | 备注                                             |
| ----------------- | ------------ | -------- | -------- | ------------------------------------------------ |
| channel_no        | 渠道商编程   | String   | 是       | Tencent、AliYun                                  |
| provider_name     | 服务商名称   | String   | 是       | 阿里云、腾讯云                                   |
| provider_no       | 服务商编号   | String   | 是       | Aly、txy                                         |
| access_key_id     | 身份标识     | String   | 是       | 腾讯：SecretId                                   |
| access_key_secret | 身份认证密钥 | String   | 是       | 腾讯：SecretKey                                  |
| endpoint          | 调用域名     | String   | 是       |                                                  |
| sdk_app_id        | 应用ID       | String   | 否       | 腾讯云必填                                       |
| region            | 地域列表     | String   | 否       | 腾讯云必填：ap-beijing、ap-guangzhou、ap-nanjing |
| remark            | 备注         | String   | 否       |                                                  |
| ext_json          | 扩展字段     | String   | 否       |                                                  |
| status            | 状态         | Bool     | 是       | 1启用 0 禁用                                     |

### 5.2 短信签名配置表

sms_sign_config(短信签名配置表)

| 字段编号      | 字段说明   | 字段类型 | 是否必填 | 备注         |
| ------------- | ---------- | -------- | -------- | ------------ |
| sign_name     | 签名名称   | String   | 是       |              |
| provider_no   | 服务商编号 | String   | 是       | Aly、txy     |
| provider_name | 服务商名称 | String   | 是       |              |
| remark        | 备注       | String   | 否       |              |
| ext_json      | 扩展字段   | String   | 否       |              |
| status        | 状态       | Bool     | 是       | 1启用 0 禁用 |

### 5.3 短信模版配置表

sms_template_config(短信模版配置表)

| 字段编号                | 字段说明       | 字段类型 | 是否必填 | 备注         |
| ----------------------- | -------------- | -------- | -------- | ------------ |
| sign_name               | 签名名称       | String   | 是       |              |
| provider_no             | 服务商编号     | String   | 是       | Aly、txy     |
| template_no             | 模版编号       | String   | 是       |              |
| template_content        | 模版内容       | String   | 是       |              |
| third_party_template_no | 第三方模版编号 | String   | 是       |              |
| remark                  | 备注           | String   | 否       |              |
| ext_json                | 扩展字段       | String   | 否       |              |
| status                  | 状态           | Bool     | 是       | 1启用 0 禁用 |

### 5.4 应用配置表

sms_app_config (应用管理配置表)

| 字段编号         | 字段说明 | 字段类型 | 是否必填 | 备注                   |
| ---------------- | -------- | -------- | -------- | ---------------------- |
| app_no           | 应用ID   | String   | 是       |                        |
| app_name         | 应用名称 | String   | 是       | Aly、txy               |
| available_number | 可用数量 | Int      | 是       | 可用短信包             |
| current_limiting | 限流数量 | Int      | 是       | 每分钟允许发送短信数量 |
| use_number       | 已用数量 | Int      | 是       |                        |
| remark           | 备注     | String   | 否       |                        |
| ext_json         | 扩展字段 | String   | 否       |                        |
| status           | 状态     | Bool     | 是       | 1启用 0 禁用           |

### 5.5 业务配置表

sms_business_config(业务配置表)

| 字段编号      | 字段说明 | 字段类型 | 是否必填 | 备注         |
| ------------- | -------- | -------- | -------- | ------------ |
| app_no        | 应用ID   | String   | 是       |              |
| business_name | 业务名称 | String   | 是       | Aly、txy     |
| business_no   | 业务编号 | String   | 是       |              |
| template_no   | 模版编号 | String   | 是       |              |
| business_desc | 业务说明 | String   | 否       |              |
| remark        | 备注     | String   | 否       |              |
| ext_json      | 扩展字段 | String   | 否       |              |
| status        | 状态     | Bool     | 是       | 1启用 0 禁用 |

### 5.6 发送日志表

sms_send_log(发送记录表)

| 字段编号     | 字段说明       | 字段类型 | 是否必填 | 备注 |
| ------------ | -------------- | -------- | -------- | ---- |
| app_no       | 应用ID         | String   | 是       |      |
| business_no  | 业务编号       | String   | 是       |      |
| status       | 状态           | Int      | 是       |      |
| fee          | 计价条数       | Int      | 是       |      |
| phone_number | 发送手机号     | String   | 是       |      |
| message      | 接口响应消息   | String   | 是       |      |
| code         | 接口响应状态码 | String   | 是       |      |
| content      | 发送内容       | String   | 是       |      |
| remark       | 备注           | String   | 否       |      |
| ext_json     | 扩展字段       | String   | 否       |      |

## 六、技术栈

[go-admin](https://doc.go-admin.dev/)

[杰子学编程 (julywhj.cn)](https://julywhj.cn/)

Mysql、Redis、

## 七、源码

关注公众号：杰子学编程，回复： "短信网关" 获取。

![qrcode_for_gh_5d871c6cb930_258](https://img-1258527903.cos.ap-beijing.myqcloud.com/img/qrcode_for_gh_5d871c6cb930_258.jpg)