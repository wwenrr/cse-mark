package handlers

import (
	"fmt"
	"gopkg.in/telebot.v3"
)

func sendErrorMsg(c telebot.Context, msg string) error {
	return send(c, msg)
}

func sendError(c telebot.Context, err error) error {
	return sendErrorMsg(c, err.Error())
}

func sendErrorArgumentValueMismatch(c telebot.Context, msg string) error {
	return sendErrorMsg(c, fmt.Sprintf("Argument number mismatch: %s", msg))
}

func sendErrorArgumentCountMismatch(c telebot.Context, needed int, actual int) error {
	return sendErrorMsg(c, fmt.Sprintf("Argument number mismatch: needed %d, actual %d", needed, actual))
}
