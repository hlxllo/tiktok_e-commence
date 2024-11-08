package model

import (
	"gorm.io/gorm"
	"tiktok_e-commence/app/checkout/biz/dal/mysql"
)

type CheckoutPo struct {
	gorm.Model
	UserId     uint32
	Firstname  string
	Lastname   string
	Email      string
	Address    []byte `gorm:"type:json"`
	CreditCard []byte `gorm:"type:json"`
}

func (table *CheckoutPo) TableName() string {
	return "checkout"
}

// 创建结算信息
func CreateCheckout(po *CheckoutPo) uint {
	mysql.DB.Create(po)
	return po.ID
}
