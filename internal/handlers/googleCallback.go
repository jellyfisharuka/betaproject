package handlers

import (
	"betaproject/internal/auth"
//	"betaproject/internal/db"
	"betaproject/internal/gooogle"
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OAuth2CallbackHandler(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth code not found"})
		return
	}
	tok, err := HandleOAuth2Callback(code, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve token from web"})
		return
	}
    email, err := gooogle.GetUserInfo(tok.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
    username := email 
	jwtToken, err := auth.GenerateToken(username, 0) // ID можно задавать позже
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}
	//err = db.RedisClient.Set(db.Ctx, fmt.Sprintf("token:%s", email), jwtToken, 0).Err()
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store token"})
	//	return
	//}

	c.SetCookie("auth_token", jwtToken, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Authorization successful!", "token": jwtToken})
}
