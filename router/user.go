package router

import (
	"code/api"
	"github.com/gin-gonic/gin"
	"net/http"
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
			rgAuthUser.GET("", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"data": []map[string]any{
						{"id": 1, "name": "zs"},
						{"id": 2, "name": "lisi"},
					},
				})
			})
		}

		rgAuthUser.GET("/:id", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"id":   1,
				"name": "zs",
			})
		})
	})
}
