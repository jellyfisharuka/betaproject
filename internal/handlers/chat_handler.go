package handlers

import (
	"betaproject/internal/models"
	"betaproject/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms/googleai"
)

type ChatHandler struct {
	Repo *repository.ChatRepository
}
type ErrorResponse struct {
    Error string `json:"error"`
}
// @Security Bearer
// CreateChat godoc
// @Summary Create a new chat
// @Description Create a new chat by providing the chat details (empty by default)
// @Tags chats
// @Accept json
// @Produce json
// @Success 201 {object} models.Chat
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /chats [post]
func (h *ChatHandler) CreateChat(c *gin.Context) {
	userID, exists := c.Get("ID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
        return
    }

    var chat models.Chat
    // Используем извлеченный userID
    chat.UserID = userID.(int) // Приводим к типу uint, если это необходимо

    // Создаем чат в репозитории
    if err := h.Repo.CreateChat(&chat); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create chat"})
        return
    }

    // Возвращаем созданный чат
    c.JSON(http.StatusCreated, chat)
}
// @Security Bearer
// GetChat godoc
// @Summary Get a chat by ID
// @Description Retrieve a chat by its ID
// @Tags chats
// @Accept json
// @Produce json
// @Param id path int true "Chat ID"
// @Success 200 {object} models.Chat
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /chats/{id} [get]
func (h *ChatHandler) GetChat(c *gin.Context) {
	idStr := c.Param("id") // Get the id as a string
    id, err := strconv.Atoi(idStr) // Convert string to int
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
        return
    }

    chat, err := h.Repo.GetChatByID(id) // Pass the integer ID
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Chat not found"})
        return
    }

    c.JSON(http.StatusOK, chat)
}

// @Security Bearer
// SendMessage godoc
// @Summary Send a message to a chat
// @Description Send a message to a specified chat, and generate an AI response
// @Tags chats
// @Accept json
// @Produce json
// @Param chatID path int true "Chat ID"
// @Param message body models.MessageSwagger true "Message content"
// @Success 201 {object} models.MessageSwagger "Message sent successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 401 {object} ErrorResponse "User not authenticated"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /chats/{chatID}/messages [post]
func (h *ChatHandler) SendMessage(c *gin.Context, llm *googleai.GoogleAI) {
    var message models.Message
    if err := c.ShouldBindJSON(&message); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, exists := c.Get("ID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
        return
    }

    message.SenderID = uint(userID.(int)) 

    c.Set("prompt", message.Prompt)

    generatedAnswer, err := GenerateHandler(c, llm)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate response"})
        return
    }

    message.Answer = generatedAnswer

    if err := h.Repo.AddMessageToChat(&message); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
        return
    }

    c.JSON(http.StatusCreated, message)
}

// @Security Bearer
// GetChat godoc
// @Summary Get all chats 
// @Description Retrieve all chats
// @Tags chats
// @Accept json
// @Produce json
// @Success 200 {object} models.Chat
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /chats [get]
func(h *ChatHandler) GetAllChats(c *gin.Context) {
    chat, err := h.Repo.GetAllChats() 
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve chats"})
        return
    }
    c.JSON(http.StatusOK, chat)
}