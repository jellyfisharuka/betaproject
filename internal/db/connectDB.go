package db

import (
	"betaproject/internal/config"
	"betaproject/internal/models"
	"context"
	"log"

	//"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	DB *gorm.DB
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func ConnectDB() {
	var err error
	//dsn := os.Getenv("DB")
	dsn := config.GetEnvConfig().DbDSN
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}
	log.Println("Successfully connected to database")

	err = DB.AutoMigrate(models.Email{}, models.Role{}, models.User{}, models.FAQ{}, models.Chat{}, models.Message{})
	if err != nil {
		panic("Failed to migrate DB schemas")
	}

}
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", 
		DB:   0,                
	})

	if err := RedisClient.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

