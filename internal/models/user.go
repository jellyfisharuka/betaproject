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
	Chats     []Chat
}
type SignupSwagger struct {
    Username  string `json:"username"`  // Уникальное имя пользователя
    Password  string `json:"password"`  // Пароль пользователя
    FirstName string `json:"firstName"` // Имя пользователя
    LastName  string `json:"lastName"`  // Фамилия пользователя
    Telephone string `json:"telephone,omitempty"`  // Номер телефона пользователя
    RoleID    int    `json:"roleID,omitempty"`    // ID роли пользователя
}
type LoginSwagger struct {
    Username string `json:"username"` // Уникальное имя пользователя
    Password string `json:"password"` // Пароль пользователя
}

type UserInfo struct {
	ID string `json:"sub"`
	Email string `json:"email"`
	Name string `json:"name"`
}