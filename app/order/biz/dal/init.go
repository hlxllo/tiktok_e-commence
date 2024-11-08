package dal

import (
	"fmt"
	"github.com/spf13/viper"
	"tiktok_e-commence/app/order/biz/dal/mysql"
)

// 读取配置并初始化数据库连接
func Init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("app/order/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("读取配置成功！")
	mysql.InitMysql()
}
