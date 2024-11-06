package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
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
	Ip := viper.GetString("server.ip")
	Port := viper.GetInt("server.port")
	ServiceName := viper.GetString("server.service_name")
	nc := &common.ClientConfig{
		Ip:          Ip,
		Port:        Port,
		ServiceName: ServiceName,
	}
	// 初始化nacos客户端
	common.InitNacosClient(nc)
	// 创建 grpc 客户端连接
	// 连接服务端，禁用安全传输
	grpcAddr := Ip + ":" + strconv.Itoa(Port)
	conn, err := grpc.NewClient(grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 建立连接
	client := model.NewUserServiceClient(conn)
	// 创建路由
	engine := gin.Default()
	// Swagger 配置
	docs.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 路由分组
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register", api.RegisterUserHandler(client))
		userGroup.POST("/login", api.LoginUserHandler(client))
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
