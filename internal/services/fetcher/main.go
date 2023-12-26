package fetcher

import (
	"github.com/rs/zerolog/log"
	"time"
)

func Execute() {
	for {
		fetchNewMarks()
		log.Info().Msg("Sleeping for 10 minutes...")
		time.Sleep(10 * time.Minute)
	}
}
