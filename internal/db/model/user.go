package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	Phone          string    `josn:"phone" gorm:"unique"`
	EmailValidated int8      `josn:"email_validated"`
	PhoneValidated int8      `json:"phone_validated"`
	LastActive     time.Time `json:"last_active"`
	Profile        string    `json:"profile"`
	Status         int8      `json:"status" gorm:"index"`
}
