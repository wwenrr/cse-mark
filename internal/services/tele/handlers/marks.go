package handlers

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
	"thuanle/cse-mark/internal/data"
	"thuanle/cse-mark/internal/services/tele/models"
	"thuanle/cse-mark/internal/validation"
)

func GetMark(c telebot.Context) error {
	course, studentId, err := args2StrStr(c)
	if err != nil {
		return err
	}

	if !validation.ValidateCourseId(course) {
		return models.NewArgValueMismatchError("course invalid")
	}

	if !validation.ValidateStudentId(studentId) {
		return sendErrorArgumentValueMismatch(c, "studentId incorrect")
	}

	log.Info().
		Str("course", course).
		Str("studentId", studentId).
		Msg("Get mark")

	msg, err := data.GetMark(course, studentId)
	if err != nil {
		return err
	}

	return sendPre(c, msg)
}
