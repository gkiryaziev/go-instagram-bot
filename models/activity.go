package models

import (
	"github.com/jinzhu/gorm"
)

// Activity struct
type Activity struct {
	gorm.Model

	Pk           string `gorm:"not null;unique"`
	MediaID      string
	MediaURL     string
	UserID       int64
	UserImageURL string
	UserName     string
	Type         string
	Comment      string
}

// TableName set table name
func (a *Activity) TableName() string {
	return "tbl_activity"
}
