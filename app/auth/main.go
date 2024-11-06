package main

import (
	"tiktok_e-commence/app/auth/biz/dal"
	"tiktok_e-commence/app/auth/biz/server"
)

func main() {
	// 读取配置并初始化数据库连接
	dal.Init()
	// 启动服务端
	go server.RunServer()
}
