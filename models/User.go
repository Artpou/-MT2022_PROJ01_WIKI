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
	Username     string `gorm:"unique;not null;size:255"`
	Password     string `gorm:"not null;size:255"`
	Role         Role   `json:"role"`
	CreationDate JSONTime
	LatestUpdate JSONTime
}

type UserWithoutPassword struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null;size:255"`
	Role         Role   `json:"role"`
	CreationDate JSONTime
	LatestUpdate JSONTime
}

func NewUser(username string, user_password string) *User {
	hash, _ := password.HashPassword(user_password)
	user := User{Username: username, Password: hash}
	user.Role = UserRole
	user.CreationDate = JSONTime(time.Now())
	user.LatestUpdate = JSONTime(time.Now())
	return &user
}

func NewUserWithoutPassword(user *User) *UserWithoutPassword {
	return &UserWithoutPassword{
		ID:           user.ID,
		Username:     user.Username,
		Role:         user.Role,
		CreationDate: user.CreationDate,
		LatestUpdate: user.LatestUpdate,
	}
}

func UpdateUser(user User, user_password string) *User {
	hash, _ := password.HashPassword(user_password)
	user.Password = hash
	user.LatestUpdate = JSONTime(time.Now())
	return &user
}
