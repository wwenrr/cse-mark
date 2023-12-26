package data

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"thuanle/cse-mark/internal/configs"
	"thuanle/cse-mark/internal/models"
	"thuanle/cse-mark/internal/services/db"
	"time"
)

func GetAllCourses() ([]*models.CourseSettingsModel, error) {
	updatedAfter := time.Now().Add(-configs.FetchMaxAge).Unix()
	return db.Instance().GetAllCourses(updatedAfter)
}

func fetchCsv(link string) ([][]string, error) {
	log.Debug().Msg("Download links...")

	// Make an HTTP GET request to the specified URL
	resp, err := http.Get(link)
	if err != nil {
		log.Error().Err(err).Msg("Error downloading link")
		return nil, errors.New("error downloading link")
	}
	defer resp.Body.Close()

	// Parse the CSV data and extract URLs
	reader := csv.NewReader(resp.Body)

	records, err := reader.ReadAll()
	if err != nil {
		log.Error().Err(err).Msg("Error parsing csv")
		return nil, errors.New("error parsing csv")
	}

	return records, nil
}

func FetchCourseMarks(course string, link string) (string, error) {
	records, err := fetchCsv(link)
	if err != nil {
		return "", err
	}

	log.Debug().
		Str("course", course).
		Strs("Flags", records[0]).
		Strs("Headers", records[1]).
		Msg("Record fetched")

	cleanData, err := cleanData(records)
	if err != nil {
		return "", err
	}

	log.Debug().Interface("cleanData", cleanData).Msg("Cleaned data")

	err = db.Instance().StoreMarks(course, *cleanData)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s: Store %d records", course, len(*cleanData)), nil
}

func cleanData(records [][]string) (*[]map[string]string, error) {
	if len(records) < 2 {
		return nil, errors.New("invalid csv file. Need at least 2 rows")
	}
	var result []map[string]string
	flags := records[0]
	headers := records[1]
	data := records[2:]

	for _, row := range data {
		value := make(map[string]string)
		for fidx, flag := range flags {
			if len(flag) > 0 {
				if flag == "id" {
					value["_id"] = row[fidx]
				} else {
					value[headers[fidx]] = row[fidx]
				}
			}
		}
		result = append(result, value)
	}

	return &result, nil
}

func ClearCourseMarks(sub string) error {
	log.Debug().
		Str("course", sub).
		Msg("Clear marks")

	err := db.Instance().ClearMarks(sub)

	return err
}
