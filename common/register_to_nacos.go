package common

import (
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type ClientConfig struct {
	Ip          string
	Port        int
	ServiceName string
}

func RegisterToNacos(nc *ClientConfig) {
	// 注册服务
	registerServiceInstance(NacosClient, vo.RegisterInstanceParam{
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
		panic("注册实例失败！")
	}
}
