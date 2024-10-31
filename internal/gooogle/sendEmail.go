package gooogle

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)
type Email struct {
	Raw string `json:"raw"`
}
func SendEmail(accessToken,from, to, subject, body string) error {
	email := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, subject, body)
	emailBase64 := base64.URLEncoding.EncodeToString([]byte(email))
	emailData := Email{Raw:emailBase64}
	jsonData, err := json.Marshal(emailData)
	if err!=nil {
		return err
	}
	url :="https://gmail.googleapis.com/gmail/v1/users/me/messages/send"
	req, err:= http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err!=nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err:= client.Do(req)
	if err!=nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send email: %s", resp.Status)
	}
	return nil
}