package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"tiktok_e-commence/app/gateway/biz/api"
	"tiktok_e-commence/app/gateway/biz/dal"
	"tiktok_e-commence/common"
	"tiktok_e-commence/docs"
)

func main() {
	dal.Init()
	Ip := viper.GetString("server.ip")
	Port := viper.GetInt("server.port")
	ServiceName := viper.GetString("server.service_name")
	common.InitNacosClient()
	common.RegisterToNacos(Ip, Port, ServiceName)
	engine := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	cartGroup := engine.Group("/cart")
	{
		serviceName := "cart-server"
		cartGroup.POST("", api.AddItemHandler(serviceName))
		cartGroup.GET("", api.GetCartHandler(serviceName))
		cartGroup.DELETE("", api.EmptyCartHandler(serviceName))
	}
	ginPort := viper.GetString("gin.port")
	log.Printf("Gin 服务监听的端口 :%v", ginPort)
	if err := engine.Run(":" + ginPort); err != nil {
		log.Fatalf("启动 Gin 服务失败: %v", err)
	}
	select {}
}
