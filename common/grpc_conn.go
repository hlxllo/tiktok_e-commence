package common

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

// 创建gRPC连接
func CreateGRPCConn(Ip string, Port int) (*grpc.ClientConn, error) {
	addr := Ip + ":" + strconv.Itoa(Port)
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(MaxSendMsgSize),
			grpc.MaxCallRecvMsgSize(MaxRecvMsgSize),
		))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
