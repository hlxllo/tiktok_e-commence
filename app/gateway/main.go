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
	dal.Init()
	Ip := viper.GetString("server.ip")
	Port := viper.GetInt("server.port")
	ServiceName := viper.GetString("server.service_name")
	common.InitNacosClient()
	common.RegisterToNacos(Ip, Port, ServiceName)
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
		productGroup.GET("/list", api.ListProductsHandler(serviceName))
		productGroup.GET("", api.GetProductHandler(serviceName))
	}
	authServer := "auth-server"
	engine.Use(middleware.AuthMiddleware(authServer))
	// 购物车
	cartGroup := engine.Group("/cart")
	{
		serviceName = "cart-server"
		cartGroup.POST("", api.AddItemHandler(serviceName))
		cartGroup.GET("", api.GetCartHandler(serviceName))
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
		orderGroup.GET("", api.ListOrderHandler(serviceName))
		orderGroup.DELETE("", api.MarkOrderPaidHandler(serviceName))
	}
	// 支付
	paymentGroup := engine.Group("/payment")
	{
		serviceName = "payment-server"
		paymentGroup.POST("", api.ChargeHandler(serviceName))
	}
	ginPort := viper.GetString("gin.port")
	log.Printf("Gin 服务监听的端口 :%v", ginPort)
	if err := engine.Run(":" + ginPort); err != nil {
		log.Fatalf("启动 Gin 服务失败: %v", err)
	}
	select {}
}
