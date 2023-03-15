package cmd

import (
	"code/conf"
	"code/global"
	"code/router"
	"code/utils"
	"fmt"
)

func Start() {
	var initErr error

	// 初始化viper
	conf.InitConfig()

	// 初始化logger
	global.Logger = conf.InitLogger()

	// 初始化数据库
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	// 初始化Redis连接
	rdClient, err := conf.InitRedis()
	global.RedisClient = rdClient
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	// 错误链处理
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}

	// 初始化gin
	router.InitRouter()
}

func Clean() {
	fmt.Println("==========Clean==========")
}
