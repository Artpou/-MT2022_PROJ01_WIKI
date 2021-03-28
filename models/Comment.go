package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type Comment struct {
	ID           uint `gorm:"primaryKey"`
	User         User
	AuthorID     uint
	Article      Article
	ArticleID    uint
	Content      string `gorm:"size:500"`
	CreationDate time.Time
	LatestUpdate time.Time
}

func NewComment(content string) *Comment {
	comment := Comment{Content: content}
	comment.CreationDate = time.Now()
	comment.LatestUpdate = time.Now()
	return &comment
}
