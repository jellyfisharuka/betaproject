package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func BetaTest(c *gin.Context) {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyBVOPL1HI_kRF2nsgByUz-EX7-YRbV6K_Q"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// The Gemini 1.5 models are versatile and work with most use cases
	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text("essay about weather."))
	if err != nil {
		log.Fatal(err)
	}

	if resp != nil {
		candidates := resp.Candidates
		if candidates != nil {
			var result []string
			for _, candidate := range candidates {
				content := candidate.Content
				if content != "" {
					result = append(result, content)
				}
			}
			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "No candidates found",
	})
}
