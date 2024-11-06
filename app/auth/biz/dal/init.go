package dal

import (
	"fmt"
	"github.com/spf13/viper"
	"tiktok_e-commence/app/auth/biz/dal/redis"
	"tiktok_e-commence/app/user/biz/dal/mysql"
)

// 读取配置并初始化数据库连接
func Init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("app/user/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("读取配置成功！")
	mysql.InitMysql()
	redis.InitRedis()
}
