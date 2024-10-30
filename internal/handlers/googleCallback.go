package handlers

import (
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

	c.JSON(http.StatusOK, gin.H{"message": "Authorization successful!", "token": tok})
}
