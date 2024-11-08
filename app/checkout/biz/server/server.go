package server

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"tiktok_e-commence/app/checkout/biz/model"
	"tiktok_e-commence/app/checkout/biz/service"
)

func RunServer() {
	// 开启端口
	serverPort := viper.GetString("server.port")
	s, _ := net.Listen("tcp", ":"+serverPort)
	// 创建grpc服务
	server := grpc.NewServer()
	// 注册服务
	model.RegisterCheckoutServiceServer(server, &service.CheckoutServer{})
	// 启动服务
	err := server.Serve(s)
	if err != nil {
		panic(err)
	}
}
