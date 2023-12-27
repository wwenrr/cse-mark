package handlers

import (
	"gopkg.in/telebot.v3"
	"thuanle/cse-mark/internal/data"
	"thuanle/cse-mark/internal/validation"
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
