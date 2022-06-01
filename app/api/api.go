package api

import (
	v1 "github.com/badfan/inno-taxi-driver-service/app/api/v1"
	"github.com/badfan/inno-taxi-driver-service/app/handlers"
	"github.com/gin-gonic/gin"
)

type ApiGroup struct {
	handler *handlers.Handler
	ApiV1   *v1.ApiV1
}

func NewApiGroup(handler *handlers.Handler, apiV1 *v1.ApiV1) *ApiGroup {
	return &ApiGroup{handler: handler, ApiV1: apiV1}
}

func (a *ApiGroup) InitRouterGroups(router *gin.Engine) {

	router.POST("/sign-in", a.handler.SignIn)
	router.POST("/sign-up", a.handler.SignUp)

	api := router.Group("/api", a.handler.Middleware)
	{
		a.ApiV1.InitApiV1Groups(api)
	}
}
