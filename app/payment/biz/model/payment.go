package model

import (
	"gorm.io/gorm"
	"tiktok_e-commence/app/payment/biz/dal/mysql"
	"tiktok_e-commence/common/model/model"
)

type PaymentPo struct {
	gorm.Model
	Amount     float32
	CreditCard []byte `gorm:"type:json"`
	OrderId    string
	UserId     uint32
}

type ChargeReqCopy struct {
	model.ChargeReq
}

func (table *PaymentPo) TableName() string {
	return "payment"
}

// 创建支付信息
func CreatePayment(po *PaymentPo) uint {
	mysql.DB.Unscoped().Create(po)
	return po.ID
}
