package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)


type JSONConfig struct {
	GeminiApiKey string `json:"gemini_api_key"`
}

var jsonConfig JSONConfig

type EnvConfig struct {
    DbDSN string 
}
var envConfig EnvConfig

type GmailConfig struct {
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
var gmailConfig GmailConfig
func LoadJSONConfig(jsonPath string) error {
	file, err := os.Open(jsonPath)
	if err!=nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(&jsonConfig)
}
func LoadEnvConfig(envPath string) error {
    err := godotenv.Load(envPath)
    if err != nil {
        log.Println("Error loading .env file, falling back to system environment variables")
    }

    envConfig.DbDSN = os.Getenv("DB")

    return nil
}
func LoadGmailConfig(gmailPath string) error {
	file, err := os.Open(gmailPath)
	if err!=nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(&gmailConfig)
}

func GetJSONConfig() *JSONConfig {
    return &jsonConfig
}

func GetEnvConfig() *EnvConfig {
    return &envConfig
}

func GetGmailConfig() *GmailConfig {
	return &gmailConfig
}


