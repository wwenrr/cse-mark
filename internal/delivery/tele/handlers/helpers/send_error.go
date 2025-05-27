package helpers

import (
	"fmt"
	"gopkg.in/telebot.v3"
)

func SendErrorMsg(c telebot.Context, msg string) error {
	return Send(c, msg)
}

func SendError(c telebot.Context, err error) error {
	return SendErrorMsg(c, err.Error())
}

func SendErrorArgumentValueMismatch(c telebot.Context, msg string) error {
	return SendErrorMsg(c, fmt.Sprintf("Argument number mismatch: %s", msg))
}

func SendErrorArgumentCountMismatch(c telebot.Context, needed int, actual int) error {
	return SendErrorMsg(c, fmt.Sprintf("Argument number mismatch: needed %d, actual %d", needed, actual))
}
