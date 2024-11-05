package model

import (
	"gorm.io/gorm"
	"tiktok_e-commence/app/user/biz/dal/mysql"
)

type User struct {
	gorm.Model
	Password string
	Email    string `gorm:"unique"`
}

// 设置表名
func (table *User) TableName() string {
	return "user"
}

// 新增用户
func CreateUser(user *User) (uint, error) {
	result := mysql.DB.Create(&user)
	// 返回插入数据主键
	return user.ID, result.Error
}

// 查询单个用户
func SelectUser(user *User) (uint, error) {
	result := mysql.DB.First(&user)
	// 返回查询用户的主键
	return user.ID, result.Error
}
