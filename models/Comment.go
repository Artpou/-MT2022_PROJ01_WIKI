package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type Comment struct {
	ID           uint 		`gorm:"primaryKey"`
	AuthorID     uint			`gorm:"not null"`
	User         User			`gorm:"foreignKey:AuthorID"`
	ArticleID    uint			`gorm:"not null"`
	Article      Article	`gorm:"foreignKey:ArticleID"`
	Content      string 	`gorm:"not null;size:500"`
	CreationDate time.Time
	LatestUpdate time.Time
}

func NewComment(content string) *Comment {
	comment := Comment{Content: content}
	comment.AuthorID = 1
	comment.CreationDate = time.Now()
	comment.LatestUpdate = time.Now()
	return &comment
}
