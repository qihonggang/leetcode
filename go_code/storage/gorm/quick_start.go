package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(localhost)/beego?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移模式
	db.AutoMigrate(&Product{})

	//创建
	db.Create(&Product{Code: "L1212", Price: 1000})

	//读取
	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "L1212")

	//更新 - 更新product的price为2000
	db.Model(&product).Update("price", 2000)

	//删除
	db.Delete(&product)
}
