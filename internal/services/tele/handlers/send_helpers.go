package handlers

import (
	"fmt"
	"gopkg.in/telebot.v3"
)

func send(c telebot.Context, message string, opts ...interface{}) error {
	if len(opts) == 0 || opts[0] == nil {
		return c.Send(message, telebot.ModeHTML, telebot.NoPreview)
	}
	return c.Send(message, append(opts, telebot.ModeHTML, telebot.NoPreview)...)
}

func sendf(c telebot.Context, format string, v ...interface{}) error {
	return send(c, fmt.Sprintf(format, v...))
}

func sendPre(c telebot.Context, message string, opts ...interface{}) error {
	return send(c, "<pre>"+message+"</pre>", opts...)
}
