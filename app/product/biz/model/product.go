package model

import (
	"gorm.io/gorm"
	"tiktok_e-commence/app/product/biz/dal/mysql"
)

type ProductPo struct {
	gorm.Model
	Name        string
	Description string
	Picture     string
	Price       float32
	Categories  []byte `gorm:"type:json"`
}

func (table *ProductPo) TableName() string {
	return "product"
}

// 根据类别查询
func SelectProductByCat(category string) []*ProductPo {
	var products []*ProductPo
	mysql.DB.Where("JSON_CONTAINS(categories, ?)", `"`+category+`"`).Find(&products)
	// 找不到返回空
	return products
}
