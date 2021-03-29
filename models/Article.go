package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type Article struct {
<<<<<<< HEAD
	ID           uint 	`gorm:"primaryKey"`
	AuthorID     uint		`gorm:"not null"`
	User         User 	`gorm:"foreignKey:AuthorID"`
=======
	ID           uint   `gorm:"primaryKey"`
	AuthorID     uint   `gorm:"not null"`
	User         User   `gorm:"foreignKey:AuthorID"  json:"-"`
>>>>>>> 0b914ca917d689ab04fe85748db34a24fb44a3c6
	Title        string `gorm:"not null;size:255"`
	Content      string `gorm:"not null;size:10000"`
	CreationDate time.Time
	LatestUpdate time.Time
}

type ArticleWithComments struct {
	Article
	Comments []Comment
}

func NewArticle(title string, content string) *Article {
	article := Article{Title: title, Content: content}
	article.AuthorID = 1
	article.CreationDate = time.Now()
	article.LatestUpdate = time.Now()
	return &article
}

func UpdateArticle(article Article, title string, content string) *Article {
	if title != "" {
		article.Title = title
	}
	if content != "" {
		article.Content = content
	}
	article.LatestUpdate = time.Now()
	return &article
}
