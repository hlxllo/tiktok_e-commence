package main

import (
	"tiktok_e-commence/app/payment/biz/dal"
	"tiktok_e-commence/app/payment/biz/dal/mysql"
	"tiktok_e-commence/app/payment/biz/model"
)

func main() {
	dal.Init()
	mysql.DB.Migrator().CreateTable(&model.PaymentPo{})
}
