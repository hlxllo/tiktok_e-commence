package common

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

// 创建gRPC连接
func CreateGRPCConn(Ip string, Port uint64) (*grpc.ClientConn, error) {
	addr := Ip + ":" + strconv.Itoa(int(Port))
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
