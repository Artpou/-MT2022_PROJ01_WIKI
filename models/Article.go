package models

import (
	"time"
	_ "github.com/jinzhu/gorm"
)

type Article struct {
	ID           uint `gorm:"primaryKey"`
	User				 User
	AuthorID     uint
	Title        string
	Content      string `gorm:"size:10000"`
	CreationDate time.Time
	LatestUpdate time.Time
}
