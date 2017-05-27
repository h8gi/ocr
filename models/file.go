package models

import (
	"github.com/jinzhu/gorm"
	"github.com/otiai10/gosseract"
)

type File struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"`
	Path string `gorm:"unique" json:"path"`
	Type string `json:"type"`
	Size int64  `json:"size"`
	Text string `json:"text"`
}

func (f *File) BeforeCreate() (err error) {
	f.Text = gosseract.Must(gosseract.Params{Src: f.Path, Languages: "jpn"})
	return err
}
