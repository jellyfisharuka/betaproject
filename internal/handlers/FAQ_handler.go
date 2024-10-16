package handlers

import (
	"betaproject/internal/db"
	"betaproject/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get FAQ Answer
// @Description Generate content based on the given prompt
// @Tags faq
// @Accept  json
// @Produce  json
// @Param question query string true "Question to get answer for"
// @Success 200 {object} map[string]string
// @Router /faq [get]
func FAQHandler(c *gin.Context) {
	question := c.Query("question")
	if question=="" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a question"})
		return
	}
	answer, err := getAnswerFromDB(question)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}
	//response := map[string]string{
	//	"question": question,
	//	"answer":   answer,
	//}
	c.JSON(http.StatusOK, answer)

}
func getAnswerFromDB(question string) (string, error) {
	var faq models.FAQ
	if err := db.DB.Where("question = ?", question).First(&faq).Error; err != nil {
		return "", err
	}
	return faq.Answer, nil
}