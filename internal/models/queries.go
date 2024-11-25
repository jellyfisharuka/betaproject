package models

import "time"

type Query struct {
	Question string
	Answer   string
	User     User
}
type Chat struct {
	ID        int      `gorm:"primaryKey"`
	UserID    int      `gorm:"not null"`
	Messages  []Message `gorm:"foreignKey:ChatID"`
}
type Message struct {
    ID        uint      `gorm:"primaryKey"`
    ChatID    uint      `gorm:"not null"`
    Prompt    string    `gorm:"type:text;not null"`  // Запрос (prompt) пользователя
    Answer    string    `gorm:"type:text;not null"`
    SenderID  uint      `gorm:"not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}


type MessageSwagger struct {
    ChatID  int    `json:"chatID"`  // The ID of the chat
    Prompt  string `json:"prompt"`  // The message content (input from user)
    SenderID uint  `json:"senderID"` // The ID of the user sending the message
}
