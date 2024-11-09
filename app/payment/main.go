package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"tiktok_e-commence/app/payment/biz/api"
	"tiktok_e-commence/app/payment/biz/dal"
	"tiktok_e-commence/app/payment/biz/server"
	"tiktok_e-commence/common"
	"tiktok_e-commence/common/model/model"
	"tiktok_e-commence/docs"
)

func main() {
	// 读取配置并初始化数据库连接
	dal.Init()
	// 启动服务端
	go server.RunServer()
	// 启动完 grpc 服务后立刻初始化nacos客户端
	Ip := viper.GetString("server.ip")
	Port := viper.GetInt("server.port")
	ServiceName := viper.GetString("server.service_name")
	// 初始化一次 nacos
	common.InitNacosClient()
	// 注册服务实例
	common.RegisterToNacos(Ip, Port, ServiceName)
	// 创建 grpc 客户端连接
	// 连接服务端，禁用安全传输
	conn, _ := common.CreateGRPCConn(Ip, Port)
	defer conn.Close()
	// 建立连接
	client := model.NewPaymentServiceClient(conn)
	// 创建路由
	engine := gin.Default()
	// Swagger 配置
	docs.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 路由分组
	paymentGroup := engine.Group("/payment")
	{
		paymentGroup.POST("", api.ChargeHandler(client))
	}
	// 启动 Gin 服务
	ginPort := viper.GetString("gin.port")
	log.Printf("Gin server listening on :%v", ginPort)
	if err := engine.Run(":" + ginPort); err != nil {
		log.Fatalf("Failed to run Gin server: %v", err)
	}
	// 阻塞
	select {}
}
