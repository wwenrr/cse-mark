package handlers

import "gopkg.in/telebot.v3"

func args2Str(c telebot.Context) (string, error) {
	args := c.Args()
	if len(args) != 1 {
		return "", sendErrorArgumentCountMismatch(c, 1, len(args))
	}
	return args[0], nil
}

func args2StrStr(c telebot.Context) (string, string, error) {
	args := c.Args()
	if len(args) != 2 {
		return "", "", sendErrorArgumentCountMismatch(c, 2, len(args))
	}
	return args[0], args[1], nil
}
