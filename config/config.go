package config

import (
	"assignment_2/structs"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/orders_by")

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(structs.Orders{})
	db.AutoMigrate(structs.Items{})
	return db
}
