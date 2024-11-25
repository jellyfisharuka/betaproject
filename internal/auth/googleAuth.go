package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
)

func SaveTokenToDB(db *sql.DB, userID string, token *oauth2.Token) error {
    _, err := db.Exec("REPLACE INTO tokens (user_id, access_token, refresh_token, expiry) VALUES (?, ?, ?, ?)",
        userID, token.AccessToken, token.RefreshToken, token.Expiry)
    return err
}
func TokenFromDB(db *sql.DB, userID string) (*oauth2.Token, error) {
    tok := &oauth2.Token{}
    var expiry time.Time

    err := db.QueryRow("SELECT access_token, refresh_token, expiry FROM tokens WHERE user_id = ?", userID).
        Scan(&tok.AccessToken, &tok.RefreshToken, &expiry)

    if err != nil {
        return nil, err 
    }
    tok.Expiry = expiry
    return tok, nil
}

func GetClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := TokenFromFile(tokFile)
	if err != nil {
			tok = GetTokenFromWeb(config)
			SaveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func GetTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
			"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
			log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
			log.Fatalf("Unable to retrieve token from web: %v", err) ///если токен не указать
	}
	return tok
}

func TokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
			return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func SaveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
			log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
func GetOAuth2Config(filePath string) (*oauth2.Config, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("unable to open client secret file: %v", err)
    }
    defer f.Close()

    var config struct {
        Web struct {
            ClientID     string   `json:"client_id"`
            ClientSecret string   `json:"client_secret"`
            RedirectURIs []string `json:"redirect_uris"`
            AuthURI      string   `json:"auth_uri"`
            TokenURI     string   `json:"token_uri"`
        } `json:"web"`
    }

    if err := json.NewDecoder(f).Decode(&config); err != nil {
        return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
    }
	log.Printf("Loaded OAuth2 Config: %+v", config.Web)

    return &oauth2.Config{
        ClientID:     config.Web.ClientID,
        ClientSecret: config.Web.ClientSecret,
        RedirectURL:  config.Web.RedirectURIs[0], 
        Scopes:       []string{
			gmail.GmailReadonlyScope,
			"https://www.googleapis.com/auth/userinfo.email", // для получения email
            "https://www.googleapis.com/auth/userinfo.profile", // для получения информации о профиле
            "https://www.googleapis.com/auth/gmail.readonly",
			"https://www.googleapis.com/auth/gmail.send",
			}, 
        Endpoint: oauth2.Endpoint{
            AuthURL:  config.Web.AuthURI,
            TokenURL: config.Web.TokenURI,
        },
    }, nil
}