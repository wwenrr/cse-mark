package tele

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"os"
	"thuanle/cse-mark/internal/configs"
	"thuanle/cse-mark/internal/services/tele/handlers"
	"thuanle/cse-mark/internal/services/tele/middlewares"
	"time"
)

var commands = []telebot.Command{
	{
		Text:        "mark",
		Description: "/mark <course> <student_id> - Get mark of course",
	},
	{
		Text:        "load",
		Description: "/load <course> <link> - Load marks of course from link",
	},
	{
		Text:        "clear",
		Description: "/clear <course> - Clear marks of course",
	},
}

func Execute() {
	pref := telebot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create bot")
		return
	}

	if err := b.SetCommands(commands); err != nil {
		log.Fatal().Err(err).Msg("setup telebot command")
		return
	}
	//b.Use(middleware.Logger())
	b.Use(middlewares.SendErrorMiddleware)
	b.Handle("/start", handlers.Hello)
	b.Handle("/mark", handlers.GetMark)

	teacherOnly := b.Group()
	teacherOnly.Use(middlewares.Teacher)
	teacherOnly.Handle("/load", handlers.TeacherLoadCourseLink)
	teacherOnly.Handle("/clear", handlers.TeacherClearCourseLink)

	adminOnly := b.Group()
	adminOnly.Use(middleware.Whitelist(configs.AdminChatIds...))
	adminOnly.Handle("/teacher", handlers.AdminSetTeacher)

	log.Info().Msg("Bot started")
	b.Start()
}
