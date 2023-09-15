package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primary_key;autoIncrement"`
	Username string `gorm:"uniqueIndex;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Phone    string `gorm:"uniqueIndex;null"`
	gorm.Model
}
