package model

import "gorm.io/gorm"

type CartPo struct {
	gorm.Model
	UserId    uint32
	ProductId uint32 `gorm:"type:bigint unsigned"`
	Quantity  int32
}

func (table *CartPo) TableName() string {
	return "cart"
}
