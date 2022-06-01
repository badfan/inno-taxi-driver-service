package v1

import "github.com/gin-gonic/gin"

func (a *ApiV1) DriverGroup(group *gin.RouterGroup) {
	driverGroup := group.Group("/driver")
	{
		driverGroup.GET("/rating", a.handler.GetDriverRating)
		driverGroup.GET("/status", a.handler.GetDriverStatus)
	}
}
