package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
)

var REDIS *redis.Client

// 初始化redis配置
func InitRedis() {
	// 配置数据库连接
	REDIS = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.Addr"),
		Password: viper.GetString("redis.Password"),
		DB:       viper.GetInt("redis.DB"),
	})
	// 测试连接
	pong, err := REDIS.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis 连接失败: %v", err)
	}
	fmt.Println("redis连接成功：", pong)
}
