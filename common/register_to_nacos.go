package common

import (
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func RegisterToNacos(Ip string, Port int, ServiceName string) {
	// 注册服务
	registerServiceInstance(NacosClient, vo.RegisterInstanceParam{
		Ip:          Ip,
		Port:        uint64(Port),
		ServiceName: ServiceName,
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
