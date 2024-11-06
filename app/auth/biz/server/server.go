package server

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"tiktok_e-commence/app/auth/biz/model"
	"tiktok_e-commence/app/auth/biz/service"
)

// 启动服务端
func RunServer() {
	// 开启端口
	serverPort := viper.GetString("server.port")
	s, _ := net.Listen("tcp", ":"+serverPort)
	// 创建grpc服务
	server := grpc.NewServer()
	// 注册服务
	model.RegisterAuthServiceServer(server, &service.AuthServer{})
	// 启动服务
	err := server.Serve(s)
	if err != nil {
		panic(err)
	}
}
