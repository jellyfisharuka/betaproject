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
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
