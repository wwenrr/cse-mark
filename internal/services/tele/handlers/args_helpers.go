package handlers

import (
	"gopkg.in/telebot.v3"
	"thuanle/cse-mark/internal/services/tele/models"
)

func args2Str(c telebot.Context) (string, error) {
	args := c.Args()
	if len(args) != 1 {
		return "", models.NewArgCountMismatchError(1, len(args))
	}
	return args[0], nil
}

func args2StrStr(c telebot.Context) (string, string, error) {
	args := c.Args()
	if len(args) != 2 {
		return "", "", models.NewArgCountMismatchError(2, len(args))
	}
	return args[0], args[1], nil
}

func args2StrDbool(c telebot.Context, arg2Def bool) (string, bool, error) {
	args := c.Args()
	if len(args) == 1 {
		return args[0], arg2Def, nil
	} else if len(args) == 2 {
		arg2 := args[1]
		arg2Value := !(arg2 == "0" || arg2 == "false" || arg2 == "f" || arg2 == "off" || arg2 == "no" || arg2 == "n")
		return args[0], arg2Value, nil
	} else {
		return "", false, models.NewArgCountMismatchError(2, len(args))
	}
}
