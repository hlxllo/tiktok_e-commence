package common

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type ClientConfig struct {
	Ip          string
	Port        int
	ServiceName string
}

func InitNacosClient(nc *ClientConfig) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("localhost", 8848,
			constant.WithContextPath("/nacos")),
	}
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
	)
	// 创建 Nacos 客户端
	client, _ := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	// 注册服务
	registerServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          nc.Ip,
		Port:        uint64(nc.Port),
		ServiceName: nc.ServiceName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})
}

// 注册服务
func registerServiceInstance(nacosClient naming_client.INamingClient, param vo.RegisterInstanceParam) {
	success, err := nacosClient.RegisterInstance(param)
	if !success || err != nil {
		panic("register Service Instance failed!")
	}
}
