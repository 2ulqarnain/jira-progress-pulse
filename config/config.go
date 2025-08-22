package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't Load .env Files !")
	}
	fmt.Println(".env file loaded !")
}
