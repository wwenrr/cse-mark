package middlewares

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
)

func SendErrorMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		err := next(c) // continue execution chain
		if err != nil {
			log.Debug().Err(err).Msg("Chat Middleware error reply")
			return c.Send(fmt.Sprintf("Error: %v", err), telebot.RemoveKeyboard)
		}
		return nil
	}
}
