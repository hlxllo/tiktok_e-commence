package server

import (
	"google.golang.org/grpc"
	"net"
	"tiktok_e-commence/app/user/biz/model"
	"tiktok_e-commence/app/user/biz/service"
)

// 启动服务端
func RunServer() {
	// 开启端口
	s, _ := net.Listen("tcp", ":8000")
	// 创建grpc服务
	server := grpc.NewServer()
	// 注册服务
	model.RegisterUserServiceServer(server, &service.UserServer{})
	// 启动服务
	err := server.Serve(s)
	if err != nil {
		panic(err)
	}
}
