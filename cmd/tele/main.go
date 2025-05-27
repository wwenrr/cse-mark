package main

import (
	"github.com/rs/zerolog/log"
	"thuanle/cse-mark/internal/infra"
)

func main() {
	//infrastructure.InitZerolog()
	//
	//internal.Load()
	//
	//tele.Run()
	//
	//defer internal.Unload()

	infra.InitZerolog()
	_ = infra.InitDotenv()

	log.Info().Msg("Initialization completed successfully")

	app, err := InitializeApp()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize application")
		return
	}

	app.TeleService.Run()

	defer app.MongoClient.Disconnect()
}
