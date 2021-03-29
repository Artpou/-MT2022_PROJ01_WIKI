package models

import (
	"time"
	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID           uint `gorm:"primaryKey"`
	Email        string `gorm:"unique;not null;size:255"`
	Username     string `gorm:"unique;not null;size:255"`
	Password     string `gorm:"not null;size:255"`
	CreationDate time.Time
	LatestUpdate time.Time
}


func NewUser(username string, password string) *User {
	hash, _ := HashPassword(password)
	user := User{Username: username, Email: username, Password: hash}
	user.CreationDate = time.Now()
	user.LatestUpdate = time.Now()
	return &user
}
