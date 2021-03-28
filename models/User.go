package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID           uint `gorm:"primaryKey"`
	Email        string
	Username     string
	Password     string
	CreationDate time.Time
	LatestUpdate time.Time
}

func NewUser(username string, password string) *User {
	user := User{Username: username, Email: username, Password: password}
	user.CreationDate = time.Now()
	user.LatestUpdate = time.Now()
	return &user
}
