package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"tiktok_e-commence/app/user/biz/api"
	"tiktok_e-commence/app/user/biz/dal"
	"tiktok_e-commence/app/user/biz/model"
	"tiktok_e-commence/app/user/biz/server"
)

func main() {
	// 读取配置并初始化数据库连接
	dal.Init()
	// 启动服务端
	go server.RunServer()
	// 创建客户端连接
	// 连接服务端，禁用安全传输
	conn, err := grpc.NewClient("127.0.0.1:8000",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 建立连接
	client := model.NewUserServiceClient(conn)
	// 创建路由并分组注册
	engine := gin.Default()
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register", api.RegisterUserHandler(client))
	}
	// 启动 Gin 服务
	log.Println("Gin server listening on :8080")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Failed to run Gin server: %v", err)
	}

}
