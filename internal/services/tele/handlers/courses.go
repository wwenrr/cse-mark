package handlers

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
	"net/url"
	"thuanle/cse-mark/internal/data"
	"thuanle/cse-mark/internal/services/tele/models"
	"thuanle/cse-mark/internal/validation"
)

func TeacherLoadCourseLink(c telebot.Context) error {
	course, link, err := args2StrStr(c)
	if err != nil {
		return err
	}

	if !validation.ValidateCourseId(course) {
		return models.NewArgValueMismatchError("course invalid")
	}

	_, err = url.ParseRequestURI(link)
	if err != nil {
		return err
	}

	chatId := c.Chat().ID
	chatUsername := c.Chat().Username

	if !data.AuthorizeModifyCourse(chatUsername, chatId, course) {
		return models.NewUnauthorizedError("cannot modify course")
	}

	log.Info().
		Int64("chatId", chatId).
		Str("chatUsername", chatUsername).
		Str("course", course).
		Str("link", link).
		Msg("Admin store marks")

	err = data.TeacherStoreCourseLink(course, link, chatId, chatUsername)
	if err != nil {
		return err
	}

	msg, err := data.FetchCourseMarks(course, link)
	if err != nil {
		return err
	}

	return sendf(c, msg)
}

func TeacherClearCourseLink(c telebot.Context) error {
	sub, err := args2Str(c)
	if err != nil {
		return err
	}

	if !validation.ValidateCourseId(sub) {
		return models.NewArgValueMismatchError("course invalid")
	}

	if !data.AuthorizeModifyCourse(c.Chat().Username, c.Chat().ID, sub) {
		return models.NewUnauthorizedError("cannot modify course")
	}

	log.Info().Str("sub", sub).Msg("Admin clear marks")

	err = data.ClearCourseMarks(sub)
	if err != nil {
		return err
	}

	return sendf(c, "%s: cleared", sub)
}
