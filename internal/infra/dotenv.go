package infra

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func InitDotenv() error {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg(".env file not found, using system environment variables")
		return err
	}

	return err
}
