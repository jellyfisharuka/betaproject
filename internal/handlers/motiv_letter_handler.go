package handlers

import (
	"fmt"
	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms/googleai"
)
// @Security Bearer
// @Security GoogleOAuth
// @Summary Create Motivational Letter
// @Description Generate a motivational letter based on user data
// @Tags letters
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param name formData string true "Name"
// @Param age formData int true "Age"
// @Param university formData string true "University"
// @Param country formData string true "Country"
// @Success 200 {object} map[string]string
// @Router /api/generate/motivational_letter [post]
func CreateMotivationalLetterHandler(c *gin.Context, llm *googleai.GoogleAI) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	university := c.PostForm("university")
	country := c.PostForm("country")

	prompt := fmt.Sprintf("My name is %s, I am %s years old. I want to apply to %s. I am from %s. Please generate a complete motivational letter for university admission without any placeholders or example instructions. Include only the final content of the letter.", name, age, university, country)

	c.Set("prompt", prompt) 
     GenerateHandler(c, llm)  

	//c.JSON(http.StatusOK, gin.H{"motivational_letter": answer})
}