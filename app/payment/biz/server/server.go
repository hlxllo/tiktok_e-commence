package server

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"tiktok_e-commence/app/payment/biz/service"
	"tiktok_e-commence/common"
	"tiktok_e-commence/common/model/model"
)

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
	model.RegisterPaymentServiceServer(server, &service.PaymentServer{})
	// 启动服务
	err := server.Serve(s)
	if err != nil {
		panic(err)
	}
}
