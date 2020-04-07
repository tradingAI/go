package model

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
	Name        string `gorm:"column:name; not null" json:"name"`
	Version     string `gorm:"column:version; not null" json:"version"`
	Description string `gorm:"column:description; not null" json:"description"`
	FileType    string `gorm:"column:file_type; not null" json:"file_type"`
	UserID      uint64 `gorm:"column:user_id; not null" json:"user_id"`
	Status      int    `gorm:"column:status; not null" json:"status"`
	Bucket      string `gorm:"column:bucket; not null" json:"bucket"`
	ObjName     string `gorm:"column:obj_name; not null" json:"obj_name"`
}

func (Model) TableName() string {
	return "models"
}

var ModelFileTypeMap map[string]string = map[string]string{
	"application/x-tar": "tar",
	"application/zip":   "zip",
}
