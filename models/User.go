package models

import (
	"time"
	"github.com/Artpou/wiki_golang/handler"
	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID           uint `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null;size:255"`
	Password     string `gorm:"not null;size:255" json:"-"`
	CreationDate time.Time
	LatestUpdate time.Time
}


func NewUser(username string, password string) *User {
	hash, _ := handler.HashPassword(password)
	user := User{Username: username, Password: hash}
	user.CreationDate = time.Now()
	user.LatestUpdate = time.Now()
	return &user
}

func UpdateUser(user User, password string) *User {
	hash, _ := handler.HashPassword(password)
	user.Password = hash
	user.LatestUpdate = time.Now()
	return &user
}
