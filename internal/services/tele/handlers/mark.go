package handlers

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
	"net/url"
	"thuanle/cse-mark/internal/data"
)

func GetMark(context telebot.Context) error {
	subject, studentId, err := args2StrStr(context)
	if err != nil {
		return err
	}

	msg, err := data.GetMark(subject, studentId)
	if err != nil {
		return err
	}

	return sendPre(context, msg)
}

func AdminLoadMark(c telebot.Context) error {
	sub, link, err := args2StrStr(c)
	if err != nil {
		return err
	}

	log.Info().Str("sub", sub).Str("link", link).Msg("Admin load mark")

	_, err = url.ParseRequestURI(link)
	if err != nil {
		return err
	}

	msg, err := data.AdminStoreMarks(sub, link)
	if err != nil {
		return err
	}

	return sendf(c, msg)
}

func AdminClearMarks(c telebot.Context) error {
	sub, err := args2Str(c)
	if err != nil {
		return err
	}

	log.Info().Str("sub", sub).Msg("Admin clear marks")

	err = data.AdminClearMarks(sub)
	if err != nil {
		return err
	}

	return sendf(c, "%s: cleared", sub)
}
