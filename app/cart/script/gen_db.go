package main

import (
	"tiktok_e-commence/app/cart/biz/dal"
	"tiktok_e-commence/app/cart/biz/dal/mysql"
	"tiktok_e-commence/app/cart/biz/model"
)

func main() {
	dal.Init()
	mysql.DB.Migrator().CreateTable(&model.CartPo{})
	// 手动添加外键约束
	mysql.DB.Exec(`
        ALTER TABLE cart
        ADD CONSTRAINT fk_product_id
        FOREIGN KEY (product_id)
        REFERENCES product(id)
        ON DELETE CASCADE ON UPDATE CASCADE;
    `)
}
