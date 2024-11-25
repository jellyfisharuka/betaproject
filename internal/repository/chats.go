package repository

import (
	"betaproject/internal/models"

	"gorm.io/gorm"
)

type ChatRepository struct {
    DB *gorm.DB
}

// Создание нового чата
func (r *ChatRepository) CreateChat(chat *models.Chat) error {
    return r.DB.Create(chat).Error
}

// Получение чата по ID
func (r *ChatRepository) GetChatByID(id int) (models.Chat, error) {
    var chat models.Chat
    err := r.DB.Preload("Messages").First(&chat, id).Error
    return chat, err
}

// Обновление чата
func (r *ChatRepository) UpdateChat(chat *models.Chat) error {
    return r.DB.Save(chat).Error
}

// Удаление чата
func (r *ChatRepository) DeleteChat(id uint) error {
    return r.DB.Delete(&models.Chat{}, id).Error
}

// Получение всех чатов
func (r *ChatRepository) GetAllChats() ([]models.Chat, error) {
    var chats []models.Chat
    err := r.DB.Preload("Messages").Find(&chats).Error
    return chats, err
}
func (r *ChatRepository) AddMessageToChat(message *models.Message) error {
    return r.DB.Create(message).Error
}