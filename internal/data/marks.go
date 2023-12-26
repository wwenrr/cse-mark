package data

import (
	"errors"
	"github.com/rs/zerolog/log"
	"thuanle/cse-mark/internal/services/db"
)

func GetMark(sub string, studentId string) (string, error) {
	msg, err := db.Instance().GetMark(sub, studentId)
	if err != nil {
		log.Info().Err(err).Msg("Get mark error")
		return "", errors.New("get mark error")
	}
	return msg, nil
}
