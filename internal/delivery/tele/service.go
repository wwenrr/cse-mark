package tele

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"thuanle/cse-mark/internal/configs"
	"thuanle/cse-mark/internal/delivery/tele/handlers"
	"thuanle/cse-mark/internal/delivery/tele/middlewares"
	"time"
)

type Service struct {
	bot *telebot.Bot
}

var commands = []telebot.Command{
	{
		Text:        "mark",
		Description: "/mark <course> <student_id> - Get mark of course",
	},
	{
		Text:        "load",
		Description: "/load <course> <link> - For teacher, load course marks from link",
	},
	{
		Text:        "clear",
		Description: "/clear - Clear query history. For teacher, clear course link",
	},
	{
		Text:        "my",
		Description: "/my - Your profile",
	},
}

func NewService(config *configs.Config,
	guestHandler *handlers.Guest, teacherHandler *handlers.Teacher, adminHandler *handlers.Admin,
	teacherOnlyMiddleware *middlewares.TeacherOnly) (*Service, error) {
	pref := telebot.Settings{
		Token:  config.TeleToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create telegram bot")
		return nil, err
	}

	if err := b.SetCommands(commands); err != nil {
		log.Fatal().Err(err).Msg("failed to set up telegram commands")
		return nil, err
	}

	b.Use(middlewares.SendErrorMiddleware)
	b.Handle("/start", guestHandler.Start)
	b.Handle("/mark", guestHandler.GetMark)
	b.Handle(telebot.OnText, guestHandler.GetMark)

	teacherOnly := b.Group()
	teacherOnly.Use(teacherOnlyMiddleware.Handle)
	teacherOnly.Handle("/my", teacherHandler.GetMyProfile)
	teacherOnly.Handle("/load", teacherHandler.LoadCourseLink)
	teacherOnly.Handle("/clear", teacherHandler.ClearCourseLink)

	adminOnly := b.Group()
	adminOnly.Use(middleware.Whitelist(config.TeleAdminChatIds...))
	adminOnly.Handle("/teacher", adminHandler.SetTeacher)

	return &Service{
		bot: b,
	}, nil
}

func (s *Service) Run() {
	log.Info().Msg("Starting telegram bot")
	s.bot.Start()
}
