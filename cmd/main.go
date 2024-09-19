package main

import (
	"context"
	"log"


	_ "betaproject/docs"
	"betaproject/internal/app"
	//"google.golang.org/api/gmail/v1"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
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
