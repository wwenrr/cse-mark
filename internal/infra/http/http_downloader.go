package http

import (
	"context"
	"encoding/csv"
	"github.com/rs/zerolog/log"
	"net/http"
	"thuanle/cse-mark/internal/configs"
	"thuanle/cse-mark/internal/domain/downloader"
)

type SimpleDownloader struct {
	Client *http.Client
	ctx    context.Context
}

func NewSimpleDownloader(config *configs.Config) downloader.Repository {
	return &SimpleDownloader{
		Client: &http.Client{
			Timeout: config.DownloaderTimeout,
		},
	}
}

func (d *SimpleDownloader) DownloadCSV(url string) ([][]string, error) {
	log.Info().Str("url", url).Msg("Downloading URL")

	// Make an HTTP GET request to the specified URL
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msg("Error downloading link")
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the CSV data and extract URLs
	reader := csv.NewReader(resp.Body)

	records, err := reader.ReadAll()
	if err != nil {
		log.Error().Err(err).Msg("Error parsing csv")
		return nil, err
	}

	return records, nil
}
