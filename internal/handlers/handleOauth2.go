package handlers

import (
	"betaproject/internal/auth"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
    "betaproject/internal/gooogle"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)
var Oauth2Config *oauth2.Config

 func InitConfig() {
	b, err := os.ReadFile("gmail.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// Инициализация oauth2Config
	Oauth2Config, err = google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
}

func HandleOAuth2Callback(code string, c *gin.Context) (*oauth2.Token, error) {
	if Oauth2Config== nil {
        return nil, fmt.Errorf("oauth2Config is not initialized")
    }
	tok, err := Oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Failed to exchange token: %v", err)
		log.Fatalf("Unable to retrieve token from web: %v", err)
		return nil, err
	}
	accessToken := tok.AccessToken
	fmt.Println(accessToken)
	userEmail, err:= gooogle.GetUserInfo(accessToken)
	if err!=nil {
		fmt.Println("error getting user info:", err)
		return nil, err
	}
	subject := "Blabla"
	body:="bla bla welcome"
	err = gooogle.SendEmail(accessToken, userEmail, subject, body)
	if err!=nil {
		fmt.Println("error sending email: ", err)
		return nil, err
	}
	fmt.Println("Email sent successfully to: ", userEmail)
	session:= sessions.Default(c)
	session.Set("access_token", tok.AccessToken)
	session.Set("refresh_token", tok.RefreshToken)
	session.Set("expiry", tok.Expiry)
	session.Save()
	
	return tok, nil
}

// LoginGoogleHandler redirects the user to the Google OAuth2 login page.
// @Summary Redirects to Google OAuth2 login
// @Description Redirects the user to Google's OAuth2 login page to initiate authorization.
// @Tags Auth
// @Produce json
// @Success 302 {string} string "Redirects to Google login"
// @Failure 500 {string} string "OAuth2 config is not initialized"
// @Router /googleLogin [get]
func LoginGoogleHandler(config *oauth2.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if config == nil {
            c.String(http.StatusInternalServerError, "OAuth2 config is not initialized.")
            return
        }
		authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		c.Redirect(http.StatusFound, authURL)
	}
}
func fetchGmailLabels(ctx context.Context, config *oauth2.Config) {
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
}
