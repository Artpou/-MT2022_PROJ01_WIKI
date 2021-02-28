package models

import (
	"time"
	_ "github.com/jinzhu/gorm"
)

type Comment struct {
	ID           uint `gorm:"primaryKey"`
	User				 User
	AuthorID     uint
	Article			 Article
	ArticleID    uint
	Content      string `gorm:"size:500"`
	CreationDate time.Time
	LatestUpdate time.Time
}
