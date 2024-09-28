package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"` //"admin", "user"
}