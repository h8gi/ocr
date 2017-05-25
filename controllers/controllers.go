package controllers

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}
