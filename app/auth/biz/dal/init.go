package dal

import (
	"fmt"
	"github.com/spf13/viper"
)

// 读取配置并初始化数据库连接
func Init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("app/auth/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("读取配置成功！")
	// TODO 暂时不需要
	//mysql.InitMysql()
	//redis.InitRedis()
}
