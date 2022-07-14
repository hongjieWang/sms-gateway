package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerMqrRouter)
}

// 需认证的路由代码
func registerMqrRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Mq{}
	r := v1.Group("/mq").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		r.GET("/describe-clusters", api.DescribeClusters)
		r.GET("/describe-page-clusters", api.PageDescribeClusters)
		r.GET("/describe-select-clusters", api.SelectClusters)
		r.GET("/describe-select-environments/:clusterId", api.SelectEnvironments)
		r.GET("/describe-environments/:clusterId", api.DescribeEnvironments)
		r.POST("/create-environments", api.CreateEnvironments)
		r.POST("/delete-environments", api.DeleteEnvironments)
		r.POST("/create-topic", api.CreateTopic)
		r.POST("/describe-topics", api.DescribeTopics)
		r.POST("/page-describe-topics", api.PageDescribeTopics)
		r.POST("/delete-topic", api.DeleteTopic)
	}
}
