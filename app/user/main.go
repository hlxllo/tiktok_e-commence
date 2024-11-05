package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"tiktok_e-commence/app/user/biz/api"
	"tiktok_e-commence/app/user/biz/dal"
	"tiktok_e-commence/app/user/biz/model"
	"tiktok_e-commence/app/user/biz/server"
	"tiktok_e-commence/common"
	"tiktok_e-commence/docs"
)

func main() {
	// 读取配置并初始化数据库连接
	dal.Init()
	// 启动服务端
	go server.RunServer()
	// 启动完 grpc 服务后立刻初始化nacos客户端
	nc := &common.ClientConfig{
		Ip:          viper.GetString("server.ip"),
		Port:        viper.GetInt("server.port"),
		ServiceName: viper.GetString("server.service_name"),
	}
	common.InitNacosClient(nc)
	fmt.Println(nc)
	// 创建客户端连接
	// 连接服务端，禁用安全传输
	grpcClient := viper.GetString("grpc.ip") + ":" + viper.GetString("grpc.port")
	fmt.Println(grpcClient)
	conn, err := grpc.NewClient(grpcClient,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 建立连接
	client := model.NewUserServiceClient(conn)
	// 创建路由并分组注册
	engine := gin.Default()
	// Swagger 配置
	docs.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register", api.RegisterUserHandler(client))
	}
	// 启动 Gin 服务
	serverPort := viper.GetString("server.port")
	fmt.Println(serverPort)
	log.Printf("Gin server listening on :%v", serverPort)
	if err := engine.Run(":" + serverPort); err != nil {
		log.Fatalf("Failed to run Gin server: %v", err)
	}
	// 阻塞
	select {}
}
