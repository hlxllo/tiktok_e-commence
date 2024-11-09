package main

import (
	"github.com/spf13/viper"
	"tiktok_e-commence/app/user/biz/dal"
	"tiktok_e-commence/app/user/biz/server"
	"tiktok_e-commence/common"
)

func main() {
	dal.Init()
	go server.RunServer()
	Ip := viper.GetString("server.ip")
	Port := viper.GetInt("server.port")
	ServiceName := viper.GetString("server.service_name")
	common.InitNacosClient()
	common.RegisterToNacos(Ip, Port, ServiceName)
	select {}
}
