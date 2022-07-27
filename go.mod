module go-admin

go 1.15

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/alibaba/sentinel-golang v0.6.1
	github.com/alibabacloud-go/darabonba-openapi v0.1.18
	github.com/alibabacloud-go/dysmsapi-20170525/v2 v2.0.16
	github.com/alibabacloud-go/tea v1.1.17
	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20190307165228-86c17b95fcd5
	github.com/bitly/go-simplejson v0.5.0
	github.com/bytedance/go-tagexpr/v2 v2.7.12
	github.com/casbin/casbin/v2 v2.37.4
	github.com/gin-gonic/gin v1.8.1
	github.com/go-admin-team/go-admin-core v1.3.8
	github.com/go-admin-team/go-admin-core/sdk v1.3.9
	github.com/google/uuid v1.2.0
	github.com/huaweicloud/huaweicloud-sdk-go-obs v3.21.12+incompatible
	github.com/mssola/user_agent v0.5.2
	github.com/nyaruka/phonenumbers v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/client_golang v1.12.2
	github.com/qiniu/go-sdk/v7 v7.11.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/shirou/gopsutil/v3 v3.22.1
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.0.0
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.7
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.440
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms v1.0.440
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdmq v1.0.330
	github.com/unrolled/secure v1.0.8
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	gorm.io/driver/mysql v1.3.5
	gorm.io/driver/postgres v1.3.8
	gorm.io/driver/sqlite v1.3.1
	gorm.io/driver/sqlserver v1.3.2
	gorm.io/gorm v1.23.8
)

replace github.com/go-admin-team/go-admin-core/sdk => github.com/hongjieWang/go-admin-core/sdk v1.3.12-0.20220727023055-ebc134c5b01d

replace github.com/go-admin-team/go-admin-core => github.com/hongjieWang/go-admin-core v1.3.9-0.20220727023055-ebc134c5b01d

//replace github.com/go-admin-team/go-admin-core/sdk => /Users/wanghongjie/GolandProjects/go-admin-core/sdk
//replace github.com/go-admin-team/go-admin-core => /Users/wanghongjie/GolandProjects/go-admin-core
