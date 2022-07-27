package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, smsDashboardRouter)
}

// registerSmsSendLogRouter
func smsDashboardRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SmsDashboard{}
	r := v1.Group("/sms-dashboard").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/count", api.GetAllCount)
		r.GET("/get-success-count", api.GetSuccessCount)
		r.GET("/get-request-count", api.GetRequestCount)
	}
}
