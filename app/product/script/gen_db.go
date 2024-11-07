package main

import (
	"tiktok_e-commence/app/product/biz/dal"
	"tiktok_e-commence/app/product/biz/dal/mysql"
	"tiktok_e-commence/app/product/biz/model"
)

func main() {
	dal.Init()
	mysql.DB.Migrator().CreateTable(&model.ProductPo{})
}
