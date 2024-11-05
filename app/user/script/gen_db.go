package main

import (
	"tiktok_e-commence/app/user/biz/dal"
	"tiktok_e-commence/app/user/biz/dal/mysql"
	"tiktok_e-commence/app/user/biz/model"
)

// 生成数据库表
func main() {
	dal.Init()
	mysql.DB.Migrator().CreateTable(&model.User{})
}
