package handlers

import (
	"betaproject/internal/models"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
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
	answer, err := generateAnswer(question.Question, question.MaxLength)
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