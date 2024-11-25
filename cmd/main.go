// @title Beta Project
// @version 1.0
// @description This is a sample server.
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @securityDefinitions.oauth2.authorizationCode googleOAuth2
// @tokenUrl https://oauth2.googleapis.com/token
// @authorizationUrl https://accounts.google.com/o/oauth2/auth
// @scope.email Access to your email
// @scope.profile Access to your profile information
// @host localhost:8080
// @BasePath /
package main

import (
	_ "betaproject/docs"
	"betaproject/internal/app"
	
	"betaproject/internal/handlers"
	"context"
	"log"
	"sync"



	//"golang.org/x/oauth2"

	
	//"google.golang.org/api/gmail/v1"
)

// swag init
//swag init -g cmd/main.go
// sudo service redis-server start
// redis-cli  GEt token:aruke  KEYS token:*
func main() {
	var wg sync.WaitGroup
	ctx := context.Background()
	handlers.InitConfig()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
	   a, err := app.NewApp(ctx)
   
	   if err != nil {
		   log.Fatalf("Error creating app: %v", err)
	   }
	   if err := a.Run(); err != nil {
		   log.Fatalf("Error running app: %v", err)
	   }
   }()

	wg.Wait()
	
}
	


