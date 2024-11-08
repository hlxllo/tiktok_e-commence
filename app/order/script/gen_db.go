package main

import (
	"tiktok_e-commence/app/order/biz/dal"
	"tiktok_e-commence/app/order/biz/dal/mysql"
	"tiktok_e-commence/app/order/biz/model"
)

func main() {
	dal.Init()
	mysql.DB.Migrator().CreateTable(&model.OrderPo{})
}
