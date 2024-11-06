package common

import (
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// 选择健康实例
func SelectHealthyInstance(serviceName string) (*model.Instance, error) {
	instances, err := NacosClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
		GroupName:   "DEFAULT_GROUP",
	})
	if err != nil {
		return nil, err
	}
	return instances, nil
}
