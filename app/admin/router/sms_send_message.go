package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSmsSendMessageRouter)
}

func registerSmsSendMessageRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SmsSendMessage{}
	r := v1.Group("/send-message")
	{
		r.POST("", api.SendMessage)
	}
}
