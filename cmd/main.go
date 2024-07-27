package main

import (
	"context"
	"flag"
	"log"
	
	"os"
     _ "betaproject/docs"
	"betaproject/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms/googleai"
	 ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

var (
	addr   = flag.String("addr", "localhost:8080", "address to serve")
	apiKey = "AIzaSyBVOPL1HI_kRF2nsgByUz-EX7-YRbV6K_Q"
)

func main() {
	// Parse flags
	flag.Parse()

	// Get the Gemini API key from the environment, if set
	if key := os.Getenv("API_KEY"); key != "" {
		apiKey = key
	}

	// Initialize the Gemini API client
	llm, err := googleai.New(context.Background(), googleai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Serve static files
	r.Static("/static", "../static")

	// Define API routes
	r.POST("/api/generate", func(c *gin.Context) {
		handlers.GenerateHandler(c, llm)
	})

	// Define the index route
	r.GET("/", handlers.IndexHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server
	log.Printf("Starting server on %s", *addr)
	log.Fatal(r.Run(*addr))
}
