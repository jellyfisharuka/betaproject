package gooogle

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type User struct {
	Email string `json:"email"`
}
func GetUserInfo(accessToken string) (string, error) {
	url := "https://www.googleapis.com/oauth2/v1/userinfo?alt=json"
	request, err := http.NewRequest("GET", url, nil)
	if err!=nil {
		return "", err
	}
	request.Header.Set("Authorization", "Bearer " + accessToken)
    client:= &http.Client{}
	resp, err:= client.Do(request)
	if err!=nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode!=http.StatusOK {
		return "", fmt.Errorf("failed to get user info: %s", resp.Status)
	}
	var user User
	if err:=json.NewDecoder(resp.Body).Decode(&user); err!=nil {
		return "", err
	}
	return user.Email, nil
}