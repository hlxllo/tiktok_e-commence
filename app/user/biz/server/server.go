package server

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"tiktok_e-commence/app/user/biz/service"
	"tiktok_e-commence/common"
	"tiktok_e-commence/common/model/model"
)

// 启动服务端
func RunServer() {
	// 开启端口
	serverPort := viper.GetString("server.port")
	s, _ := net.Listen("tcp", ":"+serverPort)
	// 创建grpc服务
	server := grpc.NewServer(
		grpc.MaxSendMsgSize(common.MaxSendMsgSize),
		grpc.MaxRecvMsgSize(common.MaxRecvMsgSize),
	)
	// 注册服务
	model.RegisterUserServiceServer(server, &service.UserServer{})
	// 启动服务
	err := server.Serve(s)
	if err != nil {
		panic(err)
	}
}
