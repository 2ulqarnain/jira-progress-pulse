package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type JiraConfig struct {
	BaseURL  string `env:"JIRA_BASE_URL,required"`
	Username string `env:"JIRA_USERNAME,required"` // Email for Jira Cloud
	APIToken string `env:"JIRA_API_TOKEN,required"`
}

type AppConfig struct {
	Port        string `env:"PORT" envDefault:"3000"`
	GobFilePath string `env:"DB_FILENAME" envDefault:"storage/data.gob"`
	Jira        *JiraConfig
}

var Cfg *AppConfig

func Load() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatalf("Error loading env variables: %v", envErr)
	}

	Cfg = &AppConfig{
		Port:        os.Getenv("PORT"),
		GobFilePath: os.Getenv("DB_FILENAME"),
		Jira: &JiraConfig{
			BaseURL:  os.Getenv("JIRA_BASE_URL"),
			Username: os.Getenv("JIRA_USERNAME"),
			APIToken: os.Getenv("JIRA_API_TOKEN"),
		},
	}
}
