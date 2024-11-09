package common

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
)

// 根据服务名称获取健康实例并创建gRPC连接
func CallService(c *gin.Context, serviceName string) (*grpc.ClientConn, error) {
	instances, err := SelectHealthyInstance(serviceName)
	if err != nil {
		HandleResponse(c, http.StatusInternalServerError, ErrDiscoverServiceFailed, nil)
		return nil, err
	}
	conn, err := CreateGRPCConn(instances.Ip, int(instances.Port))
	if err != nil {
		HandleResponse(c, http.StatusInternalServerError, ErrGRPCConnFailed, nil)
		return nil, err
	}
	return conn, nil
}
