package router

import (
	"code/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouters() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()

		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", userApi.Login)
		}

		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.POST("", userApi.AddUser)
			rgAuthUser.POST("/list", userApi.GetUserList)
			rgAuthUser.GET("/:id", userApi.GetUserByID)
			rgAuthUser.PUT("/:id", userApi.UpdateUser)
			rgAuthUser.DELETE("/:id", userApi.DeleteUserByID)
		}
	})
}
