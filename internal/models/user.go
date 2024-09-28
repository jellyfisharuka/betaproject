package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        int    `gorm:"uniqueIndex;not null"`
	Username  string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	FirstName string
	LastName  string
	Telephone string
	RoleID    int
	Role      Role
}