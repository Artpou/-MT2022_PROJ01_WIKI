package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	Pseudo    string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(pseudo string, password string) *User {
	user := User{Pseudo: pseudo, Email: pseudo, Password: password}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return &user
}
