package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type Article struct {
	ID           uint   `gorm:"primaryKey"`
	AuthorID     uint   `gorm:"not null"`
	User         User   `gorm:"foreignKey:AuthorID"`
	Title        string `gorm:"not null;size:255"`
	Content      string `gorm:"not null;size:10000"`
	CreationDate JSONTime
	LatestUpdate JSONTime
}

type ArticleWithOwner struct {
	Owner string
	Article
}

type ArticleWithComments struct {
	Article  *ArticleWithOwner
	Comments []Comment
}

func NewArticle(title string, content string) *Article {
	article := Article{Title: title, Content: content}
	article.AuthorID = 2
	article.CreationDate = JSONTime(time.Now())
	article.LatestUpdate = JSONTime(time.Now())
	return &article
}

func NewArticleWithOwner(article Article) *ArticleWithOwner {
	articleOwner := ArticleWithOwner{Owner: article.User.Username, Article: article}
	return &articleOwner
}

func NewArticleWithComments(article Article, comments []Comment) *ArticleWithComments {
	articleComments := ArticleWithComments{Article: NewArticleWithOwner(article), Comments: comments}
	return &articleComments
}

func UpdateArticle(article Article, title string, content string) *Article {
	if title != "" {
		article.Title = title
	}
	if content != "" {
		article.Content = content
	}
	article.LatestUpdate = JSONTime(time.Now())
	return &article
}
