package main

import (
	"context"
	"log"
	_ "betaproject/docs"
	"betaproject/internal/app"
	"betaproject/internal/auth"

	"fmt"
	"os"

	//"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	//"google.golang.org/api/gmail/v1"
)


func main() {
	ctx := context.Background()
	b, err := os.ReadFile("gmail.json")
	if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope) //GmailReadOnly
	if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := auth.GetClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
			log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"
	r, err := srv.Users.Labels.List(user).Do()
	if err != nil {
			log.Fatalf("Unable to retrieve labels: %v", err)
	}
	if len(r.Labels) == 0 {
			fmt.Println("No labels found.")
			return
	}
	fmt.Println("Labels:")
	for _, l := range r.Labels {
			fmt.Printf("- %s\n", l.Name)
	}

	a, err := app.NewApp(ctx)
	fmt.Println("aaallalalalla")

	if err != nil {
		log.Fatalf("Error creating app: %v", err)
	}
	if err := a.Run(); err != nil {
		log.Fatalf("Error running app: %v", err)
	}
	//r.Static("/static", "../static")

	//r.GET("/", handlers.IndexHandler)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server

	/*config := &oauth2.Config{
        ClientID:     "480618086044-ptq8von436k9gslciacgic7c8p563l5r.apps.googleusercontent.com",
        ClientSecret: "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",
        Endpoint:     google.Endpoint,
    }
	token := &oauth2.Token{
		AccessToken: "your-access-token",
		TokenType:   "Bearer",
	}

    // Создаем HTTP клиент с OAuth2 токеном
    getSMTPToken()
    config.Client(context.Background(), token)

    // Отправляем почту с помощью SMTP сервера
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"
    auth := smtp.PlainAuth("", "arukeulen@gmail.com", "", smtpHost)

    to := []string{"recipient@example.com"}
    msg := []byte("Subject: Test Email\n\nThis is a test email sent using OAuth 2.0.")

    erro := smtp.SendMail(smtpHost+":"+smtpPort, auth, "arukeulen@gmail.com", to, msg)
    if erro != nil {
        log.Fatalf("Failed to send email: %v", err)
    }

    fmt.Println("Email sent successfully!")
}
func getSMTPToken() *oauth2.Token {
    return &oauth2.Token{
        AccessToken: "https://oauth2.googleapis.com/token", // Полученный токен
    }
}*/
}
