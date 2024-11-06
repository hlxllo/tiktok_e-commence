package main

import (
	"github.com/spf13/viper"
	"tiktok_e-commence/app/auth/biz/dal"
	"tiktok_e-commence/app/auth/biz/server"
	"tiktok_e-commence/common"
)

func main() {
	// 读取配置并初始化数据库连接
	dal.Init()
	// 启动服务端
	go server.RunServer()
	// 启动完 grpc 服务后立刻初始化nacos客户端
	Ip := viper.GetString("server.ip")
	Port := viper.GetInt("server.port")
	ServiceName := viper.GetString("server.service_name")
	// 初始化一次 nacos
	common.InitNacosClient()
	// 注册服务实例
	common.RegisterToNacos(Ip, Port, ServiceName)
	// 阻塞
	select {}
}
