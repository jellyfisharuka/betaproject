package handlers

import (
	"betaproject/internal/auth"
	"betaproject/internal/db"
	"betaproject/internal/models"
	"fmt"
	"net/http"
	

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)


// @Summary Login user
// @Description Logs in an existing user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   user  body      models.LoginSwagger  true  "New user information"  Example: {"username": "exampleUser", "password": "examplePassword"}
// @Success 200    {object}  map[string]string
// @Failure 400    {object}  map[string]string
// @Failure 401    {object}  map[string]string
// @Router  /login [post]
func LoginHandler(c *gin.Context) {
	var loginUser models.User
	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	var existingUser models.User
	result := db.DB.Select("ID", "username", "password").Where("username = ?", loginUser.Username).First(&existingUser)
	if result.Error != nil || !auth.CheckPassword(existingUser.Password, loginUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := auth.GenerateToken(existingUser.Username, int(existingUser.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	fmt.Println("Generated token:", token)
	err = db.RedisClient.Set(db.Ctx, fmt.Sprintf("token:%s", existingUser.Username), token, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store token"})
		return
	}
	c.SetCookie("auth_token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"token": token})

}

// @Summary Sign up user
// @Description Signs up a new user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   user  body      models.SignupSwagger  true  "New user information"  Example: {"username": "exampleUser", "password": "examplePassword", "firstName": "John", "lastName": "Doe"}
// @Success 201    {object}  map[string]string
// @Failure 400    {object}  map[string]string
// @Failure 500    {object}  map[string]string
// @Router  /signup [post]
func SignupHandler(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	validate := validator.New()

	if err := validate.Struct(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
	if err := auth.SignupUser(db.DB, newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign up user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully"})

}

