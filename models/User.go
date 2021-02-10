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

func create(pseudo string, password string) User {
	return User{
		Pseudo:   pseudo,
		Password: password,
	}
}
