package model

import (
	"gorm.io/gorm"
	"tiktok_e-commence/app/order/biz/dal/mysql"
)

type OrderPo struct {
	gorm.Model
	UserId       uint32
	UserCurrency string
	Email        string
	Address      []byte `gorm:"type:json"`
	OrderItems   []byte `gorm:"type:json"`
}

func (table *OrderPo) TableName() string {
	return "order"
}

// 新增订单
func CreateOrder(po *OrderPo) uint {
	mysql.DB.Create(po)
	return po.ID
}

// 查询订单
func SelectOrders(queryPo *OrderPo) []*OrderPo {
	var po []*OrderPo
	mysql.DB.Where(queryPo).Find(&po)
	return po
}

// 删除订单
func DeleteOrder(po *OrderPo) uint {
	result := mysql.DB.Delete(po)
	return uint(result.RowsAffected)
}
