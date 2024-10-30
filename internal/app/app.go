package app

import (
	"betaproject/internal/config"
	"betaproject/internal/db"
	"betaproject/internal/router"

	//"betaproject/internal/router"
	"context"
	"log"
	"net/http"
	"path/filepath"
    _ "net/http/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
)

type App struct {
	router *gin.Engine
}
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}
func (a *App) initConfig(_ context.Context) error {
	jsonPath := filepath.Join("..","internal", "config", "config.json")
	gmailPath:= filepath.Join("..", "internal", "config", "gmail.json")
	envPath := filepath.Join("..", "pkg", ".env")
	err := config.LoadJSONConfig(jsonPath)
    if err != nil {
        return err
    }
    err = config.LoadEnvConfig(envPath)
    if err != nil {
        return err
    }
	err = config.LoadGmailConfig(gmailPath)
	if err != nil {
        return err
    }
	db.ConnectDB()

	return nil

}
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initRouter,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
func (a *App) Run() error {
	//address := "0.0.0.0:8080"
	address := "localhost:8080"
	log.Printf("HTTP server is running on %s", address)
	go func() {
        log.Println("Starting pprof server on :6060")
        log.Println(http.ListenAndServe("localhost:6060", nil)) // Порт для pprof
    }()
	//err := http.ListenAndServe(address, a.router)
	err:= a.router.Run(address)
	if err != nil {
		return err
	}

	return nil
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
func (a *App) initRouter(_ context.Context) error {
	store := cookie.NewStore([]byte("secret"))
	a.router = gin.Default()
	a.router.Use(sessions.Sessions("mysession", store))
	a.router.Use(CORSMiddleware())
	router.SetupRouter(a.router)
	return nil
}