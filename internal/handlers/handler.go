package handlers

import (
	"context"
	"html/template"
	"log"
	"net/http"
	

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/schema"
)

type Page struct {
	Images []string
}

var tmpl = template.Must(template.ParseFiles("../static/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
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

func GenerateHandler(w http.ResponseWriter, r *http.Request, llm *googleai.GoogleAI) {
	prompt := r.FormValue("prompt")

	if prompt == "" {
		http.Error(w, "Error: prompt is required", http.StatusBadRequest)
		return
	}

	content := []llms.MessageContent{
		{
			Role: schema.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{
				llms.TextPart(prompt),
			},
		},
	}

	_, err := llm.GenerateContent(r.Context(), content,
		llms.WithModel("gemini-1.5-flash"),
		llms.WithMaxTokens(500),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			w.Write(chunk)
			return nil
		}),
	)
	if err != nil {
		log.Printf("Error generating content: %v\n", err)
		http.Error(w, "Error: unable to generate content", http.StatusInternalServerError)
	}
}
