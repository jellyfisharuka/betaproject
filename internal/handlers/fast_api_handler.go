package handlers

import (
	"betaproject/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"log"

	//"fmt"
	//"io"
	"net/http"

	//"os"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

// @Summary GenerateContent
// @Description Generate content based on the given prompt
// @Tags generate
// @Accept  json
// @Produce  json
// @Param question body models.Question true "Question to get answer for"
// @Success 200 {object} map[string]string
// @Router /api/generate/python [post]
func GeneratePythonHandler(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request", "details": err.Error()})
		return
	}
	ctx := c.Request.Context()
	answer, err := generateChatgpt(ctx, question.Question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate answer"})
		return
	}

	c.JSON(http.StatusOK, models.AnswerResponse{Answer: answer})
	
}
func generateAnswer(question string, maxLength int) (string, error) {
	const defaultMaxLength = 50
	requestBody := models.Question{
		Question:  question,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8000/generate/", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var answerResponse models.AnswerResponse
	err = json.NewDecoder(resp.Body).Decode(&answerResponse)
	if err != nil {
		return "", err
	}

	return answerResponse.Answer, nil
}


func generateChatgpt(ctx context.Context, question string) (string, error) {
	apiKey := "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"
	client := openai.NewClient(apiKey)

	request := openai.ChatCompletionRequest{
		Model: "gpt-4o-mini", 
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user", 
				Content: question,
			},
		},
		Temperature: 0.7, 
	}

	// Выполняем запрос
	response, err := client.CreateChatCompletion(ctx, request)
	if err != nil {
		log.Printf("Error while calling OpenAI: %v", err)
		return "", err
	}

	// Возвращаем текст ответа
	return response.Choices[0].Message.Content, nil
}

