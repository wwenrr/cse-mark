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
		Description: "/mark <subject> <student_id> - Get mark of subject",
	},
	{
		Text:        "load",
		Description: "/load <subject> <link> - Load marks of subject from link",
	},
	{
		Text:        "clear",
		Description: "/clear <subject> - Clear marks of subject",
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
	b.Use(middlewares.SendErrorMiddleware)

	b.Handle("/hello", func(c telebot.Context) error {
		return c.Send("Hello!")
	})

	b.Handle("/mark", handlers.GetMark)

	adminOnly := b.Group()
	adminOnly.Use(middleware.Whitelist(configs.AdminChatIds...))

	adminOnly.Handle("/load", handlers.AdminLoadMark)
	adminOnly.Handle("/clear", handlers.AdminClearMarks)

	log.Info().Msg("Bot started")
	b.Start()
}
