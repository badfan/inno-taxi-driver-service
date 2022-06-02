package v1

import (
	"github.com/badfan/inno-taxi-driver-service/app/handlers"
	"github.com/gin-gonic/gin"
)

type ApiV1 struct {
	handler *handlers.Handler
}

func NewApiV1(handler *handlers.Handler) *ApiV1 {
	return &ApiV1{handler: handler}
}

func (a *ApiV1) InitApiV1Groups(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		a.DriverGroup(v1)
	}
}
