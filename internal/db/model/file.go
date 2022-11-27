package model

import "gorm.io/gorm"

const (
	Active = iota + 1
)

type File struct {
	gorm.Model
	FileSha256 string `json:"file_sha256" gorm:"unique"`
	FileName   string `json:"file_name"`
	FileSize   int64  `json:"file_size"`
	FileAddr   string `json:"file_addr"`
	Status     int    `json:"status" gorm:"index"`
	Ext1       int    `json:"ext1"`
	Ext2       string `json:"ext2" gorm:"type:text"`
}

func (f *File) TableName() string {
	return "meta_data"
}
