package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"tiktok_e-commence/app/gateway/biz/api"
	"tiktok_e-commence/app/gateway/biz/dal"
	"tiktok_e-commence/app/gateway/biz/middleware"
	"tiktok_e-commence/common"
	"tiktok_e-commence/docs"
)

func main() {
	// 读取配置并初始化数据库连接
	dal.Init()
	Ip := viper.GetString("server.ip")
	Port := viper.GetInt("server.port")
	ServiceName := viper.GetString("server.service_name")
	// 初始化nacos客户端
	common.InitNacosClient()
	// 注册服务实例
	common.RegisterToNacos(Ip, Port, ServiceName)
	// 创建路由
	engine := gin.Default()
	// swagger
	docs.SwaggerInfo.BasePath = ""
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	var serviceName string
	// 用户
	userGroup := engine.Group("/user")
	{
		serviceName = "user-server"
		userGroup.POST("register", api.RegisterUserHandler(serviceName))
		userGroup.POST("login", api.LoginUserHandler(serviceName))
	}
	// 商品
	productGroup := engine.Group("/product")
	{
		serviceName = "product-server"
		productGroup.POST("/list", api.ListProductsHandler(serviceName))
		productGroup.POST("", api.GetProductHandler(serviceName))
	}
	authServer := "auth-server"
	engine.Use(middleware.AuthMiddleware(authServer))
	// 购物车
	cartGroup := engine.Group("/cart")
	{
		serviceName = "cart-server"
		cartGroup.POST("", api.AddItemHandler(serviceName))
		cartGroup.POST("/get", api.GetCartHandler(serviceName))
		cartGroup.DELETE("", api.EmptyCartHandler(serviceName))
	}
	// 结算
	checkoutGroup := engine.Group("/checkout")
	{
		serviceName = "checkout-server"
		checkoutGroup.POST("", api.CheckoutHandler(ServiceName))
	}
	// 订单
	orderGroup := engine.Group("/order")
	{
		serviceName = "order-server"
		orderGroup.POST("", api.PlaceOrderHandler(serviceName))
		orderGroup.POST("/list", api.ListOrderHandler(serviceName))
		orderGroup.DELETE("", api.MarkOrderPaidHandler(serviceName))
	}
	// 支付
	paymentGroup := engine.Group("/payment")
	{
		serviceName = "payment-server"
		paymentGroup.POST("", api.ChargeHandler(serviceName))
	}
	// 启动 Gin 服务
	ginPort := viper.GetString("gin.port")
	log.Printf("Gin 服务监听的端口 :%v", ginPort)
	if err := engine.Run(":" + ginPort); err != nil {
		log.Fatalf("启动 Gin 服务失败: %v", err)
	}
	// 阻塞
	select {}
}
