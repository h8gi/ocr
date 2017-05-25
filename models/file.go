package models

import "github.com/jinzhu/gorm"

type File struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"`
	Path string `gorm:"unique" json:"path"`
	Type string `json:"type"`
	Size int64  `json:"size"`
}
