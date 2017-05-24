package models

import "github.com/jinzhu/gorm"

type File struct {
	gorm.Model
	Path string
	Type string
}
