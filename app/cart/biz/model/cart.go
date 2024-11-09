package model

import (
	"gorm.io/gorm"
	"tiktok_e-commence/app/cart/biz/dal/mysql"
	"tiktok_e-commence/common/model/model"
)

type CartPo struct {
	gorm.Model
	// 复合唯一约束
	UserId    uint32 `gorm:"uniqueIndex:idx_user_product"`
	ProductId uint32 `gorm:"type:bigint unsigned;uniqueIndex:idx_user_product"`
	Quantity  int32
}

// 服了，神奇 bug TODO
type AddItemReqCopy struct {
	model.AddItemReq
}

type GetCartReqCopy struct {
	model.GetCartReq
}

type EmptyCartReqCopy struct {
	model.EmptyCartReq
}

func (table *CartPo) TableName() string {
	return "cart"
}

// 新增购物车
func CreateCart(po *CartPo) (uint, error) {
	result := mysql.DB.Unscoped().Create(&po)
	return po.ID, result.Error
}

// 批量查询购物车
func SelectCarts(po *CartPo) []*CartPo {
	var carts []*CartPo
	mysql.DB.Where(po).Find(&carts)
	return carts
}

// 批量删除购物车
func DeleteCarts(po *CartPo) {
	mysql.DB.Where(po).Delete(&CartPo{})
}
