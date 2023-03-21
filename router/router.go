package router

import (
	_ "code/docs"
	"code/global"
	"code/middleware"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type IFnRegistRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRouters []IFnRegistRoute
)

func RegistRoute(fn IFnRegistRoute) {
	if fn == nil {
		return
	}

	gfnRouters = append(gfnRouters, fn)
}

func InitRouter() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()
	r.Use(middleware.Cors())
	//docs.SwaggerInfo.BasePath = "/api/v1"
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")
	rgAuth.Use(middleware.Auth())

	initBasePlatformRoutes()

	// 自定义验证器
	registCustValidator()

	for _, fnRegistRoute := range gfnRouters {
		fnRegistRoute(rgPublic, rgAuth)
	}

	// 集成swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	stPort := viper.GetString("server.port")

	server := &http.Server{
		Addr:    ":" + stPort,
		Handler: r,
	}

	go func() {
		global.Logger.Info("start Listen: " + stPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error("Start Server Error: " + err.Error())
			return
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error("Stop Server Error: " + err.Error())
	}

	global.Logger.Info("Stop Server Success")
}

func initBasePlatformRoutes() {
	InitUserRouters()
	InitHostRouters()
}

func registCustValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				if value != "" && strings.Index(value, "a") == 0 {
					return true
				}
			}
			return false
		})
	}
}
