package handlers

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
	"thuanle/cse-mark/internal/data"
	"thuanle/cse-mark/internal/domain/validation"
	"thuanle/cse-mark/internal/services/tele/views"
)

func Hello(c telebot.Context) error {
	chatId := c.Chat().ID
	chatUsername := c.Chat().Username

	return sendf(c, "Hello @%s (%d)", chatUsername, chatId)
}

func AdminSetTeacher(c telebot.Context) error {
	name, grant, err := args2StrDbool(c, true)
	if err != nil {
		return err
	}

	if !validation.ValidateTelegramUsername(name) {
		return sendErrorArgumentValueMismatch(c, "name incorrect")
	}

	chatName := c.Chat().Username

	err = data.AdminGrantTeacher(name, grant, chatName)
	if err != nil {
		return err
	}

	if grant {
		return sendf(c, "Set %s as teacher", name)
	} else {
		return sendf(c, "Remove %s from teacher", name)
	}
}

func GetMyProfile(c telebot.Context) error {
	chatUsername := c.Chat().Username
	if data.IsTeacher(chatUsername) {
		return GetTeacherProfile(c)
	} else {
		return GetStudentProfile(c)
	}
}

func GetStudentProfile(c telebot.Context) error {
	chatId := c.Chat().ID

	log.Info().
		Int64("chatId", chatId).
		Msg("Get student profile")

	marks, err := data.FetchMarksByChatId(chatId)
	if err != nil {
		return err
	}

	for course, mark := range marks {
		send(c, course+":\n<pre>\n"+mark+"\n</pre>")
	}

	return nil
}

func GetTeacherProfile(c telebot.Context) error {
	chatUsername := c.Chat().Username

	log.Info().
		Str("chatUsername", chatUsername).
		Msg("Get teacher profile")

	courses, err := data.FetchTeacherCourses(chatUsername)
	if err != nil {
		return err
	}

	msg := views.RenderTeacherProfile(courses)
	return sendPre(c, msg)
}

func Clear(c telebot.Context) error {
	chatUsername := c.Chat().Username

	if data.IsTeacher(chatUsername) {
		return TeacherClearCourseLink(c)
	} else {
		return StudentClearQueries(c)
	}
}

func StudentClearQueries(c telebot.Context) error {
	chatId := c.Chat().ID

	err := data.ClearStudentQueries(chatId)
	if err != nil {
		return err
	}

	return sendf(c, "Clear queries success")
}
