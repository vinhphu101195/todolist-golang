package models

import (
	"github.com/jinzhu/gorm"
)

//SetupModels ...
func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "vinhphu101195:Phu02573419@(db4free.net:3306)/todolistgolang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&TodoModel{})
	return db
}
