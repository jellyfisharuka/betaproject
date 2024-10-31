package gooogle

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func WelcomePush(userEmail string) {
	envPath := filepath.Join("..", "pkg", ".env") 
	err := godotenv.Load(envPath)
    if err != nil {
        log.Fatal("Ошибка при загрузке .env файла")
    }
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		log.Fatal("SENDGRID_API_KEY не установлен")
	}
	from := mail.NewEmail("Beta project", "arukeulen")
	subject := "Thank You For Registering"
	to := mail.NewEmail("Example User", userEmail)
	plainTextContent := "Hi "+ userEmail+ "!\n\n" +
    "Thank you for joining our beta app! We’re thrilled to have you as one of our first users.\n" +
    "We hope you enjoy exploring the features of our project.\n\n" +
    "Your feedback is incredibly important to us, so please don’t hesitate to share your thoughts and suggestions.\n" +
    "We’re committed to making our product better with your help!\n\n" +
    "We look forward to hearing from you!\n\n" +
    "Best regards,\n" +
    "Aruzhan"
	htmlContent := "<strong>Hi " + userEmail + "!</strong><br><br>" +
    "Thank you for joining our beta app! We’re thrilled to have you as one of our first users.<br>" +
    "We hope you enjoy exploring the features of our project.<br><br>" +
    "<img src='https://yt3.googleusercontent.com/82UnWmIFlgRnWam4R5tqS3qv-MaawpGx0QFLSYM5mrABFO1_XyFF7GRJLxToIU-gD9i4K_fc_w=s900-c-k-c0x00ffffff-no-rj' alt='Description of image' style='max-width: 100%; height: auto;'><br><br>" + // Замените на вашу ссылку
    "Your feedback is incredibly important to us, so please don’t hesitate to share your thoughts and suggestions.<br>" +
    "We’re committed to making our product better with your help!<br><br>" +
    "We look forward to hearing from you!<br><br>" +
    "Best regards,<br>" +
    "Aruzhan"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}