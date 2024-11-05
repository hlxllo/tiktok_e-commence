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

// 创建用户
func CreateUser(user *User) (uint, error) {
	result := mysql.DB.Create(&user)
	// 返回插入数据主键
	return user.ID, result.Error
}
