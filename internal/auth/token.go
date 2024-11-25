package auth

import (
	"betaproject/internal/db"
	"betaproject/internal/models"
	"errors"
	"fmt"
	"strings"

	//"log"
	"net/http"
	"os"

	"time"

	//"crypto/rand"
	//"encoding/hex"

	"github.com/gin-gonic/gin"
	//"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var jwtKey = []byte(os.Getenv("my_secret"))

func GenerateToken(username string, ID int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"ID":       ID,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString(jwtKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil || tokenString == "" {
		fmt.Println("No token found in cookies")
		tokenString = c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userID, ok := claims["ID"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
			c.Abort()
			return
		}

		c.Set("ID", int(userID))
		c.Next()
	}
} 

func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil || tokenString == "" {
		fmt.Println("No token found in cookies")
		tokenString = c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}
		roleID, ok := claims["roleID"].(float64)
		fmt.Println("Claims: ", claims)
		if !ok || roleID != 1 {  // Проверка, если роль не админ
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
        c.Next()
    }
}

/*func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Извлекаем токен после "Bearer "
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Проверяем токен JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Извлекаем claims из токена
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Извлекаем userID из claims
		userID, ok := claims["ID"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
			c.Abort()
			return
		}

		// Проверка токена в Redis
		redisKey := fmt.Sprintf("token:%d", int(userID)) // создаем ключ на основе userID
		storedToken, err := db.RedisClient.Get(db.Ctx, redisKey).Result()
		if err == redis.Nil {
			// Если токена нет в Redis
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in Redis"})
			c.Abort()
			return
		} else if err != nil {
			// Если произошла ошибка соединения с Redis
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		// Сравниваем токен из заголовка и из Redis
		if tokenString != storedToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token mismatch"})
			c.Abort()
			return
		}

		// Устанавливаем userID в контекст
		c.Set("ID", int(userID))
		c.Next()
	}
} */


 // Используйте ваш секретный ключ

// parseToken разбирает токен и возвращает claims, если он валиден
func parseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что используется правильный метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Извлекаем claims из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not parse claims")
	}

	// Проверяем истечение токена
	if exp, ok := claims["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return nil, fmt.Errorf("token expired")
		}
	}

	return claims, nil
}


func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("ID")

		var user models.User
		if err := db.DB.Preload("Role").First(&user, userID).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if user.Role.Name != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func SignupUser(db *gorm.DB, newUser models.User) error {
	var existingUser models.User
	result := db.Where("username = ?", newUser.Username).First(&existingUser)
	if result.Error == nil {
		return fmt.Errorf("Username already exists")
	}

	var userRole models.Role
	//if err := db.Where("name = ?", "user").First(&userRole).Error; err != nil {
	//	return fmt.Errorf("Default role not found")
	//}
	if newUser.Username=="admin" {
      if err:=db.Where("name=?","admin").First(&userRole).Error; err!=nil {
		return fmt.Errorf("Admin role not found")
	  }
	}else {
		if err:=db.Where("name=?", "user").First(&userRole).Error; err!=nil {
			return fmt.Errorf("Default role not found")
		}
	}
	
	newUser.RoleID = userRole.ID

	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		return fmt.Errorf("Failed to hash password")
	}
	newUser.Password = hashedPassword
	if err := db.Create(&newUser).Error; err != nil {
		return fmt.Errorf("Failed to create user")
	}

	return nil
}

func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return false, errors.New("token is expired")
			}
		}
		return true, nil
	} else {
		return false, err
	}
}
