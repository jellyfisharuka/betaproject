package handlers

import (
	
	"net/http"
	

	"github.com/gin-gonic/gin"
)
// LoginGoogleHandler redirects the user to the Google OAuth2 login page.
// @Summary Redirects to Google OAuth2 login
// @Description Redirects the user to Google's OAuth2 login page to initiate authorization.
// @Tags Auth
// @Produce json
// @Success 200 {string} string "Redirects to Google login"
// @Failure 500 {string} string "OAuth2 config is not initialized"
// @Router /googleLogin [get]
// @Security googleOAuth2
func RedirectPageHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "redirect.html", nil)
}
/*func RedirectPageHandler(c *gin.Context) {
	htmlFile, err := os.Open(filepath.Join( "static", "redirect.html"))
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading redirect page")
		return
	}
	defer htmlFile.Close()
	htmlContent, err := io.ReadAll(htmlFile)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading redirect page")
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", htmlContent)
} */