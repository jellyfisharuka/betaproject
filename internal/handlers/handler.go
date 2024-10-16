package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/schema"
)

// GenerateHandler godoc
// @Summary GenerateContent
// @Description Generate content based on the given prompt
// @Tags generate
// @Accept x-www-form-urlencoded
// @Produce plain
// @Param prompt formData string true "Prompt"
// @Success 200 {string} string "Generated content"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/generate [post]
func GenerateHandler(c *gin.Context, llm *googleai.GoogleAI) {
	
	prompt, exists := c.Get("prompt")

	if !exists {
		// Если в контексте нет prompt, пытаемся получить его из формы
		prompt = c.PostForm("prompt")
	}

	promptStr, ok := prompt.(string)
	if !ok || promptStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Prompt is required"})
		return
	}

	content := []llms.MessageContent{
		{
			Role: schema.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{
				llms.TextPart(promptStr),
			},
		},
	}

	_, err := llm.GenerateContent(c.Request.Context(), content,
		llms.WithModel("gemini-1.5-flash"),
		llms.WithMaxTokens(500),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			c.Writer.Write(chunk)
			return nil
		}),
	)
	if err != nil {
		log.Printf("Error generating content: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate content"})
	}
}

// IndexHandler godoc
// @Summary Index Page
// @Description Renders the index page
// @Tags index
// @Produce html
// @Success 200 {html} html "Index page"
// @Router / [get]
func IndexHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
}

