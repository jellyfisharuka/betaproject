package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/tmc/langchaingo/llms/googleai"
	"betaproject/internal/router"
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

	// Initialize the router and handlers
	r := router.NewRouter(llm)

	// Start the server
	log.Printf("Starting server on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, r))
}
