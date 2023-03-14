package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf/")

	viper.SetDefault("server.port", "8090")

	err := viper.ReadInConfig()

	if err != nil {
		panic("Load Config Error: " + err.Error())
	}

	fmt.Println(viper.GetString("server.port"))
}
