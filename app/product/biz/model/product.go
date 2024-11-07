package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ProductPo struct {
	gorm.Model
	Name        string
	Description string
	Picture     string
	Price       float32
	Categories  pq.StringArray `gorm:"type:json"`
}

func (table *ProductPo) TableName() string {
	return "product"
}
