package config

import (
	"encoding/json"
	"os"
)

type ConfigSetup struct {
	GeminiApiKey string `json:"gemini_api_key"`
}

var config ConfigSetup

func LoadConfig() error {
	file, err := os.Open("internal/config.json")
	if err!=nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(&config)
}