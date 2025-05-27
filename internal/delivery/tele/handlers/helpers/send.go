package helpers

import (
	"fmt"
	"gopkg.in/telebot.v3"
)

func Send(c telebot.Context, message string, opts ...interface{}) error {
	if len(opts) == 0 || opts[0] == nil {
		return c.Send(message, telebot.ModeHTML, telebot.NoPreview)
	}
	return c.Send(message, append(opts, telebot.ModeHTML, telebot.NoPreview)...)
}

func Sendf(c telebot.Context, format string, v ...interface{}) error {
	return Send(c, fmt.Sprintf(format, v...))
}

func SendPre(c telebot.Context, message string, opts ...interface{}) error {
	return Send(c, "<pre>"+message+"</pre>", opts...)
}
