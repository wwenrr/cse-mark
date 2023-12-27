package data

import (
	"errors"
	"github.com/rs/zerolog/log"
	"thuanle/cse-mark/internal/services/db"
)

func GetMark(course string, studentId string) (string, error) {
	msg, err := db.Instance().GetMark(course, studentId)
	if err != nil {
		log.Info().Err(err).Msg("Get mark error")
		return "", errors.New("get mark error")
	}
	return msg, nil
}

func FetchMarksByChatId(chatId int64) ([]string, error) {
	student, err := db.Instance().GetUserByIntId(chatId)
	if err != nil {
		log.Error().
			Int64("chatId", chatId).
			Err(err).
			Msg("Get user error")
		return nil, errors.New("get user not found")
	}

	if student == nil {
		return nil, errors.New("user not found")
	}
	if student.QueryId == "" {
		return nil, errors.New("studentId not found")
	}
	if len(student.QueryCourses) == 0 {
		return nil, errors.New("course not found")
	}

	var marks []string
	for _, course := range student.QueryCourses {
		mark, err := GetMark(course, student.QueryId)
		if err != nil {
			log.Error().Err(err).Msg("Get mark error")
			continue
		}
		marks = append(marks, mark)
	}
	return marks, nil
}
