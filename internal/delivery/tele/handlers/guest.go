package handlers

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
	"strings"
	"thuanle/cse-mark/internal/delivery/tele/handlers/helpers"
	"thuanle/cse-mark/internal/delivery/tele/models"
	"thuanle/cse-mark/internal/domain/course"
	"thuanle/cse-mark/internal/domain/mark"
	"thuanle/cse-mark/internal/domain/user"
)

type Guest struct {
	courseRules *course.Rules

	markRepo mark.Repository
}

func NewGuestHandler(courseRules *course.Rules, markRepo mark.Repository) *Guest {
	return &Guest{
		courseRules: courseRules,

		markRepo: markRepo,
	}
}

func (h *Guest) Start(c telebot.Context) error {
	chatId := c.Chat().ID
	chatUsername := c.Chat().Username

	return helpers.Sendf(c, "Hello @%s (%d)", chatUsername, chatId)
}

func (h *Guest) GetMark(c telebot.Context) error {
	course, studentId, err := helpers.Args2StrStr(c)
	if err != nil {
		//split c.Text() via " " and get course and studentId
		args := strings.Split(c.Text(), " ")
		if len(args) != 2 {
			return err
		}
		course = args[0]
		studentId = args[1]
	}

	if !h.courseRules.IsValidCourseId(course) {
		return models.NewArgValueMismatchError("course invalid")
	}

	if !user.IsValidStudentId(studentId) {
		return helpers.SendErrorArgumentValueMismatch(c, "studentId incorrect")
	}

	log.Info().
		Int64("chatId", c.Chat().ID).
		Str("chatName", c.Chat().Username).
		Str("course", course).
		Str("studentId", studentId).
		Msg("Get mark")

	msg, err := h.markRepo.GetMark(course, studentId)
	if err != nil {
		return err
	}

	return helpers.SendPre(c, msg)
}
