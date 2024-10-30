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
	


