package server

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

// 启动服务端
func RunServer() {
	// 开启端口
	grpcPort := viper.GetString("grpc.port")
	s, _ := net.Listen("tcp", ":"+grpcPort)
	// 创建grpc服务
	server := grpc.NewServer()
	// 注册服务
	//model.RegisterUserServiceServer(server, &service.UserServer{})
	// 启动服务
	err := server.Serve(s)
	if err != nil {
		panic(err)
	}
}
