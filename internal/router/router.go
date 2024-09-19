package router

import (
	/*"net/http"

	"github.com/gorilla/mux"
	"github.com/tmc/langchaingo/llms/googleai"
	"betaproject/internal/handlers"
	"github.com/swaggo/http-swagger"
	*/
	_ "betaproject/docs" // импортируйте свои swagger-документы
	"betaproject/internal/handlers"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tmc/langchaingo/llms/googleai"
)

// NewRouter creates and initializes the router with all routes
/*func NewRouter(llm *googleai.GoogleAI) *mux.Router {
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
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)


	return r
}
*/
var (
	apiKey = "AIzaSyBVOPL1HI_kRF2nsgByUz-EX7-YRbV6K_Q"
)
// @Beta Project
// @version 1.0
// @description This is a sample server.

// @host localhost:8080
// @BasePath /
func SetupRouter(r *gin.Engine)  {
	llm, err := googleai.New(context.Background(), googleai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.GET("/", handlers.IndexHandler)
	r.POST("/api/generate", func(c *gin.Context) {
		handlers.GenerateHandler(c, llm)
	})
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the API!")
	})
}
