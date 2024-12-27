package data

import (
	"github.com/rs/zerolog/log"
	"slices"
	"thuanle/cse-mark/internal/configs"
	"thuanle/cse-mark/internal/services/db"
)

func AuthorizeModifyCourse(user string, chatId int64, course string) bool {
	if slices.Contains(configs.AdminChatIds, chatId) {
		log.Debug().
			Str("user", user).
			Int64("chatId", chatId).
			Ints64("adminChatIds", configs.AdminChatIds).
			Msg("Admin can modify course")
		return true
	}

	courseSettings, err := db.Instance().GetCourseById(course)
	if err != nil || courseSettings == nil {
		log.Debug().
			Any("course", courseSettings).
			Err(err).
			Msg("Get course settings error or new course")
		return true
	}

	log.Debug().
		Str("course", course).
		Str("user", user).
		Int64("chatId", chatId).
		Str("byUser", courseSettings.ByUser).
		Int64("byId", courseSettings.ById).
		Err(err).
		Msg("Authorize modify course")

	return courseSettings.ByUser == user || courseSettings.ById == chatId
}
