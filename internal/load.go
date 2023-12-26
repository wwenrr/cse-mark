package internal

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"thuanle/cse-mark/internal/configs"
	"thuanle/cse-mark/internal/services/db"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	loadEnv()
	err = db.Instance().Load()
	if err != nil {
		return
	}
}

func loadEnv() {
	adminStrs := os.Getenv("ADMINS")

	err := json.Unmarshal([]byte(adminStrs), &configs.AdminChatIds)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load admins")
	}
	log.Info().Ints64("id", configs.AdminChatIds).Msg("Loaded admins")

	configs.MongoHost = os.Getenv("MONGO_HOST")
	configs.MongoPort = os.Getenv("MONGO_PORT")
	value, ok := os.LookupEnv("MONGO_DB")
	if ok {
		configs.MongoDb = value
	}
}

func Unload() {
	db.Instance().Disconnect()
}
