package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tmc/langchaingo/llms/googleai"
	"betaproject/internal/handlers"
)

// NewRouter creates and initializes the router with all routes
func NewRouter(llm *googleai.GoogleAI) *mux.Router {
	r := mux.NewRouter()

	// Serve static files
	fs := http.FileServer(http.Dir("../static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Define API routes
	r.HandleFunc("/api/generate", func(w http.ResponseWriter, r *http.Request) {
		handlers.GenerateHandler(w, r, llm)
	}).Methods("POST")

	// Define the index route
	r.HandleFunc("/", handlers.IndexHandler).Methods("GET")

	return r
}
