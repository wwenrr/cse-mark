package middlewares

import (
	"errors"
	"gopkg.in/telebot.v3"
	"thuanle/cse-mark/internal/data"
)

func Teacher(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Callback() != nil {
			defer c.Respond()
		}

		chatUsername := c.Chat().Username
		if data.IsTeacher(chatUsername) {
			return next(c)
		}
		return errors.New("you are not a teacher")
	}
}
