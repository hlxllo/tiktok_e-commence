package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"tiktok_e-commence/app/user/biz/model"
)

// 启动客户端
func RunClient() model.UserServiceClient {
	// 连接服务端，禁用安全传输
	conn, err := grpc.NewClient("127.0.0.1:8000",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 建立连接
	client := model.NewUserServiceClient(conn)
	return client
}
