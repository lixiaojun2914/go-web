package cmd

import (
	"code/conf"
	"code/global"
	"code/router"
	"fmt"
)

func Start() {
	conf.InitConfig()
	global.Logger = conf.InitLogger()
	router.InitRouter()
}

func Clean() {
	fmt.Println("==========Clean==========")
}
