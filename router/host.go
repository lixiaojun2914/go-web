package router

import (
	"code/api"
	"github.com/gin-gonic/gin"
)

func InitHostRouters() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		hostApi := api.NewHostApi()

		rgAuthHost := rgAuth.Group("host")
		{
			rgAuthHost.POST("/shutdown", hostApi.Shutdown)
		}
	})
}
