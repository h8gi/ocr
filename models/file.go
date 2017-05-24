package models

import "github.com/jinzhu/gorm"

type File struct {
	gorm.Model
	Name string
	Type string
	Size int
}
