package models

import (
	"time"

	"github.com/Artpou/wiki_golang/handler/password"
	_ "github.com/jinzhu/gorm"
)

type Role int

const (
	UserRole  = iota
	AdminRole = iota
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null;size:255" json:"username"`
	Password     string `gorm:"not null;size:255" json:"password"`
	Role         Role   `json:"role"`
	CreationDate time.Time
	LatestUpdate time.Time
}

func NewUser(username string, user_password string) *User {
	hash, _ := password.HashPassword(user_password)
	user := User{Username: username, Password: hash}
	user.Role = AdminRole
	user.CreationDate = time.Now()
	user.LatestUpdate = time.Now()
	return &user
}

func UpdateUser(user User, user_password string) *User {
	hash, _ := password.HashPassword(user_password)
	user.Password = hash
	user.LatestUpdate = time.Now()
	return &user
}
