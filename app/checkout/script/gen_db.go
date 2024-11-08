package main

import (
	"tiktok_e-commence/app/checkout/biz/dal"
	"tiktok_e-commence/app/checkout/biz/dal/mysql"
	"tiktok_e-commence/app/checkout/biz/model"
)

func main() {
	dal.Init()
	mysql.DB.Migrator().CreateTable(&model.CheckoutPo{})
}
