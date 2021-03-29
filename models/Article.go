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
	Title        string `gorm:"not null;size:255"`
	Content      string `gorm:"not null;size:10000"`
=======
	ID           uint `gorm:"primaryKey"`
	AuthorID     uint
	User         User   `gorm:"foreignKey:AuthorID" json:"-"`
	Title        string `gorm:"size:255"`
	Content      string `gorm:"size:10000"`
>>>>>>> b9687e51e84908969f60bab3eb5e129f9770b6bd
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
