package handlers

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
	"strings"
	"thuanle/cse-mark/internal/data"
	"thuanle/cse-mark/internal/domain/validation"
	"thuanle/cse-mark/internal/services/tele/models"
)

func GetMark(c telebot.Context) error {
	course, studentId, err := args2StrStr(c)
	if err != nil {
		//split c.Text() via " " and get course and studentId
		args := strings.Split(c.Text(), " ")
		if len(args) != 2 {
			return err
		}
		course = args[0]
		studentId = args[1]
	}

	if !validation.ValidateCourseId(course) {
		return models.NewArgValueMismatchError("course invalid")
	}

	if !validation.ValidateStudentId(studentId) {
		return sendErrorArgumentValueMismatch(c, "studentId incorrect")
	}

	log.Info().
		Int64("chatId", c.Chat().ID).
		Str("chatName", c.Chat().Username).
		Str("course", course).
		Str("studentId", studentId).
		Msg("Get mark")

	msg, err := data.GetMark(course, studentId)
	if err != nil {
		return err
	}

	err = data.AppendQueryHistory(c.Chat().ID, course, studentId)

	return sendPre(c, msg)
}
