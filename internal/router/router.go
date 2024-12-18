package router

import (
	/*"net/http"

	"github.com/gorilla/mux"
	"github.com/tmc/langchaingo/llms/googleai"
	"betaproject/internal/handlers"
	"github.com/swaggo/http-swagger"
	*/
	_ "betaproject/docs" // импортируйте свои swagger-документы
	"betaproject/internal/auth"
	"betaproject/internal/db"
	"betaproject/internal/handlers"
	"betaproject/internal/repository"
	"context"
	"log"
	"net/http"
	"net/http/pprof"
	"path/filepath"

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
	llm *googleai.GoogleAI
)


func SetupRouter(r *gin.Engine)  {
	var err error
    llm, err = googleai.New(context.Background(), googleai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	
	//handlers.Oauth2Config, err = auth.GetOAuth2Config("C:/Users/aruke/Desktop/golang/betaproject/cmd/gmail.json") ("./internal/config/gmail.json")
	//handlers.Oauth2Config, err = auth.GetOAuth2Config("./internal/config/gmail.json")
	log.Printf("Trying to load Gmail config from: %s", "/app/config/gmail.json")
handlers.Oauth2Config, err = auth.GetOAuth2Config("/app/config/gmail.json")
    if err != nil {
	log.Fatalf("Error loading OAuth2 configuration: %v", err)
}
log.Printf("OAuth2 Configuration: %+v", handlers.Oauth2Config)
chatRepo := &repository.ChatRepository{DB: db.DB}
    chatHandler := &handlers.ChatHandler{Repo: chatRepo}
if handlers.Oauth2Config == nil {
    log.Fatalf("OAuth2 configuration is nil")
}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.GET("/", handlers.IndexHandler)
	//r.POST("/api/generate", func(c *gin.Context) {
	//	handlers.GenerateHandler(c, llm)
	//})
	r.POST("/api/generate", func(c *gin.Context) {
		handlers.GenerateHandler(c, llm)})
	r.POST("/api/generate/python", handlers.GeneratePythonHandler)
	r.POST("/api/generate/motivational_letter", auth.AuthMiddleware(), func(c *gin.Context) {
		handlers.CreateMotivationalLetterHandler(c, llm)})
		r.POST("/api/generate/recommendation_letter", func(c *gin.Context) {
			handlers.CreateRecommendationLetterHandler(c, llm)})
	r.POST("/login", handlers.LoginHandler)
	r.POST("/signup", handlers.SignupHandler)
	r.GET("/oauth2callback", handlers.OAuth2CallbackHandler)
	r.GET("/googleLogin", handlers.LoginGoogleHandler(handlers.Oauth2Config))
	r.GET("/faq", handlers.FAQHandler)
	r.POST("/chats", auth.AuthMiddleware(), chatHandler.CreateChat)
    r.GET("/chats/:id", auth.AuthMiddleware(), chatHandler.GetChat)
    r.POST("/chats/:chatID/messages", auth.AuthMiddleware(), func(c *gin.Context) {
		chatHandler.SendMessage(c, llm) 
	})
	r.GET("/chats", auth.AuthMiddleware(), chatHandler.GetAllChats)

    pprofGroup := r.Group("/debug/pprof")
    {
        pprofGroup.GET("/", gin.WrapH(http.HandlerFunc(pprof.Index)))
        pprofGroup.GET("/cmdline", gin.WrapH(http.HandlerFunc(pprof.Cmdline)))
        pprofGroup.GET("/profile", gin.WrapH(http.HandlerFunc(pprof.Profile)))
        pprofGroup.GET("/symbol", gin.WrapH(http.HandlerFunc(pprof.Symbol)))
        pprofGroup.GET("/trace", gin.WrapH(http.HandlerFunc(pprof.Trace)))
        pprofGroup.GET("/heap", gin.WrapH(http.HandlerFunc(pprof.Handler("heap").ServeHTTP)))
        pprofGroup.GET("/goroutine", gin.WrapH(http.HandlerFunc(pprof.Handler("goroutine").ServeHTTP)))
        pprofGroup.GET("/threadcreate", gin.WrapH(http.HandlerFunc(pprof.Handler("threadcreate").ServeHTTP)))
        pprofGroup.GET("/block", gin.WrapH(http.HandlerFunc(pprof.Handler("block").ServeHTTP)))
        pprofGroup.GET("/mutex", gin.WrapH(http.HandlerFunc(pprof.Handler("mutex").ServeHTTP)))
    }
	
	//r.Static("/static", "./static")
	
	//r.GET("/redirect", handlers.RedirectPageHandler)

	//r.GET("/redirect", func(c *gin.Context) {
      //  c.File("./static/redirect.html") // Предоставляем файл напрямую
    //})
	r.LoadHTMLGlob(filepath.Join("..", "static", "*.html")) // Adjusting path to static folder

    // Serve static files
    r.Static("/static", filepath.Join("..", "static")) // Serve static files

    // Define the redirect route
    r.GET("/redirect", handlers.RedirectPageHandler)
}
