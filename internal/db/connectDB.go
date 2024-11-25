package db

import (
	"betaproject/internal/config"
	"betaproject/internal/models"
	"context"
	"log"
	"time"

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
    for i := 0; i < 5; i++ {
        RedisClient = redis.NewClient(&redis.Options{
            Addr: "redis:6379",
            DB:   0,
        })

        err := RedisClient.Ping(Ctx).Err()
        if err == nil {
            log.Println("Successfully connected to Redis")
            return
        }

        log.Printf("Redis not ready, retrying... (%d/5)", i+1)
        time.Sleep(2 * time.Second)
    }
    log.Fatalf("Could not connect to Redis after retries")
}

