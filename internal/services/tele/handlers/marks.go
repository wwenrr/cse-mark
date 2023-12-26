package handlers

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
	"net/url"
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

func TeacherLoadMarks(c telebot.Context) error {
	sub, link, err := args2StrStr(c)
	if err != nil {
		return err
	}

	if !validation.ValidateCourseId(sub) {
		return models.NewArgValueMismatchError("course invalid")
	}

	_, err = url.ParseRequestURI(link)
	if err != nil {
		return err
	}

	chatId := c.Chat().ID
	chatUsername := c.Chat().Username

	log.Info().
		Int64("chatId", chatId).
		Str("chatUsername", chatUsername).
		Str("sub", sub).
		Str("link", link).
		Msg("Admin store marks")

	err = data.TeacherStoreCourseLink(sub, link, chatId, chatUsername)
	if err != nil {
		return err
	}

	msg, err := data.FetchCourseMarks(sub, link)
	if err != nil {
		return err
	}

	return sendf(c, msg)
}

func TeacherClearMarks(c telebot.Context) error {
	sub, err := args2Str(c)
	if err != nil {
		return err
	}

	if !validation.ValidateCourseId(sub) {
		return models.NewArgValueMismatchError("course invalid")
	}

	log.Info().Str("sub", sub).Msg("Admin clear marks")

	err = data.ClearCourseMarks(sub)
	if err != nil {
		return err
	}

	return sendf(c, "%s: cleared", sub)
}
